package main

import (
	"fmt"
	"github.com/tebeka/selenium"
	"meli-product-miner/internal/webdriver"
	"time"
)

func main() {
	response, err := webdriver.CreateWebDriver()
	driver := response.Driver
	if err != nil {
		fmt.Printf("Erro ao criar o WebDriver: %v\n", err)
		return
	}
	driver.Get("https://www.google.com.br/")
	a, _ := driver.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[3]/form/div[1]/div[1]/div[1]/div/div[2]/textarea")
	a.SendKeys("hello world")
	a.Submit()

	time.Sleep(10 * time.Second)

	response.Service.Stop()
}
