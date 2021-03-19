package cmd

import (
	"errors"
	"fmt"
	"strings"

	"github.com/qianhum/github-latest/internal"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var branch string

// commitCmd represents the commit command
var commitCmd = &cobra.Command{
	Use:   "commit [repo]",
	Short: "fetch latest commit",
	Long: `Fetch latest commit information from Github repository.

Use -b to specify the branch. If no branch is provided, the default branch will be used.`,
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
		if branch == "" {
			return internal.CheckDefaultCommit(githubToken, repo[0], repo[1])
		} else {
			return internal.CheckCommit(githubToken, repo[0], repo[1], branch)
		}
	},
}

func init() {
	rootCmd.AddCommand(commitCmd)

	commitCmd.Flags().StringVarP(&branch, "branch", "b", "", "branch to fetch")
}
