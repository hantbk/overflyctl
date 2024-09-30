/*
Copyright © 2024 Ha Nguyen <captainnemot1k60@gmail.com>

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

	"github.com/hantbk/overflyctl/database"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete a server from the database",
	Long:  "Delete a server from the database by specifying its ID",
	Run: func(cmd *cobra.Command, args []string) {
		id, err := cmd.Flags().GetInt("id")
		if err != nil {
			fmt.Println("Error getting id flag:", err)
			return
		}

		if id <= 0 {
			fmt.Println("Invalid server ID. Please provide a positive integer.")
			return
		}

		err = database.DeleteServer(id)
		if err != nil {
			fmt.Println("Error deleting server:", err)
			return
		}
		fmt.Printf("Server with ID %d deleted successfully\n", id)
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().Int("id", 0, "ID of the server to delete")
	deleteCmd.MarkFlagRequired("id")
}
