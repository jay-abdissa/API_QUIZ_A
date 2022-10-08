//Filename:cmd/api/entries.go
package main

import	(
	"fmt"
	"time"
	"net/http"
	//"strconv"
	"part1.castillojadah.net/internal/data"
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
		app.notFoundResponse(w, r)
		return
	}
	//Create a new instance of my struct
	data := data.Mystruct{
		ID: id,
		CreatedAt: time.Now() ,
		Name: "Jadah Castillo",
		Year: "1",
		Contact: "1",
		Phone: "1",
		Email: "1",
		Website: "1",
		Address: "1",

	}
	err = app.writeJSON(w, http.StatusOK, envelope{ "data":data }, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)

	}
}