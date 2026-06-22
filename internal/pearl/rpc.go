package pearl

import (
	"fmt"
	"net/http"
	"time"
)

type RPCConfig struct {
	Host     string
	Port     int
	User     string
	Password string
}

func (c RPCConfig) URL() string {
	return "http://" + c.Host + ":" + fmt.Sprintf("%d", c.Port)
}

func NewHTTPClient() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second,
	}
}
