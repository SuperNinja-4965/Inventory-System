package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

var MainSiteURL string
var SitePort string
var NonHttpsPort string
var ExecPath string
var ProgramName string

func main() {
	// get the path of the program.
	var err2 error
	ExecPath, err2 = filepath.Abs(filepath.Dir(os.Args[0]))
	if err2 != nil {
		log.Fatal(err2)
	}
	// Begins the startup script.
	StartUp()
	// Define system variables
	ProgramName = "Inventory System"
	MainSiteURL = "10.0.0.2"
	SitePort = "8443"
	NonHttpsPort = "8080"
	fmt.Println("The server ip is: " + GetServerIp(0))
	openbrowser("http://" + MainSiteURL + ":" + NonHttpsPort + "/")
	openbrowser("https://" + MainSiteURL + ":" + SitePort + "/")
	//initPages
	initPages()
	//readCSV("file.de")

	go func() {
		if _, err := os.Stat(ExecPath + "/HTTPS-key/server.crt"); os.IsNotExist(err) {
			fmt.Printf("server.crt does not exist. HTTPS NOT STARTED\n")
		} else if _, err := os.Stat(ExecPath + "/HTTPS-key/server.key"); os.IsNotExist(err) {
			fmt.Printf("server.key does not exist. HTTPS NOT STARTED\n")
		} else {
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

func GetServerIp(ipNum int) string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		os.Stderr.WriteString("Oops: " + err.Error() + "\n")
		os.Exit(1)
	}

	var ips [255]string
	var i int
	i = 0
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ips[i] = ipnet.IP.String()
				i = i + 1
			}
		}
	}
	return ips[ipNum]
}

func openbrowser(url string) {
	var err error
	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}
}
