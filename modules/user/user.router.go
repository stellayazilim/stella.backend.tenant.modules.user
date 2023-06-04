package user

import "github.com/gin-gonic/gin"

func UserRouter(r *gin.RouterGroup, c UserController) {
	r.GET("", c.GetAll)
	r.POST("", c.Save)
	r.GET(":id", c.GetById)
	r.PATCH(":id", c.UpdateById)
	r.DELETE(":id", c.DeleteById)
}
