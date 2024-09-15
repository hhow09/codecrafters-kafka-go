package wireprotocol

import (
	"encoding/binary"
	"io"
)

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

type HeaderV2 struct {
	RequestApiKey uint16
	// The version of the API for the request
	RequestApiVersion uint16
	// A unique identifier for the request
	CorrelationID uint32
}

func ReadRequestHeaderV2(r io.Reader) (HeaderV2, error) {
	buf := make([]byte, 1024)
	_, err := r.Read(buf)
	if err != nil {
		return HeaderV2{}, err
	}
	msgLen := binary.BigEndian.Uint32(buf[:4])
	if msgLen > 1024 {
		return HeaderV2{}, io.ErrShortBuffer
	}
	requestApiKey := binary.BigEndian.Uint16(buf[4:6])
	requestApiVersion := binary.BigEndian.Uint16(buf[6:8])
	correlationID := binary.BigEndian.Uint32(buf[8:12])
	// TODO other fields
	return HeaderV2{
		RequestApiKey:     requestApiKey,
		RequestApiVersion: requestApiVersion,
		CorrelationID:     correlationID,
	}, nil
}
