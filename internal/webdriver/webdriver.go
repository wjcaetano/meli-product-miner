package webdriver

import (
	"fmt"
	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
	"time"
)

type (
	DriverResponse struct {
		Driver  selenium.WebDriver
		Service *selenium.Service
	}
)

const (
	port = 9515
)

func CreateWebDriver() (DriverResponse, error) {
	driverResponse := DriverResponse{}
	chromeCaps := chrome.Capabilities{
		Args: []string{
			"window-size=1920x1080",
			"--no-sandbox",
			"--disable-dev-shm-usage",
			"disable-gpu",
			// "--headless",
		},
	}

	chromeDrivePath := "/usr/local/bin/chromedriver"
	service, err := selenium.NewChromeDriverService(chromeDrivePath, port)
	if err != nil {
		return driverResponse, fmt.Errorf("failed to start the driver service: %w", err)
	}

	caps := selenium.Capabilities{
		"browserName": "chrome",
		"chromeOptions": map[string]interface{}{
			"w3c":    false,
			"args":   chromeCaps.Args,
			"binary": "",
		},
	}

	webDriver, err := selenium.NewRemote(caps, fmt.Sprintf("http://localhost:%d/wd/hub", port))
	if err != nil {
		return driverResponse, fmt.Errorf("failed to start the WebDriver: %w", err)
	}

	driverResponse = DriverResponse{
		Driver:  webDriver,
		Service: service,
	}

	webDriver.SetImplicitWaitTimeout(10 * time.Second)

	return driverResponse, nil
}
