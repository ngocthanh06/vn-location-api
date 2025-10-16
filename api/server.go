package api

import (
	"github.com/gin-gonic/gin"
	vnlocation "github.com/ngocthanh06/vn-location-api"
	"net/http"
)

func RunServer() {
	r := gin.Default()

	r.GET("/provinces", func(c *gin.Context) {
		c.JSON(http.StatusOK, vnlocation.GetProvinces())
	})

	r.GET("/wards", func(c *gin.Context) {
		provinceCodeStr := c.Query("province_code")
		wards := vnlocation.GetWardsByProvinceCode(&provinceCodeStr)
		if wards == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid or missing province_code parameter"})
			return
		}

		c.JSON(http.StatusOK, wards)
	})

	r.Run(":8080")
}
