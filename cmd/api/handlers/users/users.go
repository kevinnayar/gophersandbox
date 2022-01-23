package users

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/kevinnayar/gophersandbox/pkg/application"
	"github.com/kevinnayar/gophersandbox/pkg/logger"
	"github.com/kevinnayar/gophersandbox/pkg/middleware"

	"github.com/julienschmidt/httprouter"
)

type User struct {
	Id             string `json:"id"`
	RoleId         string `json:"roleId"`
	DisplayName    string `json:"displayName"`
	FullName       string `json:"fullName"`
	Email          string `json:"email"`
	ImgUrl         string `json:"imgUrl,omitempty"`
	UtcTimeCreated int    `json:"utcTimeCreated"`
	UtcTimeUpdated int    `json:"utcTimeUpdated"`
}

type UsersResponse struct {
	Success bool   `json:"success"`
	Data    []User `json:"data,omitempty"`
	Error   string `json:"error,omitempty"`
}

func getAll(app *application.Application) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		defer r.Body.Close()

		rows, err := app.DB.Client.Query("SELECT * FROM users")
		if err != nil {
			logger.Error.Fatal(err.Error())
		}

		var users []User

		for rows.Next() {
			var id string
			var roleId string
			var displayName string
			var fullName string
			var email string
			var imgUrlMaybe sql.NullString
			var utcTimeCreated int
			var utcTimeUpdated int

			err = rows.Scan(&id, &roleId, &displayName, &imgUrlMaybe, &fullName, &email, &utcTimeUpdated, &utcTimeCreated)
			if err != nil {
				logger.Error.Fatal(err.Error())
			}

			var imgUrl string
			if imgUrlMaybe.Valid {
				imgUrl = imgUrlMaybe.String
			}

			users = append(users, User{
				Id:             id,
				RoleId:         roleId,
				DisplayName:    displayName,
				FullName:       fullName,
				Email:          email,
				ImgUrl:         imgUrl,
				UtcTimeCreated: utcTimeCreated,
				UtcTimeUpdated: utcTimeUpdated,
			})
		}

		w.Header().Set("Content-Type", "application/json")
		res := UsersResponse{Success: true, Data: users}
		json.NewEncoder(w).Encode(res)
	}
}

func GetAll(app *application.Application) httprouter.Handle {
	mdw := []middleware.Middleware{middleware.LogRequest}

	return middleware.Chain(getAll(app), mdw...)
}
