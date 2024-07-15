package main

import (
	"fmt"

	pull_request "github.com/jpeterburs/pull-cli/internal/pull_request"
)

func main() {
	var pr pull_request.PullRequest = pull_request.New("jpeterburs", "pull-cli", 1)
	fmt.Println(pr.String())
}
