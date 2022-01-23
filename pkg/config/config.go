// pkg/config/config.go

package config

import (
	"flag"
	"fmt"
	"os"
)

type Config struct {
	dbUser  string
	dbPswd  string
	dbHost  string
	dbPort  string
	dbName  string
	apiPort string
}

func Get() *Config {
	conf := &Config{}

	flag.StringVar(&conf.dbUser, "dbuser", os.Getenv("DB_USER"), "DB user name")
	flag.StringVar(&conf.dbPswd, "dbpswd", os.Getenv("DB_PASSWORD"), "DB pass")
	flag.StringVar(&conf.dbPort, "dbport", os.Getenv("DB_PORT"), "DB port")
	flag.StringVar(&conf.dbHost, "dbhost", os.Getenv("DB_HOST"), "DB host")
	flag.StringVar(&conf.dbName, "dbname", os.Getenv("DB_NAME"), "DB name")
	flag.StringVar(&conf.apiPort, "apiPort", os.Getenv("API_PORT"), "API Port")

	flag.Parse()

	return conf
}

func (c *Config) GetDBConnStr() string {
	return c.getDBConnStr(c.dbHost, c.dbName)
}

func (c *Config) getDBConnStr(dbhost, dbname string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		c.dbUser,
		c.dbPswd,
		dbhost,
		c.dbPort,
		dbname,
	)
}

func (c *Config) GetAPIPort() string {
	return ":" + c.apiPort
}
