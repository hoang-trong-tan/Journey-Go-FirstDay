package main

import (
	"log"

	"github.com/go-rod/rod"
)

func main() {
	browser := rod.New().MustConnect()

	// Navigate to the page
	page := browser.MustPage("https://www.facebook.com/hacked")

	// Wait for the element to be ready, replace 'selector' with the actual CSS selector of the radio button
	el := page.MustElement("selector")

	// Click the radio button
	err := el.MustClick()
	if err != nil {
		log.Fatalf("Failed to click radio button: %v", err)
	}

	// Close the browser
	browser.MustClose()

}
