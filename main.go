package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	JG "github.com/joshuag1000/GoEssentials"
)

var MainSiteURL string
var SitePort string
var NonHttpsPort string
var ExecPath string
var ProgramName string
var openBrowserOnLoad bool = true
var DefaultSettings string = "// Settings file for Inventory System each setting must have a space after the colon or it will be ignored.\n\nProgram-Name: Inventory System\nHTTPS-PORT: 8443\nHTTP-PORT: 8080\n\n// options are true or false\nOpenBrowser: false"

func main() {
	//rtr := mux.NewRouter()
	// get the path of the program.
	var err2 error
	ExecPath, err2 = filepath.Abs(filepath.Dir(os.Args[0]))
	if err2 != nil {
		log.Fatal(err2)
	}
	readSettings()
	// if _, err := os.Stat(ExecPath + "/OpenBrowser.yes"); os.IsNotExist(err) {
	// 	openBrowserOnLoad = false
	// }
	// Begins the startup script.
	StartUp()
	// Define system variables
	//ProgramName = "Inventory System"
	// MainSiteURL is only used when opening browser and so can be left alone.
	MainSiteURL = "127.0.0.1"
	//SitePort = "8443"
	//NonHttpsPort = "8080"
	fmt.Println("The server ip is: " + JG.GetServerIP(0))
	//initPages
	initPages()
	//readCSV("file.de")
	//fmt.Println(ExecPath)
	go func() {
		if _, err := os.Stat(ExecPath + "/HTTPS-key/server.crt"); os.IsNotExist(err) {
			fmt.Printf("server.crt does not exist. HTTPS NOT STARTED\n")
			JG.OpenBrowser("http://" + MainSiteURL + ":" + NonHttpsPort + "/")
		} else if _, err := os.Stat(ExecPath + "/HTTPS-key/server.key"); os.IsNotExist(err) {
			fmt.Printf("server.key does not exist. HTTPS NOT STARTED\n")
			JG.OpenBrowser("http://" + MainSiteURL + ":" + NonHttpsPort + "/")
		} else {
			JG.OpenBrowser("https://" + MainSiteURL + ":" + SitePort + "/")
			// begin https server
			err_https := http.ListenAndServeTLS(":"+SitePort, ExecPath+"/HTTPS-key/server.crt", ExecPath+"/HTTPS-key/server.key", nil)
			if err_https != nil {
				log.Fatal("Web server (HTTPS): \n", err_https)
			}
		}
	}()

	// begin http server
	err_http := http.ListenAndServe(":"+NonHttpsPort, nil)
	if err_http != nil {
		log.Fatal("Web server (HTTP): ", err_http)
	}
}

func readSettings() {
	if _, err := os.Stat(ExecPath + "/settings.preferences"); os.IsNotExist(err) {
		f, _ := os.Create(ExecPath + "/settings.preferences")
		b := bufio.NewWriter(f)
		b.WriteString(DefaultSettings)
		b.Flush()
		f.Close()
	}
	readFile, err := os.Open(ExecPath + "/settings.preferences")

	if err != nil {
		panic("failed to open file: " + err.Error())
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var count int = 1
	for fileScanner.Scan() {
		//fileScanner.Text()
		value := fileScanner.Text()
		if value != "" {
			if value[0:2] != "//" {
				if count == 1 {
					ProgramName = value[14:len(value)]
					count = count + 1
				} else if count == 2 {
					SitePort = value[12:len(value)]
					count = count + 1
				} else if count == 3 {
					NonHttpsPort = value[11:len(value)]
					count = count + 1
				} else if count == 4 {
					openBrowserOnLoad, _ = strconv.ParseBool(value[13:len(value)])
					count = count + 1
				}
			}
		}
	}
	readFile.Close()
	//fmt.Printf("%s %s %s %v", ProgramName, SitePort, NonHttpsPort, openBrowserOnLoad)
}
