package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	sec "github.com/faizaldr/bang-al/security"
	// "Errors"
)

func main() {
	r := gin.Default()

	api := r.Group("/api")
	{

		// http://localhost:8080/api/pegawai_encrypt_nip?nip=17231237
		api.GET("/pegawai_encrypt_nip", func(ctx *gin.Context) {
			nip := ctx.Query("nip")

			// sec.EncryptURLSafe([]byte(nip_database), []byte("INIadalahEncryptionKey1234567890"))
			nipEncrypted, err := sec.EncryptURLSafe([]byte(nip), []byte("INIadalahEncryptionKey1234567890"))
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{"message": "failed", "error": err, "nip": nil})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"nip": nipEncrypted})
		})

		// http://localhost:8080/api/pegawai_decrypt_nip?nip=vV-64pcoWjb1uVJQG5ufTgyWo61VRW19gjITosKsBiE0qRqM
		api.GET("/pegawai_decrypt_nip", func(ctx *gin.Context) {
			nip := ctx.Query("nip")

			nipEncrypted, err := sec.DecryptURLSafe(nip, []byte("INIadalahEncryptionKey1234567890"))
			if err != nil {
				ctx.JSON(http.StatusOK, gin.H{"message": "failed", "error": err, "nip": nil})
				return
			}

			ctx.JSON(http.StatusOK, gin.H{"nip": string(nipEncrypted)})
		})
	}
	r.Run(":8080")
}
