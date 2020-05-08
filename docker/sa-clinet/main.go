package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"os"

	"github.com/docker/docker/api/types/container"

	"github.com/docker/docker/api/types"

	"github.com/docker/docker/client"
)

/*
コンテナを作ってログを出す
*/

func main() {
	ctx := context.Background()
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	defer cli.ClientVersion()

	reader, err := cli.ImagePull(ctx, "docker.io/library/golang:1.12.0", types.ImagePullOptions{})
	if err != nil {
		panic(err)
	}
	io.Copy(os.Stdout, reader)

	body, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "golang:1.12.0",
		Cmd:   []string{"sh"},
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, body.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println("container id is ", body.ID, body.Warnings)

	out, err := cli.ContainerLogs(ctx, body.ID, types.ContainerLogsOptions{
		ShowStdout: true,
	})
	if err != nil {
		panic(err)
	}
	b, err := ioutil.ReadAll(out)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(b))

}
