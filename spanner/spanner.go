package spanner

import (
	"cloud.google.com/go/spanner"
	"context"
	"errors"
	"fmt"
	"os"
	"time"
)

const healthCheckIntervalMins = 50
const numChannels = 4

func NewClient(ctx context.Context, projectID, instance, db string) (*spanner.Client, error) {
	dbPath := fmt.Sprintf("projects/%s/instances/%s/databases/%s", projectID, instance, db)
	client, err := spanner.NewClientWithConfig(ctx, dbPath,
		spanner.ClientConfig{
			SessionPoolConfig: spanner.SessionPoolConfig{
				//MinOpened:           100,
				MinOpened: 1,
				//MaxOpened:           numChannels * 100,
				MaxOpened:           1,
				MaxBurst:            10,
				WriteSessions:       0.2,
				HealthCheckWorkers:  10,
				HealthCheckInterval: healthCheckIntervalMins * time.Minute,
			},
		})
	if err != nil {
		return nil, errors.New("Failed to Create Spanner Client.")
	}
	return client, nil
}

func ConcreteNewClient(ctx context.Context) *spanner.Client {
	fmt.Println("client create start")
	SpannerProjectID, ok1 := os.LookupEnv("SPANNER_PROJECT_ID")
	SpannerInstance, ok2 := os.LookupEnv("SPANNER_INSTANCE")
	SpannerDB, ok3 := os.LookupEnv("SPANNER_DB")
	if !(ok1 && ok2 && ok3) {
		panic("環境変数がないね。")
	}
	c, err := NewClient(ctx, SpannerProjectID, SpannerInstance, SpannerDB)
	if err != nil {
		fmt.Println("client create NG")
		fmt.Errorf("%v", err)
	}
	fmt.Println("client create OK")
	return c
}
