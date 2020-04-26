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
