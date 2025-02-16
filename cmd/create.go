package cmd

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/google/go-github/v69/github"
	"github.com/jpeterburs/pull_request-cli/internal/gitutil"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a pull request",
	Run: func(cmd *cobra.Command, args []string) {
		draft, err := strconv.ParseBool(cmd.Flag("draft").Value.String())
		if err != nil {
			panic(err)
		}

		newPR := &github.NewPullRequest{
			Title: github.Ptr(cmd.Flag("title").Value.String()),
			Head:  github.Ptr(cmd.Flag("head").Value.String()),
			Base:  github.Ptr(cmd.Flag("base").Value.String()),
			Body:  github.Ptr(cmd.Flag("body").Value.String()),
			Draft: github.Ptr(draft),
		}

		token := viper.GetString("github_token") // TODO: USe GH_TOKEN and GITHUB_TOKEN env optionally
		ctx := context.Background()
		client := github.NewClientWithEnvProxy().WithAuthToken(token)

		pr, _, err := client.PullRequests.Create(ctx, "jpeterburs", "test", newPR)
		if err != nil {
			panic(err)
		}

		fmt.Printf("PR created: %s\n", pr.GetHTMLURL())
	},
}

func init() {
	base, err := gitutil.GitMainBranch()
	if errors.Is(err, gitutil.ErrNoMainBranch) {
		fmt.Println(err)
	} else if err != nil {
		panic(err)
	}

	repo, err := gitutil.FindRepo()
	if err != nil {
		panic(err)
	}

	ref, err := repo.Head()
	if err != nil {
		panic(err)
	}
	head := strings.Split(string(ref.Name()), "/")[2]

	createCmd.Flags().String("base", base, "The name of the branch you want the changes pulled into. This should be an existing branch on the current repository. You cannot submit a pull request to one repository that requests a merge to a base of another repository.")
	createCmd.Flags().StringP("body", "m", "", "The contents of the pull request.")
	createCmd.Flags().BoolP("draft", "d", false, "Indicates whether the pull request is a draft.")
	createCmd.Flags().String("head", head, "The name of the branch where your changes are implemented. For cross-repository pull requests in the same network, namespace head with a user like this: username:branch")
	// createCmd.Flags().StringArrayP("lables", "l", []string{}, "apply labels")
	// createCmd.Flags().StringArrayP("reviewers", "r", []string{}, "request review")
	createCmd.Flags().StringP("title", "t", "", "The title of the new pull request.")
	createCmd.MarkFlagRequired("title")
	rootCmd.AddCommand(createCmd)
}
