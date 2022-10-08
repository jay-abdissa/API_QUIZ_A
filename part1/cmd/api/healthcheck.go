//Filename:cmd/api/healthcheck.go
package main

import	(
	"net/http"
)
func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request)	{
	//create a map to hold our healthcheck data
	data := envelope{
		"status": "available",
		"system info": map[string]string{
			"environment": app.config.env,
			"version": version,
		},
	}
	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil{
		app.serverErrorResponse(w, r, err)
		return
	}
	//specify that we will serve our responses using json
	//w.Header().Set("Content-Type", "application/json")
	//write the json as a HTTP response body
	//w.Write([]byte(js))

}
