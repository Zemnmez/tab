package types

import (
	"io"
	"fmt"
)

func (i ID) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf(`"%s"`, i.Id)))
}

func (i *ID) UnmarshalGQL(v interface{}) {
	i.Id = string(v.(string))
}