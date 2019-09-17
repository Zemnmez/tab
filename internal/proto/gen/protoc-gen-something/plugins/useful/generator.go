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
	graphqlPkg := p.NewImport("github.com/99designs/gqlgen/graphql")
	encodingPkg := p.NewImport("encoding")
	fmtPkg := p.NewImport("fmt")
	ioutilPkg := p.NewImport("io/ioutil")
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

		p.P(`// MarshalJSON implements json.Marshaler`)
		p.P(`func (this `, ccTypeName, `) MarshalJSON() (json []byte, err error) {`)
		p.In()

		switch message.(type) {
		case *generator.Descriptor:
			p.P(`var b `, bytesPkg.Use(), ".Buffer")
			p.P(`if err=_jsonMarshaler.Marshal(&b,&this);err!=nil{return}`)
			p.P(`json = b.Bytes()`)
			p.P(`return`)

		case *generator.EnumDescriptor:
			p.P(`if _, ok := `, ccTypeName, `_name[int32(this)];!ok{return nil, ErrInvalidEnum{ EnumName: "`, ccTypeName, `", Value: this }}`)
			p.P(`return `, jsonPkg.Use(), `.Marshal(this.String())`)

		default:
			panic("i am confused!")
		}

		p.Out()
		p.P(`}`)

		p.P(`// UnmarshalJSON implements json.Unmarshaler`)
		p.P(`func (this *`, ccTypeName, `) UnmarshalJSON(json []byte) (err error) {`)
		p.In()
		switch message.(type) {
		case *generator.Descriptor:
			p.P(`return _jsonUnmarshaler.Unmarshal(`, bytesPkg.Use(), `.NewReader(json), this)`)

		case *generator.EnumDescriptor:
			p.P(`var text string`)
			p.P(`if err = `, jsonPkg.Use(), `.Unmarshal(json, &text); err != nil {return}`)

			p.P(`var ok bool`)
			p.P(`this32 := (*int32)(this)`)
			p.P(`if *this32, ok = `, ccTypeName, `_value[text];!ok{return ErrInvalidEnum{ EnumName: "`, ccTypeName, `", Value: this }}`)
			p.P(`return`)

		default:
			panic("this shouldnt happen!")
		}

		p.Out()
		p.P(`}`)

		p.P(`// MarshalGQL implements graphql.Marshaler`)
		p.P(`func (this `, ccTypeName, `) MarshalGQL(w `, ioPkg.Use(), `.Writer) {var err error;`)

		switch message.(type) {

		case *generator.Descriptor:
			p.P(`if err=_jsonMarshaler.Marshal(w,&this);err!=nil{panic(err)}`)

		case *generator.EnumDescriptor:
			p.P("b, err := this.MarshalJSON()")
			p.P("if err != nil { panic(err)}")
			p.P("if _, err = w.Write(b); err != nil { panic(err) }")

		default:
			panic("huh?")
		}
		p.P(`}`)

		// this is going to be really dumb... to avoid serious type fuckery
		// we'll need to encode back into JSON and then re parse it

		p.P(`// UnmarshalGQL implements graphql.Unmarshaler`)
		p.P(`func (this *`, ccTypeName, `) UnmarshalGQL(v interface{}) (err error) {`+
			`var newJSON []byte;if newJSON, err = `, jsonPkg.Use(), `.Marshal(v);err!=nil{return};`+
			`;return this.UnmarshalJSON(newJSON);`+
			`}`)

		switch message.(type) {
		case *generator.Descriptor:

			p.P(`// WriteTo implements io.WriterTo.`)
			p.P(`// WriteTo writes this structure as protobufs`)
			p.P(`func (this `, ccTypeName, `) WriteTo(w `, ioPkg.Use(), `.Writer) (n int64, err error) {`)
			p.P(`bt, err := this.MarshalBinary()`)
			p.P(`if err != nil { return }`)
			p.P(`nint, err := w.Write(bt)`)
			p.P(`n = int64(nint);return`)
			p.P(`}`)

			p.P("// Read implements io.Reader. It only exists")
			p.P("// as a dummy for io.Copy. It will error if called.")
			p.P("// use .WriteTo instead.")
			p.P(`func(`, ccTypeName, `)Read(b []byte)(n int, err error){`)
			p.In()
			p.P("return 0, ErrUnreadable {`", ccTypeName, "`}")
			p.Out()
			p.P("}")

			p.P(`// ReadFrom implements io.ReaderFrom.`)
			p.P(`// ReadFrom expects the structure as protobufs,`)
			p.P(`// and assumes the protobuf message should consume`)
			p.P(`// the entire reader.`)
			p.P(`func (this `, ccTypeName, `) ReadFrom(r `, ioPkg.Use(), `.Reader) (n int64, err error) {`)
			p.P(`bt, err := `, ioutilPkg.Use(), `.ReadAll(r)`)
			p.P(`n = int64(len(bt))`)
			p.P(`if err != nil { return }`)
			p.P(`err = this.UnmarshalBinary(bt)`)
			p.P(`return`)
			p.P(`}`)

			p.P(`// MarshalBinary implements encoding.BinaryMarshaler`)
			p.P(`func (this *`, ccTypeName, `) MarshalBinary() ([]byte, error) {`)
			p.In()
			p.P(`return this.Marshal()`)
			p.Out()
			p.P(`}`)

			p.P(`// UnmarshalBinary implements encoding.BinaryUnmarshaler`)
			p.P(`func (this *`, ccTypeName, `) UnmarshalBinary(b []byte) error {`)
			p.In()
			p.P(`return this.Unmarshal(b)`)
			p.Out()
			p.P(`}`)
		}

		p.P(`var _ interface {`)
		p.In()
		switch message.(type) {
		case *generator.Descriptor:
			p.P(encodingPkg.Use(), `.BinaryMarshaler`)
			p.P(encodingPkg.Use(), `.BinaryUnmarshaler`)
			p.P(ioPkg.Use(), `.WriterTo`)
			p.P(ioPkg.Use(), `.ReaderFrom`)
		}
		p.P(jsonPkg.Use(), `.Marshaler`)
		p.P(jsonPkg.Use(), `.Unmarshaler`)
		p.P(graphqlPkg.Use(), `.Unmarshaler`)
		p.P(graphqlPkg.Use(), `.Marshaler`)
		p.Out()
		p.P(`} = new(`, ccTypeName, `)`)

	}

	p.P(`//these can be set via init() to customise the (un)marshaling`)
	p.P(`var (`)
	p.In()
	p.P(`_jsonMarshaler `, jsonpbPkg.Use(), `.Marshaler`)
	p.P(`_jsonUnmarshaler `, jsonpbPkg.Use(), `.Unmarshaler`)
	p.Out()
	p.P(`)`)

	p.P("type ErrUnreadable struct { Name string }")
	p.P("func(e ErrUnreadable) Error() string {")
	p.P(`return `, fmtPkg.Use(), `.Sprintf("%s can only be read via .WriteTo", e.Name)}`)
	p.P(`type ErrInvalidEnum struct { EnumName string; Value interface{} }`)
	p.P(`func (i ErrInvalidEnum) Error() string { return fmt.Sprintf("invalid %s(%s)", i.EnumName, i.Value) }`)
}
