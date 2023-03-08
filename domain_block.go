package mastodon

import (
	"context"
	"net/http"
)

type DomainBlock struct {
	Domain   string `json:"domain"`
	Digest   string `json:"digest"`
	Severity string `json:"severity"`
}

// GetDomainBlocks returns all the rules on the current instance.
func (c *Client) GetDomainBlocks(ctx context.Context) ([]*DomainBlock, error) {
	var domainBlock []*DomainBlock
	err := c.doAPI(ctx, http.MethodGet, "/api/v1/instance/domain_blocks", nil, &domainBlock, nil)
	if err != nil {
		return nil, err
	}
	return domainBlock, nil
}
