package mastodon

import (
	"context"
	"net/http"
)

type Rule struct {
	ID   string `json:"id"`
	Text string `json:"text"`
}

// GetRules returns all the rules on the current instance.
func (c *Client) GetRules(ctx context.Context) ([]*Rule, error) {
	var rules []*Rule
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/instance/rules", nil, &rules, nil)
	if err != nil {
		return nil, err
	}
	return rules, nil
}
