package api

import (
	"regexp"
	"strings"
	"sync"
	"time"

	"codeberg.org/video-prize-ranch/rimgo/utils"
	"github.com/dustin/go-humanize"
	"github.com/microcosm-cc/bluemonday"
	"github.com/patrickmn/go-cache"
	"github.com/tidwall/gjson"
	"gitlab.com/golang-commonmark/linkify"
)

type Comment struct {
	Comments  []Comment
	User			User
	Id        string
	Comment   string
	Upvotes   int64
	Downvotes int64
	Platform  string
	CreatedAt string
	RelTime		string
	UpdatedAt string
	DeletedAt string
}

func (client *Client) FetchComments(galleryID string) ([]Comment, error) {
	cacheData, found := client.Cache.Get(galleryID + "-comments")
	if found {
		return cacheData.([]Comment), nil
	}

	data, err := utils.GetJSON("https://api.imgur.com/comment/v1/comments?client_id=" + client.ClientID + "&filter[post]=eq:" + galleryID + "&include=account,adconfig&per_page=30&sort=best")
	if err != nil {
		return []Comment{}, nil
	}

	wg := sync.WaitGroup{}
	comments := make([]Comment, 0)
	data.Get("data").ForEach(
		func(key, value gjson.Result) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()
				comments = append(comments, parseComment(value))
			}()

			return true
		},
	)
	wg.Wait()

	client.Cache.Set(galleryID + "-comments", comments, cache.DefaultExpiration)
	return comments, nil
}

var imgRe = regexp.MustCompile(`https?://i\.imgur\.com/(.*)\.(png|gif|jpg|webp)`)
var vidRe = regexp.MustCompile(`https?://i\.imgur\.com/(.*)\.(mp4|webm)`)
var vidFormatRe = regexp.MustCompile(`\.(mp4|webm)`)
var iImgurRe = regexp.MustCompile(`https?://i\.imgur\.com`)

func parseComment(data gjson.Result) Comment {
	createdTime, _ := time.Parse("2006-01-02T15:04:05Z", data.Get("created_at").String())
	createdAt := createdTime.Format("January 2, 2006 3:04 PM")
	updatedAt, _ := utils.FormatDate(data.Get("updated_at").String())
	deletedAt, _ := utils.FormatDate(data.Get("deleted_at").String())

	userAvatar := strings.ReplaceAll(data.Get("account.avatar").String(), "https://i.imgur.com", "")

	wg := sync.WaitGroup{}
	comments := make([]Comment, 0)
	data.Get("comments").ForEach(
		func(key, value gjson.Result) bool {
			wg.Add(1)

			go func() {
				defer wg.Done()
				comments = append(comments, parseComment(value))
			}()

			return true
		},
	)
	wg.Wait()

	comment := data.Get("comment").String()

	comment = strings.ReplaceAll(comment, "\n", "<br>")

	for _, match := range imgRe.FindAllString(comment, -1) {
		img := iImgurRe.ReplaceAllString(match, "")
		img = `<img src="` + img + `" class="comment__media" loading="lazy"/>`
		comment = strings.Replace(comment, match, img, 1)
	}
	for _, match := range vidRe.FindAllString(comment, -1) {
		vid := iImgurRe.ReplaceAllString(match, "")
		vid = `<video class="comment__media" controls loop preload="none" poster="` + vidFormatRe.ReplaceAllString(vid, ".webp") + `"><source type="` + strings.Split(vid, ".")[1] + `" src="` + vid + `" /></video>`
		comment = strings.Replace(comment, match, vid, 1)
	}
	for _, l := range linkify.Links(comment) {
		origLink := comment[l.Start:l.End]
		link := `<a href="` + origLink + `">` + origLink + `</a>`
		comment = strings.Replace(comment, origLink, link, 1)
	}

	p := bluemonday.UGCPolicy()
	p.AllowImages()
	p.AllowElements("video", "source")
	p.AllowAttrs("src", "tvpe").OnElements("source")
	p.AllowAttrs("controls", "loop", "preload", "poster").OnElements("video")
	p.AllowAttrs("class", "loading").OnElements("img", "video")
	p.RequireNoReferrerOnLinks(true)
	p.RequireNoFollowOnLinks(true)
	p.RequireCrossOriginAnonymous(true)
	comment = p.Sanitize(comment)

	return Comment{
		Comments: comments,
		User: User{
			Id:       data.Get("account.id").Int(),
			Username: data.Get("account.username").String(),
			Avatar:   userAvatar,
		},
		Id:        data.Get("id").String(),
		Comment:   comment,
		Upvotes:   data.Get("upvote_count").Int(),
		Downvotes: data.Get("downvote_count").Int(),
		Platform:  data.Get("platform").String(),
		CreatedAt: createdAt,
		RelTime: 	 humanize.Time(createdTime),
		UpdatedAt: updatedAt,
		DeletedAt: deletedAt,
	}
}
