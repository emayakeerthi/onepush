package handlers

import (
	"onepush-server/internal/hooks"
	"onepush-server/internal/store"

	"github.com/gin-gonic/gin"
)

type GitHubHookHandler struct {
	store  *store.Store
	github *hooks.GithubHooks
}

func NewGitHubHookHandler(store *store.Store, github *hooks.GithubHooks) *GitHubHookHandler {
	return &GitHubHookHandler{
		store:  store,
		github: github,
	}
}

func (h *GitHubHookHandler) HandleGitHubHook(c *gin.Context) {

}
