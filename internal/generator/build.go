package generator

import (
	"github.com/spf13/cobra"

	controller "github.com/auremsinistram/go-errgen/internal/generator/controller"
	usecase "github.com/auremsinistram/go-errgen/internal/generator/usecase"
)

func Build(
	cmd *cobra.Command,
	inputPath *string,
	econstPath *string,
	responsePath *string,
	errorName *string,
) {
	controller.New(
		cmd,
		usecase.New(
			inputPath,
			econstPath,
			responsePath,
			errorName,
		),
	).Expose()
}
