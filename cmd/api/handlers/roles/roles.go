package roles

import (
	"encoding/json"
	"net/http"

	"github.com/kevinnayar/gophersandbox/pkg/application"
	"github.com/kevinnayar/gophersandbox/pkg/logger"
	"github.com/kevinnayar/gophersandbox/pkg/middleware"

	"github.com/julienschmidt/httprouter"
	"github.com/lib/pq"
)

type Role struct {
	Id       string   `json:"id"`
	Name     string   `json:"name"`
	RightIds []string `json:"rightIds"`
}

type RolesResponse struct {
	Success bool   `json:"success"`
	Data    []Role `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func getAll(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		rows, err := app.DB.Client.Query("SELECT * FROM roles")
		if err != nil {
			logger.Error.Fatal(err.Error())
		}

		var roles []Role

		for rows.Next() {
			var id string
			var name string
			var rightIds []string

			err = rows.Scan(&id, &name, pq.Array(&rightIds))
			if err != nil {
				logger.Error.Fatal(err.Error())
			}

			roles = append(roles, Role{
				Id: id,

				Name:     name,
				RightIds: rightIds,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		res := RolesResponse{Success: true, Data: roles}
		json.NewEncoder(w).Encode(res)
	}
}

func GetAll(app *application.Application) httprouter.Handle {
	mdw := []middleware.Middleware{middleware.LogRequest}

	return middleware.Chain(getAll(app), mdw...)
}
