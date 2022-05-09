package handlers

import (
	"bytes"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/decadevs/shoparena/database"
	servererrors "github.com/decadevs/shoparena/handlers/serverErrors"
	"github.com/decadevs/shoparena/models"
	"github.com/gin-gonic/gin"
	validator "github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func PingHandler(c *gin.Context) {
	// healthcheck
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Server serves requests to DB with router
type Server struct {
	dB database.DB
	// Router *router.Router
}

// decode decodes the body of c into v
func (s *Server) decode(c *gin.Context, v interface{}) []string {
	if err := c.ShouldBindJSON(v); err != nil {
		errs := []string{}
		verr, ok := err.(validator.ValidationErrors)
		if ok {
			for _, fieldErr := range verr {
				errs = append(errs, servererrors.NewFieldError(fieldErr).String())
			}
		} else {
			errs = append(errs, "internal server error")
		}
		return errs
	}
	return nil
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

func (s *Server) handleUpdateUserDetails() gin.HandlerFunc {
	return func(c *gin.Context) {

		if user1, exists := c.Get("user"); exists {
			if user, ok := user1.(*models.User); ok {
				username, email := user.Username, user.Email
				if errs := s.decode(c, user); errs != nil {
					JSON(c, "", http.StatusBadRequest, nil, errs)
					return
				}

				user.Username, user.Email = username, email
				user.Model.UpdatedAt = time.Now()
				if err := s.dB.UpdateUser1(user); err != nil {
					log.Printf("update user error : %v\n", err)
					JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
					return
				}
				JSON(c, "user updated successfuly", http.StatusOK, nil, nil)
				return
			}
		}
		log.Printf("can't get user from context\n")
		JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
	}
}

//handleUpdateUserDetails1 handles update user details
func (s *Server) handleUpdateUserDetails1() gin.HandlerFunc {
	return func(c *gin.Context) {

		if userI, exists := c.Get("user"); exists {
			if user, ok := userI.(*models.User); ok {
				var update models.UpdateUser
				email := user.Email
				log.Println(user)
				if errs := s.decode(c, &update); errs != nil {
					JSON(c, "", http.StatusBadRequest, nil, errs)
					return
				}

				user.Email = email
				user.UpdatedAt = time.Now()
				if err := s.dB.UpdateUser(user.Username, &update); err != nil {
					log.Printf("update user error : %v\n", err)
					JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
					return
				}
				JSON(c, "user updated successfuly", http.StatusOK, nil, nil)
				return
			}
		}
		log.Printf("can't get user from context\n")
		JSON(c, "", http.StatusInternalServerError, nil, []string{"internal server error"})
	}
}

// handleUploadProfilePic uploads a user's profile picture
func (s *Server) handleUploadProfilePic() gin.HandlerFunc {
	return func(c *gin.Context) {

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
					log.Println(err)
					JSON(c, "", http.StatusBadRequest, nil, []string{"image not supplied"})
					return
				}
				defer file.Close()

				supportedFileTypes := map[string]bool{
					".png":  true,
					".jpeg": true,
					".jpg":  true,
				}
				fileExtension := filepath.Ext(fileHeader.Filename)
				if !supportedFileTypes[fileExtension] {
					log.Println(fileExtension)
					JSON(c, "", http.StatusBadRequest, nil, []string{fileExtension + " image file type is not supported"})
					return
				}
				// create a unique file name for the file
				tempFileName := "profile_pics/" + uuid.NewString() + fileExtension

				session, err := session.NewSession(&aws.Config{
					Region: aws.String(os.Getenv("AWS_REGION")),
					Credentials: credentials.NewStaticCredentials(
						os.Getenv("AWS_SECRET_ID"),
						os.Getenv("AWS_SECRET_KEY"),
						os.Getenv("AWS_TOKEN"),
					),
				})
				if err != nil {
					log.Printf("could not upload file: %v\n", err)
				}

				err = uploadFileToS3(session, file, tempFileName, fileHeader.Size)
				if err != nil {
					log.Println(err)
					JSON(c, "", http.StatusInternalServerError, nil, []string{"an error occured while uploading the image"})
					return
				}

				user.Image = os.Getenv("S3_BUCKET") + tempFileName
				if err = s.dB.UpdateUser1(user); err != nil {
					log.Println(err)
					JSON(c, "", http.StatusInternalServerError, nil, []string{"unable to update user's profile pic"})
					return
				}

				JSON(c, "successfully created file", http.StatusOK, gin.H{
					"imageurl": user.Image,
				}, nil)
				return
			}
		}
		JSON(c, "", http.StatusUnauthorized, nil, []string{"unable to retrieve authenticated user"})
		return
	}
}

// UploadFileToS3 saves a file to aws bucket and returns the url to the file and an error if there's any
func uploadFileToS3(s *session.Session, file multipart.File, fileName string, size int64) error {
	// get the file size and read the file content into a buffer
	buffer := make([]byte, size)
	file.Read(buffer)

	// config settings: this is where you choose the bucket,
	// filename, content-type and storage class of the file you're uploading
	_, err := s3.New(s).PutObject(&s3.PutObjectInput{
		Bucket:               aws.String(os.Getenv("S3_BUCKET_NAME")),
		Key:                  aws.String(fileName),
		ACL:                  aws.String("public-read"),
		Body:                 bytes.NewReader(buffer),
		ContentLength:        aws.Int64(int64(size)),
		ContentType:          aws.String(http.DetectContentType(buffer)),
		ContentDisposition:   aws.String("attachment"),
		ServerSideEncryption: aws.String("AES256"),
		StorageClass:         aws.String("INTELLIGENT_TIERING"),
	})
	return err
}

// MakeBucket creates a bucket.
// Inputs:
//     sess is the current session, which provides configuration for the SDK's service clients
//     bucket is the name of the bucket
// Output:
//     If success, nil
//     Otherwise, an error from the call to CreateBucket
func MakeBucket(sess *session.Session, bucket *string) error {
	svc := s3.New(sess)

	_, err := svc.CreateBucket(&s3.CreateBucketInput{
		Bucket: bucket,
	})
	if err != nil {
		return err
	}

	err = svc.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: bucket,
	})
	if err != nil {
		return err
	}
	return nil
 
}
