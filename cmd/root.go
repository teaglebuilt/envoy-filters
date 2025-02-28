import (
	"fmt"
	"github.com/spf13/cobra"
)

// Root command
var rootCmd = &cobra.Command{
	Use:   "envoy-cli",
	Short: "CLI for generating Envoy config, testing, and deploying filters",
	Long:  "A command-line tool for managing Envoy configurations, running tests, and deploying using Dagger & Helmfile",
}

// Execute runs the root command
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}

func init() {
	rootCmd.AddCommand(buildCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.AddCommand(deployCmd)
}