package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/GOLANGV1/BD"
	"github.com/GOLANGV1/permisos"
	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
)

type Tarea struct {
	Lt_id  int    `json:"id"`
	Nombre string `json:"nombre"`
	Estado int    `json:"estado"`
}

var databaseConnection *sql.DB

func main() {
	databaseConnection = BD.InitBD()
	defer databaseConnection.Close()

	r := chi.NewRouter()
	r.Use(permisos.GetCors().Handler)
	r.Get("/tareas", ObtenerTareas)
	r.Post("/", CrearTarea)
	r.Put("/tareas/{id}", ActualizarTarea)
	r.Delete("/tareas/{id}", EliminarTarea)

	http.ListenAndServe(":3000", r)
}

//Obteniendo todos el listado de tareas
func ObtenerTareas(w http.ResponseWriter, r *http.Request) {
	const query = `select * from listaTareas`
	results, err := databaseConnection.Query(query)
	catch(err)
	var tareas []*Tarea
	for results.Next() {
		tarea := &Tarea{}
		err := results.Scan(&tarea.Lt_id, &tarea.Nombre, &tarea.Estado)
		catch(err)
		tareas = append(tareas, tarea)
	}
	respondWithJSON(w, http.StatusOK, tareas)
}

//Creando un listado de tarea
func CrearTarea(w http.ResponseWriter, r *http.Request) {
	var tarea_ Tarea
	json.NewDecoder(r.Body).Decode(&tarea_)
	fmt.Println(tarea_.Estado)
	fmt.Println(tarea_.Nombre)
	const query = "insert into listaTareas(nombre, estado) values(?, ?)"
	_, err := databaseConnection.Exec(query, tarea_.Nombre, tarea_.Estado)
	catch(err)
	respondWithJSON(w, http.StatusCreated, map[string]string{"mensaje": "tarea creada"})
}

//Actulizar tarea
func ActualizarTarea(w http.ResponseWriter, r *http.Request) {
	var tarea_ Tarea
	id := chi.URLParam(r, "id")
	json.NewDecoder(r.Body).Decode(&tarea_)
	const query = "update listaTareas set estado = ? where lt_id = ?"
	_, err := databaseConnection.Exec(query, tarea_.Estado, id)
	catch(err)
	respondWithJSON(w, http.StatusCreated, map[string]string{"mensaje": "tarea actulizada"})
}

//ELIMINAR TAREA
func EliminarTarea(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	const query = "delete from listaTareas where lt_id = ?"
	_, err := databaseConnection.Exec(query, id)
	catch(err)
	mensaje := "ID " + id + " eliminado correctamente"
	respondWithJSON(w, http.StatusCreated, map[string]string{"mensaje": mensaje})
}

func respondWithJSON(w http.ResponseWriter, cod int, playload interface{}) {
	response, _ := json.Marshal(playload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(cod)
	w.Write(response)
}

func catch(err error) {
	if err != nil {
		panic(err)
	}
}
