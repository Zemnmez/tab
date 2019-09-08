package types

import (
	"io"
	"fmt"
	"encoding/json"
)

func (i ID) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf(`"%s"`, i.Id)))
}

func (i *ID) UnmarshalGQL(v interface{}) {
	i.Id = string(v.(string))
}

func (a Authorization) MarshalJSON() (b []byte, err error) {
	return json.Marshal(a.String())
}

func (a *Authorization) UnmarshalJSON(b []byte) (err error) {
	var ok bool
	a2 := (*int32)(a)
	if *a2, ok = Authorization_value[string(b)]; !ok {
		err = fmt.Errorf("invalid Authorization %+q", b)
	}

	return
}

func (a Authorization) MarshalGQL(w io.Writer) {
	w.Write([]byte(fmt.Sprintf(`"%s"`, a.String())))
}

func (a *Authorization) UnmarshalGQL(i interface{}) error {
	*(*int32)(a) = Authorization_value[i.(string)]

	return nil
}
