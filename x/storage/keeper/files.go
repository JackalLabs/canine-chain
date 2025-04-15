package keeper

import (
	"encoding/base64"
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/jackalLabs/canine-chain/v4/x/storage/types"
)

// SetFile sets a specific File in the SQLite database
func (k Keeper) SetFile(ctx sdk.Context, file types.UnifiedFile) error {
	// Start a transaction
	tx, err := k.filebase.Begin()
	if err != nil {

		_ = tx.Rollback()
		return err
	}

	// Insert into unified_files table
	_, err = tx.Exec(`
		REPLACE INTO unified_files 
		(merkle, owner, start, expires, file_size, proof_interval, proof_type, max_proofs, note)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	`, file.Merkle, file.Owner, file.Start, file.Expires, file.FileSize, file.ProofInterval,
		file.ProofType, file.MaxProofs, file.Note)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Delete existing proofs for this file
	_, err = tx.Exec(`
		DELETE FROM proofs 
		WHERE file_merkle = ? AND file_owner = ? AND file_start = ?
	`, file.Merkle, file.Owner, file.Start)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Insert new proofs
	for _, proof := range file.Proofs {
		_, err = tx.Exec(`
			INSERT INTO proofs (file_merkle, file_owner, file_start, proof)
			VALUES (?, ?, ?, ?)
		`, file.Merkle, file.Owner, file.Start, proof)
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

// GetFile returns a File from its primary key
func (k Keeper) GetFile(
	ctx sdk.Context,
	merkle []byte,
	owner string,
	start int64,
) (val types.UnifiedFile, found bool) {
	// Query file data
	row := k.filebase.QueryRow(`
		SELECT merkle, owner, start, expires, file_size, proof_interval, proof_type, max_proofs, note
		FROM unified_files
		WHERE merkle = ? AND owner = ? AND start = ?
	`, merkle, owner, start)

	// Initialize file struct
	var file types.UnifiedFile
	err := row.Scan(
		&file.Merkle, &file.Owner, &file.Start, &file.Expires, &file.FileSize,
		&file.ProofInterval, &file.ProofType, &file.MaxProofs, &file.Note,
	)
	if err != nil {
		return val, false
	}

	// Query associated proofs
	rows, err := k.filebase.Query(`
		SELECT proof FROM proofs
		WHERE file_merkle = ? AND file_owner = ? AND file_start = ?
	`, merkle, owner, start)
	if err != nil {
		return val, false
	}
	defer rows.Close()

	// Collect proofs
	var proofs []string
	for rows.Next() {
		var proof string
		if err := rows.Scan(&proof); err != nil {
			return val, false
		}
		proofs = append(proofs, proof)
	}
	file.Proofs = proofs

	return file, true
}

// RemoveFile removes a File from the database
func (k Keeper) RemoveFile(
	ctx sdk.Context,
	merkle []byte,
	owner string,
	start int64,
) error {
	// Start a transaction
	tx, err := k.filebase.Begin()
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	// Delete the file
	_, err = tx.Exec(`
		DELETE FROM unified_files 
		WHERE merkle = ? AND owner = ? AND start = ?
	`, merkle, owner, start)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()

}

func (k Keeper) GetAllFileByMerklePgWithJSONFilter(offset, limit uint64, jsonKey string, jsonValue string) (list []types.UnifiedFile, total int64) {
	if offset == 0 && limit == 0 {
		limit = 100
	}

	// Get total count with JSON filter
	countQuery := `
        SELECT COUNT(*) FROM unified_files
        WHERE json_extract(note, ?) = ?
    `
	err := k.filebase.QueryRow(countQuery, "$."+jsonKey, jsonValue).Scan(&total)
	if err != nil {
		return nil, 0
	}

	// Query with pagination and JSON filter
	rows, err := k.filebase.Query(`
        SELECT merkle, owner, start, expires, file_size, proof_interval, proof_type, max_proofs, note
        FROM unified_files
        WHERE json_extract(note, ?) = ?
        ORDER BY merkle, owner, start
        LIMIT ? OFFSET ?
    `, "$."+jsonKey, jsonValue, limit, offset)
	if err != nil {
		return nil, total
	}
	defer rows.Close()

	files := make(map[string]types.UnifiedFile)
	var fileKeys []string
	for rows.Next() {
		var file types.UnifiedFile
		err := rows.Scan(
			&file.Merkle, &file.Owner, &file.Start, &file.Expires, &file.FileSize,
			&file.ProofInterval, &file.ProofType, &file.MaxProofs, &file.Note,
		)
		if err != nil {
			continue
		}

		key := fmt.Sprintf("%s:%s:%d", base64.StdEncoding.EncodeToString(file.Merkle), file.Owner, file.Start)
		files[key] = file
		fileKeys = append(fileKeys, key)
	}

	if len(fileKeys) == 0 {
		return nil, total
	}

	// Get proofs for these files
	for _, key := range fileKeys {
		file := files[key]
		proofRows, err := k.filebase.Query(`
            SELECT proof FROM proofs
            WHERE file_merkle = ? AND file_owner = ? AND file_start = ?
        `, file.Merkle, file.Owner, file.Start)
		if err != nil {
			continue
		}

		var proofs []string
		for proofRows.Next() {
			var proof string
			if err := proofRows.Scan(&proof); err != nil {
				continue
			}
			proofs = append(proofs, proof)
		}
		proofRows.Close()

		file.Proofs = proofs
		files[key] = file
	}

	// Convert map to ordered list
	for _, key := range fileKeys {
		list = append(list, files[key])
	}

	return list, total
}

func (k Keeper) GetAllFileByMerklePg(offset, limit uint64) (list []types.UnifiedFile, total int64) {
	if offset == 0 && limit == 0 {
		limit = 100
	}

	// Get total count
	err := k.filebase.QueryRow("SELECT COUNT(*) FROM unified_files").Scan(&total)
	if err != nil {
		return nil, 0
	}

	// Query with pagination
	rows, err := k.filebase.Query(`
        SELECT merkle, owner, start, expires, file_size, proof_interval, proof_type, max_proofs, note
        FROM unified_files
        ORDER BY merkle, owner, start
        LIMIT ? OFFSET ?
    `, limit, offset)
	if err != nil {
		return nil, total
	}
	defer rows.Close()

	files := make(map[string]types.UnifiedFile)
	var fileKeys []string
	for rows.Next() {
		var file types.UnifiedFile
		err := rows.Scan(
			&file.Merkle, &file.Owner, &file.Start, &file.Expires, &file.FileSize,
			&file.ProofInterval, &file.ProofType, &file.MaxProofs, &file.Note,
		)
		if err != nil {
			continue
		}

		key := fmt.Sprintf("%s:%s:%d", base64.StdEncoding.EncodeToString(file.Merkle), file.Owner, file.Start)
		files[key] = file
		fileKeys = append(fileKeys, key)
	}

	if len(fileKeys) == 0 {
		return nil, total
	}

	// Get proofs for these files
	for _, key := range fileKeys {
		file := files[key]
		proofRows, err := k.filebase.Query(`
            SELECT proof FROM proofs
            WHERE file_merkle = ? AND file_owner = ? AND file_start = ?
        `, file.Merkle, file.Owner, file.Start)

		if err != nil {
			continue
		}

		var proofs []string
		for proofRows.Next() {
			var proof string
			if err := proofRows.Scan(&proof); err != nil {
				continue
			}
			proofs = append(proofs, proof)
		}
		proofRows.Close()

		file.Proofs = proofs
		files[key] = file
	}

	// Convert map to ordered list
	for _, key := range fileKeys {
		list = append(list, files[key])
	}

	return list, total
}

func (k Keeper) GetAllFilesWithMerklePg(merkle []byte, offset, limit uint64) (list []types.UnifiedFile, total int64) {
	if offset == 0 && limit == 0 {
		limit = 100
	}

	// Get total count
	err := k.filebase.QueryRow("SELECT COUNT(*) FROM unified_files WHERE merkle = ?", merkle).Scan(&total)
	if err != nil {
		return nil, 0
	}

	// Query with pagination
	rows, err := k.filebase.Query(`
        SELECT merkle, owner, start, expires, file_size, proof_interval, proof_type, max_proofs, note
        FROM unified_files WHERE merkle = ?
        ORDER BY merkle, owner, start
        LIMIT ? OFFSET ?
    `, merkle, limit, offset)
	if err != nil {
		return nil, total
	}
	defer rows.Close()

	files := make(map[string]types.UnifiedFile)
	var fileKeys []string
	for rows.Next() {
		var file types.UnifiedFile
		err := rows.Scan(
			&file.Merkle, &file.Owner, &file.Start, &file.Expires, &file.FileSize,
			&file.ProofInterval, &file.ProofType, &file.MaxProofs, &file.Note,
		)
		if err != nil {
			continue
		}

		key := fmt.Sprintf("%s:%s:%d", base64.StdEncoding.EncodeToString(file.Merkle), file.Owner, file.Start)
		files[key] = file
		fileKeys = append(fileKeys, key)
	}

	if len(fileKeys) == 0 {
		return nil, total
	}

	// Get proofs for these files
	for _, key := range fileKeys {
		file := files[key]
		proofRows, err := k.filebase.Query(`
            SELECT proof FROM proofs
            WHERE file_merkle = ? AND file_owner = ? AND file_start = ?
        `, file.Merkle, file.Owner, file.Start)

		if err != nil {
			continue
		}

		var proofs []string
		for proofRows.Next() {
			var proof string
			if err := proofRows.Scan(&proof); err != nil {
				continue
			}
			proofs = append(proofs, proof)
		}
		proofRows.Close()

		file.Proofs = proofs
		files[key] = file
	}

	// Convert map to ordered list
	for _, key := range fileKeys {
		list = append(list, files[key])
	}

	return list, total
}

func (k Keeper) GetAllFileByMerkle() (list []types.UnifiedFile) {

	// Query with pagination
	rows, err := k.filebase.Query(`
        SELECT merkle, owner, start, expires, file_size, proof_interval, proof_type, max_proofs, note
        FROM unified_files
        ORDER BY merkle, owner, start
    `)
	if err != nil {
		return nil
	}
	defer rows.Close()

	files := make(map[string]types.UnifiedFile)
	var fileKeys []string
	for rows.Next() {
		var file types.UnifiedFile
		err := rows.Scan(
			&file.Merkle, &file.Owner, &file.Start, &file.Expires, &file.FileSize,
			&file.ProofInterval, &file.ProofType, &file.MaxProofs, &file.Note,
		)
		if err != nil {
			continue
		}

		key := fmt.Sprintf("%s:%s:%d", base64.StdEncoding.EncodeToString(file.Merkle), file.Owner, file.Start)
		files[key] = file
		fileKeys = append(fileKeys, key)
	}

	if len(fileKeys) == 0 {
		return nil
	}

	// Get proofs for these files
	for _, key := range fileKeys {
		file := files[key]
		proofRows, err := k.filebase.Query(`
            SELECT proof FROM proofs
            WHERE file_merkle = ? AND file_owner = ? AND file_start = ?
        `, file.Merkle, file.Owner, file.Start)

		if err != nil {
			continue
		}

		var proofs []string
		for proofRows.Next() {
			var proof string
			if err := proofRows.Scan(&proof); err != nil {
				continue
			}
			proofs = append(proofs, proof)
		}
		proofRows.Close()

		file.Proofs = proofs
		files[key] = file
	}

	// Convert map to ordered list
	for _, key := range fileKeys {
		list = append(list, files[key])
	}

	return list
}

func (k Keeper) GetOpenFiles(offset, limit int) (list []types.UnifiedFile, total int64) {
	if limit == 0 {
		limit = 100
	}

	// Count total eligible files
	countQuery := `
        SELECT COUNT(*) FROM (
            SELECT f.merkle, f.owner, f.start, f.max_proofs, COUNT(p.proof) as proof_count
            FROM unified_files f
            LEFT JOIN proofs p ON f.merkle = p.file_merkle AND f.owner = p.file_owner AND f.start = p.file_start
            GROUP BY f.merkle, f.owner, f.start
            HAVING COUNT(p.proof) < f.max_proofs
        )
    `
	err := k.filebase.QueryRow(countQuery).Scan(&total)
	if err != nil {
		return nil, 0
	}

	// Get files with fewer proofs than max_proofs
	mainQuery := `
        SELECT f.merkle, f.owner, f.start, f.expires, f.file_size, 
               f.proof_interval, f.proof_type, f.max_proofs, f.note
        FROM unified_files f
        LEFT JOIN (
            SELECT file_merkle, file_owner, file_start, COUNT(proof) as proof_count
            FROM proofs
            GROUP BY file_merkle, file_owner, file_start
        ) pc ON f.merkle = pc.file_merkle AND f.owner = pc.file_owner AND f.start = pc.file_start
        WHERE COALESCE(pc.proof_count, 0) < f.max_proofs
        ORDER BY f.merkle, f.owner, f.start
        LIMIT ? OFFSET ?
    `

	rows, err := k.filebase.Query(mainQuery, limit, offset)
	if err != nil {
		return nil, total
	}
	defer rows.Close()

	files := make(map[string]types.UnifiedFile)
	var fileKeys []string

	for rows.Next() {
		var file types.UnifiedFile
		err := rows.Scan(
			&file.Merkle, &file.Owner, &file.Start, &file.Expires, &file.FileSize,
			&file.ProofInterval, &file.ProofType, &file.MaxProofs, &file.Note,
		)
		if err != nil {
			continue
		}

		key := fmt.Sprintf("%s:%s:%d", base64.StdEncoding.EncodeToString(file.Merkle), file.Owner, file.Start)
		files[key] = file
		fileKeys = append(fileKeys, key)
	}

	if len(fileKeys) == 0 {
		return nil, total
	}

	// Get proofs for the filtered files
	for _, key := range fileKeys {
		file := files[key]
		proofRows, err := k.filebase.Query(`
            SELECT proof FROM proofs
            WHERE file_merkle = ? AND file_owner = ? AND file_start = ?
        `, file.Merkle, file.Owner, file.Start)

		if err != nil {
			continue
		}

		var proofs []string
		for proofRows.Next() {
			var proof string
			if err := proofRows.Scan(&proof); err != nil {
				continue
			}
			proofs = append(proofs, proof)
		}
		proofRows.Close()

		file.Proofs = proofs
		files[key] = file
	}

	// Maintain order
	for _, key := range fileKeys {
		list = append(list, files[key])
	}

	return list, total
}

func (k Keeper) GetTotalFileSize() (totalSize int64, err error) {
	// Query sum of file_size from all files
	err = k.filebase.QueryRow(`
        SELECT COALESCE(SUM(file_size), 0) 
        FROM unified_files
    `).Scan(&totalSize)

	return totalSize, err
}

// IterateFilesByMerkle iterates through every file
func (k Keeper) IterateFilesByMerkle(ctx sdk.Context, reverse bool, fn func(key []byte, val []byte) bool) {
	// Create query with appropriate ordering
	query := `
		SELECT merkle, owner, start, expires, file_size, proof_interval, proof_type, max_proofs, note
		FROM unified_files
		ORDER BY merkle, owner, start
	`
	if reverse {
		query = `
			SELECT merkle, owner, start, expires, file_size, proof_interval, proof_type, max_proofs, note
			FROM unified_files
			ORDER BY merkle DESC, owner DESC, start DESC
		`
	}

	rows, err := k.filebase.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	// Process each file
	for rows.Next() {
		var file types.UnifiedFile
		err := rows.Scan(
			&file.Merkle, &file.Owner, &file.Start, &file.Expires, &file.FileSize,
			&file.ProofInterval, &file.ProofType, &file.MaxProofs, &file.Note,
		)
		if err != nil {
			continue
		}

		// Get proofs for this file
		proofRows, err := k.filebase.Query(`
			SELECT proof FROM proofs
			WHERE file_merkle = ? AND file_owner = ? AND file_start = ?
		`, file.Merkle, file.Owner, file.Start)
		if err != nil {
			continue
		}

		var proofs []string
		for proofRows.Next() {
			var proof string
			if err := proofRows.Scan(&proof); err != nil {
				continue
			}
			proofs = append(proofs, proof)
		}
		proofRows.Close()
		file.Proofs = proofs

		// Marshal file
		b := k.cdc.MustMarshal(&file)

		// Create key (similar to the original KeyPrefix format)
		key := types.FilesPrimaryKey(file.Merkle, file.Owner, file.Start)

		// Call the callback function
		shouldStop := fn(key, b)
		if shouldStop {
			return
		}
	}
}

// IterateAndParseFilesByMerkle iterates through every file and parses them for you
func (k Keeper) IterateAndParseFilesByMerkle(ctx sdk.Context, reverse bool, fn func(key []byte, val types.UnifiedFile) bool) {
	// Create query with appropriate ordering
	query := `
		SELECT merkle, owner, start, expires, file_size, proof_interval, proof_type, max_proofs, note
		FROM unified_files
		ORDER BY merkle, owner, start
	`
	if reverse {
		query = `
			SELECT merkle, owner, start, expires, file_size, proof_interval, proof_type, max_proofs, note
			FROM unified_files
			ORDER BY merkle DESC, owner DESC, start DESC
		`
	}

	rows, err := k.filebase.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()

	// Process each file
	for rows.Next() {
		var file types.UnifiedFile
		err := rows.Scan(
			&file.Merkle, &file.Owner, &file.Start, &file.Expires, &file.FileSize,
			&file.ProofInterval, &file.ProofType, &file.MaxProofs, &file.Note,
		)
		if err != nil {
			continue
		}

		// Get proofs for this file
		proofRows, err := k.filebase.Query(`
			SELECT proof FROM proofs
			WHERE file_merkle = ? AND file_owner = ? AND file_start = ?
		`, file.Merkle, file.Owner, file.Start)
		if err != nil {
			continue
		}

		var proofs []string
		for proofRows.Next() {
			var proof string
			if err := proofRows.Scan(&proof); err != nil {
				continue
			}
			proofs = append(proofs, proof)
		}
		proofRows.Close()
		file.Proofs = proofs

		// Create key (similar to the original KeyPrefix format)
		key := types.FilesPrimaryKey(file.Merkle, file.Owner, file.Start)

		// Call the callback function
		shouldStop := fn(key, file)
		if shouldStop {
			return
		}
	}
}

// GetAllFilesWithMerkle returns all Files that start with a specific merkle
func (k Keeper) GetAllFilesWithMerkle(ctx sdk.Context, merkle []byte) (list []types.UnifiedFile) {
	rows, err := k.filebase.Query(`
		SELECT merkle, owner, start, expires, file_size, proof_interval, proof_type, max_proofs, note
		FROM unified_files
		WHERE merkle = ?
	`, merkle)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var file types.UnifiedFile
		err := rows.Scan(
			&file.Merkle, &file.Owner, &file.Start, &file.Expires, &file.FileSize,
			&file.ProofInterval, &file.ProofType, &file.MaxProofs, &file.Note,
		)
		if err != nil {
			continue
		}

		// Get proofs for this file
		proofRows, err := k.filebase.Query(`
			SELECT proof FROM proofs
			WHERE file_merkle = ? AND file_owner = ? AND file_start = ?
		`, file.Merkle, file.Owner, file.Start)
		if err != nil {
			continue
		}

		var proofs []string
		for proofRows.Next() {
			var proof string
			if err := proofRows.Scan(&proof); err != nil {
				continue
			}
			proofs = append(proofs, proof)
		}
		proofRows.Close()
		file.Proofs = proofs

		list = append(list, file)
	}

	return list
}
