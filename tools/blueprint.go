package tools

import (
	"fmt"
	"log"
	"io/ioutil"
	"net/http"
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/Kasitaw/go-sms/configs"
)

type Body struct {
	Phone string `json:"phone,omitempty"`
	Message string `json:"message,omitempty"`
}

type SmsInterface interface {
	Send(d DataObject)
}

type DataObject struct {
	Body Body
	Config configs.Config
	Context *gin.Context
}

type Isms struct {}

type Sms123 struct {}

type Sms123Response struct {
	Status string `json:"status,omitempty"`
	Code string `json:"msgCode,omitempty"`
	Message string `json:"statusMsg,omitempty"`
}

func (s *Isms) Send(d DataObject) {
}

func (s *Sms123) Send(d DataObject) {
	var url string
	var jsonResp Sms123Response

	url = fmt.Sprintf("%s?apiKey=%s&recipients=%s&messageContent=%s", d.Config.Url, d.Config.Password, d.Body.Phone, d.Body.Message)
	resp, err := http.Get(url);
	if  err != nil {
		d.Context.JSON(403, gin.H{
			"status": "failed",
			"message": jsonResp.Message,
		})
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if  err != nil {
		log.Fatalf("Error parsing to body: %v", err)
	}

	err = json.Unmarshal(body, &jsonResp)
	if err != nil {
		log.Fatalf("Error marshalling json body: %v", err)
	}

	if jsonResp.Status == "error" {
		d.Context.JSON(403, gin.H{
			"status": "failed",
			"message": jsonResp.Message,
		})
	} else {
		d.Context.JSON(200, gin.H{
			"status": "success",
			"message": "Successfully send the SMS",
		})
	}
}



