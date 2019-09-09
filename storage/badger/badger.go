package badger

import (
	"bytes"
	"errors"
	"io"

	"github.com/dgraph-io/badger"
	"github.com/zemnmez/tab/storage"
)

const (
	O_RDONLY int = storage.O_RDONLY
	O_WRONLY int = storage.O_WRONLY
	O_RDWR   int = storage.O_RDWR
	O_APPEND int = storage.O_APPEND
	O_CREATE int = storage.O_CREATE
	O_EXCL   int = storage.O_EXCL
	O_TRUNC  int = storage.O_TRUNC
)

var _ storage.Storage = Storage{}
var _ storage.Tx = Tx{}
var _ storage.Record = &Record{}

type Storage struct{ *badger.DB }
type Tx struct{ *badger.Txn }

func NewStorage(db *badger.DB) (storage storage.Storage, err error) { return Storage{db}, nil }
func (s Storage) Tx() (txn storage.Tx, err error)                   { return Tx{s.NewTransaction(true)}, nil }
func (t Tx) Close() (err error)                                     { t.Txn.Discard(); return nil }
func (t Tx) Flush() (err error)                                     { return t.Txn.Commit() }
func (t Tx) Remove(key []byte) (err error)                          { return t.Txn.Delete(key) }
func (s Storage) Flush() (err error)                                { return s.DB.Sync() }
func (s Storage) Close() (err error)                                { return s.Flush() }

var ErrO_EXCL = errors.New("key already exists and O_EXCL was requested")

func (t Tx) Open(name []byte, flag int) (fresh storage.Record, err error) {
	var r = Record{tx: t, key: name, flag: flag}
	fresh = &r
	if err = r.Validate(); err != nil {
		return
	}

	err = r.LoadEntry()
	if O_CREATE&flag != 0 {
		if O_EXCL&flag != 0 && err != badger.ErrKeyNotFound {
			return fresh, ErrO_EXCL
		}

		if err == badger.ErrKeyNotFound {
			err = nil
		}

	}

	// if we are not creating and it doesn't exist,
	// we need to surface that error
	if err != nil {
		return
	}

	if flag&O_APPEND != 0 || flag&O_RDWR != 0 || flag&O_RDONLY != 0 {
		if err = r.LoadEntry(); err != nil {
			return
		}
	}

	if flag&O_TRUNC != 0 {
		r.Buffer.Truncate(0)
	}

	return
}

type Record struct {
	tx   Tx
	key  []byte
	flag int
	item *badger.Item
	bytes.Buffer
}

func (r *Record) Item() (item *badger.Item, err error) {
	r.item, err = r.tx.Get(r.key)
	return r.item, err
}

func (r Record) Key() (rd io.Reader, err error) { return bytes.NewReader(r.key), nil }

func (r *Record) LoadEntry() (err error) {
	// if the buffer is already loaded, skip
	if r.Buffer.Len() > 0 {
		return nil
	}

	var item *badger.Item
	item, err = r.Item()
	if err != nil {
		return
	}

	// this overwrites what was in the buffer, reslicing it
	// if necessary
	// this means we don't reallocate the backbuffer if we don't
	// need to
	bt, err := item.ValueCopy(r.Buffer.Bytes()[0:])
	if err != nil {
		return
	}

	// we then set the new (potentially resized) buffer
	// as the backbuffer of a new Buffer.
	r.Buffer = *bytes.NewBuffer(bt)

	return
}

func (r Record) Validate() (err error) {
	if err = storage.ValidateFlags(r.flag); err != nil {
		return
	}

	return
}

func (r Record) Close() error { return r.Flush() }
func (r Record) Flush() (err error) {
	if r.flag|O_RDWR != 0 || r.flag|O_WRONLY != 0 {
		if err = r.tx.Set(r.key, r.Buffer.Bytes()); err != nil {
			return
		}
	}

	return
}
