package wireprotocol

import (
	"encoding/binary"
	"fmt"
)

// https://kafka.apache.org/protocol.html#protocol_error_codes
const (
	ERR_UNSUPPORTED_VERSION = 35
)

type APIError struct {
	errorCode int
}

func (e APIError) Error() string {
	return fmt.Sprintf("APIError: %v", e.Code())
}

func NewAPIError(errorCode int) APIError {
	return APIError{errorCode: errorCode}
}

// error code is INT16
func (e APIError) Code() []byte {
	b := make([]byte, 2)
	binary.BigEndian.PutUint16(b, uint16(e.errorCode))
	return b
}

var (
	ErrReadingConn = fmt.Errorf("Error reading from connection")
)
