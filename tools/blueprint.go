package tools

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/Kasitaw/go-sms/configs"
	"github.com/gin-gonic/gin"
)

// Error code used in Isms.
const (
	e1000 = "Unknown error. Please contact the administrator."
	e1001 = "Authentication failed. Your username or password are incorrect."
	e1002 = "Account suspended or expired. Your account has been expired or suspended. Please contact the administrator."
	e1003 = "Ip not allowed. Your IP is not allowed to send SMS. Please contact the administrator."
	e1004 = "Insufficient credits. You have run our of credits. Please reload your credits."
	e1005 = "Invalid sms type. Your SMS type is not supported."
	e1006 = "Invalid body length. Your SMS body has exceed the length.Max limit = 900"
	e1007 = "Invalid hex body. Your Hex body format is wrong."
	e1008 = "Missing parameter. One or more required parameters are missing."
	e1009 = "Invalid destination number. Invalid number."
	e1012 = "Invalid message type. Message contain unicode and please use type=2 for Unicode."
	e1013 = "Invalid term and agreement. Please add agreedterm=YES in your API"
	s000 = "Successfully send message to destination number"
	s2000 = "Successfully send message to destination number"
)

// Payload received from client side.
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

type IsmsError struct {
	Code map[string]string
}

type Isms struct {
	IsmsError
}

type Sms123 struct {}

// Error response received from Sms123
type Sms123Response struct {
	Status string `json:"status,omitempty"`
	Code string `json:"msgCode,omitempty"`
	Message string `json:"statusMsg,omitempty"`
}

// Error/Success code used in Isms.
func (s *Isms) errCode() map[string]string {
	s.Code = map[string]string {
		"1000": e1000,
		"1001": e1001,
		"1002": e1002,
		"1003": e1003,
		"1004": e1004,
		"1005": e1005,
		"1006": e1006,
		"1007": e1007,
		"1008": e1008,
		"1019": e1009,
		"1012": e1012,
		"1013": e1013,
		"000": s000,
		"2000": s2000,
 	}
 	return s.Code
}

// Get error code for Isms.
func (s *Isms) getCode(code string) string {
	errCode := s.errCode()
	return errCode[code]
}

// Isms implementation of "send" method.
func (s *Isms) Send(d DataObject) {
	url := fmt.Sprintf("%s?un=%s&pwd=%s&dstno=%s&msg=%s&agreedterm=yes&type=1",
		d.Config.Url,
		url.QueryEscape(d.Config.Username),
		url.QueryEscape(d.Config.Password),
		d.Body.Phone,
		url.QueryEscape(d.Body.Message),
	)
	resp, err := http.Get(url);
	if  err != nil {
		log.Fatalf("Error request to service endpoint: %v", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if  err != nil {
		log.Fatalf("Error parsing to body: %v", err)
	}

	if smsResp := fmt.Sprintf("%s", body); smsResp != "" {
		code := strings.TrimSpace(smsResp[1:5])
		status := "failed"
		respCode := 403
		if code == "000" {
			status = "success"
			respCode = 200
		}
		d.Context.JSON(respCode, gin.H{
			"status": status,
			"message": s.getCode(code),
		})
	} else {
		d.Context.JSON(200, gin.H{
			"status": "success",
			"message": "Successfully send the SMS",
		})
	}
}

// Sms123 implementation of "send" method.
func (s *Sms123) Send(d DataObject) {
	var jsonResp Sms123Response

	url := fmt.Sprintf("%s?apiKey=%s&recipients=%s&messageContent=%s",
		d.Config.Url,
		d.Config.Password,
		d.Body.Phone,
		url.QueryEscape(d.Body.Message),
	)
	resp, err := http.Get(url);
	if  err != nil {
		log.Fatalf("Error request to service endpoint: %v", err)
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



