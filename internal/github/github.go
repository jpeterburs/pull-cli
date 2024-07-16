package github

import "fmt"

type Repository struct {
	Owner string
	Name  string
}

func (r *Repository) String() string {
	return fmt.Sprintf("%v/%v", r.Owner, r.Name)
}

type PullRequest struct {
	Number     uint32
	Repository *Repository
}

func (pr *PullRequest) String() string {
	return fmt.Sprintf("%v#%v", pr.Repository.String(), pr.Number)
}
