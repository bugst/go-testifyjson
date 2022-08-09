# `go.bug.st/testifyjson/requirejson` - unit-test JSON output in golang.

Package testifyjson is a collection of utilities and helper function for unit testing
JSON output in golang.

It is based on the excellent libraries `github.com/itchyny/gojq` and `github.com/stretchr/testify`. It provides an interface similar to `testify` but with the powerful methods available in `gojq`.

## Examples

```go
import (
    "testing"
    "go.bug.st/testifyjson/requirejson"
)

func TestJSONQuery(t *testing.T) {
    in := []byte(`
{
    "id" : 1,
    "list" : [
        10, 20, 30
    ],
    "emptylist" : []
}
`)
    requirejson.Query(t, in, ".list", "[10, 20, 30]")
    requirejson.Query(t, in, ".list.[1]", "20")

    requirejson.Contains(t, in, `{ "list": [ 30 ] }`)
    requirejson.NotContains(t, in, `{ "list": [ 50 ] }`)

    in2 := []byte(`[ ]`)
    requirejson.Empty(t, in2)
    requirejson.Len(t, in2, 0)

    in3 := []byte(`[ 10, 20, 30 ]`)
    requirejson.NotEmpty(t, in3)
    requirejson.Len(t, in3, 3)
}
```
