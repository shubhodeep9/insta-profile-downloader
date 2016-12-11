package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
)

func main() {
	//initiate browser
	bow := surf.NewBrowser()
	bow.Open("https://www.instagram.com/shubhothegreat")

	scr := bow.Find("script").Eq(6)

	extracted_react_data := []byte(scr.Text()[21 : len(scr.Text())-1])

	var parsed interface{}
	err := json.Unmarshal(extracted_react_data, &parsed)
	parse, _ := parsed.(map[string]interface{})
	fmt.Println(err, parse["country_code"])
}
