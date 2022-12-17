package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CompareTime(c *gin.Context) {
	day := time.Now()
	fmt.Println(day)

	c.JSON(http.StatusOK, gin.H{"data": day})

}
