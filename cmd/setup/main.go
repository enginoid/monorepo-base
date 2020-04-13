package main

import (
	"context"
	"fmt"
	"log"
	"os"

	container "cloud.google.com/go/container/apiv1"
	"github.com/Songmu/prompter"
	cli "github.com/urfave/cli/v2"

	"github.com/enginoid/monorepo-base/cmd/setup/internal/gcloud"
)

func main() {
	app := &cli.App{
		Name:  "setup",
		Usage: "helps you set up your monorepo",
		Action: func(c *cli.Context) error {
			ctx := context.Background()
			_, err := container.NewClusterManagerClient(ctx)
			if err != nil {
				return err
			}

			exists := gcloud.Exists()
			if !exists {
				return fmt.Errorf("gcloud needs to be installed: https://cloud.google.com/sdk/install")
			}

			project, err := gcloud.CurrentProject()
			if err != nil {
				return fmt.Errorf("failed to get current gcloud project: %v", err)
			}

			fmt.Printf("👉 The default gcloud project is: %s\n", project)
			fmt.Println("👆 This script will help you create clusters in this project.")
			fmt.Println("⚠️ To do that, Terraform and this setup script will create some resources.")

			if !prompter.YN("Do you want to continue and create resources in this project?", false) {
				fmt.Println("👍  If you want to use a different project:")
				fmt.Println("   1️⃣ (Optional) Create a new project via 'gcloud projects create [PROJECT_NAME]'.")
				fmt.Println("   2️⃣ Set it as your default project 'gcloud config set core/project [PROJECT_NAME]'.")
				fmt.Println("   3️⃣ Make sure to enable billing on the project.")
				fmt.Println("      🔗 https://console.developers.google.com/billing/linkedaccount?project=[PROJECT_NAME]")
				fmt.Println("   4️⃣ Re-run this script. It operates on the default project.")
				return nil
			}

			requiredServices := []string{
				"compute.googleapis.com",
				"container.googleapis.com",
				"cloudresourcemanager.googleapis.com",
			}

			fmt.Printf("⏳ Checking whether %s has all required services enabled...\n", project)
			enabledServices, err := gcloud.ListEnabledServices(project)
			if err != nil {
				return err
			}

			missingServices := make([]string, 0, len(requiredServices))
			for _, service := range requiredServices {
				if enabledServices[service] {
					fmt.Printf("✅ %s is enabled.\n", service)
					continue
				}

				fmt.Printf("⚠️ %s needs to be enabled.\n", service)
				missingServices = append(missingServices, service)
			}

			if len(missingServices) > 0 {
				fmt.Println()
				if prompter.YN("Do you want to enable these services?", false) {
					gcloud.EnableServices(project, missingServices)
				}

				return fmt.Errorf("can't proceed witout enabling services")
			}

			fmt.Println("🔓 The script is now going to put credentials for this project into a known location for Terraform to access it.")
			fmt.Println("🌍 Your browser is going to prompt you to authorize some scopes for these credentials.")
			if !prompter.YN("Ready?", true) {
				fmt.Println("Alternatively, you can finish this last step by running:")
				fmt.Println("   gcloud auth application-default login")
				return nil
			}

			err = gcloud.AcquireDefaultCredentials(project)
			if err != nil {
				return err
			}

			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
