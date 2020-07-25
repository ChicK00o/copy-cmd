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
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

var SetSourcePath string
var SetDestinationPath string
var SetConfigFilePath string

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("set called")
		valueSet := false
		if SetSourcePath != "" {
			viper.Set("SourcePath", SetSourcePath)
			valueSet = true
			SourcePath = SetSourcePath
		}
		if SetDestinationPath != "" {
			viper.Set("DestinationPath", SetDestinationPath)
			valueSet = true
			DestinationPath = SetDestinationPath
		}
		if SetConfigFilePath != "" {
			viper.Set("ConfigFilePath", SetConfigFilePath)
			valueSet = true
			ConfigFilePath = SetConfigFilePath
		}
		if !valueSet {
			return errors.New("Set at least 1 value in this command")
		}
		if err := viper.WriteConfig(); err != nil {
			return err
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	setCmd.Flags().StringVarP(&SetSourcePath, "source", "s", "", "Source directory to copy from")
	if err := setCmd.MarkFlagDirname("source"); err != nil {
		log.Panic(err)
	}

	setCmd.Flags().StringVarP(&SetDestinationPath, "destination", "d", "", "Destination directory to copy to")
	if err := setCmd.MarkFlagDirname("destination"); err != nil {
		log.Panic(err)
	}

	setCmd.Flags().StringVarP(&SetConfigFilePath, "config", "c", "", "Config to use for the copy operation")
	if err := setCmd.MarkFlagFilename("config"); err != nil {
		log.Panic(err)
	}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
