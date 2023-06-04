package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/stellayazilim/stella.backend.tenants/entities"
)

// interface
type UserController interface {
	GetAll(*gin.Context)
	Create(*gin.Context)
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
func (c *userController) Create(ctx *gin.Context) {

	user := entities.User{}
	err := ctx.ShouldBindJSON(&user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	c.userService.Create(user)

	ctx.AbortWithStatus(200)
}

// get single user by id
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

// update user by id
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

// delete user by id
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
