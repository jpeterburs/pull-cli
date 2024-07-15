package pull_request

import (
	"testing"

	repository "github.com/jpeterburs/pull-cli/internal/repository"
)

func TestNew(t *testing.T) {
	owner := "jpeterburs"
	name := "test"
	number := uint32(1)
	pr := New(owner, name, number)

	if pr.number != number {
		t.Errorf("expected number %v, got %v", number, pr.number)
	}

	if pr.title != "" {
		t.Errorf("expected title %v, got %v", "", pr.title)
	}

	if pr.head_ref != "" {
		t.Errorf("expected head_ref %v, got %v", "", pr.head_ref)
	}

	if pr.base_ref != "" {
		t.Errorf("expected base_ref %v, got %v", "", pr.base_ref)
	}

	if pr.repo.String() != repository.New(owner, name).String() {
		t.Errorf("expected repo %v, got %v", repository.New(owner, name).String(), pr.repo.String())
	}
}

func TestPullRequest_String(t *testing.T) {
	owner := "jpeterburs"
	name := "test"
	number := uint32(1)
	pr := New(owner, name, number)

	expected := "jpeterburs/test#1"
	if pr.String() != expected {
		t.Errorf("expected %v, got %v", expected, pr.String())
	}
}
