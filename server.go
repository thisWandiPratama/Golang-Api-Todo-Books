package main

import (
	"golang_api_todo_books/config"
	"golang_api_todo_books/controllers"
	"golang_api_todo_books/middleware"
	"golang_api_todo_books/repository"
	"golang_api_todo_books/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var (
	db             *gorm.DB                   = config.SetupDatabaseConnection()
	userRepository repository.UserRepository  = repository.NewUserRepository(db)
	bookRepository repository.BookRepository  = repository.NewBookRepository(db)
	jwtService     service.JWTService         = service.NewJWTService()
	userService    service.UserService        = service.NewUserService(userRepository)
	bookService    service.BookService        = service.NewBookService(bookRepository)
	authService    service.AuthService        = service.NewAuthService(userRepository)
	authController controllers.AuthController = controllers.NewAuthController(authService, jwtService)
	userController controllers.UserController = controllers.NewUserController(userService, jwtService)
	bookController controllers.BookController = controllers.NewBookController(bookService, jwtService)
)

func main() {

	defer config.CloseDatabaseConnection(db)
	r := gin.Default()

	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/login", authController.Login)
			auth.POST("/register", authController.Register)
		}

		userRoutes := v1.Group("/user", middleware.AuthorizeJWT(jwtService))
		{
			userRoutes.GET("/profile", userController.Profile)
			userRoutes.PUT("/profile", userController.Update)
		}

		bookRoutes := v1.Group("/books", middleware.AuthorizeJWT(jwtService))
		{
			bookRoutes.GET("/", bookController.All)
			bookRoutes.POST("/", bookController.Insert)
			bookRoutes.GET("/:id", bookController.FindByID)
			bookRoutes.PUT("/:id", bookController.Update)
			bookRoutes.DELETE("/:id", bookController.Delete)
		}
	}

	r.Run()
}
