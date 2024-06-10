package agent

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
	dimage "github.com/docker/docker/api/types/image"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"time"
)

type Docker struct {
	cli *client.Client
}

func NewDocker() *Docker {
	d := &Docker{}
	var err error
	if d.cli, err = client.NewClientWithOpts(); err != nil {
		panic(err)
	}
	return d
}

func (d *Docker) PullImage(image string) error {
	_, err := d.cli.ImagePull(context.Background(), image, dimage.PullOptions{})
	return err
}

func (d *Docker) FindContainerByUUID(uuid string) (types.Container, error) {
	searchme := filters.NewArgs()
	searchme.Add("Name", fmt.Sprintf("/%s", uuid))
	boxes, err := d.cli.ContainerList(context.Background(), container.ListOptions{
		Filters: searchme,
	})
	if err != nil {
		return types.Container{}, err
	}
	if len(boxes) == 0 {
		return types.Container{}, fmt.Errorf("container not found")
	}
	return boxes[0], nil
}

func (d *Docker) WatchContainer(uuid string) (running chan bool) {
	running = make(chan bool)
	go func() {
		for {
			c, err := d.FindContainerByUUID(uuid)
			if err != nil {
				running <- false
				return
			}
			if c.State == "running" {
				running <- true
				return
			}
			<-time.After(1 * time.Second)
		}
	}()
	return running
}

func (d *Docker) StopContainer(id string) error {
	return d.cli.ContainerStop(context.Background(), id, container.StopOptions{})
}

func (d *Docker) StartContainer(cmd []string, image string, binds []string, wispuuid string) (id string, err error) {
	box, err := d.cli.ContainerCreate(
		context.Background(),
		&container.Config{
			Cmd:   cmd,
			Image: image,
		},
		&container.HostConfig{
			Binds:         binds,
			NetworkMode:   container.NetworkMode(fmt.Sprintf("container:/%s", wispuuid)),
			RestartPolicy: container.RestartPolicy{Name: container.RestartPolicyAlways},
			AutoRemove:    true,
		},
		&network.NetworkingConfig{},
		nil,
		fmt.Sprintf("%s-sidecar", wispuuid),
	)
	return box.ID, err
}
