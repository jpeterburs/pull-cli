package gitutil

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"slices"

	"github.com/go-git/go-git/v5"
	"github.com/go-git/go-git/v5/plumbing"
)

func FindRepo() (*git.Repository, error) {
	dir, err := os.Getwd()
	if err != nil {
		return nil, fmt.Errorf("error getting current directory: %w", err)
	}

	for {
		repo, err := git.PlainOpen(dir)
		if err == nil {
			return repo, nil
		}

		if err == git.ErrRepositoryNotExists {
			parent := filepath.Dir(dir)
			if parent == dir {
				return nil, fmt.Errorf("fatal: not a git repository (or any of the parent directories): .git")
			}

			dir = parent
			continue
		}

		return nil, err
	}
}

var mainBranches = []string{"main", "trunk", "mainline", "default", "stable", "master"}

var ErrNoMainBranch = errors.New("no main branch found, defaulting to master")

func GitMainBranch() (string, error) {
	repo, err := FindRepo()
	if err != nil {
		return "", err
	}

	refs, err := repo.References()
	if err != nil {
		return "", fmt.Errorf("failed to list references: %w", err)
	}

	var foundBranches []string

	err = refs.ForEach(func(ref *plumbing.Reference) error {
		name := ref.Name().String()
		for _, branch := range mainBranches {
			if slices.Contains([]string{
				"ref/heads/" + branch,
				"refs/remotes/origin/" + branch,
				"refs/remotes/upstream/" + branch,
			}, name) {
				foundBranches = append(foundBranches, branch)
			}
		}

		return nil
	})

	if err != nil {
		return "", fmt.Errorf("error iterating refs: %w", err)
	}

	if len(foundBranches) > 0 {
		return foundBranches[0], nil
	}

	return "master", ErrNoMainBranch
}
