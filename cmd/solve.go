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
	tourist "github.com/xydinesh/tourist/pkg"
)

var TSPFile string
var Temperature float64
var Beta float64
var Iterations int
var Goal float32
var Output int
var Solution bool

// solveCmd represents the solve command
var solveCmd = &cobra.Command{
	Use:   "solve",
	Short: "Solve TSP for given instance",
	Long:  `Solve Traveling Salesman Problem (TSP) for a given problem instance provided as a text file.`,
	Run: func(cmd *cobra.Command, args []string) {
		tsp_instance, err := tourist.ReadDataFile(TSPFile)
		if err != nil {
			fmt.Errorf("An error happened, %s", err)
		}
		r := tourist.GenerateRandomRoute(tsp_instance.Dimension)
		s := tourist.StopConditon{Goal: Goal, Iterations: Iterations, Output: Output}
		nr := tsp_instance.ComputeOptimalRoute(&r, Temperature, Beta, &s)
		cost := tsp_instance.GetRouteCost(&nr)
		fmt.Printf("Final Cost: %f\n", cost)
		if Solution == true {
			fmt.Printf("Final Solution: %v\n", nr.NodeOrder)
		}
	},
}

func init() {
	solveCmd.Flags().StringVarP(&TSPFile, "file", "f", "", "File with TSP instance")
	solveCmd.Flags().Float64VarP(&Temperature, "temp", "t", 10.0, "Initial Temperature for the instance")
	solveCmd.Flags().Float64VarP(&Beta, "beta", "b", 0.99981, "Beta value for reducing temperature")
	solveCmd.Flags().Float32VarP(&Goal, "goal", "g", 0.0, "Cost value of the solution")
	solveCmd.Flags().IntVarP(&Iterations, "iterations", "i", 1000, "Iterations to run")
	solveCmd.Flags().IntVarP(&Output, "output", "o", 1000, "How often to see an output")
	solveCmd.Flags().BoolVarP(&Solution, "solution", "s", false, "Print final solution")
	RootCmd.AddCommand(solveCmd)
}
