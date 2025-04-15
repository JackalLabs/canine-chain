package keeper

import "database/sql"

// CreateTablesIfNotExist initializes the file database schema
func CreateTablesIfNotExist(db *sql.DB) error {
	// Create unified_files table if it doesn't exist
	_, err := db.Exec(`
	CREATE TABLE IF NOT EXISTS unified_files (
		merkle BLOB NOT NULL,
		owner TEXT NOT NULL,
		start INTEGER NOT NULL,
		expires INTEGER NOT NULL,
		file_size INTEGER NOT NULL,
		proof_interval INTEGER NOT NULL,
		proof_type INTEGER NOT NULL,
		max_proofs INTEGER NOT NULL,
		note TEXT,
		PRIMARY KEY (merkle, owner, start)
	)`)
	if err != nil {
		return err
	}

	// Create proofs table if it doesn't exist
	_, err = db.Exec(`
	CREATE TABLE IF NOT EXISTS proofs (
		file_merkle BLOB NOT NULL,
		file_owner TEXT NOT NULL,
		file_start INTEGER NOT NULL,
		proof TEXT NOT NULL,
		PRIMARY KEY (file_merkle, file_owner, file_start, proof),
		FOREIGN KEY (file_merkle, file_owner, file_start) 
			REFERENCES unified_files (merkle, owner, start) ON DELETE CASCADE
	)`)
	if err != nil {
		return err
	}

	// Create index for expires column if it doesn't exist
	_, err = db.Exec(`
	CREATE INDEX IF NOT EXISTS idx_unified_files_expires 
	ON unified_files (expires)
	`)

	return err
}
