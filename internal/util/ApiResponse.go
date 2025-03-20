package util

import "github.com/gofiber/fiber/v2"


type ApiResponse struct {
	Message   string `json:"message"`
	Status    int    `json:"status"`
	Data      any    `json:"data,omitempty"`
	Errors    any    `json:"errors,omitempty"`
	IsSuccess bool   `json:"isSuccess"`
}

// Error implements error.
func (a ApiResponse) Error() string {
	panic("unimplemented")
}

func SuccessApiResponse(c *fiber.Ctx , message string, status int, data any) error {
	return c.Status(status).JSON(ApiResponse{
		Message:   message,
		Status:    status,
		Data:      data,
		IsSuccess: true,
	})
}

func SuccessApiResponseWithoutData(c *fiber.Ctx , message string, status int) error {
	return c.Status(status).JSON(ApiResponse{
		Message:   message,
		Status:    status,
		IsSuccess: true,
	})
}
func ErrorApiResponse(c *fiber.Ctx , message string, status int, errors any) error {
	return c.Status(status).JSON(ApiResponse{
		Message:   message,
		Status:    status,
		Errors:      errors,
		IsSuccess: false,
	})
}
func ErrorApiResponseWithoutData(c *fiber.Ctx, message string, status int) error {
	return c.Status(status).JSON(ApiResponse{
		Message:   message,
		Status:    status,
		IsSuccess: false,
	})
}
