package types

import "strconv"

func (d *ActiveDeals) IsVerified(height int64, proofWindow int64) bool {
	start, err := strconv.ParseInt(d.Startblock, 10, 64)
	if err != nil {
		return false
	}

	lifetime := height - start

	windowStart := (height - lifetime%proofWindow) - proofWindow
	if windowStart < 0 {
		windowStart = 0
	}

	if d.LastProof >= windowStart {
		return true
	}

	return false
}
