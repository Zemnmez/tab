package storage

import (
	"bytes"
	"crypto/rand"
	"errors"
	"fmt"
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

// Table is a special type of Txn that appends a table prefix to
// keys. It also provides iteration support
type Table struct {
	prefix []byte
	ITxn
}

func (t Table) KeyPrefix() (prefix []byte)            { return t.prefix }
func (t Table) DerivedKey(key []byte) (newKey []byte) { return append(t.KeyPrefix(), key...) }
func (t Table) Open(key []byte, flag int) (IRecord, error) {
	return t.ITxn.Open(t.DerivedKey(key), flag)
}
func (t Table) Remove(key []byte) (err error) { return t.ITxn.Remove(t.DerivedKey(key)) }

type ITxn interface {
	Open(key []byte, flag int) (IRecord, error)
	Remove(key []byte) (err error)

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

// NewTables returns a set of tables. It panics if the table
// names are not unique.
func NewTables(names ...io.WriterTo) (valid tables) {
	valid = make(tables, len(names))
	var uniques = make(map[string]bool, len(names))

	for i, name := range names {
		var buf bytes.Buffer
		if _, err := io.Copy(&buf, name); err != nil {
			return
		}

		bt := buf.Bytes()

		if len(bt) == 0 {
			panic("empty table name")
		}

		if _, ok := uniques[string(bt)]; ok {
			panic(fmt.Sprintf("table %s already exists", bt))
		}

		([][]byte)(valid)[i] = bt
	}

	return
}

type Storage struct {
	IStorage
	Tables tables
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

type CopiableWriterTo struct {
	io.WriterTo
}

func (c CopiableWriterTo) Read(b []byte) (n int, err error) { panic("only WriteTo is allowed") }

type CopiableReaderFrom struct {
	io.ReaderFrom
}

func (c CopiableReaderFrom) Write(b []byte) (n int, err error) {
	panic("only ReadFrom is allowed")
}

// Get retrieves the value by the given key
func (t Txn) Get(key io.WriterTo, val io.ReaderFrom, flags ...int) (err error) {
	var buf bytes.Buffer
	if _, err = key.WriteTo(&buf); err != nil {
		return
	}

	var r IRecord
	if r, err = t.ITxn.Open(buf.Bytes(), or(flags...)|O_RDONLY); err != nil {
		return
	}

	defer r.Close()

	if _, err = io.Copy(CopiableReaderFrom{val}, r); err != nil {
		return
	}

	if err = r.Flush(); err != nil {
		return
	}

	return
}

// Put stores the value with the given key
func (t Txn) Put(key, val io.WriterTo, flags ...int) (err error) {
	var buf bytes.Buffer
	if _, err = key.WriteTo(&buf); err != nil {
		return
	}

	var r IRecord
	if r, err = t.ITxn.Open(buf.Bytes(), or(flags...)|O_WRONLY); err != nil {
		return
	}

	defer r.Close()

	if _, err = io.Copy(r, CopiableWriterTo{val}); err != nil {
		return
	}

	if err = r.Flush(); err != nil {
		return
	}

	return
}

// Post stores the value, generating a key as it is added.
// The key is passed rand.Reader to read entropy from for a uuid.
func (t Txn) Post(key interface {
	io.WriterTo
	io.ReaderFrom
}, val io.WriterTo, flags ...int) (err error) {
	if _, err = key.ReadFrom(rand.Reader); err != nil {
		return
	}

	var buf bytes.Buffer
	if _, err = key.WriteTo(&buf); err != nil {
		return
	}

	var r IRecord
	if r, err = t.ITxn.Open(buf.Bytes(), or(flags...)|O_WRONLY); err != nil {
		return
	}

	defer r.Close()

	if _, err = io.Copy(r, CopiableWriterTo{val}); err != nil {
		return
	}

	if err = r.Flush(); err != nil {
		return
	}

	return
}

// Delete deletes the value at the given key
func (t Txn) Delete(key io.WriterTo) (err error) {
	var buf bytes.Buffer
	if _, err = key.WriteTo(&buf); err != nil {
		return
	}

	return t.ITxn.Remove(buf.Bytes())
}

func (t Txn) WithTable(name io.WriterTo) (handle Table, err error) {
	var buf bytes.Buffer
	if _, err = io.Copy(&buf, CopiableWriterTo{name}); err != nil {
		return
	}

	handle.prefix = buf.Bytes()

	if !t.Parent.Tables.Contains(handle.prefix) {
		err = fmt.Errorf("%+q is not a registered table name", handle.prefix)
		return
	}

	handle.ITxn = t
	return
}
