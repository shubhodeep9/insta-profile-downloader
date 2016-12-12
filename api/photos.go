package main

import (
	"encoding/json"
	"fmt"
	_ "github.com/PuerkitoBio/goquery"
	"github.com/headzoo/surf"
)

var ct int = 0

func page_iterator(uri string) {
	//initiate browser
	bow := surf.NewBrowser()
	bow.Open(uri)

	scr := bow.Find("script").Eq(6)

	extracted_react_data := []byte(scr.Text()[21 : len(scr.Text())-1])

	var parsed interface{}
	json.Unmarshal(extracted_react_data, &parsed)
	parse, _ := parsed.(map[string]interface{})["entry_data"].(map[string]interface{})["ProfilePage"].([]interface{})[0].(map[string]interface{})["user"].(map[string]interface{})["media"].(map[string]interface{})

	//posts
	posts, _ := parse["nodes"].([]interface{})

	// for i := range posts {
	// 	post_uri, _ := posts[i].(map[string]interface{})["display_src"].(string)
	// 	fmt.Println(post_uri)
	// }

	ct = ct + len(posts)
	//next_page
	page_info, _ := parse["page_info"].(map[string]interface{})
	hasnextpage, _ := page_info["has_next_page"].(bool)

	fmt.Println(hasnextpage)

	if hasnextpage {
		end_cursor := page_info["end_cursor"].(string)
		fmt.Println(page_info)
		page_iterator(uri + "/?max_id=" + end_cursor)
	}

}

func main() {
	page_iterator("https://www.instagram.com/rishi_raj95")
}
