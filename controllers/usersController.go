package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/BuilderSong/gin-json-crud/initializers"
	"github.com/BuilderSong/gin-json-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {

	//get password and email from request body
	var body struct {
		Email    string `json:"Email" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email or password missing",
		})
		return
	}

	//hash the password

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.MinCost)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password hash issue",
		})
		return
	}

	//save the email and password to user table

	user := models.User{Email: body.Email, Password: string(hashedPassword)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "database saving issue",
		})
		return
	}

	//return response to user

	c.JSON(200, gin.H{})

}

func Login(c *gin.Context) {
	//get email and password from request body
	var body struct {
		Email    string `json:"Email" binding:"required"`
		Password string `json:"Password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email or password missing",
		})
		return
	}

	//find email in db
	var user models.User
	initializers.DB.Where("email = ?", body.Email).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "email is wrong",
		})
		return
	}

	//compare password in db

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "password is wrong",
		})
		return
	}

	//generate jwt and set cookie to user
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":   user.ID,
		"expire": time.Now().Add(time.Hour * 6).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "jwt token generating error",
		})
		return
	}

	//return response
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*6, "", "", false, true)

	c.JSON(200, gin.H{})
}

func Logout(c *gin.Context) {
	_, exists := c.Get("user")
	if exists {
		c.SetCookie("Authorization", "", 0, "", "", false, true)
		c.JSON(http.StatusOK, gin.H{
			"message": "logged out",
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "user not logged in",
		})
	}
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{
		"message": "validate successed",
		"user":    user.(models.User).Email,
	})
}
