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
	"errors"
	"fmt"
	"github.com/ChicK00o/copy-cmd/userFunctions"
	"log"

	"github.com/spf13/cobra"
)

// executeWithOverrideCmd represents the executeWithOverride command
var executeWithOverrideCmd = &cobra.Command{
	Use:   "executeWithOverride",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("executeWithOverride called")
		if (SourcePath == "" && ExeSourcePath == "") ||
			(DestinationPath == "" && ExeDestinationPath == "") ||
			(ConfigFilePath == "" && ExeConfigFilePath == "") {
			return errors.New("Missing Data to execute, check 'show' command")
		}
		source := ""
		destination := ""
		config := ""

		if ExeSourcePath != "" {
			source = ExeSourcePath
		} else {
			source = SourcePath
		}

		if ExeDestinationPath != "" {
			destination = ExeDestinationPath
		} else {
			destination = DestinationPath
		}

		if ExeConfigFilePath != "" {
			config = ExeConfigFilePath
		} else {
			config = ConfigFilePath
		}

		if err := userFunctions.ExecuteCopy(source, destination, config); err != nil {
			return err
		}

		return nil
	},
}

var ExeSourcePath string
var ExeDestinationPath string
var ExeConfigFilePath string

func init() {
	rootCmd.AddCommand(executeWithOverrideCmd)

	executeWithOverrideCmd.Flags().StringVarP(&ExeSourcePath, "source", "s", "", "Source directory to copy from")
	if err := executeWithOverrideCmd.MarkFlagDirname("source"); err != nil {
		log.Panic(err)
	}

	executeWithOverrideCmd.Flags().StringVarP(&ExeDestinationPath, "destination", "d", "", "Destination directory to copy to")
	if err := executeWithOverrideCmd.MarkFlagDirname("destination"); err != nil {
		log.Panic(err)
	}

	executeWithOverrideCmd.Flags().StringVarP(&ExeConfigFilePath, "config", "c", "", "Config to use for the copy operation")
	if err := executeWithOverrideCmd.MarkFlagFilename("config"); err != nil {
		log.Panic(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// executeWithOverrideCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// executeWithOverrideCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
