package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v63/github"
)

func main() {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		log.Fatal("Unauthorized: No token present")
	}

	ctx := context.Background()
	client := github.NewClientWithEnvProxy().WithAuthToken(token)

	pr, _, err := client.PullRequests.Get(ctx, "jpeterburs", "pull-cli", 1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(pr.GetTitle())
}
