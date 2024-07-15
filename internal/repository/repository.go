package repository

type Repository struct {
	owner string
	name  string
}

func New(owner string, name string) Repository {
	return Repository{
		owner: owner,
		name:  name,
	}
}

func (r Repository) String() string {
	return r.owner + "/" + r.name
}
