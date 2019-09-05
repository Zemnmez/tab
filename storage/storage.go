package storage

type Storage interface {
	Txn() Txn
	Commit(Txn) error
}

type Txn interface {
	Set(value interface{}) error
	Get(value interface{}) error
}



type ID interface {
	String() string
}

type Gettable interface{}

type GetPartial struct {
	IDs []ID
}

func Get(id ...ID) GetPartial {
	return GetPartial{IDs: id}
}

func (g GetPartial) Into(values ...Gettable) (err error) { panic("unimplemented") }
