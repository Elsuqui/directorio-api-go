package router

import (
	"github.com/gin-gonic/gin"
)

// Define struct to declare all groups of routes divided by modules
type Routes struct {
	videosRoutes VideoRoutes
}

// Function to initialize all api routes
func (r *Routes) Boot(server *gin.Engine) {
	r.videosRoutes.InitializeRoutes(server)
}
