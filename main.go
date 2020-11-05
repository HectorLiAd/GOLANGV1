package main

import (
	"fmt"

	"github.com/GOLANGV1/BD"
	"github.com/go-chi/chi"
)

type Persona struct {
	PERSONA_ID      int    `JSON:"ID_PERSONA"`
	PERSONA_NOMBRE  string `JSON:"NOMBRE"`
	PERSONA_APE_PAT string `JSON:"APELLIDO_PATERNO"`
	PERSONA_APE_MAT string `JSON:"APELLIDO_MATERNO"`
	PERSONA_EDAD    int    `JSON:"EDAD"`
}

func main() {
	// r := gin.Default()
	// |r.GET("/ping", func(c *gin.Context) {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "hola HOLA",
	// 	})
	// })
	// r.Run(":8087")
	connection := BD.InitBD()
	defer connection.Close()
	r := chi.NewRouter()
	fmt.Println(r)
}
