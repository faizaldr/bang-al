package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	sec "github.com/faizaldr/bang-al/security"
	// "Errors"
)

func main() {
	r := gin.Default()

	// http://localhost:8080/api/pegawai?nip=17231237
	api := r.Group("/api")
	{
		api.GET("/pegawai", func(ctx *gin.Context) {
			nip := ctx.Query("nip")

			nipEncrypted, err := sec.EncryptURLSafe([]byte(nip), []byte("INIadalahEncryptionKey1234567890"))
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{"message": "failed", "error": err, "nip": nil})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"nip": nipEncrypted})
		})
	}
	r.Run(":8080")
}
