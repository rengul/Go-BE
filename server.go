package main

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

type signInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	ID       string
	username string
	password string
}

type AuthClaims struct {
	jwt.StandardClaims
	User *User `json:"user"`
}

type signInResponse struct {
	Token string `json:"token"`
}

func SignIn(c *gin.Context) {
	inp := new(signInput)

	if err := c.BindJSON(inp); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	pwd := sha1.New()
	pwd.Write([]byte(inp.Password))
	pwd.Write([]byte("salt"))
	//password := fmt.Sprintf("%x", pwd.Sum(nil))

	// user, err := a.userRepo.GetUser(ctx, username, password)
	// if err != nil {
	// 	return "", auth.ErrUserNotFound
	// }

	user := &User{
		ID:       "1",
		username: inp.Username,
		password: fmt.Sprintf("%x", pwd.Sum(nil)),
	}
	log.Println(user.username)
	log.Println()
	log.Println(time.Now())
	log.Println(time.Now().Add(86400))
	seconds := 86400 * time.Second
	claims := AuthClaims{
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(seconds)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, signInResponse{Token: tokenString})
}

func getConsumption(c *gin.Context) {

	c.Header("Content-Type", "application/json")
	type Consumption struct {
		LastUpdate time.Time `json:"lastupdate"`
		Heating    float64   `json:"heating"`
	}

	type Response struct {
		Status      int           `json:"status"`
		Message     string        `json:"message"`
		Consumption []Consumption `json:"consumption"`
	}

	db, err := sql.Open("mysql", "gouser:G1nW3bUs3r!@tcp(192.168.178.35:3306)/home?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, err := db.Query(`SELECT 
	DATE(lastupdate) AS lastupdate,
	MAX(heating) - MIN(heating) AS heating	
FROM 
	consumption
WHERE 
	MONTH(lastupdate) = MONTH(CURRENT_DATE()) 
	AND YEAR(lastupdate) = YEAR(CURRENT_DATE())    
GROUP BY 
	DATE(lastupdate)
ORDER BY 
	lastupdate;`)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var consumptions []Consumption
	for rows.Next() {
		var consumption Consumption
		if err := rows.Scan(&consumption.LastUpdate, &consumption.Heating); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		consumptions = append(consumptions, consumption)
	}

	var response Response
	response.Status = 200
	response.Message = "OK"
	response.Consumption = consumptions

	if err := rows.Err(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}

func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "submit",
	})
}

func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "read",
	})
}

func main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/consumption", getConsumption)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
		v1.POST("/login", SignIn)
	}

	router.Run(":8080")
}
