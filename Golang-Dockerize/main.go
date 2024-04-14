package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func main() {

	args := os.Args
	if len(args) < 2 {
		fmt.Println("Usage: main <folder-path>")
		return
	}
	codeDir := args[1]

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Docker start err: %v", err)
	}
	defer cli.Close()

	imageName := "docker.io/library/golang"

	reader, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		log.Fatalf("%s image err: %v", imageName, err)
	}
	defer reader.Close()
	io.Copy(os.Stdout, reader)
	fmt.Printf("%s image are well\n", imageName)

	fmt.Println("Container creating...")
	resp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "golang",
		Cmd:   []string{"sh", "-c", "go run *.go"},
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: codeDir,
				Target: "/go",
			},
		},
	}, nil, nil, "")
	if err != nil {
		log.Fatalf("Container create err: %v", err)
	}

	fmt.Println("Container starting...")
	if err := cli.ContainerStart(ctx, resp.ID, container.StartOptions{}); err != nil {
		log.Fatalf("Container start err: %v", err)
	}

	fmt.Println("Container waiting...")
	statusCh, errCh := cli.ContainerWait(ctx, resp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			log.Fatalf("Container waiting err: %v", err)
		}
	case <-statusCh:
	}

	fmt.Println("Container logs waiting...")
	out, err := cli.ContainerLogs(ctx, resp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		log.Fatalf("Container logs err: %v", err)
	}
	defer out.Close()
	fmt.Println("------------------")
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	fmt.Println("------------------")

}
