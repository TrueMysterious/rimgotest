package api

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/tidwall/gjson"
)

type User struct {
	Id        int64
	Bio       string
	Username  string
	Points    int64
	Cover     string
	Avatar    string
	CreatedAt string
}

type Submission struct {
	Id        string
	Title     string
	Link      string
	Cover     Media
	Points    int64
	Upvotes   int64
	Downvotes int64
	Comments  int64
	Views     int64
	IsAlbum   bool
}

func (client *Client) FetchUser(username string) (User, error) {
	cacheData, found := client.Cache.Get(username + "-user")
	if found {
		return cacheData.(User), nil
	}

	res, err := http.Get("https://api.imgur.com/account/v1/accounts/" + username + "?client_id=" + utils.Config.ImgurId)
	if err != nil {
		return User{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return User{}, err
	}

	data := gjson.Parse(string(body))

	createdTime, _ := time.Parse(time.RFC3339, data.Get("created_at").String())

	user := User{
		Id:        data.Get("id").Int(),
		Bio:       data.Get("bio").String(),
		Username:  data.Get("username").String(),
		Points:    data.Get("reputation_count").Int(),
		Cover:     strings.ReplaceAll(data.Get("cover_url").String(), "https://imgur.com", ""),
		Avatar:    strings.ReplaceAll(data.Get("avatar_url").String(), "https://i.imgur.com", ""),
		CreatedAt: createdTime.Format("January 2, 2006"),
	}

	client.Cache.Set(username + "-user", user, 1*time.Hour)
	return user, nil
}

func (client *Client) FetchSubmissions(username string, sort string, page string) ([]Submission, error) {
	cacheData, found := client.Cache.Get(username + "-submissions")
	if found {
		return cacheData.([]Submission), nil
	}

	data, err := utils.GetJSON("https://api.imgur.com/3/account/" + username + "/submissions/" + page + "/" + sort + "?album_previews=1&client_id=" + utils.Config.ImgurId)
	if err != nil {
		return []Submission{}, err
	}

	submissions := []Submission{}

	wg := sync.WaitGroup{}
	data.Get("data").ForEach(
		func(key, value gjson.Result) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()

				coverData := value.Get("images.#(id==\"" + value.Get("cover").String() + "\")")
				cover := Media{}
				if coverData.Exists() {
					cover = Media{
						Id:          coverData.Get("id").String(),
						Description: coverData.Get("description").String(),
						Type:        strings.Split(coverData.Get("type").String(), "/")[0],
						Url:         strings.ReplaceAll(coverData.Get("link").String(), "https://i.imgur.com", ""),
					}
				} else {
					cover = Media{
						Id:          value.Get("id").String(),
						Description: value.Get("description").String(),
						Type:        strings.Split(value.Get("type").String(), "/")[0],
						Url:         strings.ReplaceAll(value.Get("link").String(), "https://i.imgur.com", ""),
					}
				}

				id := value.Get("id").String()

				submissions = append(submissions, Submission{
					Id:    id,
					Link:  "/a/" + id,
					Title: value.Get("title").String(),
					Cover: cover,
					Points:    value.Get("points").Int(),
					Upvotes:   value.Get("ups").Int(),
					Downvotes: value.Get("downs").Int(),
					Comments:  value.Get("comment_count").Int(),
					Views:     value.Get("views").Int(),
					IsAlbum:   value.Get("is_album").Bool(),
				})
			}()

			return true
		},
	)
	wg.Wait()

	client.Cache.Set(username + "-submissions", submissions, 15*time.Minute)
	return submissions, nil
}
