package main

import (
	"cloud.google.com/go/storage"
	"context"
	"encoding/base64"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/iam/v1"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
	"os"
	"testing"
	"time"
)

func Test_ADCのチェック(t *testing.T)  {
	projectID := os.Getenv("PROJECT_ID")
	ctx := context.Background()

	// storageのclientを生成する
	storageClient, err := storage.NewClient(ctx)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("client : %v", storageClient)

	// ADCをチェックする
	sa, err := storageClient.ServiceAccount(ctx, projectID)
	t.Logf("SA : %v [%v]", sa, err)
}

var(
	iamService *iam.Service
)

func Test_署名URL(t *testing.T){
	//projectID := os.Getenv("PROJECT_ID")
	sa := os.Getenv("SA")
	ctx := context.Background()

	cli, err := google.DefaultClient(ctx, iam.CloudPlatformScope)
	if err != nil {
		t.Fatalf("%v", err)
	} else {
		t.Logf("cli : %+v", cli)
	}
	iamService, err = iam.New(cli)
	if err != nil {
		t.Fatalf("%v", err)
	}

	expires, _ := time.Parse(time.RFC3339, "2020-03-18T10:00:00-05:00")

	opt := &storage.SignedURLOptions{
		GoogleAccessID: sa,
		Method:         "GET",
		SignBytes:func(b []byte) ([]byte, error) {
			resp, err := iamService.Projects.ServiceAccounts.SignBlob(
				"projects/-/serviceAccounts/" + sa,
				&iam.SignBlobRequest{BytesToSign: base64.StdEncoding.EncodeToString(b)},
			).Context(ctx).Do()
			if err != nil {
				return nil, err
			}
			return base64.StdEncoding.DecodeString(resp.Signature)
		},
		Expires:expires,
		Scheme:storage.SigningSchemeV4,
	}

	url, err := storage.SignedURL("owncloud-backend", "test.txt", opt)

	t.Logf("URL : %v [%v]", url, err)
}

func Test_ADC(t *testing.T) {
	projectID := os.Getenv("PROJECT_ID")
	ctx := context.Background()

	creds, err := google.FindDefaultCredentials(ctx, storage.ScopeReadOnly)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("Creds %v", string(creds.JSON))
	client, err := storage.NewClient(ctx, option.WithCredentials(creds))
	if err != nil {
		t.Fatal(err)
	}
	t.Log("Buckets:")
	it := client.Buckets(ctx, projectID)
	for {
		battrs, err := it.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		t.Log(battrs.Name)
	}
}





