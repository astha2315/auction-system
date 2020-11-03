package route

import (
	"auction-system/controllers"

	"github.com/labstack/echo"
)

func StudentService(e *echo.Echo) {

	//
	e.GET("/students/:studentId", controllers.StudentDetails)

}
