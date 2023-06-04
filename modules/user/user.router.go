package user

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func UserRouter(r *gin.RouterGroup, c UserController) {

	fmt.Println("router registered")
	r.GET("", c.GetAll)
	r.POST("", c.Save)
	r.GET(":id", c.GetById)
	r.PATCH(":id", c.UpdateById)
	r.DELETE(":id", c.DeleteById)
}
