package main

import (
	"context"
	"fmt"
	"io/ioutil"

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
	//
	//reader, err := cli.ImagePull(ctx, "docker.io/library/alpine", types.ImagePullOptions{})
	//if err != nil {
	//	panic(err)
	//}
	//io.Copy(os.Stdout, reader)
	//
	body, err := cli.ContainerCreate(ctx, &container.Config{
		Image: "alpine",
		Cmd:   []string{"echo", "hello, world"},
		Tty:   true,
	}, nil, nil, "")
	if err != nil {
		panic(err)
	}

	if err := cli.ContainerStart(ctx, body.ID, types.ContainerStartOptions{}); err != nil {
		panic(err)
	}

	fmt.Println("container id is ", body.ID, body.Warnings)

	cid, err := cli.ContainerWait(ctx, body.ID)
	if err != nil {
		panic(err)
	}

	fmt.Println("cid is ", cid)

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

func exist(s []string, target string) bool {
	for i := range s {
		if target == s[i] {
			return true
		}
	}

	return false
}
