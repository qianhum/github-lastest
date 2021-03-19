package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/qianhum/github-latest/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// tagCmd represents the tag command
var tagCmd = &cobra.Command{
	Use:   "tag [repo]",
	Short: "fetch latest tag",
	Long:  `Fetch latest tag information from Github repository.`,
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
		return internal.CheckTag(githubToken, repo[0], repo[1])
	},
}

func init() {
	rootCmd.AddCommand(tagCmd)
}
