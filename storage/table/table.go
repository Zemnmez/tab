// Package table implements a table-like abstraction on a Storage KV store.
package table

type Storage struct { storage.Storage }

func (s Storage) Get(key Key, val io.ReaderFrom, flags ...int) { panic("todo ") }
func (s Storage) Put(key Key, val io.WriterTo, flags ...int) { panic("todo") }
func (s Storage) Post(key Key, val io.WriterTo, flags ...int) { panic("todo") }

type Key interface {
	GetTable() io.WriterTo
	GetID() io.WriterTo
	io.WriterTo
}