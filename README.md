# protoc-gen-parquet

A protoc plugin that generates parquet schemas from protobuf files. See [examples](./examples).

## Installation

```bash
go get github.com/atoulme/protoc-gen-parquet
```

Alternatively clone the repo and build the plugin:

```bash
go build -o bin/protoc-gen-parquet .

export PATH=$PWD/bin:$PATH
```

## Usage

Generate parquet schema:

```bash
protoc \
    --parquet_out=no_unsigned=true,go_file=true,max_recursion=5,output_messages=Person;Address:./ \
    --parquet_opt=paths=source_relative \
    --go_opt=paths=source_relative \
    --go_out=./ \
    examples/person.proto
```

Re-generate parquet_options stubs:

```bash
protoc \
    --go_opt=paths=source_relative \
    --go_out=./ \
    parquet_options/parquet_options.proto
```

## Flags

`no_unsigned` (bool): Avoid unsigned integers and use the corresponding intenger instead.

`timestamp_int96` (bool): Fields extended as timestamps (see `timestamp_type` in [parquet options](parquet_options/parquet_options.proto)
can be defined as `INT96` instead of `INT64` to ensure compatibility with all Hive and Presto versions.

`go_file` (bool): An additional `.go` file containing the schema as a string constant will be generated. It makes it easier to import
a versioned schema into a Go application.

`max_recursion` (int): An optional recursion level - a number of times the schema will repeat the same element.

`output_messages` (string): Output semicolon-separated message objects, instead of just the ones with a table name annotation

## Parquet Annotations

The following annotations are *not* implemented.

- `(DATE)`
- `(UUID)`
- `(MAP)`, `(MAP_KEY_VALUE)`
- `(STRING)`  (all UTF8 by default)

## Well-known Protobuf types

Reference: https://developers.google.com/protocol-buffers/docs/reference/google.protobuf.

The following types are supported:

- [x] `google.protobuf.Timestamp`

## Compatibility

It's tested against the new protobuf API `google.golang.org/protobuf` or version `1.4.0` of the legacy API `github.com/golang/protobuf`.
