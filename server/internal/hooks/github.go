package hooks

import "onepush-server/internal/store"

type GithubHooks struct {
	store *store.Store
}

func NewGithubHooks(store *store.Store) *GithubHooks {
	return &GithubHooks{
		store: store,
	}
}
