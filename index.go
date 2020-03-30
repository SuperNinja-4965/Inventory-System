package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type MainIndexPage struct {
	Data        template.HTML
	ProjectName string
}

//initPages
func initPages() {
	http.HandleFunc("/", indexPage)
	CatagoryPageDefine()
	initNew()
	InitSearch()
	// Making the assets folder work.
	// Location of local file
	//fs := http.FileServer(http.Dir(ExecPath + "/html/assets/"))
	//location on server when hosted
	//http.Handle("/assets/", http.StripPrefix("/assets/", fs))
	http.HandleFunc("/assets/css/styles.css", stylesCss)
	http.HandleFunc("/assets/css/styles2.css", styles2Css)
}

func stylesCss(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	fmt.Fprint(w, cssIndex)
}

func styles2Css(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/css")
	fmt.Fprint(w, cssTwo)
}

func indexPage(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		var searchIndex string = r.FormValue("search")
		p := MainIndexPage{Data: template.HTML(BeginSearch(searchIndex, "all")), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
		return
	}
	if r.FormValue("search") != "" {
		var searchIndex string = r.FormValue("search")
		p := MainIndexPage{Data: template.HTML(BeginSearch(searchIndex, "all")), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
		return
	}
	GetCatagories()
	var cats string
	//fmt.Println(len(catagories))
	if len(catagories) != 0 {
		for i := 0; i <= len(catagories)-1; i++ {
			var count int = 1
			csvfile, err := os.Open(ExecPath + "/data/" + catagories[i] + ".csv")
			if err != nil {
				log.Fatalln("Couldn't open the csv file", err)
			}
			// Parse the file
			r := csv.NewReader(csvfile)
			//r := csv.NewReader(bufio.NewReader(csvfile))
			// Iterate through the records
			for {
				// Read each record from csv
				_, err := r.Read()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatal(err)
				}
				count = count + 1
			}
			csvfile.Close()
			count = count - 1
			if count == -1 {
				count = 0
			}
			cats = cats + ItemView("/category/"+catagories[i], catagories[i], "Amount of items: "+strconv.Itoa(count))
		}
		p := MainIndexPage{Data: template.HTML(cats), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
	} else {
		cats = cats + ItemView("", "NO cats found.", "0")
		p := MainIndexPage{Data: template.HTML(cats), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
	}
}

func ItemView(link string, name string, details string) string {
	return "<li class=\"folders\"><a href=\"" + link + "\" title=\"files/\"" + name + "\" class=\"folders\"><span class=\"icon folder full\"></span><span class=\"name\">" + name + "</span><span class=\"details\">" + details + "</span></a></li>"
}

var catagories []string

// GetCatagories - gets all of the catagories in /data
func GetCatagories() {
	catagories = nil
	//catagories = append(catagories, "Blah", "Blah")
	//catagories = append(catagories, "Blah")
	count := 0
	err := filepath.Walk(ExecPath+"/data/",
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			count = count + 1
			//fmt.Println(count)
			if count > 1 {
				switch runtime.GOOS {
				case "linux":
					path = strings.ReplaceAll(path, ExecPath+"/data/", "")
				case "windows":
					path = strings.ReplaceAll(path, ExecPath+"\\data\\", "")
				case "darwin":
					path = strings.ReplaceAll(path, ExecPath+"/data/", "")
				default:
					path = strings.ReplaceAll(path, ExecPath+"/data/", "")
				}
				path = strings.ReplaceAll(path, ".csv", "")
				if path == "" {
				} else if path == "\n" {
				} else {
					//fmt.Println(path, info.Size())
					//fmt.Printf("length: %d\n", len(catagories))
					catagories = append(catagories, path)
				}
			}
			return nil
		})
	if err != nil {
		log.Println(err)
	} else {
		//fmt.Println("Loaded Catagories")
	}
}
