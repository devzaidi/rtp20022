package rtp

import (
	"fmt"
	"time"
)

// BusinessID generates a Business Message Identifier value.
//
// Format: BYYYYMMDDbbbbbbbbbbbGAAAnnnnnnnnnnn
// Pos. 01-01 – Prefix ‘B’
// Pos. 02-09 - File creation date in format YYYYMMDD
// Pos. 10-20 - Participant ID (11 characters)
// Pos. 21-21 - Message generation source (“B” if generated by a TCH FI, "H" generated by RTP)
// Pos. 22-24 - Discretionary bank field (3 digit alpha numeric that can be used to serialize, identify specific systems, etc.)
// Pos. 25-35 - Message serial number (11 numeric characters)
//
// Example: B20171112021200201A1BHEA00000000011
func BusinessID(ts time.Time, participantID, bankField string) string {
	timestamp := ts.Format("20060102")
	partID := fmt.Sprintf("%011.11s", participantID)
	bank := fmt.Sprintf("%3.3s", bankField)
	serial := NumericSerialNumber(11)

	return fmt.Sprintf("B%s%sB%s%11.11s", timestamp, partID, bank, serial)
}
