package wireprotocol

import "encoding/binary"

// ResponseHeaderV0 sends a response header
// first 4 bytes: response size
// last 4 bytes: correlation id
func ResponseHeaderV0(correlationID uint32) []byte {
	len := 8
	b := make([]byte, len)
	binary.BigEndian.PutUint32(b[:4], uint32(len))
	binary.BigEndian.PutUint32(b[4:], correlationID)
	return b
}
