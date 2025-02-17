package gitutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRemoteString(t *testing.T) {
	remote := Remote{
		Owner: "jpeterburs",
		Repo:  "pull_request-cli",
	}

	assert.Equal(t, remote.String(), "jpeterburs/pull_request-cli")
}
