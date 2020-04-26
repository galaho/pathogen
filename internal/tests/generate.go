// Copyright (c) 2020 Joshua Mark Rutherford
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package tests

import (
	"encoding/hex"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"
	"time"
)

var (
	// Random provides a random source for generative test data.
	Random = rand.New(rand.NewSource(time.Now().Unix()))
)

// MustGenerate generates and returns a random value of a type or fails a test.
func MustGenerate(tipe reflect.Type, test *testing.T) reflect.Value {
	value, ok := quick.Value(tipe, Random)
	if !ok {
		test.Errorf("unable to generate random value of type [%s]", tipe)
	}
	return value
}

// MustGenerateBytes generates a random by slice or fails a test.
func MustGenerateBytes(test *testing.T) []byte {
	return MustGenerate(reflect.TypeOf([]byte{}), test).Interface().([]byte)
}

// MustGenerateHex generates a random hexadecimal string or fails a test.
func MustGenerateHex(test *testing.T) string {
	return hex.EncodeToString(MustGenerateBytes(test))
}

// MustGenerateString generates a random string or fails a test.
func MustGenerateString(test *testing.T) string {
	return MustGenerate(reflect.TypeOf(""), test).Interface().(string)
}
