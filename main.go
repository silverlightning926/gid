package main

import (
	"fmt"
	"io"
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

func copyFile(src, dst string) error {
	sourceFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer sourceFile.Close()

	destFile, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer destFile.Close()

	_, err = io.Copy(destFile, sourceFile)
	return err
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

var useCmd = &cobra.Command{
	Use:     "use",
	Short:   "Use Selected Git Profile",
	Aliases: []string{"switch"},
	Args:    cobra.ExactArgs(1),
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

		red := color.New(color.FgHiRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		grey := color.New(color.FgHiBlack).SprintFunc()

		if len(entries) <= 0 {
			fmt.Printf("%s\n", red("No Profiles Found"))
			return
		}

		for _, entry := range entries {
			filename := entry.Name()
			profileName := strings.TrimSuffix(filename, ".gitconfig")

			if profileName == args[0] {
				homeDir, err := os.UserHomeDir()
				if err != nil {
					fmt.Println(err)
				}

				srcPath := filepath.Join(configPath, "profiles", filename)
				dstPath := filepath.Join(homeDir, ".gitconfig")

				err = copyFile(srcPath, dstPath)
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Printf("Using Profile: %s %s\n", green(profileName), grey("("+filename+")"))
				return
			}
		}

		fmt.Printf("%s %s\n", red("Profile Not Found:"), args[0])

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(useCmd)
}

func main() {
	rootCmd.Execute()
}
