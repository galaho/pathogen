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

package version

import (
	"fmt"

	"github.com/galaho/pathogen/versions"
	"github.com/spf13/cobra"
)

// Command returns a command that prints the version.
func Command() *cobra.Command {

	command := &cobra.Command{
		Use:   "version",
		Short: "Print the version and commit",
		Run: func(command *cobra.Command, args []string) {
			fmt.Printf("pathogen %s (%s)\n", versions.Version(), versions.Commit())
		},
	}

	return command
}
