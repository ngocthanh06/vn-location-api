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

	r.GET("/wards/", func(c *gin.Context) {
		provinceCode := c.Query("province_code")

		if provinceCode == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "province_code is required"})
			return
		}

		c.JSON(http.StatusOK, vnlocation.GetWardsByProvinceCode(provinceCode))
	})

	r.Run(":8080")
}
