package tronics

import (
	"fmt"
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo"
	"gopkg.in/go-playground/validator.v9"
)

var e = echo.New()
var v = validator.New()

func init() {
	err := cleanenv.ReadEnv(&cfg)
	fmt.Printf("%v", cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configurations")
	}
}

func serverMessage(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		fmt.Println("inside custom middleware")
		return next(c)
	}
}

func Start() {
	//e.Use(serverMessage) //for all endpoints e.Pre can be used as well .pre() will always execute first before .use(middleware) and middleware declared in GET/POST/PUT/DELETE method
	e.GET("/products", getProducts, serverMessage)
	e.GET("/products/:id", getProduct)
	e.POST("/products", createProduct)
	e.PUT("/products/:id", updateProduct)
	e.DELETE("/products/:id", deleteProduct)

	e.Logger.Print(fmt.Sprintf("Listening on port %s", cfg.Port))
	e.Logger.Fatal(e.Start(fmt.Sprintf("localhost:%s", cfg.Port)))
}
