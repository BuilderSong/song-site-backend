package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/BuilderSong/gin-json-crud/initializers"
	"github.com/BuilderSong/gin-json-crud/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func AuthRequired(c *gin.Context) {
	//get cookie from request head
	fmt.Println(c.Request.Cookies())
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//decode and validate cookie
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		//check the expire date
		if float64(time.Now().Unix()) > claims["expire"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//find the user with token sub
		var user models.User
		initializers.DB.First(&user, claims["user"])

		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		//attach to req
		c.Set("user", user)

	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	//continue
	c.Next()
}
