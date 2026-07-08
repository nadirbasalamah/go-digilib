package constant

const PORT = "PORT"

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

const (
	CLOUDINARY_URL     = "CLOUDINARY_URL"
	RAJAONGKIR_API_KEY = "RAJAONGKIR_API_KEY"
	AI_API_KEY         = "AI_API_KEY"
	AI_MODEL           = "AI_MODEL"
)

var ALLOWED_EXTENSIONS map[string]bool = map[string]bool{
	".png":  true,
	".jpg":  true,
	".jpeg": true,
}
