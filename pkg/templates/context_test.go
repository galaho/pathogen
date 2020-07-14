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

package templates

import (
	"os"
	"testing"
	"time"

	"github.com/galaho/pathogen/internal/tests"
	. "github.com/smartystreets/goconvey/convey"
)

func TestContext(t *testing.T) {

	Convey("When Context", t, func() {

		variableName := tests.MustGenerateString(t)
		variableValue := tests.MustGenerateString(t)

		context := &Context{
			Variables: map[string]string{
				variableName: variableValue,
			},
		}

		Convey(".Functions is invoked", func() {

			functions := context.Functions()

			Convey("it returns a map that contains", func() {

				Convey("an environment function", func() {

					function, exists := functions["environment"].(func(string) string)

					name := tests.MustGenerateHex(t)
					value := tests.MustGenerateHex(t)

					os.Setenv(name, value)

					So(exists, ShouldBeTrue)

					Convey("that returns an empty string when invoked for an unset environment variable", func() {
						So(function(tests.MustGenerateHex(t)), ShouldBeEmpty)
					})

					Convey("that returns value when invoked for a set environment variable", func() {
						So(function(name), ShouldEqual, value)
					})
				})

				Convey("a now function", func() {

					function, exists := functions["now"].(func(string) string)

					So(exists, ShouldBeTrue)

					Convey("that returns the same time when invoked multiple times", func() {
						So(function(time.RFC3339), ShouldEqual, function(time.RFC3339))
					})
				})

				Convey("a variable function", func() {

					function, exists := functions["variable"].(func(string) (string, error))

					So(exists, ShouldBeTrue)

					Convey("that when invoked", func() {

						Convey("and the variable does not exist", func() {

							value, err := function(tests.MustGenerateString(t))

							Convey("returns an empty string", func() {
								So(value, ShouldBeZeroValue)
							})

							Convey("returns a non-nil error", func() {
								So(err, ShouldNotBeNil)
							})
						})

						Convey("and the variable exists", func() {

							value, err := function(variableName)

							Convey("returns the value", func() {
								So(value, ShouldEqual, variableValue)
							})

							Convey("returns a nil error", func() {
								So(err, ShouldBeNil)
							})
						})
					})
				})
			})
		})
	})
}
