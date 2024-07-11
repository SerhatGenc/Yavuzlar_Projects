package docker

import (
	"context"
	cde "docCalc/codetester"
	"docCalc/structs"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/docker/docker/client"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
)

func DockerUp(req structs.Request, test structs.SingleTest) bool {
	inp, _ := strconv.ParseFloat(test.Input, 64)
	cde.CodeMaker(req, inp)
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	codeDir := wd + "/TestCode/"

	ctx := context.Background()
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Docker start err: %v", err)
	}
	defer cli.Close()

	imageName := "docker.io/library/golang"

	reader, err := cli.ImagePull(ctx, imageName, image.PullOptions{})
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
	if err != nil {
		log.Fatalf("Container logs err: %v", err)
	}
	defer out.Close()

	logOutput := readLogs(out)

	newstr2 := strings.Replace(logOutput, "\n", "", -1)
	newstr2 = newstr2[8:]
	fmt.Println("------------------")
	fmt.Println("sonuc:")
	fmt.Println(newstr2)
	fmt.Println("------------------")

	outputCodeFloat, err := strconv.ParseFloat(newstr2, 64)
	if err != nil {
		fmt.Println(err)
	}

	otp, _ := strconv.ParseFloat(test.Output, 64)
	fmt.Println("beklenen:\n" + test.Output)
	if outputCodeFloat == otp {
		return true
	} else {
		return false
	}
}

func readLogs(reader io.Reader) string {
	buf := new(strings.Builder)
	_, err := io.Copy(buf, reader)
	if err != nil {
		log.Fatalf("Read logs err: %v", err)
	}
	return buf.String()
}
