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

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List all Linux servers",
	Long:  `List all Linux servers that are installed with Overfly Cloud Agent for Backup/Restore`,
	Run: func(cmd *cobra.Command, args []string) {
		servers, err := database.ListServers()
		if err != nil {
			fmt.Println("Error listing servers:", err)
			return
		}

		if len(servers) == 0 {
			fmt.Println("No servers found.")
			return
		}

		fmt.Println("List of servers:")
		for _, server := range servers {
			fmt.Printf("ID: %d, Name: %s, IP: %s, Username: %s, Password: %s, SSH Key: %s\n",
				server.ID, server.Name, server.IP, server.Username, server.Password, server.SSHKeyPath)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
