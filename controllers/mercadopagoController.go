package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PeticiónMp(c *gin.Context) {
	apiURL := "https://api.mercadopago.com/v1/"

	accessToken := "*************************************"

	req, err := http.NewRequest("GET", apiURL, nil)
	if err != nil {
		fmt.Println("Error al crear la solicitud:", err)
		return
	}

	req.Header.Add("Authorization", "Bearer "+accessToken)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error al hacer la solicitud HTTP:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error al leer la respuesta:", err)
		return
	}

	fmt.Println(string(body))
}

func Preciopostal(codigoPostal int, c *gin.Context) int {

	precios := map[int]int{
		5000: 1320,
		5500: 1500,
		6200: 2000,
		9000: 1240,
		9407: 1800,
		3300: 4000,
	}

	precio, encontrado := precios[codigoPostal]
	if encontrado {
		return precio
	}
	return -1
}
