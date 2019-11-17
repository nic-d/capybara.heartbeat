package cmd

import (
	"capybara-heartbeat/api"
	"errors"
	"github.com/urfave/cli/v2"
)

// @todo - get the token from the args and pass that to the authenticate method
func Auth(c *cli.Context) error {
	res := api.Authenticate("helloworld")

	if res != true {
		return errors.New("Failed to authenticate")
	}

	return nil
}
