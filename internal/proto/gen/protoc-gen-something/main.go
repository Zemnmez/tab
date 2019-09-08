package main

import (
	"github.com/gogo/protobuf/vanity"
	"github.com/gogo/protobuf/vanity/command"
	"github.com/zemnmez/tab/internal/proto/gen/protoc-gen-gqlgen/plugins/useful"
)

func main() {
	req := command.Read()
	files := req.GetProtoFile()
	files = vanity.FilterFiles(files, vanity.NotGoogleProtobufDescriptorProto)

	vanity.ForEachFile(files, vanity.TurnOnMarshalerAll)
	vanity.ForEachFile(files, vanity.TurnOnSizerAll)
	vanity.ForEachFile(files, vanity.TurnOnUnmarshalerAll)

	vanity.ForEachFieldInFilesExcludingExtensions(vanity.OnlyProto2(files), vanity.TurnOffNullableForNativeTypesWithoutDefaultsOnly)
	vanity.ForEachFile(files, vanity.TurnOffGoUnrecognizedAll)
	vanity.ForEachFile(files, vanity.TurnOffGoUnkeyedAll)
	vanity.ForEachFile(files, vanity.TurnOffGoSizecacheAll)

	vanity.ForEachFile(files, vanity.TurnOffGoEnumPrefixAll)
	vanity.ForEachFile(files, vanity.TurnOffGoEnumStringerAll)
	vanity.ForEachFile(files, vanity.TurnOnEnumStringerAll)

	vanity.ForEachFile(files, vanity.TurnOnEqualAll)
	vanity.ForEachFile(files, vanity.TurnOnGoStringAll)
	vanity.ForEachFile(files, vanity.TurnOffGoStringerAll)
	vanity.ForEachFile(files, vanity.TurnOnStringerAll)

	command.Write(command.Generate(req))
	command.Write(command.GeneratePlugin(req, useful.New(), "_useful.pb.go"))
}