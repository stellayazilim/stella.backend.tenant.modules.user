package user

import "github.com/gin-gonic/gin"

func InitalizeUserModule(r *gin.RouterGroup) {

	// stack
	UserRouter(r, Controller())

}
