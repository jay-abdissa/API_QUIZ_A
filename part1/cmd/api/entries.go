//Filename:cmd/api/entries.go
package main

import	(
	"fmt"
	"time"
	"net/http"
	//"strconv"
	//"part1.castillojadah.net/internal/data"
	//"github.com/julienschmidt/httprouter"
)
//create entry handler for the POST /v1/entries endpoint
func (app *application) createEntryHandler (w http.ResponseWriter, r *http.Request){
	
	//our target decode destination
	type input struct {
		ID int64 `json:"id"`
		CreatedAt time.Time `json:"createdat"`
		Name string `json:"name"`
		Year string `json:"year"`
		Contact string `json:"contact"`
		Phone string `json:"phone"`
		Email string `json:"email"`
		Website string `json:"website"`
		Address string `json:"address"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
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