package controllers

import (
	"fmt"
	"net/http"
	"kakao/models"
	"github.com/labstack/echo"
	df "github.com/meinside/dialogflow-go"
)

//Dialogflow set info
const (
	token     = "--------------------------------" //Client access token
	sessionId = "--------------------------------"
	timezone  = "Asia/Seoul"
)
func KakaoKeyboard() echo.HandlerFunc {
	return func(c echo.Context) error {
		kb := models.Keyboard{
			Type: "text",
		}
		return c.JSON(http.StatusOK, kb)
	}
}

func KakaoMessage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var msg models.KakaoMessage
		req := new(models.KakaoMessageRequest)
		if err := c.Bind(req); err != nil {
			return err
		}

		client := df.NewClient(token)
		client.Verbose = true // for verbose messages

		// query text
		if response, err := client.QueryText(df.QueryRequest{
			Query:     []string{req.Content},
			SessionId: sessionId,
			Language:  df.Korean,
			Timezone: timezone,

		}); err == nil {
			fmt.Printf(">>> response = %+v\n", response)
			msg = models.KakaoMessage{
				//Message: map[string]string{"text": "User_key="+request.User_key+" Type="+request.Type+" Content="+request.Content + " entity.Name="+response.Result.Fulfillment.Speech},
				Message: map[string]string{"text": string(response.Result.Fulfillment.Speech)},
				Keyboard: map[string]string{"type": "text"},
			}
		} else {
			fmt.Printf("*** error: %s\n", err)
		}
		return c.JSON(http.StatusOK, msg)
	}
}
