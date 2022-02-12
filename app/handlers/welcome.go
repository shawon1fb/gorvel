package handlers

import (
	"net/http"

	"github.com/daison12006013/gorvel/pkg/facade/request"
	"github.com/daison12006013/gorvel/pkg/response"
)

func Home(w http.ResponseWriter, r *http.Request) {
	// If we're properly writing a response from the http.ResponseWriter,
	// therefore no need to write the header as Status 200 "OK",
	// although it is still good to write it at first, and override
	// underneath if there are conditional cases that you want to filter-out
	w.WriteHeader(http.StatusOK)

	// let's extend the request
	request := request.Parse(r)

	// prepare the data
	data := map[string]interface{}{
		"title": "Gorvel Rocks!",
	}

	// this is api request
	if request.IsJson() && request.WantsJson() {
		response.Json(w, data)
	}

	// render the template
	response.View(
		w,
		// this example below, we're telling the compiler
		// to parse the base.html first, and then parse the welcome.html
		// therefore the defined "body" should render accordingly
		[]string{"base.html", "welcome.html"},
		data,
	)
}
