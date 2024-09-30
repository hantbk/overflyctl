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
	"github.com/hantbk/overflyctl/ssh"
	"github.com/spf13/cobra"
)

var connectCmd = &cobra.Command{
	Use:   "connect",
	Short: "Connect to a server and copy SSH key",
	Long:  `Connect to a server from the database, copy the SSH key, and establish an SSH session`,
	Run: func(cmd *cobra.Command, args []string) {
		serverName, _ := cmd.Flags().GetString("server")

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

		// Ensure the SSH key path is absolute
		if targetServer.SSHKeyPath[:2] == "~/" {
			homeDir, err := os.UserHomeDir()
			if err != nil {
				fmt.Println("Error getting home directory:", err)
				return
			}
			targetServer.SSHKeyPath = filepath.Join(homeDir, targetServer.SSHKeyPath[2:])
		}

		// Copy SSH key
		publicKeyPath := targetServer.SSHKeyPath + ".pub"
		err = ssh.CopySSHKey(publicKeyPath, targetServer.IP, targetServer.Username, targetServer.Password)
		if err != nil {
			fmt.Println("Error copying SSH key:", err)
			return
		}

		fmt.Println("SSH key copied successfully")

		// Establish SSH connection
		err = ssh.InteractiveSession(targetServer.IP, targetServer.SSHKeyPath, targetServer.Username)
		if err != nil {
			fmt.Println("Error establishing SSH connection:", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(connectCmd)

	connectCmd.Flags().String("server", "", "Server name")
	connectCmd.Flags().String("username", "root", "SSH username")

	connectCmd.MarkFlagRequired("server")
}
