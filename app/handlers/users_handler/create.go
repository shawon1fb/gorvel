package users_handler

import (
	"net/http"

	"github.com/gorilla/csrf"
	"github.com/lucidfy/lucid/app/models/users"
	"github.com/lucidfy/lucid/app/validations"
	"github.com/lucidfy/lucid/pkg/engines"
	"github.com/lucidfy/lucid/pkg/errors"
	"github.com/lucidfy/lucid/pkg/lucid"
)

func create(ctx lucid.Context) *errors.AppError {
	engine := ctx.Engine()
	router := ctx.Router()
	ses := engine.GetSession()
	req := engine.GetRequest()
	res := engine.GetResponse()

	bUrl, _ := router.Get("users.lists").URL()

	data := map[string]interface{}{
		"title":          "Create Form",
		"record":         &users.Model{},
		"isCreate":       true,
		csrf.TemplateTag: csrf.TemplateField(engine.(engines.NetHttpEngine).HttpRequest),

		//> to retrieve the flashes from Store() or somewhere
		//> that redirects back to Create()
		"success": ses.GetFlash("success"),
		"error":   ses.GetFlash("error"),
		"fails":   ses.GetFlashMap("fails"),
		"inputs":  ses.GetFlashMap("inputs"),

		"base_url": bUrl,
	}

	if req.WantsJson() {
		return res.Json(data, http.StatusOK)
	}

	return res.View(
		[]string{"base", "users/show"},
		data,
	)
}

func store(ctx lucid.Context) *errors.AppError {
	engine := ctx.Engine()
	ses := engine.GetSession()
	req := engine.GetRequest()
	res := engine.GetResponse()
	url := engine.GetURL()

	//> validate the inputs
	if validator := req.Validator(validations.Users().Create()); validator != nil {
		if req.WantsJson() {
			return res.Json(map[string]interface{}{
				"fails": validator.ValidationError,
			}, http.StatusUnauthorized)
		}

		ses.PutFlashMap("fails", validator.ValidationError)
		ses.PutFlashMap("inputs", req.All())
		return url.RedirectPrevious()
	}

	//> prepare message and status
	message := "Successfully Created!"
	status := http.StatusOK

	//> create user based on all request inputs
	data, app_err := users.Create(req.All())
	if app_err != nil {
		return app_err
	}

	//> for api based
	if req.WantsJson() {
		return res.Json(map[string]interface{}{
			"success": message,
			"data":    data,
		}, status)
	}

	//> for form based, just redirect
	ses.PutFlash("success", message)
	return url.RedirectPrevious()
}
