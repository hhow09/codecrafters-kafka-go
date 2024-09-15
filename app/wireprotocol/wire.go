package wireprotocol

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
)

var (
	validAPIVerRange = [2]uint16{0, 4}
)

// ResponseHeaderV0 sends a response header
// first 4 bytes: response size
// last 4 bytes: correlation id
// and the body of the response follows
func ResponseV0(correlationID uint32, body []byte) []byte {
	len := 8
	b := make([]byte, len)
	binary.BigEndian.PutUint32(b[:4], uint32(len))
	binary.BigEndian.PutUint32(b[4:], correlationID)
	b = append(b, body...)
	fmt.Printf("correlationID: %v\n", correlationID)
	fmt.Printf("ResponseV0: %v\n", b)
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
		return HeaderV2{}, errors.Join(ErrReadingConn, err)
	}
	msgLen := binary.BigEndian.Uint32(buf[:4])
	if msgLen > 1024 {
		return HeaderV2{}, io.ErrShortBuffer
	}
	h := HeaderV2{}
	h.RequestApiKey = binary.BigEndian.Uint16(buf[4:6])
	h.RequestApiVersion = binary.BigEndian.Uint16(buf[6:8])
	h.CorrelationID = binary.BigEndian.Uint32(buf[8:12])

	// validate header
	if h.RequestApiKey < validAPIVerRange[0] || h.RequestApiKey > validAPIVerRange[1] {
		return h, NewAPIError(ERR_UNSUPPORTED_VERSION)
	}
	// TODO other fields
	return h, nil
}
