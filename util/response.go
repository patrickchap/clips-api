package util

import "github.com/gin-gonic/gin"

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func ErrorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
func RandomEmailTest() string {
	return RandomUserName() + "@test.com"
}
