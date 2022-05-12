package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/decadevs/shoparena/database"
	"github.com/decadevs/shoparena/models"
	"github.com/decadevs/shoparena/services"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	DB database.DB
}

// Server serves requests to DB with router
type Server struct {
	dB database.DB
	// Router *router.Router
} 

func PingHandler(c *gin.Context) {
	// healthcheck
	c.JSON(200, gin.H{
		"message": "pong",
	})
}


func JSON(c *gin.Context, message string, status int, data interface{}, errs []string) {
	responsedata := gin.H{
		"message": message,
		"data":    data,
		"errors":  errs,
		"status":  http.StatusText(status),
	}

	c.JSON(status, responsedata)
}

func (h *Handler) UpdateProfileHandler(c *gin.Context){
	details := &struct{
       Name string `json:"user_name"`
	   Email string `json:"email"`
	   NewEmail string `json:"new_email"`

	}{}
				if errs := c.BindJSON(details); errs != nil {
					JSON(c, "", http.StatusBadRequest, nil, []string{"errs"})
					return
				}
				username := details.Name
				buyerDetail, err := h.DB.FindBuyerByEmail(details.Email)
				if err != nil{
					log.Println("buyer not found", err)
					JSON(c, "", http.StatusBadRequest, nil, []string{"internal server error"})
					return
				}
				log.Print(buyerDetail.User)
				buyerDetail.User.Username = username
				buyerDetail.User.Email = details.NewEmail
				if err := h.DB.UpdateUser(buyerDetail, details.Email); err != nil {
					log.Printf("update user error : %v\n", err)
					JSON(c, "", http.StatusInternalServerError, nil, []string{"update user error"})
					return
				}
				JSON(c, "user updated successfuly", http.StatusOK, nil, nil)
				
		// if user1, exists := c.Get("user"); exists {
		// }
		
	}


func (h *Handler) SearchProductHandler(c *gin.Context) {
	//Equivalent to param
	categoryName := c.Query("category")
	lowerPrice := c.Query("lower-price")
	upperPrice := c.Query("upper-price")
	name := c.Query("name")

	product, err := h.DB.SearchProduct(lowerPrice, upperPrice, categoryName, name)
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


// handleUploadProfilePic uploads a user's profile picture
func (h *Handler) UploadImageHandler(c *gin.Context) {
    
        if userI, exists := c.Get("user"); exists {
            if user, ok := userI.(*models.User); ok {
                const maxSize = int64(2048000) // allow only 2MB of file size
                r := c.Request
                err := r.ParseMultipartForm(maxSize)
                if err != nil {
                    log.Printf("parse image error: %v\n", err)
                    JSON(c, "", http.StatusBadRequest, nil, []string{"image too large"})
                    return
                }
                file, fileHeader, err := r.FormFile("profile_picture")
                if err != nil {
                    log.Println("error getting profile picture", err)
                    JSON(c, "", http.StatusBadRequest, nil, []string{"image not supplied"})
                    return
                }
                defer file.Close()
                fileExtension, ok := services.CheckSupportedFile(strings.ToLower(fileHeader.Filename))
                log.Println(filepath.Ext(strings.ToLower(fileHeader.Filename)))
                fmt.Println(fileExtension)
                if ok {
                    log.Println(fileExtension)
                    JSON(c, "", http.StatusBadRequest, nil, []string{fileExtension + " image file type is not supported"})
                    return
                }
                session, tempFileName, err := services.PreAWS(fileExtension, "profile_picture")
                if err != nil {
                    log.Printf("could not upload file: %v\n", err)
                }
                url, err := h.DB.UploadFileToS3(session, file, tempFileName, fileHeader.Size)
                if err != nil {
                    log.Println(err)
                    JSON(c, "", http.StatusInternalServerError, nil, []string{"an error occured while uploading the image"})
                    return
                }
                user.Image = url
                err = h.DB.UpdateUserImageURL(user.Username, user.Image)
                if err != nil {
                    log.Println(err)
                    JSON(c, "", http.StatusInternalServerError, nil, []string{"an error occured while uploading the image"})
                    return
                }
                JSON(c, "successfully created file", http.StatusOK, gin.H{
                    "imageurl": user.Image,
                }, nil)
                return
            }
        }
        JSON(c, "", http.StatusUnauthorized, nil, []string{"unable to retrieve authenticated user"})
        
}

