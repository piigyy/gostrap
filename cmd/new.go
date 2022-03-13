/*
	Author Muhamamd Ilham <hi@muhammadilham.xyz>
*/
package cmd

import (
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

const (
	RENAME_SCRIPT = "https://raw.githubusercontent.com/piigyy/gostrap/main/script/rename.sh"
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [project-name]",
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
		modulePlaceholder, _ := cmd.Flags().GetString("placeholder")

		if templateSource == "" {
			templateSource = Config.Template
		}

		if modulePlaceholder == "" {
			modulePlaceholder = Config.GoModulePlaceholder
		}

		projectName := args[0]
		wd, _ := os.Getwd()
		dir := fmt.Sprintf("%s/%s", wd, projectName)

		fmt.Printf("project name: %s\n", projectName)
		fmt.Printf("tempate: %s\n", templateSource)
		fmt.Printf("Go Mod Name: %s\n", gomoduleName)

		if err := clearDirectory(dir); err != nil {
			fmt.Printf("error clearing directory: %v\n", err)
			return
		}

		fmt.Printf("clone template directory: %s\n", cloneTemplateDirectory(templateSource, dir))
		fmt.Printf("re init git:\n%s\n", reInitGit(dir))

		if modulePlaceholder == "" {
			fmt.Println("module placeholder is empty")
			fmt.Println("finished init you project!")
			return
		}

		if err := replaceModuleName(dir, gomoduleName, modulePlaceholder); err != nil {
			fmt.Printf("error replacing Go module name: %v\n", err)
			return
		}

		fmt.Println("finished init you project!")
	},
}

func init() {
	rootCmd.AddCommand(newCmd)
	newCmd.Flags().StringP("template", "t", "", "chose project structure template")
	newCmd.Flags().StringP("module", "m", "", "your go mod name")
	newCmd.Flags().StringP("placeholder", "p", "", "current go module name in your project structure")
	newCmd.Flags().BoolP("default", "d", false, "set template flag as default")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func reInitGit(dir string) string {
	gitDir := fmt.Sprintf("%s/.git", dir)
	os.RemoveAll(gitDir)
	gitInitCmd := exec.Command("git", "init")
	gitInitCmd.Dir = dir
	output, _ := gitInitCmd.CombinedOutput()
	return string(output)
}

func cloneTemplateDirectory(template, dir string) string {
	fmt.Printf("cloning template (%s) into %s\n", template, dir)
	gitCloneCmd := exec.Command("git", "clone", "--single-branch", template, dir)
	outputRaws, _ := gitCloneCmd.CombinedOutput()
	return string(outputRaws)
}

func clearDirectory(dir string) error {
	return os.RemoveAll(dir)
}

func replaceModuleName(dir, moduleName, placeholder string) error {
	log.Println("masuk sinii ga?")
	renameFilename := fmt.Sprintf("%s/rename.sh", dir)
	if err := downloadWGETScript(dir, renameFilename); err != nil {
		return err
	}

	if err := os.Chmod(renameFilename, 0777); err != nil {
		return err
	}

	rename := exec.Command("./rename.sh", placeholder, moduleName)
	rename.Dir = dir
	output, err := rename.Output()
	if err != nil {
		return nil
	}

	fmt.Printf("rename module: %s\n", string(output))
	os.RemoveAll(renameFilename)
	return nil
}

func downloadWGETScript(dir, filename string) error {
	wget := exec.Command("wget", "-O", filename, RENAME_SCRIPT)
	output, err := wget.CombinedOutput()
	if err != nil {
		return err
	}

	fmt.Printf("wget: %s\n", string(output))
	return nil
}
