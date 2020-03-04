package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func LoadItemPage(w http.ResponseWriter, r *http.Request, ItemURL string, ItemCategory string) {
	ItemName, ItemValue, ItemAmount, ItemInUse, ItemTotal, ItemNotes := readCSVItemSearch(ItemCategory, ItemURL)
	fmt.Fprintln(w, ItemName+" "+ItemValue+" "+strconv.Itoa(ItemAmount)+" "+strconv.Itoa(ItemInUse)+" "+strconv.Itoa(ItemTotal)+" "+ItemNotes)
}
