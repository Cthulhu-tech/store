package auth

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
)

func Vkontakte(c echo.Context) error {

	code := c.QueryParam("code")

	if len(code) > 0 {

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

		if len(data.AT) == 0 || len(data.ID) == 0 || len(data.EXPIRES) == 0 || len(data.Email) == 0 {

			return invalidQueryParam()

		}

		fmt.Println(data.AT, data.USER_ID, data.EXPIRES, data.EMAIL)

		return c.Redirect(302, os.Getenv("URL_SERVER"))

	}

	return invalidQueryParam()

}

func invalidQueryParam() {

	return c.JSON(402, &Message{message: "Invalid query parameters"})

}
