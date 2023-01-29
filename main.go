package main

import (
	//"fmt"
	"log"

	"github.com/playwright-community/playwright-go"

	"os"
    "github.com/joho/godotenv"
)


func main() {
    if err := godotenv.Load();err != nil {
        log.Fatal("Error loading .env file: %v", err)
    }
	log.Printf("F660A_HOST=%v F660A_ID=%s F660A_PW=%s", os.Getenv("F660A_HOST"), os.Getenv("F660A_ID"), os.Getenv("F660A_PW"))
	log.Printf("ORBI_HOST=%v ORBI_ID=%s ORBI_PW=%s", os.Getenv("ORBI_HOST"), os.Getenv("ORBI_ID"), os.Getenv("ORBI_PW"))


	_, err := playwright.Run()
	if err != nil {
		log.Fatalf("Could not start playwright: %v", err)
	}
	//browser, err := pw.Chromium.Launch()
	//if err != nil {
	//	log.Fatalf("could not launch browser: %v", err)
	//}
	//page, err := browser.NewPage()
	//if err != nil {
	//	log.Fatalf("could not create page: %v", err)
	//}
	//if _, err = page.Goto("https://news.ycombinator.com"); err != nil {
	//	log.Fatalf("could not goto: %v", err)
	//}
	//entries, err := page.QuerySelectorAll(".athing")
	//if err != nil {
	//	log.Fatalf("could not get entries: %v", err)
	//}
	//for i, entry := range entries {
	//	titleElement, err := entry.QuerySelector("td.title > span > a")
	//	if err != nil {
	//		log.Fatalf("could not get title element: %v", err)
	//	}
	//	title, err := titleElement.TextContent()
	//	if err != nil {
	//		log.Fatalf("could not get text content: %v", err)
	//	}
	//	fmt.Printf("%d: %s\n", i+1, title)
	//}
	//if err = browser.Close(); err != nil {
	//	log.Fatalf("could not close browser: %v", err)
	//}
	//if err = pw.Stop(); err != nil {
	//	log.Fatalf("could not stop Playwright: %v", err)
	//}
}
