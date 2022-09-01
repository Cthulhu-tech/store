package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Cthulhu-tech/store/src/utils/message"
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

	if data.AT == "" || data.USER_ID <= 0 || data.EXPIRES <= 0 || !isEmailValid(data.EMAIL) {

		return invalidQueryParam(c)

	}

	defer func() {

		go smtpEmail(data.EMAIL, "confirm", c)

	}()

	return message.JSON(c, 202, "Please, check your email address")

}

func invalidQueryParam(c echo.Context) error {

	return message.JSON(c, 402, "Invalid query parameters")

}
