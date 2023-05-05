package controllers

import (
	"net/http"
	"valdson/api-gin/database"
	"valdson/api-gin/models"

	"github.com/gin-gonic/gin"
)

func ShowAllStudents(c *gin.Context) {
	var alunos []models.Aluno
	database.DB.Find(&alunos)
	c.JSON(200, alunos)
}

func Greetings(c *gin.Context) {
	nome := c.Params.ByName("nome")
	c.JSON(200, gin.H{
		"API diz": "E aí " + nome + ", tudo beleza?",
	})
}

func CreateNewStudent(c *gin.Context) {
	var aluno models.Aluno
	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}

	database.DB.Create(&aluno)
	c.JSON(http.StatusOK, aluno)

}

func SearchStudentID(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.First(&aluno, id)
	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)
}

func DeleteStudent(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")

	database.DB.Delete(&aluno, id)

	c.JSON(http.StatusOK, gin.H{"data": "Aluno deletado com sucesso"})

}

func EditStudent(c *gin.Context) {
	var aluno models.Aluno
	id := c.Params.ByName("id")
	database.DB.First(&aluno, id)

	if err := c.ShouldBindJSON(&aluno); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&aluno).UpdateColumns(aluno)
	c.JSON(http.StatusOK, aluno)
}

func SerchStudentCPF(c *gin.Context) {
	var aluno models.Aluno
	cpf := c.Param("cpf")
	database.DB.Where(&models.Aluno{CPF: cpf}).First(&aluno)

	if aluno.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Not found": "Aluno não encontrado",
		})
		return
	}

	c.JSON(http.StatusOK, aluno)

}
