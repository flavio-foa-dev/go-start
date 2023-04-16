package routes

import (
	"api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequest() {
	r := gin.Default()
	r.GET("/api/student", controllers.GetStudentAll)
	r.GET("/api/student/:id", controllers.GetStudentbyId)
	r.GET("/api/student/cpf/:cpf", controllers.GetStudentByCPF)

	r.POST("/api/student", controllers.CREATEStudent)

	r.PUT("/api/student/:id", controllers.UpdateStudent)

	r.DELETE("/api/student/:id", controllers.DestroyStudent)
	r.Run(":8000")
}
