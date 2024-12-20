package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateClaims(ctx *gin.Context, userID string) bool {
	claims, exists := ctx.Get("claims")
	if !exists {
		ctx.JSON(http.StatusForbidden, "Claims not found in context")
		return false
	}	

	claimsMap, ok := claims.(map[string]interface{})
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Invalid format for claims"})
		return false
	}

	subClaim, ok := claimsMap["sub"].(string)
	if !ok {
		ctx.JSON(http.StatusForbidden, gin.H{"error": "Sub claim not found or not a string"})
		return false
	}

	if subClaim != userID {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Only owners of this video can update"})
		return false
	}
	
	return true
}
