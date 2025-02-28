package cmd

import (
	"context"
	"fmt"

	"dagger.io/dagger"
	"github.com/spf13/cobra"
)

var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Build the WASM filter",
	Run: func(cmd *cobra.Command, args []string) {
		err := buildWasm()
		if err != nil {
			fmt.Println("Error building WASM filter:", err)
		} else {
			fmt.Println("✅ WASM filter built successfully")
		}
	},
}

func buildWasm() error {
	ctx := context.Background()
	client, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	builder := client.Container().
		From("tinygo/tinygo:latest").
		WithMountedDirectory("/src", client.Host().Directory(".")).
		WithWorkdir("/src").
		WithExec([]string{"tinygo", "build", "-o", "output/filter.wasm", "-target=wasi", "main.go"})

	_, err = builder.File("output/filter.wasm").Export(ctx, "./output/filter.wasm")
	return err
}
