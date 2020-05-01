package main

import (
	"github.com/Kasitaw/go-sms/configs"
	"github.com/Kasitaw/go-sms/tools"
	"github.com/gin-gonic/gin"
)

func main() {
	var provider configs.Provider
	var driver configs.Config
	var sms tools.SmsInterface
	var dataObject tools.DataObject

	// Set to release/production mode
	gin.SetMode(gin.ReleaseMode)

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()
	router.POST("/send", func(c *gin.Context) {
		phone := c.PostForm("phone")
		message := c.PostForm("message")

		body := tools.Body{
			Phone: phone,
			Message: message,
		}

		provider = configs.Parse()
		driver = configs.GetDriver(provider.Drivers, provider.Default)

		// Decide sms provider based on default setting
		switch provider.Default {
		case "sms123":
			sms = &tools.Sms123{}
		default:
			sms = &tools.Isms{}
		}

		dataObject = tools.DataObject{
			Body: body,
			Config: driver,
			Context: c,
		}
		// Send message to destination number.
		sms.Send(dataObject)
	})

	router.Run(":2346")
}