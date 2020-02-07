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

package cmd

import (
	"github.com/galaho/pathogen/cmd/version"
	"github.com/spf13/cobra"
)

// Pathogen returns a command that generates filesystem entries from templates.
func Pathogen() *cobra.Command {

	command := &cobra.Command{
		Use:   "pathogen",
		Short: "Generate filesystem entries from templates",
		Long:  "A command line utility for generating filesystem entries from templates.",
	}

	command.AddCommand(version.Command())

	return command
}
