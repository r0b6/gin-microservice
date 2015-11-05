package app

import (
	"github.com/gin-gonic/gin"
    "regexp"
)

func ValidCityMiddleware() gin.HandlerFunc {
  return func(c *gin.Context) {
    city := c.Param("city")

    isValid, err := regexp.MatchString(`^[a-zA-Z]+(\,[a-zA-Z]+)?$`, city)

    if err != nil || !isValid {
        c.JSON(404, gin.H{
                "message": "City not valid",
        })
        c.Abort()
    } else {
    	c.Next()
    }
  }
}