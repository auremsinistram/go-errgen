package main

import (
	"fmt"
	"os"

	"github.com/auremsinistram/go-errgen/internal/generator"
	"github.com/auremsinistram/go-errors"
	"github.com/auremsinistram/go-toolkit/tools"
	"github.com/spf13/cobra"
)

var (
	inputPath    string
	econstPath   string
	responsePath string
	errorName    string
)

func main() {
	debug := tools.GetenvBool("DEBUG", false)

	rootCmd := &cobra.Command{
		Use:   "errgen",
		Short: "Generate error code constants and localization payloads from a JSON specification",
		Long: `Tool errgen reads a JSON array of error definitions and generates:
  - A Go source file with:
    * Integer error constants (hashed from the "name" field)
    * A map (ErrData) for status codes and default messages
    * A []byte constant (ErrRes) for unexpected errors (uses --error as fallback)
  - A JSON file mapping numeric error codes to localized messages (en_US, ru_RU, etc.)

The generated Go file is used internally for error handling.
The JSON file is served by an API endpoint to provide client-side error localization.`,
	}

	rootCmd.SilenceErrors = true

	flags := rootCmd.Flags()

	flags.StringVarP(
		&inputPath, "input", "i", "",
		"Path to input JSON file with error definitions (array of {status, name, messages})",
	)

	flags.StringVarP(
		&econstPath, "econst", "c", "",
		"Output path for the generated .go file (error constants, ErrData map, and ErrRes []byte)",
	)

	flags.StringVarP(
		&responsePath, "response", "r", "",
		"Output path for the generated .json file (maps error codes to localized messages for client API responses)",
	)

	flags.StringVarP(
		&errorName, "error", "e", "",
		"Name of the error used to generate the ErrRes []byte constant for fallback/unexpected error responses",
	)

	generator.Build(
		rootCmd,
		&inputPath,
		&econstPath,
		&responsePath,
		&errorName,
	)

	if err := rootCmd.Execute(); err != nil {
		var message string

		if debug {
			message = err.Error()
		} else {
			var e errors.Error

			if errors.As(errors.GetRoot(err), &e) {
				message = e.Description()
			} else {
				message = err.Error()
			}
		}

		fmt.Fprintln(os.Stderr, "Error:", message)

		os.Exit(1)
	}
}
