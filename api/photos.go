package main

import (
	"encoding/json"
	"fmt"
	"github.com/headzoo/surf"
)

var ct int = 0
var uri string
var urls []string

func page_iterator(url string) {
	//initiate browser
	bow := surf.NewBrowser()
	bow.Open(url)

	scr := bow.Find("script").Eq(6)

	extracted_react_data := []byte(scr.Text()[21 : len(scr.Text())-1])

	var parsed interface{}
	json.Unmarshal(extracted_react_data, &parsed)
	parse, _ := parsed.(map[string]interface{})["entry_data"].(map[string]interface{})["ProfilePage"].([]interface{})[0].(map[string]interface{})["user"].(map[string]interface{})["media"].(map[string]interface{})

	//posts
	posts, _ := parse["nodes"].([]interface{})

	for i := range posts {
		post_uri, _ := posts[i].(map[string]interface{})["display_src"].(string)
		urls = append(urls, post_uri)
	}

	//next_page
	page_info, _ := parse["page_info"].(map[string]interface{})
	hasnextpage, _ := page_info["has_next_page"].(bool)

	if hasnextpage {
		end_cursor := page_info["end_cursor"].(string)
		page_iterator(uri + "/?max_id=" + end_cursor)
	}

}

func main() {
	uri = "https://www.instagram.com/rishi_raj95"
	page_iterator(uri)
	fmt.Println(len(urls))
}
