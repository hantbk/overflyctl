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
	"os"
	"path/filepath"

	"github.com/hantbk/overflyctl/database"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new Linux server",
	Long:  `Add a new Linux server to the database`,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		ip, _ := cmd.Flags().GetString("ip")
		username, _ := cmd.Flags().GetString("username")
		password, _ := cmd.Flags().GetString("password")
		sshKeyPath, _ := cmd.Flags().GetString("ssh-key")

		// Expand the tilde to the user's home directory
		if sshKeyPath[:2] == "~/" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Error getting home directory:", err)
				return
			}
			sshKeyPath = filepath.Join(homeDir, sshKeyPath[2:])
		}

		err := database.AddServer(name, ip, username, password, sshKeyPath)
		if err != nil {
			fmt.Println("Error adding server:", err)
			return
		}

		fmt.Println("Server added successfully")
	},
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().String("name", "", "Server name")
	addCmd.Flags().String("ip", "", "Server IP address")
	addCmd.Flags().String("username", "", "SSH username")
	addCmd.Flags().String("password", "", "SSH password")
	addCmd.Flags().String("ssh-key", "", "Path to SSH key")

	addCmd.MarkFlagRequired("name")
	addCmd.MarkFlagRequired("ip")
	addCmd.MarkFlagRequired("username")
	addCmd.MarkFlagRequired("password")
	addCmd.MarkFlagRequired("ssh-key")
}
