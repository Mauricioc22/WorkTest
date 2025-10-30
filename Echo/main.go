package main 
import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main(){
	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}, time=${latency_human}\n",
	}))

	e.GET("/", func(c echo.Context) error{
		// return c.String(http.StatusOK, "Hola Mundo!")
		return c.JSON(http.StatusOK, map[string]string{
			"message": "Hola desde en servidor",
		})
	})

	e.GET("/saludo", func(c echo.Context) error{
		return c.String(http.StatusOK, "Saludos desde Echo")
		
	})
	

	e.Start(":8080")
}