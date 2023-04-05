package pages

import (
	"codeberg.org/video-prize-ranch/rimgo/api"
	"codeberg.org/video-prize-ranch/rimgo/utils"
)

var ApiClient *api.Client

func InitializeApiClient() {
    ApiClient = api.NewClient(utils.Config.ImgurId)
}