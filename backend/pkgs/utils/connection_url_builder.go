package utils

import (
	"errors"
	"fmt"

	"github.com/Teemo4621/Basic-Webchat/configs"
)

func ConnectionURLBuilder(stuff string, config configs.Config) (string, error) {
	var url string

	switch stuff {
	case "fiber":
		url = config.App.Host + ":" + config.App.Port
	case "postgres":
		url = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.PostgreSQL.Host,
			config.PostgreSQL.Port,
			config.PostgreSQL.Username,
			config.PostgreSQL.Password,
			config.PostgreSQL.Database,
			config.PostgreSQL.SSLMode,
		)
	default:
		errMsg := fmt.Sprintf("error, connection url builder doesn't know the %s", stuff)
		return "", errors.New(errMsg)
	}

	return url, nil
}
