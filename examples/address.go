package examples

const PARQUET_SCHEMA_ADDRESS = `message address {
  optional binary street (UTF8);
  optional binary number (UTF8);
  optional int32 zip_code;
}
`
