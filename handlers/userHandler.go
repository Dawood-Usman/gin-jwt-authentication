package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/dawood-usman/go-ops/config"
	"github.com/dawood-usman/go-ops/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func SignUp(c *gin.Context) {
	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	bcryptedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
		return
	}

	user.Password = string(bcryptedPassword)

	if err := config.DB.Create(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, user)
}
func Login(c *gin.Context) {
	var userLoginBody struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if err := c.ShouldBindJSON(&userLoginBody); err != nil {
		c.JSON(http.StatusBadGateway, gin.H{"error": "Invalid request body"})
		return
	}

	var user models.User
	if err := config.DB.First(&user, "email = ?", userLoginBody.Email).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email!"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userLoginBody.Password)); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Password!"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Couldn't create jwt token!"})
		return
	}

	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	c.JSON(http.StatusOK, gin.H{"status": "User Logged In Successfully!"})
}

func Validate(c *gin.Context) {
	user := c.MustGet("user").(models.User)

	c.JSON(http.StatusOK, gin.H{user.SubDomain + ".dawoodworld.com": user})
}