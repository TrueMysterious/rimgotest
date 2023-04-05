package utils

import (
	"log"
	"os"
	"time"
)

type config struct {
	Port            string
	Addr            string
	ImgurId         string
	FiberPrefork    bool
	ImageCache      bool
	CleanupInterval time.Duration
	CacheDir        string
	Privacy         map[string]interface{}
}

var Config config

func LoadConfig() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("RIMGU_PORT") != "" {
		port = os.Getenv("RIMGU_PORT")
	}

	addr := "0.0.0.0"
	if os.Getenv("ADDRESS") != "" {
		addr = os.Getenv("ADDRESS")
	}
	if os.Getenv("RIMGU_ADDRESS") != "" {
		addr = os.Getenv("RIMGU_ADDRESS")
	}

	imgurId := "546c25a59c58ad7"
	if os.Getenv("IMGUR_CLIENT_ID") != "" {
		imgurId = os.Getenv("IMGUR_CLIENT_ID")
	}
	if os.Getenv("RIMGU_IMGUR_CLIENT_ID") != "" {
		imgurId = os.Getenv("RIMGU_IMGUR_CLIENT_ID")
	}

	imageCache := os.Getenv("IMAGE_CACHE") == "true"

	cleanupInterval, err := time.ParseDuration(os.Getenv("IMAGE_CACHE_CLEANUP_INTERVAL"))
	if err != nil && imageCache {
		log.Fatal("invalid configuration: invalid duration for IMAGE_CACHE_CLEANUP_INTERVAL")
	}

	cacheDir := os.Getenv("IMAGE_CACHE_DIR")
	if cacheDir == "" && imageCache {
		log.Fatal("invalid configuration: no IMAGE_CACHE_DIR")
	}

	Config = config{
		Port:            port,
		Addr:            addr,
		ImgurId:         imgurId,
		FiberPrefork:    os.Getenv("FIBER_PREFORK") == "true",
		ImageCache:      imageCache,
		CleanupInterval: cleanupInterval,
		CacheDir:        cacheDir,
		Privacy: map[string]interface{}{
			"set":           os.Getenv("PRIVACY_NOT_COLLECTED") != "",
			"policy":        os.Getenv("PRIVACY_POLICY"),
			"message":       os.Getenv("PRIVACY_MESSAGE"),
			"country":       os.Getenv("PRIVACY_COUNTRY"),
			"provider":      os.Getenv("PRIVACY_PROVIDER"),
			"cloudflare":    os.Getenv("PRIVACY_CLOUDFLARE") == "true",
			"not_collected": os.Getenv("PRIVACY_NOT_COLLECTED") == "true",
			"ip":            os.Getenv("PRIVACY_IP") == "true",
			"url":           os.Getenv("PRIVACY_URL") == "true",
			"device":        os.Getenv("PRIVACY_DEVICE") == "true",
			"diagnostics":   os.Getenv("PRIVACY_DIAGNOSTICS") == "true",
		},
	}
}
