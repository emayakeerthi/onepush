package handlers

import (
	"fmt"
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
	var data any
	if err := c.ShouldBind(&data); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	fmt.Printf("Received GitHub hook data: %+v\n", data)
}
