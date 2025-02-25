package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"task/database"
	"task/models"
	"time"
)

func Signup(c *gin.Context) {
	var Body struct {
		Email    string
		Password string
	}

	if c.Bind(&Body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Bad Request",
		})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(Body.Password), 10)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to hash password",
		})
		return
	}

	user := models.User{Email: Body.Email, Password: string(hash)}
	result := database.Db.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, models.Report{
			Status:  "Error",
			Message: "Failed to create user",
		})
		return
	}
	c.JSON(http.StatusOK, models.Report{
		Status:  "Success!",
		Message: "User created successfully",
	})
}

func Login(c *gin.Context) {
	var Body struct {
		Email    string
		Password string
	}
	if c.Bind(&Body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Bad Request",
		})
		return
	}

	var user models.User
	database.Db.First(&user, "email = ?", Body.Email)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"Error": "Invalid email or password",
		})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(Body.Password))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Invalid email or password",
		})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": "Failed to create token",
		})

		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "/", "", false, true)

	c.JSON(http.StatusOK, models.Report{
		Status:  "Success!",
		Message: "User logged successfully",
	})
}

func Validate(c *gin.Context) {
	user, _ := c.Get("user")

	c.JSON(http.StatusOK, gin.H{
		"Message": user,
	})
}
