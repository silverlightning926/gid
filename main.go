package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
	builtBy = "unknown"
)

// Helper Functions

func getCurrentConfigPath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return filepath.Join(homeDir, ".gitconfig"), nil
}

func getAvailableConfigPaths() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return []string{}, err
	}

	profilesDir := filepath.Join(homeDir, ".config", "gid", "profiles")
	entries, err := os.ReadDir(profilesDir)
	if err != nil {
		return []string{}, err
	}

	var paths []string
	for _, entry := range entries {
		fullPath := filepath.Join(profilesDir, entry.Name())
		paths = append(paths, fullPath)
	}

	return paths, nil

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

func fileSHA256(filepath string) (string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	hash := sha256.New()

	if _, err := io.Copy(hash, file); err != nil {
		return "", fmt.Errorf("failed to read file: %w", err)
	}

	hashBytes := hash.Sum(nil)
	return fmt.Sprintf("%x", hashBytes), nil
}

// Cobra

var rootCmd = &cobra.Command{
	Use:     "git-id",
	Aliases: []string{"gid"},
	Version: version,
	Short:   "A CLI Tool To Manage Git Profiles",
	Long:    `git-id is a command-line tool for managing multiple Git profiles. It allows you to easily switch between different Git configurations for different projects or contexts.`,
}

var statusCmd = &cobra.Command{
	Use:     "status",
	Aliases: []string{"now"},
	Short:   "Find the currently set config",
	Long:    "Find which of the configs has been currently set to active",
	Run: func(cmd *cobra.Command, args []string) {
		currentConfigPath, err := getCurrentConfigPath()
		if err != nil {
			fmt.Println(err)
			return
		}

		currentConfigPathHash, err := fileSHA256(currentConfigPath)
		if err != nil {
			fmt.Println(err)
			return
		}

		availableConfigPaths, err := getAvailableConfigPaths()
		if err != nil {
			fmt.Println(err)
			return
		}

		hiBlue := color.New(color.FgHiBlue).SprintFunc()
		red := color.New(color.FgHiRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		grey := color.New(color.FgHiBlack).SprintFunc()

		for _, availableConfigPath := range availableConfigPaths {

			availableConfigPathHash, err := fileSHA256(availableConfigPath)
			if err != nil {
				fmt.Println(err)
				return
			}

			if currentConfigPathHash == availableConfigPathHash {
				fileName := filepath.Base(availableConfigPath)
				profileName := strings.TrimSuffix(fileName, ".gitconfig")

				fmt.Printf("%s %s %s\n", hiBlue("Currently Using Profile"), green(profileName), grey("("+fileName+")"))
				return
			}
		}

		fmt.Printf("%s\n", red("The Current Profile Is Unknown"))
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List Git Profiles",
	Long:  "List all available Git profiles stored in your configuration directory.",
	Run: func(cmd *cobra.Command, args []string) {
		paths, err := getAvailableConfigPaths()
		if err != nil {
			fmt.Println(err)
			return
		}

		hiBlue := color.New(color.FgHiBlue).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		grey := color.New(color.FgHiBlack).SprintFunc()

		fmt.Printf("%s %s\n",
			hiBlue("Profiles"),
			(grey("(" + strconv.Itoa(len(paths)) + " Found" + ")")),
		)

		for _, path := range paths {
			fileName := filepath.Base(path)
			fmt.Printf("  • %s %s\n", green(strings.TrimSuffix(fileName, ".gitconfig")), grey("("+fileName+")"))
		}
	},
}

var useCmd = &cobra.Command{
	Use:     "use",
	Short:   "Use Selected Git Profile",
	Long:    "Switch to the specified Git profile by copying it to ~/.gitconfig.",
	Aliases: []string{"switch"},
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		currentConfigPath, err := getCurrentConfigPath()
		if err != nil {
			fmt.Println(err)
		}

		availableConfigPaths, err := getAvailableConfigPaths()
		if err != nil {
			fmt.Println(err)
		}

		hiBlue := color.New(color.FgHiBlue).SprintFunc()
		red := color.New(color.FgHiRed).SprintFunc()
		green := color.New(color.FgGreen).SprintFunc()
		grey := color.New(color.FgHiBlack).SprintFunc()

		if len(availableConfigPaths) <= 0 {
			fmt.Printf("%s\n", red("No Profiles Found"))
			return
		}

		for _, path := range availableConfigPaths {
			fileName := filepath.Base(path)
			profileName := strings.TrimSuffix(fileName, ".gitconfig")

			if profileName == args[0] {
				err = copyFile(path, currentConfigPath)
				if err != nil {
					fmt.Println(err)
					return
				}

				fmt.Printf("%s %s %s\n", hiBlue("Switched To Using Profile"), green(profileName), grey("("+fileName+")"))
				return
			}
		}

		fmt.Printf("%s %s\n", red("Profile Not Found:"), args[0])

	},
}

func init() {
	rootCmd.AddCommand(statusCmd)
	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(useCmd)
}

func main() {
	rootCmd.Execute()
}
