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
	"testing"

	"github.com/galaho/pathogen/internal/tests"
	"github.com/galaho/pathogen/pkg/repositories"

	. "github.com/smartystreets/goconvey/convey"
)

func TestFileResolver(t *testing.T) {

	Convey("When FileResolver", t, func() {

		Convey(".Resolve is invoked", func() {

			variables := []repositories.Variable{
				repositories.Variable{
					Name:        "name",
					Description: "your name",
					Value:       "",
				},
				repositories.Variable{
					Name:        "quest",
					Description: "your quest",
					Value:       "",
				},
				repositories.Variable{
					Name:        "color",
					Description: "your favorite colour",
					Value:       "",
				},
			}

			Convey("with an invalid file", func() {

				resolver := NewFileResolver(tests.MustGenerateHex(t))
				actual, err := resolver.Resolve(variables)

				Convey("it returns a nil variable map", func() {
					So(actual, ShouldBeNil)
				})

				Convey("it returns a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("with an improperly formatted file", func() {

				resolver := NewFileResolver("testdata/invalid.yaml")
				actual, err := resolver.Resolve(variables)

				Convey("it returns a nil variable map", func() {
					So(actual, ShouldBeNil)
				})

				Convey("it returns a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("with a properly formatted file", func() {

				resolver := NewFileResolver("testdata/valid.yaml")

				Convey("and an empty variable slice", func() {

					variables := []repositories.Variable{}
					expected := map[string]string{}
					actual, err := resolver.Resolve(variables)

					Convey("it returns an empty variable map", func() {
						So(actual, ShouldResemble, expected)
					})

					Convey("it returns a nil error", func() {
						So(err, ShouldBeNil)
					})

				})

				Convey("and a non-empty variable slice", func() {

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
	})
}
