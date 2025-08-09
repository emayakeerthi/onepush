package main

import (
	"onepush-server/config"
	"onepush-server/handlers"
	"onepush-server/internal/hooks"
	"onepush-server/internal/store"
	"onepush-server/utils"
	"strconv"

	"github.com/gin-gonic/gin"
)

func HandleSystemHealth(c *gin.Context) {
	healthStatus := utils.CheckSystemHealth()
	c.JSON(200, healthStatus)
}

func main() {
	router := gin.Default()

	config := config.NewConfig()
	store := store.NewStore()
	githubHooks := hooks.NewGithubHooks(store)

	router.GET("/health", HandleSystemHealth)

	githubHandler := handlers.NewGitHubHookHandler(store, githubHooks)
	router.POST("/github/webhook", githubHandler.HandleGitHubHook)

	router.Run(":" + strconv.Itoa(config.ServerPort))
}
