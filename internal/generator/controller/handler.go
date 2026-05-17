package generator

import (
	"github.com/auremsinistram/go-errors"
	"github.com/spf13/cobra"
)

func (c *controller) Generate() CommandFunc {
	return func(cmd *cobra.Command, args []string) error {
		if err := c.usecase.Generate(); err != nil {
			return errors.Wrap(err, "controller - Generate - #1")
		}

		return nil
	}
}
