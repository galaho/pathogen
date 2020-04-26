// Copyright 2020 Joshua Mark Rutherford
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package templates

import (
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
