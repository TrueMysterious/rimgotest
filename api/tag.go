package api

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	
	"github.com/patrickmn/go-cache"
	"github.com/tidwall/gjson"
)

type Tag struct {
	Tag          string
	Display      string
	Sort         string
	PostCount    int64
	Posts        []Submission
	Background   string
	BackgroundId string
}

func (client *Client) FetchTag(tag string, sort string, page string) (Tag, error) {
	cacheData, found := client.Cache.Get(tag + sort + page + "-tag")
	if found {
		return cacheData.(Tag), nil
	}

	req, err := http.NewRequest("GET", "https://api.imgur.com/post/v1/posts/t/"+tag, nil)
	if err != nil {
		return Tag{}, err
	}

	q := req.URL.Query()
	q.Add("client_id", client.ClientID)
	q.Add("include", "cover")
	q.Add("page", page)

	switch sort {
	case "newest":
		q.Add("filter[window]", "week")
		q.Add("sort", "-time")
	case "best":
		q.Add("filter[window]", "all")
		q.Add("sort", "-top")
	case "popular":
	default:
		q.Add("filter[window]", "week")
		q.Add("sort", "-viral")
		sort = "popular"
	}

	req.URL.RawQuery = q.Encode()

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return Tag{}, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return Tag{}, err
	}

	data := gjson.Parse(string(body))

	wg := sync.WaitGroup{}
	posts := make([]Submission, 0)
	data.Get("posts").ForEach(
		func(key, value gjson.Result) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()
				posts = append(posts, Submission{
					Id:    value.Get("id").String(),
					Title: value.Get("title").String(),
					Link:  strings.ReplaceAll(value.Get("url").String(), "https://imgur.com", ""),
					Cover: Media{
						Id:   value.Get("cover_id").String(),
						Type: value.Get("cover.type").String(),
						Url:  strings.ReplaceAll(value.Get("cover.url").String(), "https://i.imgur.com", ""),
					},
					Points:    value.Get("point_count").Int(),
					Upvotes:   value.Get("upvote_count").Int(),
					Downvotes: value.Get("downvote_count").Int(),
					Comments:  value.Get("comment_count").Int(),
					Views:     value.Get("view_count").Int(),
					IsAlbum:   value.Get("is_album").Bool(),
				})
			}()

			return true
		},
	)

	wg.Wait()

	tagData := Tag{
		Tag:          tag,
		Display:      data.Get("display").String(),
		Sort:         sort,
		PostCount:    data.Get("post_count").Int(),
		Posts:        posts,
		Background:   "/" + data.Get("background_id").String() + ".webp",
	}

	client.Cache.Set(tag + sort + page + "-tag", tagData, cache.DefaultExpiration)
	return tagData, nil
}
