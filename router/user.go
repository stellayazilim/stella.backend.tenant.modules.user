package router

import (
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenants/modules/user"
)

func userRouter(r *gin.Engine) {

	c := user.Controller()
	// stack endpoints here
	r.GET("/users", c.GetAll)
	r.POST("/users", c.Save)
	r.GET("/users/:id", c.GetById)
	r.PATCH("/users/:id", c.UpdateById)
	r.DELETE("/users/:id", c.DeleteById)
}
