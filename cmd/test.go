package cmd

import (
	"context"
	"fmt"

	"dagger.io/dagger"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run unit and integration tests",
	Run: func(cmd *cobra.Command, args []string) {
		err := runTests()
		if err != nil {
			fmt.Println("Tests failed:", err)
		} else {
			fmt.Println("✅ Tests passed")
		}
	},
}

func runTests() error {
	ctx := context.Background()
	client, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	tester := client.Container().
		From("golang:latest").
		WithMountedDirectory("/src", client.Host().Directory(".")).
		WithWorkdir("/src").
		WithExec([]string{"go", "test", "./..."})

	_, err = tester.ExitCode(ctx)
	return err
}
