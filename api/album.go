package api

import (
	"io"
	"net/http"
	"strings"
	"time"

	"codeberg.org/video-prize-ranch/rimgo/types"
	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

var albumCache = cache.New(1*time.Hour, 15*time.Minute)

func FetchAlbum(albumID string) (types.Album, error) {
	cacheData, found := albumCache.Get(albumID + "-album")
	if found {
		return cacheData.(types.Album), nil
	}

	res, err := http.Get("https://api.imgur.com/post/v1/albums/" + albumID + "?client_id=" + viper.GetString("RIMGU_IMGUR_CLIENT_ID") + "&include=media%2Caccount")
	if err != nil {
		return types.Album{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.Album{}, err
	}

	data := gjson.Parse(string(body))

	album := types.Album{}
	if data.Get("shared_with_community").Bool() || res.StatusCode == 404 {
		album, err = FetchPosts(albumID)
	} else {
		album, err = ParseAlbum(data)
	}

	albumCache.Set(albumID + "-album", album, cache.DefaultExpiration)
	return album, err
}

func FetchPosts(albumID string) (types.Album, error) {
	cacheData, found := albumCache.Get(albumID + "-posts")
	if found {
		return cacheData.(types.Album), nil
	}

	res, err := http.Get("https://api.imgur.com/post/v1/posts/" + albumID + "?client_id=" + viper.GetString("RIMGU_IMGUR_CLIENT_ID") + "&include=media%2Caccount%2Ctags")
	if err != nil {
		return types.Album{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.Album{}, err
	}

	data := gjson.Parse(string(body))
	album, err := ParseAlbum(data)
	if err != nil {
		return types.Album{}, err
	}

	albumCache.Set(albumID + "-posts", album, cache.DefaultExpiration)
	return album, nil
}

func FetchMedia(mediaID string) (types.Album, error) {
	cacheData, found := albumCache.Get(mediaID + "-media")
	if found {
		return cacheData.(types.Album), nil
	}
	
	res, err := http.Get("https://api.imgur.com/post/v1/media/" + mediaID + "?client_id=" + viper.GetString("RIMGU_IMGUR_CLIENT_ID") + "&include=media%2Caccount")
	if err != nil {
		return types.Album{}, err
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return types.Album{}, err
	}

	data := gjson.Parse(string(body))
	album, err := ParseAlbum(data)
	if err != nil {
		return types.Album{}, err
	}

	albumCache.Set(mediaID + "-media", album, cache.DefaultExpiration)
	return album, nil
}

func ParseAlbum(data gjson.Result) (types.Album, error) {
	media := make([]types.Media, 0)
	data.Get("media").ForEach(
		func(key gjson.Result, value gjson.Result) bool {
			url := value.Get("url").String()
			url = strings.ReplaceAll(url, "https://i.imgur.com", "")

			media = append(media, types.Media{
				Id:          value.Get("id").String(),
				Name:        value.Get("name").String(),
				MimeType:    value.Get("mime_type").String(),
				Type:        value.Get("type").String(),
				Title:       value.Get("metadata.title").String(),
				Description: value.Get("metadata.description").String(),
				Url:         url,
			})

			return true
		},
	)

	tags := make([]types.Tag, 0)
	data.Get("tags").ForEach(
		func(key gjson.Result, value gjson.Result) bool {
			tags = append(tags, types.Tag{
				Tag: value.Get("tag").String(),
				Display: value.Get("display").String(),
				Background: "/" + value.Get("background_id").String() + ".webp",
			})
			return true
		},
	)

	createdAt, err := utils.FormatDate(data.Get("created_at").String())
	if err != nil {
		return types.Album{}, err
	}

	album := types.Album{
		Id:                  data.Get("id").String(),
		Title:               data.Get("title").String(),
		SharedWithCommunity: data.Get("shared_with_community").Bool(),
		Views:               data.Get("view_count").Int(),
		Upvotes:             data.Get("upvote_count").Int(),
		Downvotes:           data.Get("downvote_count").Int(),
		Comments:            data.Get("comment_count").Int(),
		CreatedAt:           createdAt,
		Media:               media,
		Tags:								 tags,
	}

	account := data.Get("account")
	if account.Raw != "" {
		album.User = types.User{
			Id: account.Get("id").Int(),
			Username: account.Get("username").String(),
			Avatar: strings.ReplaceAll(account.Get("avatar_url").String(), "https://i.imgur.com", ""),
		}
	}

	return album, nil
}
