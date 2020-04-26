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
	"time"
)

// Context represents a context.
type Context struct {

	// Variables define the variables for the context.
	Variables map[string]string
}

// Functions returns the function map for the context.
func (c *Context) Functions() map[string]interface{} {
	return map[string]interface{}{
		"now":      now(),
		"variable": variable(c.Variables),
	}
}

// now returns a function that returns the current time as a formatted string.
func now() func(string) string {
	present := time.Now()
	return func(format string) string {
		return present.Format(format)
	}
}

// variable returns a function that returns a variable from a map.
func variable(variables map[string]string) func(string) (string, error) {
	return func(name string) (string, error) {
		value, exists := variables[name]
		if !exists {
			return "", fmt.Errorf("error fetching variable with name [%s]", name)
		}
		return value, nil
	}
}
