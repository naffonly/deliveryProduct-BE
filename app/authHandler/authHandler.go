package authHandler

import (
	"deliveryProduct/model/domain"
	"deliveryProduct/utils"
	"deliveryProduct/utils/token"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"net/http"
)

var (
	g               = galidator.New()
	customizer      = g.Validator(domain.RegistrationDTO{})
	customizerLogin = g.Validator(domain.LoginDto{})
)

type Handler struct {
	DB *gorm.DB
}

func NewHandlerUser(db *gorm.DB) Handler {
	return Handler{DB: db}
}
func (h *Handler) Register(c *gin.Context) {
	var payload domain.RegistrationDTO
	var user domain.User

	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": customizer.DecryptErrors(err),
		})
		c.Abort()
		return
	}

	if rs := h.DB.Where("email=?", payload.Email).First(&user); rs.RowsAffected > 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Email already exist",
		})
		return
	}

	if payload.Password != payload.PasswordConfirm {
		c.JSON(http.StatusBadGateway, gin.H{
			"message": "password not match",
		})
		c.Abort()
		return
	}

	hashedPassword, err := utils.HashPassword(payload.Password)
	if err != nil {
		c.JSON(http.StatusBadGateway, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	newPayload := domain.User{
		Username:      payload.Username,
		Email:         payload.Email,
		Password:      hashedPassword,
		AddressOffice: payload.AddressOffice,
	}
	h.DB.Create(&newPayload)
	c.JSON(http.StatusOK, gin.H{
		"Message": "Success Register",
		"data":    newPayload,
	})
}
func (h *Handler) Login(c *gin.Context) {
	var payload domain.LoginDto

	err := c.BindJSON(&payload)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": customizerLogin.DecryptErrors(err),
		})
	}

	newPayload := domain.User{
		Username: payload.Username,
		Password: payload.Password,
	}

	loginToken, loginErr := h.loginCheck(newPayload.Username, newPayload.Password)

	if loginErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "username or password is incorrect."})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"Message": "Success Login",
		"token":   loginToken,
	})

}

func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (h *Handler) loginCheck(username string, password string) (string, error) {

	payload := domain.User{}
	err := h.DB.Model(domain.User{}).Where("username =? ", username).Take(&payload).Error
	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, payload.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(payload.ID)

	if err != nil {
		return "", err
	}

	return token, nil
}
