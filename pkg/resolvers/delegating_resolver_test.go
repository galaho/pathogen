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

package resolvers

import (
	"errors"
	"reflect"
	"testing"

	"github.com/galaho/pathogen/internal/tests"
	"github.com/galaho/pathogen/pkg/repositories"

	. "github.com/smartystreets/goconvey/convey"
)

// mockResolver implements a mock variable resolver for tests.
type mockResolver struct {
	variables map[string]string
	err error
}

// Resolve resolves variables.
func (r *mockResolver) Resolve(_ []repositories.Variable) (map[string]string, error) {
	return r.variables, r.err
}

func TestDelegatingResolver(t *testing.T) {

	Convey("When DelegatingResolver", t, func() {

		variables := tests.MustGenerate(reflect.TypeOf(map[string]string{}), t).Interface().(map[string]string)

		Convey(".Resolve is invoked", func() {

			Convey("and all resolvers error", func() {

				resolver := NewDelegatingResolver(
					&mockResolver{variables: variables, err: errors.New(tests.MustGenerateHex(t))},
					&mockResolver{variables: variables, err: errors.New(tests.MustGenerateHex(t))},
					&mockResolver{variables: variables, err: errors.New(tests.MustGenerateHex(t))},
				)

				actual, err := resolver.Resolve(nil)

				Convey("it returns a nil variable map", func() {
					So(actual, ShouldBeNil)
				})

				Convey("it returns a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("and no resolvers error", func() {

				resolver := NewDelegatingResolver(
					&mockResolver{variables: variables, err: nil},
					&mockResolver{variables: tests.MustGenerate(reflect.TypeOf(map[string]string{}), t).Interface().(map[string]string), err: nil},
					&mockResolver{variables: tests.MustGenerate(reflect.TypeOf(map[string]string{}), t).Interface().(map[string]string), err: nil},
				)

				actual, err := resolver.Resolve(nil)

				Convey("it returns the variables from the first resolver", func() {
					So(actual, ShouldEqual, variables)
				})

				Convey("it returns a nil error", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
