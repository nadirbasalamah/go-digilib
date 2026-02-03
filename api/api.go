package api

import (
	"go-digilib/api/handlers"
	"go-digilib/api/middlewares"
	"go-digilib/auth"
	"go-digilib/books"
	"go-digilib/categories"
	"go-digilib/users"

	"github.com/cloudinary/cloudinary-go/v2"
	echojwt "github.com/labstack/echo-jwt/v5"
	"github.com/labstack/echo/v5"
	"github.com/labstack/echo/v5/middleware"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

func NewEcho(repository *gorm.DB, cld *cloudinary.Cloudinary, jwtConfig middlewares.JWTConfig) *echo.Echo {
	var (
		e                 = echo.New()
		categories        = categories.New(repository)
		books             = books.New(repository)
		auth              = auth.New(repository)
		users             = users.New(repository)
		authHandler       = handlers.NewAuth(auth, jwtConfig)
		usersHandler      = handlers.NewUsers(users, cld)
		categoriesHandler = handlers.NewCategories(categories)
		booksHandler      = handlers.NewBooks(books, cld)
	)

	e.Validator = &middlewares.CustomValidator{
		Validator: middlewares.InitValidator(),
	}

	logger, _ := zap.NewProduction()
	loggerConfig := middleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c *echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	}

	logMiddleware := middlewares.LoggerConfig{Config: loggerConfig}
	jwtMiddleware := jwtConfig.Init()

	e.Use(logMiddleware.Init())

	authRoutes := e.Group("/api/v1/auth")

	authRoutes.POST("/register", authHandler.Register)
	authRoutes.POST("/login", authHandler.Login)

	userRoutes := e.Group("/api/v1", echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)

	userRoutes.GET("/profile", usersHandler.GetProfile)
	userRoutes.PATCH("/profile/edit", usersHandler.EditProfile)

	categoryRoutes := e.Group("/api/v1", echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)

	categoryRoutes.GET("/categories", categoriesHandler.GetAll)
	categoryRoutes.GET("/categories/:id", categoriesHandler.GetByID)
	categoryRoutes.POST("/categories", categoriesHandler.Create, middlewares.VerifyAdmin)
	categoryRoutes.PATCH("/categories/:id", categoriesHandler.Update, middlewares.VerifyAdmin)
	categoryRoutes.DELETE("/categories/:id", categoriesHandler.Delete, middlewares.VerifyAdmin)

	bookRoutes := e.Group("/api/v1", echojwt.WithConfig(jwtMiddleware), middlewares.VerifyToken)
	bookRoutes.GET("/books", booksHandler.GetAll)
	bookRoutes.GET("/books/:id", booksHandler.GetByID)
	bookRoutes.GET("/books/category/:id", booksHandler.GetByCategory)
	bookRoutes.POST("/books", booksHandler.Create, middlewares.VerifyAdmin)
	bookRoutes.PATCH("/books/:id", booksHandler.Update, middlewares.VerifyAdmin)
	bookRoutes.DELETE("/books/:id", booksHandler.Delete, middlewares.VerifyAdmin)

	return e
}
