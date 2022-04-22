package utils

import "os"

var Config map[string]interface{}

func LoadConfig() {
	port := "3000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	if os.Getenv("RIMGU_PORT") != "" {
		port = os.Getenv("RIMGU_PORT")
	}

	fiberPrefork := false
	if os.Getenv("FIBER_PREFORK") == "true" {
		fiberPrefork = true
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

	Config = map[string]interface{}{
		"port": port,
		"addr": addr,
		"imgurId": imgurId,
		"fiberPrefork": fiberPrefork,
	}
}