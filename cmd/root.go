/*
	Author Muhamamd Ilham <hi@muhammadilham.xyz>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/piigyy/gostrap/entity"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile string
	Config  entity.Configuration
)

const (
	CURRENT_VERSION = "v0.0.3-alpha"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gostrap",
	Short: getShortDesc(),
	Long:  getLongDesc(),
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gostrap.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".gostrap" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".gostrap")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
		if err := viper.Unmarshal(&Config); err != nil {
			fmt.Fprintln(os.Stderr, "failed marshalling to entity.Configuration:", err)
		}
	}
}

func getShortDesc() string {
	return fmt.Sprintf("GoStrap (%s) CLI Is Here To Help You Bootstraping Your Go Project", CURRENT_VERSION)
}

func getLongDesc() string {
	return fmt.Sprintf(`GoStrap (%s) is a CLI
that help faster your development time by bootstrapping
your go project so you don't really need to start from
"scratch."
	`, CURRENT_VERSION)
}
