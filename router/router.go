package router

import "github.com/gin-gonic/gin"

func StartRouter(port string) {
	r := gin.Default()

	// stack routers here
	userRouter(r)

	r.Run(port)
}
