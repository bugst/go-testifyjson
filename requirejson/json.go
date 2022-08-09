//
// Copyright 2022 Cristian Maglie. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//

package requirejson

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/itchyny/gojq"
	"github.com/stretchr/testify/require"
)

// Query performs a test on a given json output. A jq-like query is performed
// on the given jsonData and the result is compared with the expected output.
// If the output doesn't match the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func Query(t *testing.T, jsonData []byte, jqQuery string, jsonExpected string, msgAndArgs ...interface{}) {
	var data interface{}
	require.NoError(t, json.Unmarshal(jsonData, &data))
	var expected interface{}
	require.NoError(t, json.Unmarshal([]byte(jsonExpected), &expected))
	q, err := gojq.Parse(jqQuery)
	require.NoError(t, err)
	i := q.Run(data)
	v, ok := i.Next()
	require.True(t, ok)
	require.IsType(t, expected, v)
	require.Equal(t, expected, v, msgAndArgs...)
}

// Contains check if the json object is a subset of the jsonData.
// If the output doesn't match the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func Contains(t *testing.T, jsonData []byte, jsonObject string, msgAndArgs ...interface{}) {
	var data interface{}
	require.NoError(t, json.Unmarshal(jsonData, &data))
	q, err := gojq.Parse("contains(" + jsonObject + ")")
	require.NoError(t, err)
	i := q.Run(data)
	v, ok := i.Next()
	require.True(t, ok)
	require.IsType(t, true, v)
	if !v.(bool) {
		msg := fmt.Sprintf("json data does not contain: %s", jsonObject)
		require.FailNow(t, msg, msgAndArgs...)
	}
}

// NotContains check if the json object is NOT a subset of the jsonData.
// If the output match the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func NotContains(t *testing.T, jsonData []byte, jsonObject string, msgAndArgs ...interface{}) {
	var data interface{}
	require.NoError(t, json.Unmarshal(jsonData, &data))
	q, err := gojq.Parse("contains(" + jsonObject + ")")
	require.NoError(t, err)
	i := q.Run(data)
	v, ok := i.Next()
	require.True(t, ok)
	require.IsType(t, true, v)
	if v.(bool) {
		msg := fmt.Sprintf("json data contains: %s", jsonObject)
		require.FailNow(t, msg, msgAndArgs...)
	}
}

// Len check if the size of the json object match the given value.
// If the lenght doesn't match the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func Len(t *testing.T, jsonData []byte, expectedLen int, msgAndArgs ...interface{}) {
	var data interface{}
	require.NoError(t, json.Unmarshal(jsonData, &data))
	q, err := gojq.Parse("length")
	require.NoError(t, err)
	i := q.Run(data)
	v, ok := i.Next()
	require.True(t, ok)
	require.IsType(t, expectedLen, v)
	if v.(int) != expectedLen {
		msg := fmt.Sprintf("json data length does not match: expected=%d, actual=%d", expectedLen, v.(int))
		require.FailNow(t, msg, msgAndArgs...)
	}
}

// Empty check if the size of the json object is zero.
// If the lenght is not zero the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func Empty(t *testing.T, jsonData []byte, msgAndArgs ...interface{}) {
	var data interface{}
	require.NoError(t, json.Unmarshal(jsonData, &data))
	q, err := gojq.Parse("length")
	require.NoError(t, err)
	i := q.Run(data)
	v, ok := i.Next()
	require.True(t, ok)
	require.IsType(t, 0, v)
	if v.(int) != 0 {
		require.FailNow(t, "json data is not empty", msgAndArgs...)
	}
}

// NotEmpty check if the size of the json object is greater than zero.
// If the lenght is not greater than zero the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func NotEmpty(t *testing.T, jsonData []byte, msgAndArgs ...interface{}) {
	var data interface{}
	require.NoError(t, json.Unmarshal(jsonData, &data))
	q, err := gojq.Parse("length")
	require.NoError(t, err)
	i := q.Run(data)
	v, ok := i.Next()
	require.True(t, ok)
	require.IsType(t, 0, v)
	if v.(int) == 0 {
		require.FailNow(t, "json data is empty", msgAndArgs...)
	}
}
