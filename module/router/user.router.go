package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenants.modules.user/module/controller"
)

func UserRouter(r *gin.RouterGroup, c controller.UserController) {
	r.GET("", c.GetAll)
	r.POST("", c.Create)
	r.GET(":id", c.GetById)
	r.PATCH(":id", c.UpdateById)
	r.DELETE(":id", c.DeleteById)
}
