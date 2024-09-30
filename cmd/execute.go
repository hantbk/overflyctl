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
	"github.com/hantbk/overflyctl/ssh"
	"github.com/spf13/cobra"
)

var executeCmd = &cobra.Command{
	Use:   "execute",
	Short: "Execute a command on a remote server",
	Long:  `Execute a command on a remote server using the stored SSH information`,
	Run: func(cmd *cobra.Command, args []string) {
		serverName, _ := cmd.Flags().GetString("server")
		command, _ := cmd.Flags().GetString("command")

		servers, err := database.ListServers()
		if err != nil {
			fmt.Println("Error listing servers:", err)
			return
		}

		var targetServer database.Server
		for _, server := range servers {
			if server.Name == serverName {
				targetServer = server
				break
			}
		}

		if targetServer.Name == "" {
			fmt.Println("Server not found")
			return
		}

		output, err := ssh.ExecuteCommand(targetServer.IP, targetServer.Username, targetServer.Password, targetServer.SSHKeyPath, command)
		if err != nil {
			fmt.Println("Error executing command:", err)
			return
		}

		fmt.Println("Command output:")
		fmt.Println(output)
	},
}

func init() {
	rootCmd.AddCommand(executeCmd)

	executeCmd.Flags().String("server", "", "Server name")
	executeCmd.Flags().String("command", "", "Command to execute")

	executeCmd.MarkFlagRequired("server")
	executeCmd.MarkFlagRequired("command")
}
