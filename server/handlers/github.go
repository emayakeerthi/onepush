package handlers

import (
	"onepush-server/internal/hooks"
	storePkg "onepush-server/internal/store"

	"github.com/gin-gonic/gin"
)

type GitHubHookHandler struct {
	store  *storePkg.Store
	github *hooks.GithubHooks
}

func NewGitHubHookHandler(store *storePkg.Store, github *hooks.GithubHooks) *GitHubHookHandler {
	return &GitHubHookHandler{
		store:  store,
		github: github,
	}
}

func (h *GitHubHookHandler) HandleGitHubHook(c *gin.Context) {
}
