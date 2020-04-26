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

package variables

import (
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFileResolver(t *testing.T) {

	Convey("When FileResolver", t, func() {

		Convey(".Resolve is invoked", func() {

			Convey("with an empty variable slice", func() {

				reader := strings.NewReader("orange\n")
				writer := &strings.Builder{}
				resolver := NewIOResolver(reader, writer)

				variables := []Variable{}
				expected := map[string]string{}
				actual, err := resolver.Resolve(variables)

				Convey("it returns an empty variable map", func() {
					So(actual, ShouldResemble, expected)
				})

				Convey("it returns a nil error", func() {
					So(err, ShouldBeNil)
				})
			})

			Convey("with a non-empty variable slice", func() {

				reader := strings.NewReader("Sir Galahad\nTo seek the Holy Grail\nBlue\n")
				writer := &strings.Builder{}
				resolver := NewIOResolver(reader, writer)

				variables := []Variable{
					Variable{
						Name:        "name",
						Description: "your name",
						Value:       "",
					},
					Variable{
						Name:        "quest",
						Description: "your quest",
						Value:       "",
					},
					Variable{
						Name:        "color",
						Description: "your favorite colour",
						Value:       "",
					},
				}

				expected := map[string]string{
					"name":  "Sir Galahad",
					"quest": "To seek the Holy Grail",
					"color": "Blue",
				}

				actual, err := resolver.Resolve(variables)

				Convey("it returns the expected variable map", func() {
					So(actual, ShouldResemble, expected)
				})

				Convey("it returns a nil error", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
