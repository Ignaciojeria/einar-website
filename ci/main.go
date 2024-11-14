package main

import (
	"context"
	"fmt"
	"os"

	run "cloud.google.com/go/run/apiv2"
	"cloud.google.com/go/run/apiv2/runpb"
	"dagger.io/dagger"
	"github.com/joho/godotenv"
)

const GCR_SERVICE_URL = "projects/einar-404623/locations/us-central1/services/einar"
const GCR_PUBLISH_ADDRESS = "gcr.io/einar-404623/einar"

func main() {
	godotenv.Load()

	// create Dagger client
	ctx := context.Background()
	daggerClient, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stderr))
	if err != nil {
		panic(err)
	}
	defer daggerClient.Close()

	// get working directory on host
	source := daggerClient.Host().Directory(".", dagger.HostDirectoryOpts{
		Exclude: []string{"ci"},
	})

	// build application
	builder := daggerClient.Container(dagger.ContainerOpts{Platform: "linux/amd64"}).
		From("golang:1.23.3").
		WithDirectory("/src", source).
		WithWorkdir("/src").
		WithEnvVariable("CGO_ENABLED", "0").
		WithExec([]string{"go", "build", "-o", "myapp"})

	// Validate the build by executing the binary
	_, err = builder.WithExec([]string{"./myapp"}).ExitCode(ctx)
	if err != nil {
		panic("Build validation failed: " + err.Error())
	}
	fmt.Println("Build validation passed!")

	// add binary to alpine base
	prodImage := daggerClient.Container(dagger.ContainerOpts{Platform: "linux/amd64"}).
		From("alpine").
		WithFile("/bin/myapp", builder.File("/src/myapp")).
		WithEntrypoint([]string{"/bin/myapp"})

	// Test the container locally
	testContainer := prodImage.WithExec([]string{"ping", "-c", "1", "localhost"})
	_, err = testContainer.ExitCode(ctx)
	if err != nil {
		panic("Container ping failed: " + err.Error())
	}
	fmt.Println("Container ping test passed!")

	// publish container to Google Container Registry
	addr, err := prodImage.Publish(ctx, GCR_PUBLISH_ADDRESS)
	if err != nil {
		panic(err)
	}

	// print ref
	fmt.Println("Published at:", addr)

	// create Google Cloud Run client
	gcrClient, err := run.NewServicesClient(ctx)
	if err != nil {
		panic(err)
	}
	defer gcrClient.Close()

	// define service request
	gcrRequest := &runpb.UpdateServiceRequest{
		Service: &runpb.Service{
			Name: GCR_SERVICE_URL,
			Template: &runpb.RevisionTemplate{
				Containers: []*runpb.Container{
					{
						Image: addr,
						Ports: []*runpb.ContainerPort{
							{
								Name:          "http1",
								ContainerPort: 80,
							},
						},
					},
				},
			},
		},
	}

	// update service
	gcrOperation, err := gcrClient.UpdateService(ctx, gcrRequest)
	if err != nil {
		panic(err)
	}

	// wait for service request completion
	gcrResponse, err := gcrOperation.Wait(ctx)
	if err != nil {
		panic(err)
	}

	// print ref
	fmt.Println("Deployment for image", addr, "now available at", gcrResponse.Uri)
}
