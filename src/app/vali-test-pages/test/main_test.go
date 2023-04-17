package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/guilhermeonrails/api-go-gin/controllers"
	"github.com/guilhermeonrails/api-go-gin/database"
	"github.com/guilhermeonrails/api-go-gin/models"
	"github.com/stretchr/testify/assert"
)

var ID int

func RotasTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()

	return router
}

func CreateAlunoMock() {
	aluno := models.Aluno{Nome: "Aluno teste", CPF: "12345678900", RG: "098765432"}
	database.DB.Create(&aluno)
	ID = int(aluno.ID)

}

func DestroyAlunoMock() {
	var aluno models.Aluno
	database.DB.Delete(&aluno, ID)
}

func TestSimples(t *testing.T) {
	t.Fatal("should not be called")
}

func TestStatusCode(t *testing.T) {
	r := RotasTeste()
	r.GET("/:nome", controllers.Saudacao)

	req, _ := http.NewRequest("GET", "/flavio", nil)
	result := httptest.NewRecorder()
	r.ServeHTTP(result, req)

	// if result.Code != 200 {
	// 	t.Fatalf("Status error: %d", result.Code)
	// }

	assert.Equal(t, http.StatusOK, result.Code, "should equal result")

	mockResult := `{"API diz:":"E ai flavio, tudo beleza?"}`
	resultBody, _ := ioutil.ReadAll(result.Body)

	assert.Equal(t, mockResult, string(resultBody))
	// fmt.Println(mockResult, "mock")
	// fmt.Println(string(resultBody), "result")
}

func TestGetAlunosAll(t *testing.T) {
	database.ConectaComBancoDeDados()
	CreateAlunoMock()
	defer DestroyAlunoMock()
	r := RotasTeste()
	r.GET("/alunos", controllers.ExibeTodosAlunos)

	req, _ := http.NewRequest("GET", "/alunos", nil)
	result := httptest.NewRecorder()
	r.ServeHTTP(result, req)

	assert.Equal(t, http.StatusOK, result.Code)
	// fmt.Println(result.Body)

}

func TestGetAlunosByCPF(t *testing.T) {
	database.ConectaComBancoDeDados()
	CreateAlunoMock()
	defer DestroyAlunoMock()

	r := RotasTeste()
	r.GET("/alunos/cpf/:cpf", controllers.BuscaAlunoPorCPF)

	req, _ := http.NewRequest("GET", "/alunos/cpf/12345678900", nil)
	result := httptest.NewRecorder()
	r.ServeHTTP(result, req)

	assert.Equal(t, result.Code, http.StatusOK)

}

func TestGetAlunosById(t *testing.T) {
	database.ConectaComBancoDeDados()
	CreateAlunoMock()
	defer DestroyAlunoMock()

	r := RotasTeste()
	r.GET("/alunos/:id", controllers.BuscaAlunoPorID)

	pathURL := "/alunos/" + strconv.Itoa(int(ID))
	req, _ := http.NewRequest("GET", pathURL, nil)
	result := httptest.NewRecorder()
	r.ServeHTTP(result, req)

	var alunoMock models.Aluno
	json.Unmarshal(result.Body.Bytes(), &alunoMock)
	// fmt.Println(alunoMock.Nome, alunoMock.ID)

	assert.Equal(t, "Aluno teste", alunoMock.Nome, "message")
	assert.Equal(t, "12345678900", alunoMock.CPF)
	assert.Equal(t, "098765432", alunoMock.RG)

}

func TestDeleteAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CreateAlunoMock()

	r := RotasTeste()
	r.DELETE("/alunos/:id", controllers.DeletaAluno)

	pathUrl := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("DELETE", pathUrl, nil)
	result := httptest.NewRecorder()

	r.ServeHTTP(result, req)

	assert.Equal(t, http.StatusOK, result.Code, "deleting")
}

func TestUpdateAluno(t *testing.T) {
	database.ConectaComBancoDeDados()
	CreateAlunoMock()
	defer DestroyAlunoMock()

	r := RotasTeste()
	r.PATCH("/alunos/:id", controllers.EditaAluno)

	aluno := models.Aluno{Nome: "Aluno teste", CPF: "12345678910", RG: "098765432"}
	parseAlunoJson, _ := json.Marshal(aluno)

	pathUrl := "/alunos/" + strconv.Itoa(ID)
	req, _ := http.NewRequest("PATCH", pathUrl, bytes.NewBuffer(parseAlunoJson))
	result := httptest.NewRecorder()

	r.ServeHTTP(result, req)

	var alunosMockUpdate models.Aluno
	json.Unmarshal(result.Body.Bytes(), &alunosMockUpdate)

	assert.Equal(t, "12345678910", alunosMockUpdate.CPF)
}
