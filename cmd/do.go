/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/vncsb/tasks/tasksdb"
)

// doCmd represents the do command
var doCmd = &cobra.Command{
	Use:   "do [task id]",
	Short: "Mark task as complete",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		idString := args[0]
		id, err := strconv.Atoi(idString)
		if err != nil {
			fmt.Println("Task ID required")
			return
		}

		description, err := tasksdb.DoTask(id)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("You have completed the \"%v\" task.", description)
	},
}

func init() {
	rootCmd.AddCommand(doCmd)
}
