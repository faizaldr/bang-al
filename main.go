package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/faizaldr/bang-al/security/crypto"
)

func main() {
	r := gin.Default()

	// http://localhost:8080/api/pegawai?nip=17231237
	api := r.Group("/api")
	{
		api.GET("/pegawai", func(ctx *gin.Context) {
			nip := ctx.Query("nip")

			nipEncrypted := crypto.EncryptURLSafe(nip, "INIadalahEncryptionKey")

			ctx.JSON(http.StatusOK, gin.H{"nip": nipEncrypted})
		})
	}
	r.Run(":8080")
}
