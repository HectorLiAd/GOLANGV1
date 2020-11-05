package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GOLANGV1/BD"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

type Persona struct {
	PERSONA_ID      int    `JSON:"ID_PERSONA"`
	PERSONA_NOMBRE  string `JSON:"NOMBRE"`
	PERSONA_APE_PAT string `JSON:"APELLIDO_PATERNO"`
	PERSONA_APE_MAT string `JSON:"APELLIDO_MATERNO"`
	PERSONA_EDAD    int    `JSON:"EDAD"`
}

var databaseConnection *sql.DB

func main() {
	databaseConnection = BD.InitBD()
	defer databaseConnection.Close()

	r := chi.NewRouter()

	r.Get("/personas", func(w http.ResponseWriter, r *http.Request) {
		const sql = `SELECT * FROM PERSONA;`
		result, err := databaseConnection.Query(sql)
		catch(err)

		var personas []*Persona
		for result.Next() {
			persona := &Persona{}
			err = result.Scan(
				&persona.PERSONA_ID,
				&persona.PERSONA_NOMBRE,
				&persona.PERSONA_APE_PAT,
				&persona.PERSONA_APE_MAT,
				&persona.PERSONA_EDAD)
			catch(err)
			personas = append(personas, persona)
		}
		respondwithJSON(w, http.StatusOK, personas)
	})
	http.ListenAndServe(":3000", r)
	fmt.Println(r)
}

func respondwithJSON(w http.ResponseWriter, code int, playload interface{}) {
	response, _ := json.Marshal(playload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
