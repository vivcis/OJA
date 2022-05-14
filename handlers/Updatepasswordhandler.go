package handlers

import (
	"github.com/decadevs/shoparena/database"
	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
)

type Handler struct {
	DB   database.DB
	Mail database.Mailer
}

func PingHandler(c *gin.Context) {
	// healthcheck
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (h *Handler) SearchProductHandler(c *gin.Context) {
	//Equivalent to param
	category := c.Query("category")
	lowerPrice := c.Query("lower-price")
	upperPrice := c.Query("upper-price")
	name := c.Query("name")

	product, err := h.DB.SearchProduct(lowerPrice, upperPrice, category, name)
	if err != nil {
		log.Println("handler error in search product", err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	if len(product) == 0 {
		c.JSON(http.StatusInternalServerError, "no such product")
		return
	}

	c.JSON(http.StatusFound, product)
}

func (h *Handler) SendForgotPasswordEMailHandler(c *gin.Context) {
	// crete a password reset struct and initialize it

	var forgotPassword models.ResetPasswordRequest

	err := c.BindJSON(&forgotPassword)
	if err != nil {
		c.JSON(400, gin.H{"message": "please fill all fields"})
		c.Abort()
		return
	}
	buyer, err := h.DB.FindBuyerByEmail(forgotPassword.Email)
	// return 404 status if there is an error getting the buyer
	if err != nil {
		c.JSON(404, gin.H{"message": "buyer not found"})
		c.Abort()
		return
	}

	// generate token that'll be used to reset the password
	secretString := os.Getenv("JWTSECRET")
	resetToken, err := h.Mail.GenerateNonAuthToken(buyer.Email, secretString)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "something went wrong"})
		c.Abort()
		return
	}

	// the link to be clicked in order to perform password reset
	link := "http://localhost:8085/reset-password?reset_token=" + *resetToken
	// define the body of the email
	body := "Here is your reset <a href='" + link + "'>link</a>"
	html := "<strong>" + body + "</strong>"

	//intialize the email sendout
	privateAPIKey := os.Getenv("MAILGUN_API_KEY")
	yourDomain := os.Getenv("DOMAIN_STRING")
	err = h.Mail.SendMail("forgot Password", html, buyer.Email, privateAPIKey, yourDomain)

	//if email was sent return 200 status code
	if err == nil {
		c.JSON(200, gin.H{"message": "please check your email for password reset link"})
		c.Abort()
		return
	} else {
		log.Println(err)
		c.JSON(500, gin.H{"message": "something went wrong while trying to send you a mail, please try again"})
		c.Abort()
		return
	}

}

func (h *Handler) ForgotPasswordResetHandler(c *gin.Context) {
	// instance of the forgot password request
	var password models.ForgotPassword
	// requires all fields aand returns an error if some fields are missing
	err := c.BindJSON(&password)
	if err != nil {
		// Return response when we some fields are missing
		log.Println(err)
		c.JSON(406, gin.H{"message": "please provide all fields"})
		c.Abort()
		return
	}

	// verifies if the new password string and the string confirmation string match
	if password.NewPassword != password.ConfirmNewPassword {
		// Return response when passwords mismatch
		c.JSON(400, gin.H{"message": "password mismatch"})
		c.Abort()
		return
	}

	// gets the reset token from the dynamic route
	resetToken, _ := c.GetQuery("reset_token")

	// decodes token to get user email
	secretString := os.Getenv("JWTSECRET")
	userEmail, err := h.Mail.DecodeToken(resetToken, secretString)
	if err != nil {
		// Return response when we get an error while decoding the token
		log.Println(err)
		c.JSON(500, gin.H{"message": "Something wrong happened, Please try again later"})
		c.Abort()
		return
	}
	// check the db for a buyer with the email decoded from the token
	buyer, err := h.DB.FindBuyerByEmail(userEmail)
	if err != nil {
		// Return response when we get an error while fetching buyer
		log.Println(err)
		c.JSON(500, gin.H{"message": "buyer not found"})
		c.Abort()
		return
	}

	// generates a new password hash from the password string in the password reset request
	newPasswordHash, err := bcrypt.GenerateFromPassword([]byte(password.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		// Return response when we get an error while hatching the password
		log.Println(err)
		c.JSON(500, gin.H{"message": "internal server error"})
		c.Abort()
		return
	}

	// resets the buyer password hash by email in the database
	_, err = h.DB.BuyerResetPassword(buyer.Email, string(newPasswordHash))
	if err != nil {
		// Return response if we are not able to update user password
		c.JSON(500, gin.H{"message": "Somehting happened while updating your password try again"})
		c.Abort()
		return
	}

	// returns a response after successfully resetting the password
	c.JSON(200, gin.H{"message": "password reset successful"})
	c.Abort()
	return
}

func (h *Handler) SellerUpdatePasswordHandler(c *gin.Context) {
	var password models.PasswordResetReq
	email := c.Param("email")
	seller, err := h.DB.FindSellerByEmail(email)
	if err != nil {
		c.JSON(404, gin.H{"message": "User not found"})
		c.Abort()
		return
	}
	err = c.BindJSON(&password)
	if err != nil {
		c.JSON(406, gin.H{"message": "Provide all Relevant Fields"})
		c.Abort()
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(seller.PasswordHash), []byte(password.OldPassword))
	if err != nil {
		c.JSON(500, gin.H{"message": "OldPassword is incorrect"})
		c.Abort()
		return
	}

	if password.NewPassword != password.ConfirmNewPassword {
		c.JSON(400, gin.H{"message": "Password Mismatch, Please make sure newpassword & confirmNewPassword match"})
		c.Abort()
		return
	}
	passwordhash, err := bcrypt.GenerateFromPassword([]byte(password.NewPassword), bcrypt.DefaultCost)
	_, err = h.DB.SellerUpdatePassword(seller.PasswordHash, string(passwordhash))
	if err != nil {
		log.Println(err)
		return
	}
	c.JSON(200, gin.H{"message": "successfully reset password"})
}

func (h *Handler) BuyerUpdatePasswordHandler(c *gin.Context) {
	var password models.PasswordResetReq

	email := c.Param("email")

	err := c.BindJSON(&password)
	if err != nil {
		c.JSON(406, gin.H{"message": "Provide all Relevant Fields"})
		c.Abort()
		return
	}

	if password.NewPassword != password.ConfirmNewPassword {
		c.JSON(400, gin.H{"message": "Password Mismatch"})
		c.Abort()
		return
	}

	buyer, err := h.DB.FindBuyerByEmail(email)
	if err != nil {
		c.JSON(500, gin.H{"message": "Internal Server Error"})
		c.Abort()
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(buyer.PasswordHash), []byte(password.OldPassword))
	if err != nil {
		c.JSON(500, gin.H{"message": "OldPassword is incorrect"})
		c.Abort()
		return
	}

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(password.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = h.DB.BuyerUpdatePassword(buyer.PasswordHash, string(passwordHash))
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{"message": "internal server error"})
		c.Abort()
		return
	}
	c.JSON(200, gin.H{"message": "successfully reset password"})
}
