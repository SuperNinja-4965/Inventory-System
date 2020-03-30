package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func InitSearch() {
	http.HandleFunc("/search/", Search)
	//http.HandleFunc("/AdvancedSearch/", AdvancedSearch)
}

type NewSearch struct {
	Search string
	InCat  string
}

var searchOptionsHTML string = "<center><h1 style=\"color:white;\">New Search</h1><div class=\"container\"><br><form method=\"POST\"><table><thead><tr><th><h2 style=\"color:white;\">Type</h2></th><th><h2 style=\"color:white;\">Value</h2></th>                                </tr>                            </thead>                            <tbody>                                <tr>                                    <td>                                        <h3 style=\"color:white;\">Search For:</h3></td><td><input type=\"textbox\" style=\"display: inline-block;\" placeholder=\"Search For\" name=\"search\" id=\"search\"></td></tr><tr><td><h3 style=\"color:white;\">In:</h3></td><td><input list=\"categoriesList\" type=\"textbox\" style=\"display: inline-block;\" id=\"cat\" name=\"cat\" placeholder=\"In\"></td></tr></tbody></table><br><input type=\"submit\" value=\"Search\"> </form></div></center>"

func Search(w http.ResponseWriter, r *http.Request) {
	// Show search results from search
	Search := NewSearch{
		Search: r.FormValue("search"),
		InCat:  r.FormValue("cat"),
	}
	if Search.Search == "" {
		// Show AdvancedSearch options
		p := MainIndexPage{Data: template.HTML(searchOptionsHTML + generateDatalist()), ProjectName: ProgramName}
		t, _ := template.New("indexTemplate").Parse(PageIndex)
		t.Execute(w, p)
		//tmpl.Execute(w, nil)
		return
	}
	if Search.InCat == "" {
		Search.InCat = "all"
	}
	p := MainIndexPage{Data: template.HTML(BeginSearch(Search.Search, Search.InCat)), ProjectName: ProgramName}
	t, _ := template.New("indexTemplate").Parse(PageIndex)
	t.Execute(w, p)
}

func BeginSearch(Search string, where string) string {
	if where == "all" {
		// search all categories
		GetCatagories()
		var buildHTML string
		for i := 0; i <= len(catagories)-1; i++ {
			// for each category
			foundItems = nil
			foundVals = nil
			foundAmounts = nil
			readCSVItemSearchLOWERCASE(catagories[i], Search)
			foundErrors = false
			for f := 0; f <= len(foundItems)-1; f++ {
				// for each category
				buildHTML = buildHTML + ItemView("/category/"+catagories[i]+"/"+foundItems[f], foundItems[f], "Value: "+foundVals[f]+", Amount: "+strconv.Itoa(foundAmounts[f]))
			}
		}
		if buildHTML == "" {
			return "<center><h1 style=\"color:red;\">No items were found.<h1><br><h3 style=\"color:white;\">Did you know if you leave the In box blank or type \"all\" the program will search all categories</h3><center><br><br>" + searchOptionsHTML + generateDatalist()
		}
		return buildHTML
	}
	// for each category
	var buildHTML string
	foundItems = nil
	foundVals = nil
	foundAmounts = nil
	readCSVItemSearchLOWERCASE(where, Search)
	if foundErrors == true {
		foundErrors = false
		return "<center><h1 style=\"color:red;\">There was an error with your search.<h1><br><h3 style=\"color:white;\">Did you know if you leave the In box blank or type \"all\" the program will search all categories</h3><center><br><br>" + searchOptionsHTML + generateDatalist()
	} else {
		for i := 0; i <= len(foundItems)-1; i++ {
			// for each category
			buildHTML = buildHTML + ItemView("/category/"+where+"/"+foundItems[i], foundItems[i], "Value: "+foundVals[i]+", Amount: "+strconv.Itoa(foundAmounts[i]))
		}
		if buildHTML == "" {
			return "<center><h1 style=\"color:red;\">>No items were found.<h1><br><h3 style=\"color:white;\">Did you know if you leave the In box blank or type \"all\" the program will search all categories</h3><center><br><br>" + searchOptionsHTML + generateDatalist()
		}
		return buildHTML
	}
}

func generateDatalist() string {
	GetCatagories()
	var buildTemp string = "<datalist id=\"categoriesList\">"
	for i := 0; i <= len(catagories)-1; i++ {
		buildTemp = buildTemp + "<option value=\"" + catagories[i] + "\">"
	}
	return buildTemp + "</datalist>"
}
