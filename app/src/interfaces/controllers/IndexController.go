
package controllers

import (
	// "fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)


type IndexController struct {

}


func NewIndexController() *IndexController {
	return &IndexController{}
}


func (controller *IndexController) Get(c *gin.Context) {

	test := c.Query("test")
	tests := c.QueryArray("tests")
	// tests := c.QueryMap("tests")

	// fmt.Print("\n=========\n", tests, "\n=========\n")

	c.JSON(http.StatusOK, gin.H{
		"statusCode": http.StatusOK,
		"method": "GET",
		"contentType": c.GetHeader("Content-Type"),
		"test": test,
		"tests": tests,
	})
}


func (controller *IndexController) Post(c *gin.Context) {

	test := c.PostForm("test")
	tests := c.PostFormArray("tests")

	c.JSON(http.StatusCreated, gin.H{
		"statusCode": http.StatusCreated,
		// "statusCode": http.StatusBadGateway,
		"method": "POST",
		"contentType": c.GetHeader("Content-Type"),
		"test": test,
		"tests": tests,
	})
}


func (controller *IndexController) Put(c *gin.Context) {

	test := c.PostForm("test")
	tests := c.PostFormArray("tests")

	c.JSON(200, gin.H{
		"statusCode": http.StatusOK,
		"method": "PUT",
		"contentType": c.GetHeader("Content-Type"),
		"test": test,
		"tests": tests,
	})
}


func (controller *IndexController) Patch(c *gin.Context) {

	test := c.PostForm("test")
	tests := c.PostFormArray("tests")

	c.JSON(200, gin.H{
		"statusCode": http.StatusOK,
		"method": "PATCH",
		"contentType": c.GetHeader("Content-Type"),
		"test": test,
		"tests": tests,
	})
}


func (controller *IndexController) Delete(c *gin.Context) {

	test := c.PostForm("test")
	tests := c.PostFormArray("tests")

	c.JSON(200, gin.H{
		"statusCode": http.StatusOK,
		"method": "DELETE",
		"contentType": c.GetHeader("Content-Type"),
		"test": test,
		"tests": tests,
	})
}

