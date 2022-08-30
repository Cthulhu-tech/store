package auth

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"fmt"
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

		fmt.Println(data.AT, data.USER_ID, data.EXPIRES, data.EMAIL)

		return c.Redirect(302, os.Getenv("URL_SERVER"))

	}

	return c.JSON(402, &Error{message: "Invalid query parameters"})

}
