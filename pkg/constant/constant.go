package constant

const (
	DB_HOST     = "DB_HOST"
	DB_PORT     = "DB_PORT"
	DB_USERNAME = "DB_USERNAME"
	DB_PASSWORD = "DB_PASSWORD"
	DB_NAME     = "DB_NAME"
)

const (
	DISTRICT_ID = "DISTRICT_ID"
)

var ALLOWED_EXTENSIONS map[string]bool = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
}
