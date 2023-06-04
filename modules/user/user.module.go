package user

import "github.com/gin-gonic/gin"

func UserModule(r *gin.RouterGroup) {

	// stack
	UserRouter(r, Controller())

}
