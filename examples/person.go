package examples

const person = `message person {
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
  optional int64 created_at (TIMESTAMP_MILLIS);
  optional int64 updated_at (TIMESTAMP_MILLIS);
  optional int64 generated_at (TIMESTAMP_MICROS);
}
`
