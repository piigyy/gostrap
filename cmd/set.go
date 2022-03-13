/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/spf13/cobra"
)

// setCmd represents the set command
var setCmd = &cobra.Command{
	Use:   "set [config] [value]",
	Short: "Set GoStrap Configuration",
	Long: `Set GoStrap Configuration like template and placeholder
Current available config:
- template
- gomoduleplaceholder`,
	Example: "cobra-clit set template https://github.com/golang-standards/project-layout",
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) > 2 {
			return fmt.Errorf("expected 2 arguments but found: %d", len(args))
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := setConfig(args[0], args[1]); err != nil {
			fmt.Printf("error set config: %v\n", err)
			return
		}

		fmt.Println("set config success!")
	},
}

func init() {
	rootCmd.AddCommand(setCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// setCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// setCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func setConfig(field, value string) error {
	fieldMapping := map[string]string{
		"template":            "Template",
		"gomoduleplaceholder": "GoModulePlaceholder",
	}

	structField := reflect.ValueOf(&Config).Elem().FieldByName(fieldMapping[strings.ToLower(field)])
	structField.SetString(value)
	return Config.Update()
}
