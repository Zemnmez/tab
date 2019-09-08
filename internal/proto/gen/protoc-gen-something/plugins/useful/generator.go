
package useful

import (
	"github.com/gogo/protobuf/protoc-gen-gogo/generator"
)

type useful struct {
	*generator.Generator
	generator.PluginImports
}

func New() *useful {
	return &useful{}
}

func (p *useful) Name() string {
	return "json"
}

func (p *useful) Init(g *generator.Generator) {
	p.Generator = g
}


func (p *useful) Generate(file *generator.FileDescriptor) {
	p.PluginImports = generator.NewPluginImports(p.Generator)

	bytesPkg := p.NewImport("bytes")
	ioPkg := p.NewImport("io")
	jsonPkg := p.NewImport("encoding/json")

	type NamedType interface {
		TypeName() []string
	}

	var toExtend []NamedType
	for _, f := range file.Messages() {
		toExtend = append(toExtend, f)
	}

	for _, e := range file.Enums() {
		toExtend = append(toExtend, e)
	}

	jsonpbPkg := p.NewImport("github.com/gogo/protobuf/jsonpb")
	for _, message := range toExtend {
		ccTypeName := generator.CamelCaseSlice(message.TypeName())
		p.P(`func (this `, ccTypeName, `) MarshalJSON() (json []byte, err error) {`)
		p.In()
		p.P(`var b `, bytesPkg.Use(), ".Buffer")
		p.P(`if err=_jsonMarshaler.Marshal(&b,&this);err!=nil{return}`)
		p.P(`json = b.Bytes()`)
		p.P(`return`)
		p.Out()
		p.P(`}`)

		p.P(`func (this *`, ccTypeName, `) UnmarshalJSON(json []byte) (err error) {`+
		`return _jsonUnmarshaler.Unmarshal(`, bytesPkg.Use(), `.NewReader(json), this)`+
		`}`)

		p.P(`func (this `, ccTypeName, `) MarshalGQL(w `, ioPkg.Use(), `.Writer) {var err error;`+
		`if err=_jsonMarshaler.Marshal(w,&this);err!=nil{panic(err)}`+
		`}`)

		// this is going to be really dumb... to avoid serious type fuckery
		// we'll need to encode back into JSON and then re parse it
		p.P(`func (this *`, ccTypeName, `) UnmarshalGQL(v interface{}) (err error) {`+
		`var newJSON []byte;if newJSON, err = `, jsonPkg.Use(), `.Marshal(v);err!=nil{return};`+
		`;return this.UnmarshalJSON(newJSON);`+
		`}`)
	}

	p.P(`//these can be set via init() to customise the (un)marshaling`)
	p.P(`var (`)
	p.In()
	p.P(`_jsonMarshaler `, jsonpbPkg.Use(), `.Marshaler`)
	p.P(`_jsonUnmarshaler `, jsonpbPkg.Use(), `.Unmarshaler`)
	p.Out()
	p.P(`)`)

}