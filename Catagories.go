package main

import (
	"html/template"
	"net/http"
	"strconv"
	"strings"
)

func categoryPage(w http.ResponseWriter, r *http.Request, CatURL string) {
	//fmt.Println(CatURL)
	if CatURL == "" {
		http.Redirect(w, r, "/", 301)
	}
	if string(CatURL[len(CatURL)-1]) == "/" {
		temp := reverse(CatURL)
		CatURL = reverse(strings.Replace(temp, "/", "", 1))
	}
	//fmt.Println(CatURL)
	if strings.Contains(CatURL, "/") == true {
		itemcatsplit := strings.Split(CatURL, "/")
		LoadItemPage(w, r, itemcatsplit[1], itemcatsplit[0])
	} else {
		var found bool = false
		for i := 0; i <= len(catagories)-1; i++ {
			if found == false {
				if CatURL == catagories[i] {
					found = true
					readCSV(CatURL)
					var cats string
					//fmt.Println(len(catagories))
					if len(items) != 0 {
						for i := 0; i <= len(items)-1; i++ {
							cats = cats + ItemView("/category/"+CatURL+"/"+items[i], items[i], "Value: "+Value[i]+", Amount: "+strconv.Itoa(ItemsTotal[i]))
						}
						p := MainIndexPage{Catagories: template.HTML(cats), ProjectName: ProgramName, Table: template.HTML("")}
						t, _ := template.ParseFiles(ExecPath + "/html/index.html")
						t.Execute(w, p)
					} else {
						cats = cats + ItemView("", "NO items found.", "0")
						p := MainIndexPage{Catagories: template.HTML(cats), ProjectName: ProgramName, Table: template.HTML("")}
						t, _ := template.ParseFiles(ExecPath + "/html/index.html")
						t.Execute(w, p)
					}
				}
			}
		}
		if found == false {
			var cats string
			cats = cats + ItemView("", "Category NOT found.", "0")
			p := MainIndexPage{Catagories: template.HTML(cats), ProjectName: ProgramName, Table: template.HTML("")}
			t, _ := template.ParseFiles(ExecPath + "/html/index.html")
			t.Execute(w, p)
		}
	}
}

func CatagoryPageDefine() {
	http.HandleFunc("/category/", makeHandlercat(categoryPage))
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
		if urll == "/category/" {
			http.Redirect(w, r, "/", 301)
		} else if urll == "/category" {
			http.Redirect(w, r, "/", 301)
		} else {
			urll = strings.Replace(urll, "/category/", "", -1)
			//fmt.Println(urll)
			fn(w, r, urll)
		}
	}
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
