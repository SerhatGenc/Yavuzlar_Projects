package functions

import (
	stc "Calc/Structs"
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/mount"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

func runDocker(requests stc.RunRequest, w http.ResponseWriter, y int) string {
	tmpDir := "/home/elijah"
	tmpFile := filepath.Join(tmpDir, "tmpfile.go")
	if err := os.WriteFile(tmpFile, []byte(AddString(requests.Code, requests.Fonk, y)), 0644); err != nil {
		http.Error(w, fmt.Sprintf("Error writing temp file: %v", err), http.StatusInternalServerError)
		return "Run docker error 1"
	}

	ctx := context.Background()

	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatalf("Error :NewCLientWithOpt: %v", err)
	}
	defer cli.Close()
	contImage := "docker.io/library/golang"

	read, err := cli.ImagePull(ctx, contImage, image.PullOptions{})
	if err != nil {
		log.Fatalf("Error :ImagePull: %v", err)
	}
	defer read.Close()

	io.Copy(os.Stdout, read)

	path := "/home/elijah/tmpfile.go"
	fmt.Println(path)
	createResp, err := cli.ContainerCreate(ctx, &container.Config{
		Image: contImage,
		Cmd:   []string{"go", "run", "/app" + path},
	}, &container.HostConfig{
		Mounts: []mount.Mount{
			{
				Type:   mount.TypeBind,
				Source: path,
				Target: "/app" + path,
			},
		},
	},
		nil, nil, "")
	if err != nil {
		log.Fatalf("Error :ContainerCreate: %v", err)
	}

	if err := cli.ContainerStart(ctx, createResp.ID, container.StartOptions{}); err != nil {
		log.Fatalf("Error :ContainerStart: %v", err)
	}

	statusCh, errCh := cli.ContainerWait(ctx, createResp.ID, container.WaitConditionNotRunning)
	select {
	case err := <-errCh:
		if err != nil {
			log.Fatalf("Error :ContainerWait: %v", err)
		}
	case <-statusCh:
	}

	output, err := cli.ContainerLogs(ctx, createResp.ID, container.LogsOptions{ShowStdout: true, ShowStderr: true})
	if err != nil {
		log.Fatalf("Error :ContainerLogs: %v", err)
	}
	defer output.Close()

	var outputCode bytes.Buffer
	var hata bytes.Buffer
	_, err = stdcopy.StdCopy(&outputCode, &hata, output)
	if err != nil {
		log.Fatalf("Error :Stdout: %v", err)
	}

	outputCodeStr := outputCode.String()

	fmt.Println("Output:", outputCodeStr)
	return outputCodeStr
}

func AddString(code string, funcname string, control int) string {
	data, err := os.ReadFile("temple.txt")
	if err != nil {
		fmt.Println("File read error:", err)
		return ""
	}

	text := string(data)
	text = text + code
	text = strings.Replace(text, "tempfonk", funcname, -1)
	text = strings.Replace(text, "tempvar", strconv.Itoa(control), -1)
	fmt.Println(text)
	return text

}
