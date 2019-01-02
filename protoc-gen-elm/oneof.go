package main

import (
	"fmt"

	"github.com/golang/protobuf/protoc-gen-go/descriptor"
)

func (fg *FileGenerator) GenerateOneofDefinition(inFile *descriptor.FileDescriptorProto, prefix string, inMessage *descriptor.DescriptorProto, oneofIndex int) error {
	inOneof := inMessage.GetOneofDecl()[oneofIndex]

	// TODO: Prefix with message name to avoid collisions.
	msgName := prefix + inMessage.GetName()
	oneofName := oneofType(inOneof)
	oneofType := fmt.Sprintf("%s_%s", msgName, oneofName)

	fg.P("")
	fg.P("")
	fg.P("type %s", oneofType)
	{
		fg.In()

		leading := "="
		{
			oneofVariantName := fmt.Sprintf("%s_%s", elmTypeName(msgName), oneofUnspecifiedValue(inOneof))
			fg.P("%s %s", leading, oneofVariantName)
			leading = "|"
		}
		for _, inField := range inMessage.GetField() {
			if inField.OneofIndex != nil && inField.GetOneofIndex() == int32(oneofIndex) {
				oneofVariantName := fmt.Sprintf("%s_%s_%s", elmTypeName(msgName), elmTypeName(inOneof.GetName()), elmTypeName(inField.GetName()))
				oneofArgumentType := fieldElmType(inFile, inField)
				fg.P("%s %s %s", leading, oneofVariantName, oneofArgumentType)

				leading = "|"
			}
		}
		fg.Out()
	}

	return nil
}

func (fg *FileGenerator) GenerateOneofDecoder(inFile *descriptor.FileDescriptorProto, prefix string, inMessage *descriptor.DescriptorProto, oneofIndex int) error {
	inOneof := inMessage.GetOneofDecl()[oneofIndex]

	// TODO: Prefix with message name to avoid collisions.
	msgName := prefix + inMessage.GetName()
	oneofName := oneofType(inOneof)
	oneofType := fmt.Sprintf("%s_%s", msgName, oneofName)
	decoderName := oneofDecoderName(inOneof)

	fg.P("")
	fg.P("")
	fg.P("%s : JD.Decoder %s", decoderName, oneofType)
	fg.P("%s =", decoderName)
	{
		fg.In()
		fg.P("JD.lazy <| \\_ -> JD.oneOf")
		{
			fg.In()

			leading := "["
			for _, inField := range inMessage.GetField() {
				if inField.OneofIndex != nil && inField.GetOneofIndex() == int32(oneofIndex) {
					oneofVariantName := fmt.Sprintf("%s_%s_%s", elmTypeName(msgName), elmTypeName(inOneof.GetName()), elmTypeName(inField.GetName()))
					decoderName := fieldDecoderName(inFile, inField)
					fg.P("%s JD.map %s (JD.field %q %s)", leading, oneofVariantName, inField.GetJsonName(), decoderName)
					leading = ","
				}
			}
			fg.P("%s JD.succeed %s", leading, fmt.Sprintf("%s_%s", elmTypeName(msgName), oneofUnspecifiedValue(inOneof)))
			fg.P("]")
			fg.Out()
		}
		fg.Out()
	}

	return nil
}

func (fg *FileGenerator) GenerateOneofEncoder(inFile *descriptor.FileDescriptorProto, prefix string, inMessage *descriptor.DescriptorProto, oneofIndex int) error {
	inOneof := inMessage.GetOneofDecl()[oneofIndex]

	// TODO: Prefix with message name to avoid collisions.
	msgName := prefix + inMessage.GetName()

	oneofName := oneofType(inOneof)
	oneofType := fmt.Sprintf("%s_%s", msgName, oneofName)
	encoderName := oneofEncoderName(inOneof)
	argName := "v"

	fg.P("")
	fg.P("")
	fg.P("%s : %s -> Maybe ( String, JE.Value )", encoderName, oneofType)
	fg.P("%s %s =", encoderName, argName)
	{
		fg.In()
		fg.P("case %s of", argName)
		{
			fg.In()

			valueName := "x"
			{
				oneofVariantName := fmt.Sprintf("%s_%s", elmTypeName(msgName), oneofUnspecifiedValue(inOneof))
				fg.P("%s ->", oneofVariantName)
				fg.In()
				fg.P("Nothing")
				fg.Out()
			}
			// TODO: Evaluate them in reverse order, as per
			// https://developers.google.com/protocol-buffers/docs/proto3#oneof
			for _, inField := range inMessage.GetField() {
				if inField.OneofIndex != nil && inField.GetOneofIndex() == int32(oneofIndex) {
					oneofVariantName := fmt.Sprintf("%s_%s_%s", elmTypeName(msgName), elmTypeName(inOneof.GetName()), elmTypeName(inField.GetName()))
					e := fieldEncoderName(inFile, inField)
					fg.P("%s %s ->", oneofVariantName, valueName)
					fg.In()
					fg.P("Just ( %q, %s %s )", inField.GetJsonName(), e, valueName)
					fg.Out()
				}
			}
			fg.Out()
		}
		fg.Out()
	}

	return nil
}

func oneofDecoderName(inOneof *descriptor.OneofDescriptorProto) string {
	typeName := elmTypeName(inOneof.GetName())
	return decoderName(typeName)
}

func oneofEncoderName(inOneof *descriptor.OneofDescriptorProto) string {
	typeName := elmTypeName(inOneof.GetName())
	return encoderName(typeName)
}

func oneofType(inOneof *descriptor.OneofDescriptorProto) string {
	return elmTypeName(inOneof.GetName())
}

func oneofUnspecifiedValue(inOneof *descriptor.OneofDescriptorProto) string {
	return fmt.Sprintf("%s_Unspecified", elmTypeName(inOneof.GetName()))
}
