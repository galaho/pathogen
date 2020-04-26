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
	"fmt"
	"testing"

	"github.com/galaho/pathogen/internal/tests"
	. "github.com/smartystreets/goconvey/convey"
)

func TestTemplates(t *testing.T) {

	Convey("When templates", t, func() {

		Convey(".Render is invoked", func() {

			variableName := tests.MustGenerateHex(t)
			variableValue := tests.MustGenerateHex(t)

			context := &Context{
				Variables: map[string]string{
					variableName: variableValue,
				},
			}

			source := fmt.Sprintf("{{ variable \"%s\" }}", variableName)

			Convey("with an invalid source", func() {

				source = "{{ nothing }}"

				rendered, err := Render(source, context)

				Convey("it returns an empty string", func() {
					So(rendered, ShouldBeZeroValue)
				})

				Convey("it returns a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("with an invalid context", func() {

				context = &Context{}

				rendered, err := Render(source, context)

				Convey("it returns an empty string", func() {
					So(rendered, ShouldBeZeroValue)
				})

				Convey("it returns a non-nil error", func() {
					So(err, ShouldNotBeNil)
				})
			})

			Convey("with a valid source and context", func() {

				rendered, err := Render(source, context)

				Convey("it returns the rendered template", func() {
					So(rendered, ShouldEqual, variableValue)
				})

				Convey("it returns a nil error", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
