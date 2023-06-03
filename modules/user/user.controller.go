package user

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenants/entities"
)

// interface
type UserController interface {
	GetAll(*gin.Context)
	Save(*gin.Context)
	GetById(*gin.Context)
	UpdateById(*gin.Context)
	DeleteById(*gin.Context)
}

// controller
type userController struct {
	userService UserService
}

// constructor
func Controller() UserController {
	return &userController{
		userService: Service(),
	}
}

// get all users
func (c *userController) GetAll(ctx *gin.Context) {
	u := c.userService.GetAll()
	ctx.JSON(200, u)
}

// save user
func (c *userController) Save(ctx *gin.Context) {

	user := entities.User{}
	ctx.BindJSON(&user)

	c.userService.Save(user)

	ctx.AbortWithStatus(200)
}

func (c *userController) GetById(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)

	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	user, err := c.userService.GetById(id)

	if err != nil {
		ctx.AbortWithError(404, err)
		return
	}
	ctx.JSON(200, user)
}

func (c *userController) UpdateById(ctx *gin.Context) {

	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	data := entities.User{}
	ctx.BindJSON(&data)

	err = c.userService.UpdateById(id, data)

	if err != nil {
		ctx.AbortWithStatus(404)
		return
	}
	ctx.AbortWithStatus(200)
}

func (c *userController) DeleteById(ctx *gin.Context) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.AbortWithError(400, err)
		return
	}

	err = c.userService.DeleteById(id)

	if err != nil {
		ctx.AbortWithError(404, err)
		return
	}

	ctx.AbortWithStatus(200)
}
