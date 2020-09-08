package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"

	hiveOpts "github.com/simo7/protoc-gen-gluecatalog/hive_options"
	parquetOpts "github.com/simo7/protoc-gen-parquet/parquet_options"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	flags          flag.FlagSet
	noUnsigned     = flags.Bool("no_unsigned", false, "use correspondant integer in place of unsigned integer")
	timestampInt96 = flags.Bool("timestamp_int96", false, "use INT96 for Hive timestamps")
)

var protoToParquet = map[string]string{
	"bool":      "boolean",
	"int32":     "int32",
	"int64":     "int64",
	"uint32":    "uint32",
	"uint64":    "uint64",
	"float":     "float",
	"float64":   "double",
	"double":    "double",
	"string":    "binary",
	"[]byte":    "binary",
	"enum":      "binary",
	"message":   "group",
	"timestamp": "int64",
	"int96":     "int96",
}

func main() {
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		for _, f := range gen.Files {
			if f.Generate {
				generateFile(gen, f)
			}
		}
		return nil
	})
}

func getIndent(indentLevel int) string {
	return strings.Repeat("  ", indentLevel)
}

func generateFile(gen *protogen.Plugin, file *protogen.File) {
	path, _ := filepath.Split(file.GeneratedFilenamePrefix)

	for _, message := range file.Messages {
		opts := message.Desc.Options()

		if !opts.ProtoReflect().IsValid() {
			continue
		}

		optValue := proto.GetExtension(opts, parquetOpts.E_ParquetMessageOpts)
		tableName := optValue.(*parquetOpts.ParquetMessageOptions).GetTableName()

		if tableName == "" {
			continue
		}

		filename := path + tableName + ".schema"
		g := gen.NewGeneratedFile(filename, file.GoImportPath)

		g.P("message " + tableName + " {")
		for _, field := range message.Fields {
			generateField(g, field.Desc, 1)
		}
		g.P("}")
	}
}

func generateField(g *protogen.GeneratedFile, field protoreflect.FieldDescriptor, indentLevel int) {
	fieldName := string(field.Name())
	protoKind := field.Kind().String()

	opts := field.Options()
	if opts.ProtoReflect().IsValid() {
		optValue := proto.GetExtension(opts, hiveOpts.E_HiveFieldOpts)
		protoKind = optValue.(*hiveOpts.HiveFieldOptions).GetTypeOverride()
	}

	if protoKind == "timestamp" && *timestampInt96 {
		protoKind = "int96"
	}

	fieldType := protoToParquet[protoKind]
	if fieldType == "" {
		log.Fatalf("type %v for field %v was not recognized", protoKind, fieldName)
	}

	if strings.HasPrefix(fieldType, "uint") && *noUnsigned == true {
		fieldType = strings.Replace(fieldType, "uint", "int", 1)
	}

	isRepeatedMessage := false
	if field.IsList() {
		isRepeatedMessage = true
	}

	lineEnd := ";"
	if protoKind == "message" {
		lineEnd = " {"
	}

	annotation := ""
	if protoKind == "string" || protoKind == "enum" {
		annotation = "(UTF8)"
	}

	if isRepeatedMessage {
		g.P(fmt.Sprintf("%s optional group %s {", getIndent(indentLevel), fieldName))
	} else {
		g.P(fmt.Sprintf("%s optional %s %s %s %s",
			getIndent(indentLevel), fieldType, fieldName, annotation, lineEnd,
		))
	}

	if isRepeatedMessage {
		indentLevel++
		g.P(fmt.Sprintf("%s repeated group list {", getIndent(indentLevel)))
		indentLevel++

		if protoKind == "message" {
			g.P(fmt.Sprintf("%s optional group element {", getIndent(indentLevel)))
		} else {
			g.P(fmt.Sprintf("%s optional %s %s %s %s",
				getIndent(indentLevel), fieldType, fieldName, annotation, lineEnd,
			))
		}
	}

	if protoKind == "message" {
		fds := field.Message().Fields()
		for i := 0; i < fds.Len(); i++ {
			field := fds.Get(i)
			generateField(g, field, indentLevel+1)
		}
		g.P(fmt.Sprintf("%s}", getIndent(indentLevel)))
	}

	if isRepeatedMessage {
		g.P(fmt.Sprintf("%s}", getIndent(indentLevel-1)))
		g.P(fmt.Sprintf("%s}", getIndent(indentLevel-2)))
	}
}
