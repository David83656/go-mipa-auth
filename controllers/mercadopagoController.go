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
