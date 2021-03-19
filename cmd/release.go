package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/qianhum/github-latest/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// releaseCmd represents the release command
var releaseCmd = &cobra.Command{
	Use:   "release [repo]",
	Short: "fetch latest release",
	Long:  `Fetch latest release information from Github repository.`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a repo argument, like \"spf13/cobra\"")
		}
		if internal.IsValidRepo(args[0]) {
			return nil
		}
		return fmt.Errorf("invalid repo specified: %s. valid repo should look like \"spf13/cobra\"", args[0])
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		githubToken := viper.GetString("GITHUB_TOKEN")
		if githubToken == "" {
			return errors.New("can not find Github token. Did you set it up in $HOME/.github-latest.yaml?")
		}
		repo := strings.Split(args[0], "/")
		return internal.CheckRelease(githubToken, repo[0], repo[1])
	},
}

func init() {
	rootCmd.AddCommand(releaseCmd)
}
