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

// JQObject is an object that abstract operations on JSON data.
type JQObject struct {
	t    *testing.T
	data interface{}
}

func (obj *JQObject) String() string {
	r, err := json.Marshal(obj.data)
	require.NoError(obj.t, err)
	return string(r)
}

// Parse creates a new JQObect from the given jsonData.
// If jsonData is not a valid json the test fails.
func Parse(t *testing.T, jsonData []byte, msgAndArgs ...interface{}) *JQObject {
	var data interface{}
	require.NoError(t, json.Unmarshal(jsonData, &data), msgAndArgs...)
	return &JQObject{t: t, data: data}
}

// Query performs a query on the given JQObject and returns the first result. If the query
// produces no result the test will fail.
func (obj *JQObject) Query(jqQuery string) *JQObject {
	q, err := gojq.Parse(jqQuery)
	require.NoError(obj.t, err)
	iter := q.Run(obj.data)
	data, ok := iter.Next()
	require.True(obj.t, ok)
	return &JQObject{t: obj.t, data: data}
}

// MustEqual tests if the JQObject equals jsonExpected. If the check
// is not successful the test will fail. If msgAndArgs are provided they
// will be used to explain the error.
func (obj *JQObject) MustEqual(jsonExpected string, msgAndArgs ...interface{}) {
	jsonActual, err := json.Marshal(obj.data)
	require.NoError(obj.t, err)
	require.JSONEq(obj.t, jsonExpected, string(jsonActual), msgAndArgs...)
}

// MustContain tests if jsonExpected is a subset of the JQObject. If the check
// is not successful the test will fail. If msgAndArgs are provided they
// will be used to explain the error.
func (obj *JQObject) MustContain(jsonExpected string, msgAndArgs ...interface{}) {
	v := obj.Query("contains(" + jsonExpected + ")").data
	require.IsType(obj.t, true, v)
	if !v.(bool) {
		msg := fmt.Sprintf("json data does not contain: %s", jsonExpected)
		require.FailNow(obj.t, msg, msgAndArgs...)
	}
}

// MustNotContain tests if the JQObject does not contain jsonExpected. If the check
// is not successful the test will fail. If msgAndArgs are provided they
// will be used to explain the error.
func (obj *JQObject) MustNotContain(jsonExpected string, msgAndArgs ...interface{}) {
	v := obj.Query("contains(" + jsonExpected + ")").data
	require.IsType(obj.t, true, v)
	if v.(bool) {
		msg := fmt.Sprintf("json data contains: %s", jsonExpected)
		require.FailNow(obj.t, msg, msgAndArgs...)
	}
}

// LengthMustEqualTo tests if the size of JQObject match expectedLen. If the check
// is not successful the test will fail. If msgAndArgs are provided they
// will be used to explain the error.
func (obj *JQObject) LengthMustEqualTo(expectedLen int, msgAndArgs ...interface{}) {
	v := obj.Query("length").data
	require.IsType(obj.t, expectedLen, v)
	if v.(int) != expectedLen {
		msg := fmt.Sprintf("json data length does not match: expected=%d, actual=%d", expectedLen, v.(int))
		require.FailNow(obj.t, msg, msgAndArgs...)
	}
}

// MustBeEmpty test if the size of JQObject is 0. If the check
// is not successful the test will fail. If msgAndArgs are provided they
// will be used to explain the error.
func (obj *JQObject) MustBeEmpty(msgAndArgs ...interface{}) {
	v := obj.Query("length").data
	require.IsType(obj.t, 0, v)
	if v.(int) != 0 {
		require.FailNow(obj.t, "json data is not empty", msgAndArgs...)
	}
}

// MustNotBeEmpty test if the size of JQObject is not 0. If the check
// is not successful the test will fail. If msgAndArgs are provided they
// will be used to explain the error.
func (obj *JQObject) MustNotBeEmpty(msgAndArgs ...interface{}) {
	v := obj.Query("length").data
	require.IsType(obj.t, 0, v)
	if v.(int) == 0 {
		require.FailNow(obj.t, "json data is empty", msgAndArgs...)
	}
}

// Query performs a test on a given json output. A jq-like query is performed
// on the given jsonData and the result is compared with jsonExpected.
// If the output doesn't match the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func Query(t *testing.T, jsonData []byte, jqQuery string, jsonExpected string, msgAndArgs ...interface{}) {
	data := Parse(t, jsonData)
	v := data.Query(jqQuery)
	v.MustEqual(jsonExpected, msgAndArgs...)
}

// Contains check if the json object is a subset of the jsonData.
// If the output doesn't match the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func Contains(t *testing.T, jsonData []byte, jsonObject string, msgAndArgs ...interface{}) {
	Parse(t, jsonData).MustContain(jsonObject, msgAndArgs...)
}

// NotContains check if the json object is NOT a subset of the jsonData.
// If the output match the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func NotContains(t *testing.T, jsonData []byte, jsonObject string, msgAndArgs ...interface{}) {
	Parse(t, jsonData).MustNotContain(jsonObject, msgAndArgs...)
}

// Len check if the size of the json object match the given value.
// If the lenght doesn't match the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func Len(t *testing.T, jsonData []byte, expectedLen int, msgAndArgs ...interface{}) {
	Parse(t, jsonData).LengthMustEqualTo(expectedLen, msgAndArgs...)
}

// Empty check if the size of the json object is zero.
// If the lenght is not zero the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func Empty(t *testing.T, jsonData []byte, msgAndArgs ...interface{}) {
	Parse(t, jsonData).MustBeEmpty(msgAndArgs...)
}

// NotEmpty check if the size of the json object is not zero.
// If the lenght is zero the test fails. If msgAndArgs are provided they
// will be used to explain the error.
func NotEmpty(t *testing.T, jsonData []byte, msgAndArgs ...interface{}) {
	Parse(t, jsonData).MustNotBeEmpty(msgAndArgs...)
}
