package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenants.modules.user/entity"
	"github.com/stellayazilim/stella.backend.tenants.modules.user/helpers"
	"github.com/stellayazilim/stella.backend.tenants.modules.user/serializers"
)

// create user
func (c *userController) Create(ctx *gin.Context) {

	body := helpers.ParseBody[serializers.UserCreateRequest](ctx)

	if _, exist := c.userService.GetByTenantIdAndEmail(entity.User{TenantId: "sdfsduhfksefj", Email: body.Email}); exist {
		ctx.JSON(http.StatusAccepted, gin.H{"error": "User with email already exist"})
		return
	}

	err := c.userService.Create(body.SerializeUserCreate())

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.AbortWithStatus(200)

}
