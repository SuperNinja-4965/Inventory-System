package main

import (
	"html/template"
	"net/http"
	"strings"
)

func CatagoryPage(w http.ResponseWriter, r *http.Request, CatURL string) {
	for i := 0; i <= len(catagories)-1; i++ {
		if CatURL == catagories[i] {
			readCSV(CatURL)
			var cats string
			//fmt.Println(len(catagories))
			if len(items) != 0 {
				for i := 0; i <= len(items)-1; i++ {
					cats = cats + ItemView("/"+items[i], items[i], "1")
				}
				p := MainIndexPage{Catagories: template.HTML(cats), ProjectName: ProgramName}
				t, _ := template.ParseFiles(ExecPath + "/html/index.html")
				t.Execute(w, p)
			} else {
				cats = cats + ItemView("", "NO ITEMS FOUND", "0")
				p := MainIndexPage{Catagories: template.HTML(cats), ProjectName: ProgramName}
				t, _ := template.ParseFiles(ExecPath + "/html/index.html")
				t.Execute(w, p)
			}
		}
	}
}

func CatagoryPageDefine() {
	http.HandleFunc("/catagory/", makeHandlercat(CatagoryPage))
}

func makeHandlercat(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// if strings.Contains(r.URL.Path, ".html") {
		// 	if strings.Contains(r.URL.Path, "index.html") {
		// 		length := len(r.URL.Path)
		// 		length = length - 11
		// 		runes := []rune(r.URL.Path)
		// 		http.Redirect(w, r, "../../"+string(runes[0:length]), 307)
		// 	} else {
		// 		length := len(r.URL.Path)
		// 		length = length - 5
		// 		runes := []rune(r.URL.Path)
		// 		http.Redirect(w, r, "../../"+string(runes[0:length]), 307)
		// 	}
		// 	return
		// }
		// m := validPath.FindStringSubmatch(r.URL.Path)
		// if m == nil {
		// 	http.Redirect(w, r, "../../../errorPages/not_found", 307)
		// 	fmt.Printf("404 Error at %s @ %s\n", r.URL.Path, dateTimeFormatted())
		// 	return
		// }
		//fmt.Printf("%s", r.URL.Path)
		//fn(w, r, m[2])
		urll := r.URL.Path
		urll = strings.Replace(urll, "/catagory/", "", -1)
		fn(w, r, urll)
	}
}
