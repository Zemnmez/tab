package storage

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
