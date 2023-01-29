package main

import (
	"log"
	"os"
	"github.com/playwright-community/playwright-go"
	"github.com/joho/godotenv"
)

type AuthInfo struct {
	Id       string
	Password string
	Hostname string
}

func loadEnvs() [2]AuthInfo {
	var ret [2]AuthInfo
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}
	f660aAuth := AuthInfo{
		Id:       os.Getenv("F660A_ID"),
		Password: os.Getenv("F660A_PW"),
		Hostname: os.Getenv("F660A_HOST"),
	}
	orbiAuth := AuthInfo{
		Id:       os.Getenv("ORBI_ID"),
		Password: os.Getenv("ORBI_PW"),
		Hostname: os.Getenv("ORBI_HOST"),
	}
	log.Printf("\n f660aAuth: %v \n orbiAuth: %v", f660aAuth, orbiAuth)
	ret[0] = f660aAuth
	ret[1] = orbiAuth
	return ret
}
func screenShot(p playwright.Page) {
	p.WaitForLoadState()
	if _, err := p.Screenshot(playwright.PageScreenshotOptions{
		Path: playwright.String("screenshot.png"),
	}); err != nil {
		log.Fatalf("could not create screenshot: %v", err)
	}
}
func main() {
	au := loadEnvs()
	if err := playwright.Install(); err != nil {
		log.Fatal("Error Installing playwright: %v", err)
	}
	pw, err := playwright.Run()
	if err != nil {
		log.Fatalf("Could not start playwright: %v", err)
	}
	browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
		Headless: playwright.Bool(false),
	})
	if err != nil {
		log.Fatalf("Could not launch browser: %v", err)
	}
	page, err := browser.NewPage()
	if err != nil {
		log.Fatalf("Could not create page: %v", err)
	}
	log.Printf("STEP:LOGIN goto http://" + au[0].Hostname + "/")
	if _, err = page.Goto("http://" + au[0].Hostname + "/"); err != nil {
		log.Fatalf("Could not goto: %v", err)
	}
	page.WaitForLoadState()
	log.Printf("login...")
	if err := page.Type("#Frm_Username", au[0].Id); err != nil {
		log.Fatalf("Operation Missed: %v", err)
	}
	if err := page.Type("#Frm_Password", au[0].Password); err != nil {
		log.Fatalf("Operation Missed: %v", err)
	}
	if err := page.Press("#Frm_Password", "Enter"); err != nil {
		log.Fatalf("Operation Missed: %v", err)
	}

	log.Printf("STEP:KANRI navigate to kanri IFRAME...")
	page.WaitForTimeout(1 * 1000) // TODO: fix someday
	log.Printf("goto KANRI page http://" + au[0].Hostname + "/getpage.gch?pid=1002&nextpage=manager_dev_conf_t.gch")
	if _, err = page.Goto("http://" + au[0].Hostname + "/getpage.gch?pid=1002&nextpage=manager_dev_conf_t.gch"); err != nil {
		log.Fatalf("Could not goto: %v", err)
	}

	log.Printf("STEP:REBOOT Reboot/Submit")
	page.Click("#Submit1")
	page.WaitForTimeout(1 * 1000) // TODO: fix someday

	log.Printf("STEP:CONFIRM REBOOT OK")
	//FIXME:wakaranai
	//page.On("Dialog","OK")
	//browser.ExpectedDialog.Accept("OK")
	//page.ExpectedDialog.Accept("OK")

	log.Printf("STEP:CLICKED!!!!!!!")
	//screenShot(page)
	if err = browser.Close(); err != nil {
		log.Fatalf("Could not close browser: %v", err)
	}
	if err = pw.Stop(); err != nil {
		log.Fatalf("Could not stop Playwright: %v", err)
	}
}
