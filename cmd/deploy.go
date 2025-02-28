package cmd

import (
	"context"
	"fmt"

	"dagger.io/dagger"
	"github.com/spf13/cobra"
)

var deployCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deploy Envoy using Helmfile",
	Run: func(cmd *cobra.Command, args []string) {
		err := deployEnvoy()
		if err != nil {
			fmt.Println("Deployment failed:", err)
		} else {
			fmt.Println("✅ Envoy deployed successfully")
		}
	},
}

func deployEnvoy() error {
	ctx := context.Background()
	client, err := dagger.Connect(ctx)
	if err != nil {
		return err
	}
	defer client.Close()

	deployer := client.Container().
		From("alpine/helm:latest").
		WithMountedDirectory("/src", client.Host().Directory(".")).
		WithWorkdir("/src").
		WithExec([]string{"helmfile", "apply"})

	_, err = deployer.ExitCode(ctx)
	return err
}
