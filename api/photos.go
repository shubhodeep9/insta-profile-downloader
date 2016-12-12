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
	json.Unmarshal(extracted_react_data, &parsed)
	parse, _ := parsed.(map[string]interface{})["entry_data"].(map[string]interface{})["ProfilePage"].([]interface{})[0].(map[string]interface{})["user"].(map[string]interface{})["media"].(map[string]interface{})

	//posts
	posts, _ := parse["nodes"].([]interface{})

	for i := range posts {
		post_uri, _ := posts[i].(map[string]interface{})["display_src"].(string)
		fmt.Println(post_uri)
	}
}
