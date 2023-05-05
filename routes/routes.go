package routes

import (
	"valdson/api-gin/controllers"

	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.GET("/alunos", controllers.ShowAllStudents)
	r.GET("/:nome", controllers.Greetings)
	r.POST("/alunos", controllers.CreateNewStudent)
	r.GET("/alunos/:id", controllers.SearchStudentID)
	r.DELETE("/alunos/:id", controllers.DeleteStudent)
	r.PATCH("/alunos/:id", controllers.EditStudent)
	r.GET("/alunos/cpf/:cpf", controllers.SerchStudentCPF)
	r.Run(":5000")
}
