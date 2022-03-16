package api

import (
	"io"
	"net/http"
	"strings"
	"sync"
	"time"

	"codeberg.org/video-prize-ranch/rimgo/types"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

var userCache = cache.New(30*time.Minute, 15*time.Minute)

func FetchUser(username string) (types.User, error) {
	cacheData, found := userCache.Get(username)
	if found {
		return cacheData.(types.User), nil
	}

	res, err := http.Get("https://api.imgur.com/account/v1/accounts/" + username + "?client_id=" + viper.GetString("RIMGU_IMGUR_CLIENT_ID"))
	if err != nil {
		return types.User{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.User{}, err
	}

	data := gjson.Parse(string(body))

	createdTime, _ := time.Parse(time.RFC3339, data.Get("created_at").String())

	user := types.User{
		Id:        data.Get("id").Int(),
		Bio:       data.Get("bio").String(),
		Username:  data.Get("username").String(),
		Points:    data.Get("reputation_count").Int(),
		Cover:     strings.ReplaceAll(data.Get("cover_url").String(), "https://imgur.com", ""),
		Avatar:    strings.ReplaceAll(data.Get("avatar_url").String(), "https://i.imgur.com", ""),
		CreatedAt: createdTime.Format("January 2, 2006"),
	}

	userCache.Set(username, user, 1*time.Hour)
	return user, nil
}

func FetchSubmissions(username string, sort string, page string) ([]types.Submission, error) {
	cacheData, found := userCache.Get(username + "-submissions")
	if found {
		return cacheData.([]types.Submission), nil
	}

	res, err := http.Get("https://api.imgur.com/3/account/" + username + "/submissions/" + page + "/" + sort + "?album_previews=1&client_id=" + viper.GetString("RIMGU_IMGUR_CLIENT_ID"))
	if err != nil {
		return []types.Submission{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []types.Submission{}, err
	}

	data := gjson.Parse(string(body))

	submissions := []types.Submission{}

	wg := sync.WaitGroup{}
	data.Get("data").ForEach(
		func(key, value gjson.Result) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()

				coverData := value.Get("images.#(id==\"" + value.Get("cover").String() + "\")")
				cover := types.Media{}
				if coverData.Exists() {
					cover = types.Media{
						Id:          coverData.Get("id").String(),
						Description: coverData.Get("description").String(),
						Type:        strings.Split(coverData.Get("type").String(), "/")[0],
						Url:         strings.ReplaceAll(coverData.Get("link").String(), "https://i.imgur.com", ""),
					}
				} else {
					cover = types.Media{
						Id:          value.Get("id").String(),
						Description: value.Get("description").String(),
						Type:        strings.Split(value.Get("type").String(), "/")[0],
						Url:         strings.ReplaceAll(value.Get("link").String(), "https://i.imgur.com", ""),
					}
				}

				id := value.Get("id").String()

				submissions = append(submissions, types.Submission{
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

	userCache.Set(username + "-submissions", submissions, 15*time.Minute)
	return submissions, nil
}
