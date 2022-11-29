package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	fox "github.com/tebeka/selenium/firefox"
)

func main() {

	fmt.Println("=======================")
	fmt.Println("     THE BEGINNING")
	fmt.Println(">----------------------")

	// Connect to server

	caps := selenium.Capabilities{"browserName": "firefox"}
	firefoxCaps := fox.Capabilities{
		Args: []string{"--ignore-certificate-errors"},
	}
	caps.AddFirefox(firefoxCaps)

	driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		panic(err)
	}
	fmt.Println(driver.Capabilities())
	defer driver.Close()
	defer driver.Quit()

	// Get on first website
	err = driver.Get("https://mfarez.leluke.com")
	if err != nil {
		fmt.Println(err)
	}
	screenSource, err := driver.Screenshot()
	if err != nil {
		fmt.Println(err)
	}

	// Get on second website
	err = driver.Get("https://51.159.123.1/~mfarez")
	if err != nil {
		fmt.Println(err)
	}

	screenTarget, err := driver.Screenshot()
	if err != nil {
		fmt.Println(err)
	}
	err = ScreenHandler(screenSource, screenTarget)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
	fmt.Println(">----------------------")
	fmt.Println("        THE END")
	fmt.Println("=======================")
}
