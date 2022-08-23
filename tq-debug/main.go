package main

import (
	cloudtasks "cloud.google.com/go/cloudtasks/apiv2"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	taskspb "google.golang.org/genproto/googleapis/cloud/tasks/v2"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var signalChan = make(chan os.Signal, 1)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8081"
		log.Printf("defaulting to port %s", port)
	}

	g := gin.Default()

	g.GET("/add", put)
	g.POST("/push", push)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: g,
	}

	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("listening on port %s", port)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}()

	sig := <-signalChan
	log.Printf("%s signal caught", sig)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		log.Printf("server shutdown failed: %+v", err)
	}
	log.Print("server exited")
}

func put(c *gin.Context) {
	projectID := os.Getenv("PROJECT_ID")
	locationID := os.Getenv("LOCATION")
	queue := os.Getenv("QUEUE")
	url := os.Getenv("URL")

	_, err := createHTTPTask(c.Request.Context(), projectID, locationID, queue, url+"/tasks/exec", "hello")
	if err != nil {
		fmt.Printf("%+v\n", err)
		c.String(http.StatusBadRequest, "")
		return
	}

	c.String(http.StatusOK, "OK")
}

func push(c *gin.Context) {
	for k, v := range c.Request.Header {
		fmt.Printf("key: %s .. value: %+v\n", k, v)
	}
	b, _ := io.ReadAll(c.Request.Body)
	fmt.Println("来たよ ", string(b))
}

func createHTTPTask(ctx context.Context, projectID, locationID, queueID, url, message string) (*taskspb.Task, error) {

	client, err := cloudtasks.NewClient(ctx)
	if err != nil {
		return nil, fmt.Errorf("NewClient: %v", err)
	}
	defer client.Close()

	queuePath := fmt.Sprintf("projects/%s/locations/%s/queues/%s", projectID, locationID, queueID)

	req := &taskspb.CreateTaskRequest{
		Parent: queuePath,
		Task: &taskspb.Task{
			MessageType: &taskspb.Task_HttpRequest{
				HttpRequest: &taskspb.HttpRequest{
					HttpMethod: taskspb.HttpMethod_POST,
					Url:        url,
				},
			},
		},
	}
	req.Task.GetHttpRequest().Body = []byte(message)

	createdTask, err := client.CreateTask(ctx, req)
	if err != nil {
		return nil, fmt.Errorf("cloudtasks.CreateTask: %v", err)
	}
	return createdTask, nil
}
