package api

import (
	"time"

	"github.com/patrickmn/go-cache"
)

type Client struct {
    ClientID string
    Cache    *cache.Cache
}

func NewClient(clientId string) (*Client) {
    client := Client{
        ClientID: clientId,
        Cache: cache.New(15*time.Minute, 15*time.Minute),
    }
    
    return &client
}