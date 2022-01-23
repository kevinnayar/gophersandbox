// pkg/app/app.go

package application

import (
	"github.com/kevinnayar/gophersandbox/pkg/config"
	"github.com/kevinnayar/gophersandbox/pkg/db"
)

type Application struct {
	DB  *db.DB
	Cfg *config.Config
}

func Get() (*Application, error) {
	cfg := config.Get()
	db, err := db.Get(cfg.GetDBConnStr())

	if err != nil {
		return nil, err
	}

	return &Application{
		DB:  db,
		Cfg: cfg,
	}, nil
}
