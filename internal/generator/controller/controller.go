package generator

import (
	"github.com/spf13/cobra"

	domain "github.com/auremsinistram/go-errgen/internal/generator/domain"
)

type CommandFunc func(cmd *cobra.Command, args []string) error

type controller struct {
	cmd *cobra.Command

	usecase domain.Usecase
}

func New(
	cmd *cobra.Command,
	usecase domain.Usecase,
) domain.Controller {
	return &controller{
		cmd:     cmd,
		usecase: usecase,
	}
}

func (c *controller) Expose() {
	c.cmd.RunE = c.Generate()
}
