package storage

import (
	"bytes"
	"crypto/rand"
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

// ValidateFlags is a helper function that
// checks that a combination of flags is valid
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

type IIterator interface {
	Record() IRecord
	Scan() bool
	Err() error
	io.Closer
}

type ITxn interface {
	Open(key io.Reader, flag int) (IRecord, error)
	Remove(key io.Reader) (err error)

	Flush() (err error)
	io.Closer
}

type IStorage interface {
	Txn() (ITxn, error)
	io.Closer
	Flush() (err error)
}

// A Record represents an Open()'ed database record.
// the interface represents the *minimum* implemented by
// a Record. It's also suggested that other io methods are
// implemented as appropriate
type IRecord interface {
	io.ReadWriteCloser
	Key() (io.Reader, error)
	Flush() (err error)
}

type Txn struct {
	Parent Storage
	ITxn
}

type tables [][]byte

func (t tables) Contains(tableName []byte) bool {
	for _, realTableName := range ([][]byte)(t) {
		if bytes.Equal(tableName, realTableName) {
			return true
		}
	}

	return false
}

type Storage struct {
	IStorage
}

// Returns a new transaction (Txn)
func (s Storage) Txn() (New Txn, err error) {
	New.ITxn, err = s.IStorage.Txn()
	New.Parent = s
	return
}

func or(ints ...int) (together int) {
	for _, i := range ints {
		together |= i
	}
	return
}

// Get retrieves the value by the given key
func (t Txn) Get(key io.Reader, val io.Writer, flags ...int) (err error) {
	var r IRecord
	if r, err = t.ITxn.Open(key, or(flags...)|O_RDONLY); err != nil {
		return
	}

	defer r.Close()

	if _, err = io.Copy(val, r); err != nil {
		return
	}

	if err = r.Flush(); err != nil {
		return
	}

	return
}

// Put stores the value with the given key
func (t Txn) Put(key, val io.Reader, flags ...int) (err error) {

	var r IRecord
	if r, err = t.ITxn.Open(key, or(flags...)|O_WRONLY); err != nil {
		return
	}

	defer r.Close()

	if _, err = io.Copy(r, val); err != nil {
		return
	}

	if err = r.Flush(); err != nil {
		return
	}

	return
}

// Post stores the value, generating a key as it is added.
// The key is passed rand.Reader to read entropy from for a uuid.
func (t Txn) Post(key io.ReadWriter, val io.Reader, flags ...int) (err error) {
	if _, err = io.Copy(key, rand.Reader); err != nil {
		return
	}

	var r IRecord
	if r, err = t.ITxn.Open(key, or(flags...)|O_WRONLY); err != nil {
		return
	}

	defer r.Close()

	if _, err = io.Copy(r, val); err != nil {
		return
	}

	if err = r.Flush(); err != nil {
		return
	}

	return
}

// Delete deletes the value at the given key
func (t Txn) Delete(key io.Reader) (err error) {
	return t.ITxn.Remove(key)
}
