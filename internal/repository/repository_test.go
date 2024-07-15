package repository

import (
	"testing"
)

func TestNew(t *testing.T) {
	owner := "jpeterburs"
	name := "test"
	repo := New(owner, name)

	if repo.owner != owner {
		t.Errorf("expected owner %v, got %v", owner, repo.owner)
	}

	if repo.name != name {
		t.Errorf("expected name %v, got %v", name, repo.name)
	}
}

func TestRepository_String(t *testing.T) {
	owner := "jpeterburs"
	name := "test"
	repo := New(owner, name)

	expected := "jpeterburs/test"
	if repo.String() != expected {
		t.Errorf("expected %v, got %v", expected, repo.String())
	}
}
