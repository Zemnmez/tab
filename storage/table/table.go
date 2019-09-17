// Package table implements a table-like abstraction on a Storage KV store.
//
// A table, strictly speaking can be thought of as a kind of composite key.
// Instead of the key being strictly composed of entropy, the key is composed
// of some prefix plus entropy, where the prefix identifies the table itself.
//
// table.Storage can be used to wrap any storage.IStorage. When this happens,
// all keys are asserted to be `table.Key`s.
package table

type Key interface {
	// a key should be able to unmarshal itself
	// from bytes
	io.ReadWriter
	
	// and be able to uniquely identify this table
	Table() io.Writer
}

type Storage struct {
	storage.IStorage
	Tables map[string]bool
}

type ErrInvalidKey struct {
	Got reflect.Type
}

func (e ErrInvalidKey) Error() string { return fmt.Sprintf("using table but key %s is not table.Key", e.Got)}

type ErrInvalidTable struct {
	Name string
}

func (e ErrInvalidTable) Error() string { return fmt.Sprintf("using table but attempting to work with table %s, which has not been defined", e.Name) }

type Txn struct { storage.ITxn; Storage Storage }

func (s Storage) Txn() (ITxn, error) { return Txn { ITxn: s.IStorage.ITxn(); storage: s } }

func (t Txn) Open(key io.Reader, flag int) (rec IRecord, err error) {
	newKey, err := makeKey(key)
	if err != nil { return }
	return t.ITxn.Open(&newKey, flag)
}

func (t Txn) Remove(key io.Reader) (err error) {
	newKey, err := makeKey(key)
	if err != nil { return }
	return t.ITxn.Remove(&newKey, flag)
}

func makeKey(from io.Reader) (key bytes.Buffer, err error) {
	var ok bool
	var k Key
	if k, ok = from.(Key); !ok {
		err = ErrInvalidKey { reflect.TypeOf(key) }
		return
	}

	tableN, err := io.Copy(&key, k.Table())
	if err != nil { return }

	tableBytes := key.Bytes()[:tableN]

	if _, ok := t.Storage.Tables[string(tableBytes)]; !ok {
		err = ErrInvalidTable { string(tableBytes) }
		return
	}
	_, err := io.Copy(&key, k.ID())
	if err != nil { return }

	return
}

func (t Txn) Remove (key io.Reader) (err error) {

}