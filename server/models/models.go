package models

type Test interface {
	HandleGitHubHook(event any) error
}
