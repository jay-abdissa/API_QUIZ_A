//Filename:cmd/api/entries.go
package main

import	(
	"fmt"
	"net/http"
	//"strconv"

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
	//display the school id
	fmt.Fprintln(w, "Show the details for the entry %d\n", id)
}