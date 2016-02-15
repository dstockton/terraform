package rightscale

import (
	"log"

	"github.com/rightscale/rsc/cm15"
	"github.com/rightscale/rsc/rsapi"
)

type Config struct {
	RefreshToken string
	AccountID    int
}

// Client() returns a new client for accessing RightScale.
func (c *Config) Client() (*cm15.API, error) {
	auth := rsapi.NewOAuthAuthenticator(c.RefreshToken, c.AccountID)
	client := cm15.New("us-3.rightscale.com", auth)

	log.Printf("[INFO] RightScale Client configured for account: %s", c.AccountID)

	return client, nil
}
