/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

const (
	DEFAULT_SOURCE = "https://github.com/golang-standards/project-layout"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new",
	Short: "create new golang template project",
	Long:  `create new golang template project`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("please provide your project name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		templateSource, _ := cmd.Flags().GetString("template")
		gomoduleName, _ := cmd.Flags().GetString("module")
		dir, _ := os.Getwd()
		projectName := args[0]
		directory := fmt.Sprintf("%s/%s", dir, projectName)

		fmt.Printf("project name: %s\n", projectName)
		fmt.Printf("tempate: %s\n", templateSource)
		fmt.Printf("Go Mod Name: %s\n", gomoduleName)
		fmt.Printf("Creating project %s(%s) to directory %s\n", projectName, gomoduleName, directory)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("template", "t", "https://github.com/golang-standards/project-layout", "chose project structire template")
	newCmd.Flags().StringP("module", "m", "gostrap", "your go mod name")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
