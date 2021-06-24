package config

import "fmt"

const (
	DATABASE_USERNAME string = "root"
	DATABASE_PASSWORD string = nil
	DATABASE_NAME     string = "budreck"
	DATABASE_HOST     string = "localhost"
	DATABASE_PORT     int32  = 5432
)

func DatabaseConnectionURL() string {
	return fmt.Sprintf(
		"%s://%s:%s@%s:%d/%s?sslmode=disable",
		DATABASE_HOST,
		DATABASE_USERNAME,
		DATABASE_PASSWORD,
		DATABASE_HOST,
		DATABASE_PORT,
		DATABASE_NAME,
	)
}
