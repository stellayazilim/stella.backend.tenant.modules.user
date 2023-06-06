package module

import (
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenants.modules.user/database"
	"github.com/stellayazilim/stella.backend.tenants.modules.user/module/controller"
	"github.com/stellayazilim/stella.backend.tenants.modules.user/module/router"
)

func MainModule(e *gin.Engine) {

	// stack modules here
	database.InitalizeDatabase()

	// stack router groups here
	router.UserRouter(e.Group("/users"), controller.CreateController())

}
