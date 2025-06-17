package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)


type Apaixonada struct{
	ID string `json:"id"`
	Nome string `json:"nome"`
	Nacionalidade string `json:"nacionalidade"`
	Cidade string `json:"cidade"`
	País string `json:"pais"`
}

 var apaixonadas = []Apaixonada{
	{    
		ID: "1",
		Nome: "Ani Silva Costa",
		Nacionalidade: "Guineense",
		Cidade: "curitiba",
		País: "Brasil",
	},
	{   
		ID: "2",
		Nome: "Márcia Panzo",
		Nacionalidade: "Angolana",
		Cidade: "Florianopolis",
		País: "Brasil",
	},
	{   
		ID: "3",
		Nome: "Esther Oliveira",
		Nacionalidade: "Brasileira",
		Cidade: "Cuiabá",
		País: "Brasil",
	},
	{
		ID: "4",
		Nome: "Rafaela Oliveira",
		Nacionalidade: "Brasileira",
		Cidade: "Cuiabá",
		País: "Brasil",
	},
	{
		ID: "5",
		Nome: "Avelina Sila",
		Nacionalidade: "Guineense",
		Cidade: "São Francisco de Conde",
		País: "Brasil",

	},
	{
		ID: "6",
		Nome: "Hortênciaa Fernandes",
		Nacionalidade: "Guineense",
		Cidade: "Três Laras",
		País: "Brasil",
	},
	{
		ID: "7",
		Nome: "Lara Santos",
		Nacionalidade: "Brasileira",
		Cidade: "Maraacanau",
		País: "Brasil",
	},
}

//getApaixonadas retorna todas as apixonadas
func getApaixonadas(c *gin.Context) {
	c.IndentedJSON(200, apaixonadas)
}

//AddApaixonada adiciona uma nova apixonada
func AddApaixonada(c *gin.Context) {
	var newApaixonada Apaixonada
	if err := c.BindJSON(&newApaixonada); err != nil {
		return
	}
	apaixonadas = append(apaixonadas, newApaixonada)
	c.IndentedJSON(201, newApaixonada)
}

//DeleteApaixonada deleta uma apixonada
func DeleteApaixonada(c *gin.Context) {
	id := c.Param("id")
	for i, a := range apaixonadas {
		if a.ID == id {
			apaixonadas = append(apaixonadas[:i], apaixonadas[i+1:]...)
			break
		}
	}
	c.IndentedJSON(http.StatusNoContent, apaixonadas)
}

//getApaixonadaByID retorna uma apixonada
func getApaixonadaByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range apaixonadas {
		if a.ID == id {
			c.IndentedJSON(200, a)
			return
		}
	}
	c.IndentedJSON(404, gin.H{"message": "apaixonada not found"})
}

func main() {
	println("Bem vindo a api do golang com framework gin!")

	router := gin.Default()
	router.GET("/apaixonadas", getApaixonadas)
	router.GET("/apaixonadas/:id", getApaixonadaByID)
	router.POST("/apaixonadas", AddApaixonada)
	router.DELETE("/apaixonadas/:id", DeleteApaixonada)
	router.Run(":8080")
}