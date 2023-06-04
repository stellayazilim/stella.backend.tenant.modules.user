package modules

import (
	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenants/modules/database"
	"github.com/stellayazilim/stella.backend.tenants/modules/user"
)

func MainModule(e *gin.Engine) {
	// stack modules here
	database.DatabaseModule()
	user.UserModule(e.Group("/users"))

}
