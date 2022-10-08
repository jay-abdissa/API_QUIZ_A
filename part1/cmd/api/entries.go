//Filename:cmd/api/entries.go
package main

import	(
	"fmt"
	"time"
	"net/http"
	//"strconv"
	"part1.castillojadah.net/internals/data"
	"part1.castillojadah.net/internals/validator"
	//"github.com/julienschmidt/httprouter"
)
//create entry handler for the POST /v1/entries endpoint
func (app *application) createEntryHandler (w http.ResponseWriter, r *http.Request){
	
	//our target decode destination
	var input struct{
		Name string `json:"name"`
		Year string `json:"year"`
		Contact string `json:"contact"`
		Phone string `json:"phone"`
		Email string `json:"email"`
		Website string `json:"website"`
		Address string `json:"address"`
	}

	err := app.readJSON(w, r, &input )
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	entries := &data.Mystruct{
		Name: input.Name,
		Contact: input.Contact,
		Phone: input.Phone,
		Email: input.Email,
		Website: input.Website,
		Address: input.Address,
	}
	//initialize a new validator instance
	v := validator.New()
	//check the map to determine if there were any validation errors
	if data.ValidateEntries(v,entries); !v.Valid(){
		app.failedValidationResponse(w,r,v.Errors)
		return
	}
	//Display the request
	fmt.Fprintf(w, "%+v\n", input)

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

func (app *application) showRandomHandler(w http.ResponseWriter, r *http.Request){
	
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Display the random string
	num := int(id) //Convert int64 id to int
	fmt.Fprintf(w, "show random string for %d\n", num)
	fmt.Fprintln(w, app.Tools.GenerateRandomString(num))
}

