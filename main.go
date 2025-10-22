package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// http://localhost:8080/api/pegawai?nip=17231237
	api := r.Group("/api")
	{
		api.GET("/pegawai", func(ctx *gin.Context) {
			nip := ctx.Query("nip")

			ctx.JSON(http.StatusOK, gin.H{"nip": nip})
		})
	}
	r.Run(":8080")
}
