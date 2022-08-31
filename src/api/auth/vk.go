package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Vkontakte(c echo.Context) error {

	code := c.QueryParam("code")

	if len(code) == 0 {

		return invalidQueryParam(c)

	}

	res, err := http.Get(os.Getenv("vk_get_email") + code)

	if err != nil {

		return err

	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {

		return err

	}

	data := &VK_DATA{}

	json.Unmarshal(body, &data)

	if data.AT == "" || data.USER_ID <= 0 || data.EXPIRES <= 0 || data.EMAIL == "" {

		return invalidQueryParam(c)

	}

	smtpEmail(data.EMAIL, "confirm", c)

	return c.JSON(202, &Message{message: "Please, check your email address"})

}

func invalidQueryParam(c echo.Context) error {

	return c.JSON(402, &Message{message: "Invalid query parameters"})

}
