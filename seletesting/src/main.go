package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func main() {

	fmt.Println("=======================")
	fmt.Println("     THE BEGINNING")
	fmt.Println(">----------------------")

	// Connect to server
	caps := selenium.Capabilities{"browserName": "chrome"}
	chromeCaps := chrome.Capabilities{
		Args: []string{"--ignore-certificate-errors"},
	}
	caps.AddChrome(chromeCaps)

	driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		panic(err)
	}
	fmt.Println(driver.Capabilities())
	defer driver.Quit()

	// Get on first website
	err = driver.Get("http://mfarez.leluke.com")
	if err != nil {
		fmt.Println(err)
	}
	screenSource, err := driver.Screenshot()
	if err != nil {
		fmt.Println(err)
	}

	// Get on second website
	err = driver.Get("http://51.159.123.1/~mfarez")
	if err != nil {
		fmt.Println(err)
	}

	screenTarget, err := driver.Screenshot()
	if err != nil {
		fmt.Println(err)
	}
	err = ScreenHandler(screenSource, screenTarget, "mfarez")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(err)
	fmt.Println(">----------------------")
	fmt.Println("        THE END")
	fmt.Println("=======================")
}
