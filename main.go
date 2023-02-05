package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/playwright-community/playwright-go"
)

type AuthInfo struct {
	Id       string
	Password string
	Hostname string
}

var DryRun string

func loadEnvs() [2]AuthInfo {
	var ret [2]AuthInfo
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file: %v", err)
	}
	DryRun = os.Getenv("DRY_RUN")
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
	auth := loadEnvs()
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
	if err != nil {
		log.Fatalf("Could not create page: %v", err)
	}
	page, err := browser.NewPage()

	// orbi
	log.Printf("START to reboot Orbi")
	restartOrbi(page, auth[1])

	// f660a
	log.Printf("START to reboot f660a")
	restartF660a(page, auth[0])

	if err := browser.Close(); err != nil {
		log.Fatalf("Could not close browser: %v", err)
	}
	if err := pw.Stop(); err != nil {
		log.Fatalf("Could not stop Playwright: %v", err)
	}
}
func restartOrbi(p playwright.Page, ai AuthInfo) {
	// NG: couln't use httpCredential
	//pw, err := playwright.Run()
	//if err != nil {
	//	log.Fatalf("Could not start playwright: %v", err)
	//}
	//op :=playwright.BrowserNewContextOptions {
	//	HttpCredentials: &playwright.BrowserNewContextOptionsHttpCredentials {
	//		Username: &ai.Id,
	//		Password: &ai.Password,
	//	},
	//}
	//browser, err := pw.Chromium.Launch(playwright.BrowserTypeLaunchOptions{
	//	Headless: playwright.Bool(false),
	//})
	//browser.NewContext(op)
	//p, err := browser.NewPage()

	log.Printf("STEP: goto " + ai.Hostname)
	//if _, err := p.Goto("http://" + ai.Id + ":" + ai.Password +"@"+ ai.Hostname + "/adv_index.htm"); err != nil {
	if _, err := p.Goto("http://" + ai.Id + ":" + ai.Password + "@" + ai.Hostname + "/reboot.htm"); err != nil {
		log.Fatalf("Could not goto: %v", err)
	}
	p.WaitForLoadState()
	//log.Printf("STEP: CLICK Reboot")
	//if err := p.Press("#reboot", "Enter"); err != nil {
	//	log.Fatalf("Operation Missed: %v", err)
	//}

	log.Printf("STEP: Confirm Reboot")
	if DryRun == "true" {
		log.Println("dry-run")
		if err := p.Press("#main > table > tbody > tr:nth-child(4) > td > input:nth-child(3)", "Enter"); err != nil {
			log.Fatalf("Operation Missed: %v", err)
		}
	} else {
		log.Println("restart!")
		if err := p.Press("#main > table > tbody > tr:nth-child(4) > td > input:nth-child(2)", "Enter"); err != nil {
			log.Fatalf("Operation Missed: %v", err)
		}
	}
	//p.WaitForLoadState()
	//p.WaitForTimeout(10 * 1000) // TODO: fix someday
}

func restartF660a(p playwright.Page, ai AuthInfo) {
	log.Printf("STEP: LOGIN goto http://" + ai.Hostname + "/")
	if _, err := p.Goto("http://" + ai.Hostname + "/"); err != nil {
		log.Fatalf("Could not goto: %v", err)
	}
	p.WaitForLoadState()
	//p.WaitForTimeout(10 * 1000) // TODO: fix someday
	log.Printf("STEP: LOGIN with authInfo")
	if err := p.Type("#Frm_Username", ai.Id); err != nil {
		log.Fatalf("Operation Missed: %v", err)
	}
	if err := p.Type("#Frm_Password", ai.Password); err != nil {
		log.Fatalf("Operation Missed: %v", err)
	}
	if err := p.Press("#Frm_Password", "Enter"); err != nil {
		log.Fatalf("Operation Missed: %v", err)
	}
	//p.WaitForLoadState()
	p.WaitForTimeout(1 * 1000) // TODO: fix someday

	//log.Printf("STEP: GOTO KANRI iframe page http://" + ai.Hostname + "/getp.gch?pid=1002&nextpage=manager_dev_conf_t.gch")
	log.Printf("STEP: GOTO KANRI iframe page http://" + ai.Hostname + "/template.gch?pid=1002&nextpage=manager_dev_conf_t.gch")
	//http://192.168.1.1/template.gch?pid=1002&nextpage=manager_dev_conf_t.gch
	if _, err := p.Goto("http://" + ai.Hostname + "/template.gch?pid=1002&nextpage=manager_dev_conf_t.gch"); err != nil {
		log.Fatalf("Could not goto: %v", err)
	}
	p.WaitForLoadState()

	log.Printf("STEP: REBOOT Reboot/Submit")
	// NG: coudn't click js.alert
	// p.Click("#Submit1")
	// NOTE: click dialog "OK"
	//p.On("dialog", func(dialog playwright.Dialog) {
	//	log.Printf(dialog.Message())
	//	dialog.Accept()
	//})
	//screenShot(page)

	// NOTE: Execute JS when cliked Reboot->OK
	p.EvaluateHandle("jslDisable(\"Submit1\",\"Submit2\");", struct{}{})
	p.EvaluateHandle("setValue(\"IF_ACTION\", \"devrestart\");", struct{}{})
	p.EvaluateHandle("setValue(\"flag\", \"1\");", struct{}{})
	p.EvaluateHandle("DisableALL();", struct{}{})
	if DryRun == "true" {
		log.Println("dry-run")
	} else {
		log.Println("restart!")
		p.EvaluateHandle("getObj(\"fSubmit\").submit();", struct{}{})
	}
	p.WaitForLoadState()
	//p.WaitForTimeout(10 * 1000) // TODO: fix someday
}
