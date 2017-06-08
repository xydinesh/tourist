// Copyright © 2017 Dinesh Weerapurage <xydinesh@gmail.com>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var TSPFile string

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve TSP for given instance",
	Long:  `Solve Traveling Salesman Problem (TSP) for a given problem instance provided as a text file.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("solve called")
		fmt.Printf("%s\n", TSPFile)
	},
}

func init() {
	solveCmd.Flags().StringVarP(&TSPFile, "tsp", "t", "", "File with TSP instance")
	RootCmd.AddCommand(solveCmd)
}
