package mastodon

import (
	"context"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"time"
)

type Rate struct {
	Remaining int64     `json:"remaining"`
	MaxLimit  int64     `json:"max_limit"`
	Reset     time.Time `json:"reset"`
}

// GetRate returns the API rate.
func (c *Client) GetRate(ctx context.Context) (*Rate, error) {
	u, err := url.Parse(c.Config.Server)
	if err != nil {
		return nil, err
	}
	u.Path = path.Join(u.Path, "/api/v1/notifications")
	req, err := http.NewRequest("GET", u.String(), nil)
	req = req.WithContext(ctx)
	if c.Config.AccessToken != "" {
		req.Header.Set("Authorization", "Bearer "+c.Config.AccessToken)
	}
	if err != nil {
		return nil, err
	}

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}

	header := res.Header
	remaining, _ := strconv.ParseInt(header["X-Ratelimit-Remaining"][0], 10, 64)
	maxLimit, err := strconv.ParseInt(header["X-Ratelimit-Limit"][0], 10, 64)
	if err != nil {
		return nil, err
	}
	resetStr := header["X-Ratelimit-Reset"][0]
	resetTimestamp, _ := time.Parse(time.RFC3339, resetStr)

	rate := Rate{
		Remaining: remaining,
		MaxLimit:  maxLimit,
		Reset:     resetTimestamp,
	}
	return &rate, nil
}
