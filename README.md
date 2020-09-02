# protoc-gen-parquet

A protoc plugin that generates parquet schemas from protobuf files. See [examples](./examples).

## Installation

```bash
go get github.com/simo7/protoc-gen-parquet
```

Alternatively clone the repo and build the plugin:

```bash
go build -o bin/protoc-gen-parquet .

export PATH=$PWD/bin:$PATH
```

## Usage

```bash
protoc \
    --parquet_out=no_unsigned=true:./ \
    --parquet_opt=paths=source_relative \
    examples/person.proto
```

## Flags

`no_unsigned` (bool): Avoid unsigned integers and use the corresponding intenger instead.

`timestamp_int96` (bool): Fields extended as timestamps using [hive options](https://github.com/simo7/protoc-gen-gluecatalog/blob/master/hive_options/hive_options.proto)
can be defined as `INT96` instead of `INT64` to ensure compatibility with all Hive and Presto versions.

## Parquet Annotations

The following annotations are *not* implemented.

- `(TIMESTAMP(MILLIS, true))`, `(TIMESTAMP(MICROS, true))`, `(TIMESTAMP(NANOS, true))`
- `(DATE)`
- `(UUID)`
- `(MAP)`, `(MAP_KEY_VALUE)`
- `(STRING)`  (all UTF8 by default)

## Compatibility

It's tested against the new protobuf API `google.golang.org/protobuf` or version `1.4.0` of the legacy API `github.com/golang/protobuf`.
