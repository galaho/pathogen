/*
 * Copyright 2020 Joshua Mark Rutherford
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

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
