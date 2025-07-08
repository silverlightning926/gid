package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func getConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".config", "gid"), nil
}

func getConfigFiles(dir string) ([]os.DirEntry, error) {
	entries, err := os.ReadDir(filepath.Join(dir, "profiles"))
	if err != nil {
		return nil, err
	}

	return entries, nil
}

var rootCmd = &cobra.Command{
	Use:   "git-id",
	Short: "A CLI Tool To Manage Git Profiles",
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Git Profiles",
	Run: func(cmd *cobra.Command, args []string) {
		configPath, err := getConfigPath()
		if err != nil {
			fmt.Println(err)
			return
		}

		entries, err := getConfigFiles(configPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		hiBlue := color.New(color.FgHiBlue).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		grey := color.New(color.FgHiBlack).SprintFunc()

		fmt.Printf("%s %s\n",
			hiBlue("Profiles"),
			(grey("(" + strconv.Itoa(len(entries)) + " Found" + ")")),
		)

		for _, entry := range entries {
			filename := entry.Name()
			profileName := strings.TrimSuffix(filename, ".gitconfig")

			fmt.Printf("  • %s %s\n", green(profileName), grey("("+filename+")"))
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}

func main() {
	rootCmd.Execute()
}
