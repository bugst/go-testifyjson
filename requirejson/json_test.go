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

func TestJSONFlowParadigm(t *testing.T) {
	in := requirejson.Parse(t, []byte(`
{
	"id" : 1,
	"list" : [
		10, 20, 30
	],
	"emptylist" : []
}
`))

	in.Query(".list").MustEqual("[10, 20, 30]")
	in.Query(".list.[1]").MustEqual("20")

	in.MustContain(`{ "list": [ 30 ] }`)
	in.MustNotContain(`{ "list": [ 50 ] }`)

	in.Query(".list | length").MustEqual("3")

	in2 := requirejson.Parse(t, []byte(`[ ]`))
	in2.MustBeEmpty()
	in2.LengthMustEqualTo(0)

	in3 := requirejson.Parse(t, []byte(`[ 10, 20, 30 ]`))
	in3.MustNotBeEmpty()
	in3.LengthMustEqualTo(3)
}

func TestJSONAssertions(t *testing.T) {
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

	requirejson.Parse(t, in).Query(".list").ArrayMustContain("20")
}
