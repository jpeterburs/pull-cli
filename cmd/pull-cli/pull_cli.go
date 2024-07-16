package main

import (
	"fmt"

	"github.com/jpeterburs/pull-cli/internal/github"
)

func main() {
	pr := github.PullRequest{
		Number: 1,
		Repository: &github.Repository{
			Owner: "jpeterburs",
			Name:  "pull-cli",
		},
	}
	fmt.Println(pr.String())
}
