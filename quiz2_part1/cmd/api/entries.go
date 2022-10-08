//Filename:cmd/api/entries.go
package main

import	(
	"fmt"
	"time"
	"net/http"
	//"strconv"
	//"quiz2_part1.castillojadah.net/internal/data"
	//"github.com/julienschmidt/httprouter"
)
//create entry handler for the POST /v1/entries endpoint
func (app *application) createEntryHandler (w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Create a new entry...")
}
//create entry handler for the GET /v1/entries/:id endpoint
func (app *application) showEntryHandler (w http.ResponseWriter, r *http.Request){
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	//Create a new instance of my struct
	mystruct := data.Mystruct{
		ID: id,
		CreatedAt: time.Now() ,
		Name: "Jadah Castillo",
		Year: "",
		Contact: "",
		Phone: "",
		Email: "",
		Website: "",
		Address: "",

	}
	err.writeJSON(w, http.StatusOK, mystruct, nil)
	if err != nil {
		app.logger.Println(err)
		http.error(w, "The server could not process your request", http.StatusInternalServerError)

	}
}