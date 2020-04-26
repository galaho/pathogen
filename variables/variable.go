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

// Variable represents a variable for substitution in a template.
type Variable struct {

	// Name defines the name of the variable.
	Name string `json:"name" mapstructure:"name" yaml:"name"`

	// Description defines the description the variable.
	Description string `json:"description" mapstructure:"description" yaml:"description"`

	// Value defines the value of the variable.
	Value string `json:"value" mapstructure:"value" yaml:"value"`
}
