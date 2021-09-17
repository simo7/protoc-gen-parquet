package examples

const Address_parquet_schema = `message address {
  optional binary street (UTF8);
  optional binary number (UTF8);
  optional int32 zip_code;
}
`
