package services

import (
	"errors"
	"myapp/config"
	"myapp/internal/email"
	"myapp/internal/model"
	"myapp/internal/util"
	"myapp/internal/validation"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreateUser(c *fiber.Ctx) error {
	// Start overall time measurement

	// Measure time taken to parse the request body
	cu := new(validation.CreateUserRequest)
	err := c.BodyParser(cu)

	if err != nil {
		return util.ErrorApiResponse(c, "Error parsing request body", 400, err.Error())
	}	
	
	// Measure time taken for validation
	validate := validator.New()
	e := validate.Struct(cu)

	if e != nil {
		errors := make(map[string]string)
		for _, v := range e.(validator.ValidationErrors){
			errors[v.Field()] = v.Tag()
		}
		return util.ErrorApiResponse(c, "Error validating request body", 400 , errors)
	}

	// Measure time taken for hashing the password
	hashedPassword , _ := util.HashPassword(cu.Password)

	// Prepare user for database insertion
	user := model.User{
		Name:     cu.Name,
		Email:    cu.Email,
		Password: hashedPassword,
	}

	// Measure time taken to save user to DB
	db := config.GetDB()
	result := db.Create(&user)

	// If there's an error during DB insertion, handle it
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) ||  result.Error.Error() == "ERROR: duplicate key value violates unique constraint \"idx_users_email\" (SQLSTATE 23505)"  {
			return util.ErrorApiResponseWithoutData(c, "User already exists", 400)
		}
		return util.ErrorApiResponseWithoutData(c, "Error creating user", 500)
	}

	// End overall time measurement

	// Log the time taken for each step
	
	// SendMail(user.Email , )
	go email.SendEmail(user.Email ,"Welcome to our platform" , email.GenerateWelcomeEmail(user.Name , "http://youtube.com"))
	// Return the successful response
	go CreateEmailRecord(user.Id , user.Email)
	return c.Status(201).JSON(util.ApiResponse{
		Message:   "Success",
		Status:    201,
		IsSuccess: true,
	})
}




func CheckPassword(c *fiber.Ctx) error {
	db := config.GetDB()
	var user model.User
	result := db.Where("email = ?", c.Query("email")).First(&user)
	if result.Error != nil {
		return util.ErrorApiResponse(c, "Error getting user", 500, result.Error)
	}
	match, _ := util.VerifyPassword(user.Password, c.Query("password"))
	if match {
		return util.SuccessApiResponse(c, "Password matches", 200, nil)
	}
	return util.ErrorApiResponseWithoutData(c, "Password does not match", 400)
}




func GetUsers(c *fiber.Ctx) error {
	db := config.GetDB()
	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		return util.ErrorApiResponse(c, "Error getting users", 500, result.Error)
	}
	return util.SuccessApiResponse(c, "Users retrieved successfully", 200, users)
}




func CreateEmailRecord(user uint , sendto string ){
	// var email model.EmailRecord
	em := model.EmailRecord{
		UserID: user,
		SendTo: sendto,
		Status: model.Delivered,
	} 
	db := config.GetDB()
	db.Create(&em)
}


func SendEmailToUser(c *fiber.Ctx)error{
	var cu validation.SendEmailRequest
	err := c.BodyParser(&cu)

	if err != nil {
		return util.ErrorApiResponse(c, "Error parsing request body", 400, err.Error())
	}	
	
	// Measure time taken for validation
	validate := validator.New()
	e := validate.Struct(cu)
	if e != nil {
		errors := make(map[string]string)
		for _, v := range e.(validator.ValidationErrors){
			errors[v.Field()] = v.Tag()
		}
		return util.ErrorApiResponse(c, "Error validating request body", 400 , errors)
	}
	er :=model.EmailRecord{
		UserID: cu.UserID,
		SendTo: cu.Email,
		Status: model.Delivered,
	}
	db:=config.GetDB()
	db.Create(&er)
	go email.SendEmail(cu.Email , cu.Subject  , cu.Content)
	return util.SuccessApiResponseWithoutData(c , "Email will send after sometime" , 200)
}
func SendEmailToMultipleUser(c *fiber.Ctx) error {
	// Parse the request body into the SendEmailRequest struct (Array/Slice)
	var emailRequests []validation.SendEmailRequest
	err := c.BodyParser(&emailRequests)
	if err != nil {
			return util.ErrorApiResponse(c, "Error parsing request body", 400, err.Error())
	}

	// Validate the request
	validate := validator.New()
	for _, req := range emailRequests {
			e := validate.Struct(req)
			if e != nil {
					errors := make(map[string]string)
					for _, v := range e.(validator.ValidationErrors) {
							errors[v.Field()] = v.Tag()
					}
					return util.ErrorApiResponse(c, "Error validating request body", 400, errors)
			}
	}

	// Process each email request
	db := config.GetDB()
	for _, req := range emailRequests {
			// Send the email asynchronously
			go email.SendEmail(req.Email, req.Subject, req.Content)

			// Store the email record in the database
			er := model.EmailRecord{
					UserID:  req.UserID,
					SendTo:  req.Email,
					Status:  model.Delivered,
			}
			db.Create(&er)
	}

	// Return success response
	return util.SuccessApiResponseWithoutData(c, "Emails will be sent after some time", 200)
}
