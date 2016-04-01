// Author hoenig

package master

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/shoenig/subspace/core/common/stream"
	"github.com/shoenig/subspace/core/config"
	"github.com/shoenig/toolkit"
)

// A Client is a wrapper for an http client that attempts to communicate
// with all of the known masters, with hopes of connecting to one of them.
type Client struct {
	masters config.Masters
	client  *http.Client
}

// NewClient creates a new http client for communicating with one of the available masters.
func NewClient(masters config.Masters) *Client {
	return &Client{
		masters: masters,
		client: &http.Client{
			Timeout: 30 * time.Second,
		},
	}
}

// CreateStream is used for sending a Stream creation request.
func (c *Client) CreateStream(creation stream.Creation) error {
	js, err := creation.JSON()
	if err != nil {
		return err
	}
	return c.doPOST("/v1/stream/create", js)
}

// Publish is used to announce the availability of a new Bundle.
func (c *Client) Publish(bundle stream.Bundle) error {
	js, err := bundle.JSON()
	if err != nil {
		return err
	}
	return c.doPOST("/v1/stream/publish", js)
}

// attempt to POST some json to the masters
func (c *Client) doPOST(endpoint, body string) error {
	for _, master := range c.masters {
		r := strings.NewReader(body)
		resp, err := c.client.Post(master.API(endpoint), "application/json", r)
		if err != nil {
			log.Println("post to", master, "failed:", err)
			continue
		}
		defer toolkit.Drain(resp.Body)
		// confirm the result?

		return nil
	}

	return fmt.Errorf("client failed to post to any master: %v", c.masters)
}
