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

var dbPath string

var rootCmd = &cobra.Command{
	Use:   "overfly",
	Short: "Overfly Cloud Command line",
	Long:  `Overfly Cloud Command line is a tool to interact with Overfly Cloud API`,

	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Error getting home directory:", err)
		os.Exit(1)
	}
	dbPath = filepath.Join(homeDir, ".overflyctl.db")

	// Check if the database file exists
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		fmt.Println("Database file does not exist. Creating a new one.")
		// Init database only if it doesn't exist
		err = database.InitDB(dbPath)
		if err != nil {
			fmt.Println("Error initializing database:", err)
			os.Exit(1)
		}
		fmt.Println("Database initialized successfully")
	} else {
		fmt.Println("Using existing database file:", dbPath)
	}

	// Open the database connection without reinitializing
	err = database.OpenDB(dbPath)
	if err != nil {
		fmt.Println("Error opening database:", err)
		os.Exit(1)
	}

	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
