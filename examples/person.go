package main

const parquetSchema = `message person {
  optional binary name (UTF8);
  optional int32 age;
  optional group addresses (LIST) {
    repeated group list {
      optional binary element (UTF8);
    }
  }
  optional group phones (LIST) {
    repeated group list {
      optional group element {
        optional int32 number;
        optional binary type (UTF8);
        optional group carriers (LIST) {
          repeated group list {
            optional binary element (UTF8);
          }
        }
      }
    }
  }
  optional int64 entry_timestamp (TIMESTAMP_MILLIS);
`
