package storage

import (
	"encoding"
	"errors"
	"io"
	"os"
)

// to be used with Tx.Open()
const (
	O_RDONLY int = os.O_RDONLY // open read only
	O_WRONLY int = os.O_WRONLY // open write only
	O_RDWR   int = os.O_RDWR   // open read-write

	O_APPEND int = os.O_APPEND // append data
	O_CREATE int = os.O_CREATE // create the Record
	O_EXCL   int = os.O_EXCL   // for use with O_CREATE; Record must not already exist
	O_TRUNC  int = os.O_TRUNC  // truncate (delete contents of) Record
)

var ErrExclusiveOpen = errors.New("only one of O_RDONLY, O_WRONLY or O_RDWR may be specified")

func ValidateFlags(toValidate int) (err error) {
	var ctr uint
	for _, flag := range []int{O_RDONLY, O_WRONLY, O_RDWR} {
		if toValidate&flag != 0 {
			ctr++
		}

		if ctr > 1 {
			return ErrExclusiveOpen
		}
	}

	return
}

type Storable interface {
	encoding.BinaryMarshaler
	encoding.BinaryUnmarshaler
}

type Tx interface {
	Open(key []byte, flag int) (Record, error)
	Remove(key []byte) (err error)

	Flush() (err error)
	io.Closer
}

type Storage interface {
	Tx() (Tx, error)
	io.Closer
	Flush() (err error)
}

// A Record represents an Open()'ed database record.
// the interface represents the *minimum* implemented by
// a Record. It's also suggested that other io methods are
// implemented as appropriate
type Record interface {
	io.ReadWriteCloser
	Key() (io.Reader, error)
}
