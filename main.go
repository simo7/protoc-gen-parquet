package main

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"unicode"

	parquetOpts "github.com/simo7/protoc-gen-parquet/parquet_options"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

var (
	flags          flag.FlagSet
	noUnsigned             = flags.Bool("no_unsigned", false, "use correspondant integer in place of unsigned integer")
	timestampInt96         = flags.Bool("timestamp_int96", false, "use INT96 for Hive timestamps")
	gofile                 = flags.Bool("go_file", false, "output a .go file containing the schema as a string constant")
	outputMessages         = flags.String("output_messages", "", "output semicolon-separated message objects, instead of just the ones with a table name annotation")
	maxRecursionLevel      = flags.Int("max_recursion", 0, "max recursion depth, disabled by default")
	outputMessagesList     []string
)

var protoToParquet = map[string]string{
	"bool":             "boolean",
	"int32":            "int32",
	"int64":            "int64",
	"uint32":           "uint32",
	"uint64":           "uint64",
	"fixed32":          "uint32",
	"fixed64":          "uint64",
	"float":            "float",
	"float64":          "double",
	"double":           "double",
	"string":           "binary",
	"[]byte":           "binary",
	"bytes":            "binary",
	"enum":             "binary",
	"message":          "group",
	"sfixed32":         "int32",
	"sfixed64":         "int64",
	"sint32":           "int32",
	"sint64":           "int64",
	"timestamp_millis": "int64",
	"timestamp_micros": "int64",
	"timestamp_nanos":  "int64",
	"int96":            "int96",
}

var protoAnnotations = map[string]string{
	"string":           " (UTF8)",
	"enum":             " (UTF8)",
	"timestamp_millis": " (TIMESTAMP_MILLIS)",
	"timestamp_micros": " (TIMESTAMP_MICROS)",
	"timestamp_nanos":  " (TIMESTAMP_NANOS)",
}

func main() {
	protogen.Options{
		ParamFunc: flags.Set,
	}.Run(func(gen *protogen.Plugin) error {
		outputMessagesList = strings.Split(*outputMessages,";")
		for _, f := range gen.Files {
			if f.Generate {
				generateFile(gen, f)
			}
		}
		return nil
	})
}

func getIndent(indentLevel int) string {
	return strings.Repeat(" ", indentLevel*2)
}

func generateFile(gen *protogen.Plugin, file *protogen.File) {
	path, _ := filepath.Split(file.GeneratedFilenamePrefix)

	for _, message := range file.Messages {
		var tableName string
		messageName := string(message.Desc.Name())
		found := false
		for _, match := range outputMessagesList {
			if match == messageName {
				found = true
				break
			}
		}
		if found {
			tableName = strings.ToLower(messageName)
		} else {

			opts := message.Desc.Options()

			if !opts.ProtoReflect().IsValid() {
				continue
			}

			optValue := proto.GetExtension(opts, parquetOpts.E_MessageOpts)
			tableName = optValue.(*parquetOpts.MessageOptions).GetTableName()
			if tableName == "" {
				continue
			}
		}

		filename := path + tableName
		g := gen.NewGeneratedFile(filename+".schema", file.GoImportPath)

		g.P("message " + tableName + " {")
		for _, field := range message.Fields {
			fieldMap := map[protoreflect.FieldDescriptor]bool{}
			generateField(g, field.Desc, 1, fieldMap, 0)
		}
		g.P("}")

		if *gofile {
			g2 := gen.NewGeneratedFile(filename+".go", file.GoImportPath)
			content, _ := g.Content()

			g2.P("package " + file.GoPackageName)
			body := fmt.Sprintf("const PARQUET_SCHEMA_%s=`%s`", ToUpperSnakeCase(tableName), content)
			g2.P(body)
		}
	}
}

func generateField(g *protogen.GeneratedFile, field protoreflect.FieldDescriptor, indentLevel int, fieldMap map[protoreflect.FieldDescriptor]bool, recursionLevel int) {

	fieldName := string(field.Name())
	if _, ok := fieldMap[field]; ok {
		recursionLevel++
		if (*maxRecursionLevel != 0) && (recursionLevel >= *maxRecursionLevel) {
			return
		}
	} else {
		fieldMap[field] = true
	}

	protoKind := field.Kind().String()

	opts := field.Options()
	if opts.ProtoReflect().IsValid() {
		optValue := proto.GetExtension(opts, parquetOpts.E_FieldOpts)
		timestampType := optValue.(*parquetOpts.FieldOptions).GetTimestampType().String()
		if timestampType != "UNSPECIFIED" {
			protoKind = strings.ToLower(timestampType)
		}
	}

	if strings.HasPrefix(protoKind, "timestamp_") && *timestampInt96 {
		protoKind = "int96"
	}

	fieldType := protoToParquet[protoKind]
	if fieldType == "" {
		log.Fatalf("type %v for field %v was not recognized", protoKind, fieldName)
	}

	if strings.HasPrefix(fieldType, "uint") && *noUnsigned == true {
		fieldType = strings.Replace(fieldType, "uint", "int", 1)
	}

	lineEnd := ";"
	if protoKind == "message" {
		lineEnd = " {"

		fds := field.Message().Fields()
		if fds.Len() > 0 {
			if isProtoTimestamp(fds.Get(0)) {
				g.P(fmt.Sprintf("%soptional int64 %s (TIMESTAMP_NANOS);",
					getIndent(indentLevel), string(field.Name())))
				return
			}
		}
	}

	annotation := protoAnnotations[protoKind]

	if field.IsList() {
		g.P(fmt.Sprintf("%soptional group %s (LIST) {", getIndent(indentLevel), fieldName))
		indentLevel++
		g.P(fmt.Sprintf("%srepeated group list {", getIndent(indentLevel)))
		indentLevel++

		if protoKind == "message" {
			g.P(fmt.Sprintf("%soptional group element {", getIndent(indentLevel)))
		} else {
			g.P(fmt.Sprintf("%soptional %s element%s",
				getIndent(indentLevel), fieldType, annotation+lineEnd,
			))
		}

	} else {
		g.P(fmt.Sprintf("%soptional %s %s%s",
			getIndent(indentLevel), fieldType, fieldName, annotation+lineEnd,
		))
	}

	if protoKind == "message" {
		fds := field.Message().Fields()
		for i := 0; i < fds.Len(); i++ {
			fd := fds.Get(i)
			generateField(g, fd, indentLevel+1, fieldMap, recursionLevel)
		}
		g.P(fmt.Sprintf("%s}", getIndent(indentLevel)))
	}

	if field.IsList() {
		g.P(fmt.Sprintf("%s}", getIndent(indentLevel-1)))
		g.P(fmt.Sprintf("%s}", getIndent(indentLevel-2)))
	}
}

func isProtoTimestamp(fd protoreflect.FieldDescriptor) bool {
	return strings.HasPrefix(string(fd.FullName()), "google.protobuf.Timestamp.")
}

func ToUpperSnakeCase(s string) string {
	var res = make([]rune, 0, len(s))
	var p = '_'
	for i, r := range s {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			res = append(res, '_')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				res = append(res, '_', unicode.ToLower(r))
			} else {
				res = append(res, unicode.ToLower(r))
			}
		} else {
			res = append(res, unicode.ToLower(r))
		}

		p = r
	}
	return strings.ToUpper(string(res))
}
