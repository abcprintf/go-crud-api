package router

import (
	"net/http"

	db "abcpirntf/go-crud-api/db"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users", GetUsers)
	r.GET("/users/:id", GetUser)
	r.POST("/users", PostUser)
	r.PUT("/users/:id", PutUser)
	r.DELETE("/users/:id", DeleteUser)
	return r
}

func GetUsers(ctx *gin.Context) {
	res, err := db.GetUsers()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"users": res,
	})
}

func GetUser(ctx *gin.Context) {
	id := ctx.Param("id")
	res, err := db.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}

func PostUser(ctx *gin.Context) {
	var user db.User
	err := ctx.Bind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	res, err := db.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"user": res,
	})
}

func PutUser(ctx *gin.Context) {
	var updatedUser db.User
	err := ctx.Bind(&updatedUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	id := ctx.Param("id")
	dbUser, err := db.GetUser(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	dbUser.Username = updatedUser.Username
	dbUser.Password = updatedUser.Password

	res, err := db.UpdateUser(dbUser)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"user": res,
	})
}

func DeleteUser(ctx *gin.Context) {
	id := ctx.Param("id")
	err := db.DeleteUser(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"error": err.Error(),
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"message": "user deleted successfully",
	})
}
