package github

import (
	"testing"
)

func TestRepository_String(t *testing.T) {
	owner := "jpeterburs"
	name := "pull-cli"
	repository := Repository{
		Owner: owner,
		Name:  name,
	}

	expected := "jpeterburs/pull-cli"
	if repository.String() != expected {
		t.Errorf("expected %v, got %v", expected, repository.String())
	}
}

func TestPullRequest_String(t *testing.T) {
	number := 1
	owner := "jpeterburs"
	name := "pull-cli"
	pr := PullRequest{
		Number: uint32(number),
		Repository: &Repository{
			Owner: owner,
			Name:  name,
		},
	}

	expected := "jpeterburs/pull-cli#1"
	if pr.String() != expected {
		t.Errorf("expected %v, got %v", expected, pr.String())
	}
}
