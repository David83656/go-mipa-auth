package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/David83656/go-mipa-auth/initializers"
	"github.com/David83656/go-mipa-auth/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

//INICIO DE SESIÓN

func SignUp(c *gin.Context) {
	var body struct {
		Email    string
		Password string
		Name     string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error ": "Failed hash the password",
		})
	}

	user := models.User{Email: body.Email, Password: string(hash), Name: body.Name}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create the user",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Success": "Success creating the user!",
	})

}

func Login(c *gin.Context) {
	var body struct {
		Email    string
		Password string
		Name     string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to read body",
		})
		return
	}
	var user models.User
	initializers.DB.First(&user, "email = ?", body.Email)
	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid email or password",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	fmt.Println(tokenString, err)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Failed to create token",
		})
		return
	}
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{})

}

func Validate(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"Message": "Logged in!",
	})
}

func Preciopostal(codigoPostal int, c *gin.Context) int {

	precios := map[int]int{
		5000: 100,
		5500: 150,
		6200: 200,
		9000: 120,
		9407: 180,
		3300: 90,
	}

	precio, encontrado := precios[codigoPostal]
	if encontrado {
		return precio
	}
	return -1
}
