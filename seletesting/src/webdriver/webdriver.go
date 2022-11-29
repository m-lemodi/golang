package webdriver

import (
	"github.com/tebeka/selenium"
)

func NewInstance() (selenium.WebDriver, error) {

	caps := selenium.Capabilities{"browserName": "firefox"}
	driver, err := selenium.NewRemote(caps, "http://localhost:4444")
	if err != nil {
		panic(err)
	}
	return driver, err

}
