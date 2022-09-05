//
// Copyright 2022 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package requirejson_test

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

	requirejson.Query(t, in, ".list | length", "3")
}
