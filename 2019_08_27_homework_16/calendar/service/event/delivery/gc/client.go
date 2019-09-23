package gc

import (
	"context"
	"log"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Client struct {
	gc EventServiceClient
}

func NewClient() *Client {
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	return &Client{gc: NewEventServiceClient(cc)}
}

func (c *Client) Run(logger *logrus.Logger) {
	ev := &Event{
		ID:          0,
		Name:        "First event",
		Description: "This is my first event with grpc",
		Date:        "2019-08-22",
	}

	ev, err := c.Create(ev)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.WithFields(logrus.Fields{
		"event": ev.ID,
		"name":  ev.Name,
		"date":  ev.Date,
		"desc":  ev.Description,
	}).Info("was created!!!")

	ev, err = c.Read(ev.ID)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.WithFields(logrus.Fields{
		"event": ev.ID,
		"name":  ev.Name,
		"date":  ev.Date,
		"desc":  ev.Description,
	}).Info("was read!!!")

	ev.Name = "Updated First Event"

	updated, err := c.Update(ev)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Infof("update was %v", updated)

	deleted, err := c.Delete(ev.ID)
	if err != nil {
		logger.Fatalln(err)
	}

	logger.Infof("delete was %v", deleted)
}

func (c *Client) Create(ev *Event) (*Event, error) {
	resp, err := c.gc.Create(context.Background(), &CreateRequest{
		Event: ev,
	})

	if err != nil {
		return nil, err
	}

	return resp.Event, nil
}

func (c *Client) Read(eventID uint64) (*Event, error) {
	resp, err := c.gc.Read(context.Background(), &ReadRequest{
		Event_ID: eventID,
	})

	if err != nil {
		return nil, err
	}

	return resp.Event, nil
}

func (c *Client) Update(ev *Event) (bool, error) {
	resp, err := c.gc.Update(context.Background(), &UpdateRequest{
		Event: ev,
	})
	if err != nil {
		return false, err
	}

	return resp.Updated, nil
}

func (c *Client) Delete(eventID uint64) (bool, error) {
	resp, err := c.gc.Delete(context.Background(), &DeleteRequest{
		Event_ID: eventID,
	})
	if err != nil {
		return false, err
	}

	return resp.Deleted, nil
}
