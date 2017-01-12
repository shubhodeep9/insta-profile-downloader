package api

import (
	"encoding/json"
	"github.com/headzoo/surf"
	"strconv"
)

var ct int = 0
var urls []PhotoStruct

type Photos struct {
	Count int
	Photo []PhotoStruct
}

type PhotoStruct struct {
	Title string
	Url   string
}

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

	if len(posts) > 0 {

		for i := range posts {
			post_uri, _ := posts[i].(map[string]interface{})["display_src"].(string)
			date, _ := posts[i].(map[string]interface{})["date"].(int)
			photo := PhotoStruct{
				Title: strconv.Itoa(date),
				Url:   post_uri,
			}
			urls = append(urls, photo)
		}

		//next_page
		page_info, _ := parse["page_info"].(map[string]interface{})
		hasnextpage, _ := page_info["has_next_page"].(bool)

		if hasnextpage {
			end_cursor := page_info["end_cursor"].(string)
			page_iterator(uri + "/?max_id=" + end_cursor)
		}
	}

}

func GetPhotos(uri string) *Photos {
	page_iterator(uri)
	return &Photos{
		Count: len(urls),
		Photo: urls,
	}
}
