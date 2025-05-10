package main

import (
	"os"
	"strconv"

	"github.com/Teemo4621/Basic-Webchat/configs"
	"github.com/Teemo4621/Basic-Webchat/modules/servers"
	"github.com/Teemo4621/Basic-Webchat/pkgs/databases"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("error loading .env file")
	}

	cfg := new(configs.Config)
	cfg.App.Host = os.Getenv("FIBER_HOST")
	cfg.App.Port = os.Getenv("FIBER_PORT")

	cfg.PostgreSQL.Host = os.Getenv("DB_HOST")
	cfg.PostgreSQL.Port = os.Getenv("DB_PORT")
	cfg.PostgreSQL.Username = os.Getenv("DB_USERNAME")
	cfg.PostgreSQL.Password = os.Getenv("DB_PASSWORD")
	cfg.PostgreSQL.Database = os.Getenv("DB_DATABASE")
	cfg.PostgreSQL.SSLMode = os.Getenv("DB_SSL_MODE")

	cfg.JWT.Secret = os.Getenv("JWT_SECRET")

	expire, err := strconv.Atoi(os.Getenv("JWT_EXPIRE"))
	if err != nil {
		panic(err)
	}
	cfg.JWT.Expire = expire

	cfg.JWT.RefreshSecret = os.Getenv("JWT_REFRESH_SECRET")
	refreshExpire, err := strconv.Atoi(os.Getenv("JWT_REFRESH_EXPIRE"))
	if err != nil {
		panic(err)
	}
	cfg.JWT.RefreshExpire = refreshExpire

	db, err := databases.NewPostgresConnection(*cfg)
	if err != nil {
		panic(err)
	}

	if err := databases.Migrate(db); err != nil {
		panic(err)
	}

	server := servers.NewServer(cfg, db)
	server.Start()
}
