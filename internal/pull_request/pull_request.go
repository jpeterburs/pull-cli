package pull_request

import (
	"strconv"

	repository "github.com/jpeterburs/pull-cli/internal/repository"
)

type PullRequest struct {
	number   uint32
	title    string
	head_ref string
	base_ref string
	repo     repository.Repository
}

func New(owner string, name string, number uint32) PullRequest {
	return PullRequest{
		number:   number,
		title:    "",
		head_ref: "",
		base_ref: "",
		repo:     repository.New(owner, name),
	}
}

func (pr *PullRequest) String() string {
	return pr.repo.String() + "#" + strconv.FormatUint(uint64(pr.number), 10)
}
