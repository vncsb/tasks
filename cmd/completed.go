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
	"time"

	"github.com/spf13/cobra"
	"github.com/vncsb/tasks/tasksdb"
)

// completedCmd represents the completed command
var completedCmd = &cobra.Command{
	Use:   "completed",
	Short: "List tasks completed today",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := tasksdb.ListTasks(func(t tasksdb.Task) bool {
			if t.Done {
				return completedToday(t)
			}
			return false
		})

		if err != nil {
			fmt.Println(err)
			return
		}

		if len(tasks) == 0 {
			fmt.Println("No completed tasks yet, get to work!")
			return
		}

		fmt.Println("You have finished the following tasks today:")
		for _, task := range tasks {
			fmt.Printf("%v. %v\n", task.ID, task.Description)
		}
	},
}

func completedToday(task tasksdb.Task) bool {
	y1, m1, d1 := task.DateCompleted.Date()
	y2, m2, d2 := time.Now().Date()
	return y1 == y2 && m1 == m2 && d1 == d2
}

func init() {
	rootCmd.AddCommand(completedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// completedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// completedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
