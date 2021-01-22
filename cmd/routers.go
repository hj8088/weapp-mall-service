package cmd

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path/filepath"
)

func configureServerHandler() (http.Handler, error) {
	gin.SetMode(gin.DebugMode)
	router := gin.New()

	router.Static("/resource", filepath.Join(filepath.Dir(os.Args[0]), "resource"))

	registerAPIRouter(router)
	return router, nil
}
