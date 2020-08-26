package main

import (
	"flag"
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
	flags      flag.FlagSet
	noUnsigned = flags.Bool("no_unsigned", false, "use correspondant integer in place of unsigned integer")
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
	"timestamp": "int96",
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

func lineFmt(indentLevel int, args ...string) string {
	line := strings.Repeat("  ", indentLevel)
	for i, arg := range args {
		if arg == "" {
			continue
		}
		if i > 0 && i < len(args)-1 {
			line += " "
		}
		line += arg
	}
	return line
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

	fieldType := protoToParquet[protoKind]
	if fieldType == "" {
		log.Fatalf("type %v for field %v was not recognized", protoKind, fieldName)
	}

	if strings.HasPrefix(fieldType, "uint") && *noUnsigned == true {
		fieldType = strings.Replace(fieldType, "uint", "int", 1)
	}

	isRepeatedMessage := false
	if protoKind == "message" && field.IsList() {
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
		annotation = "(LIST)"
	}

	g.P(lineFmt(
		indentLevel,
		"optional",
		fieldType,
		fieldName,
		annotation,
		lineEnd,
	))

	if isRepeatedMessage {
		indentLevel++
		g.P(lineFmt(indentLevel, "repeated group list {"))
		indentLevel++
		g.P(lineFmt(indentLevel, "optional group element {"))
	}

	if protoKind == "message" {
		fds := field.Message().Fields()
		for i := 0; i < fds.Len(); i++ {
			field := fds.Get(i)
			generateField(g, field, indentLevel+1)
		}
		g.P(lineFmt(indentLevel, "}"))
	}

	if isRepeatedMessage {
		g.P(lineFmt(indentLevel-1, "}"))
		g.P(lineFmt(indentLevel-2, "}"))
	}
}
