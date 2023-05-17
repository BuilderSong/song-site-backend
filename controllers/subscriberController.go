package controllers

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/BuilderSong/gin-json-crud/initializers"
	"github.com/BuilderSong/gin-json-crud/models"
	"github.com/gin-gonic/gin"
	"gopkg.in/gomail.v2"
)

func Subscribe(c *gin.Context) {
	//get data off req body
	var body struct {
		Email string
		Name  string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	subscriber := models.Subscriber{Email: body.Email, Name: body.Name}

	if err := initializers.DB.Create(&subscriber).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "subscriber saved to db",
		"posts":   subscriber})
}

func UnSubscribe(c *gin.Context) {
	//get data off req body
	var body struct {
		Email string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if result := initializers.DB.Where("email = ?", body.Email).Delete(&models.Subscriber{}); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	c.JSON(200, gin.H{
		"message": "deleted ok",
	})
}

func SendEmails(c *gin.Context) {
	var body struct {
		ID       int
		Topic    string
		Title    string
		Abstract string
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	type Subscribers struct {
		Name  string
		Email string
	}

	var subscribers []Subscribers

	result := initializers.DB.Table("subscribers").Select("email, name").Find(&subscribers)

	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read from db"})
		return
	}

	fmt.Println(subscribers)

	d := gomail.NewDialer("smtp.gmail.com", 587, os.Getenv("EMAIL"), os.Getenv("EMAIL_PD"))
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	s, err := d.Dial()
	if err != nil {
		panic(err)
	}

	m := gomail.NewMessage()

	blog_url := os.Getenv("DOMAIN") + "/blog/" + strconv.Itoa(body.ID)
	unsubscribe_url := os.Getenv("DOMAIN") + "/unsubscribe"

	for _, r := range subscribers {
		m.SetHeader("From", os.Getenv("EMAIL"))
		m.SetAddressHeader("To", r.Email, r.Name)
		m.SetHeader("Subject", fmt.Sprintf("%s from Song's Site", body.Title))
		htmlBody := fmt.Sprintf("<html><body>Hello %s!<br><br>How are you doing recently? Below is another blog from <a href=%s>Song's Site.</a><br><br>Title: %s<br><br>Abstract: %s<br><br><a href=%s >Click here to read more</a><br><br><a href=%s>Click here to unsubsctibe</a><br><br><br>Regards,<br><a href=%s>Song's Site</a></body></html>", r.Name, os.Getenv("DOMAIN"), body.Title, body.Abstract, blog_url, unsubscribe_url, os.Getenv("DOMAIN"))
		m.SetBody("text/html", htmlBody)

		if err := gomail.Send(s, m); err != nil {
			log.Printf("Could not send email to %q: %v", r.Email, err)
		}

		m.Reset()
	}

}
