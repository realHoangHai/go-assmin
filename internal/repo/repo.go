package repo

import "github.com/google/wire"

var ProviderRepoSet = wire.NewSet(NewRepo)

type Repo struct {
}

func NewRepo() IRepo {
	return &Repo{}
}
