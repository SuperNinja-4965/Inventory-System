package main

import (
	"html/template"
	"net/http"
	"strconv"
)

func LoadItemPage(w http.ResponseWriter, r *http.Request, ItemURL string, ItemCategory string) {
	ItemName, ItemValue, ItemAmount, ItemInUse, ItemTotal, ItemNotes := readCSVItemSearch(ItemCategory, ItemURL)
	prepare := "<center><h1 style=\"color:white;\">" + ItemName + " - Information</h1><div class=\"container\"><br><table><thead><tr><th>Name</th><th>Value</th><th>Amount available</th><th>Amount in use</th><th>Total amount</th><th>Notes</th></tr></thead><tbody><tr><td>" + ItemName + "</td><td>" + ItemValue + "</td><td>" + strconv.Itoa(ItemAmount) + "</td><td>" + strconv.Itoa(ItemInUse) + "</td><td>" + strconv.Itoa(ItemTotal) + "</td><td>" + ItemNotes + "</td></tr></tbody></table></div></center>"
	//fmt.Fprintln(w, ItemName+" "+ItemValue+" "+strconv.Itoa(ItemAmount)+" "+strconv.Itoa(ItemInUse)+" "+strconv.Itoa(ItemTotal)+" "+ItemNotes)
	p := MainIndexPage{Data: template.HTML(prepare), ProjectName: ProgramName}
	t, _ := template.New("indexTemplate").Parse(PageIndex)
	t.Execute(w, p)
}
