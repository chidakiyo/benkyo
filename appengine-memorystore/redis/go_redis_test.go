package redis

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/go-redis/redis/v7"
	"os"
	"testing"
	"time"
)

var client *redis.Client

const (
	key   = "hoge"
	value = "hoge-value"
)

func TestMain(m *testing.M) {
	client = NewClient()
	ret := m.Run()
	os.Exit(ret)
}

func NewClient() *redis.Client {
	redisHost := os.Getenv("REDISHOST")
	redisPort := os.Getenv("REDISPORT")
	client := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisHost, redisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	//pong, err := client.Ping().Result()
	//fmt.Println(pong, err)
	// Output: PONG <nil>
	return client // TODO error適当
}

// 通常のSET/GET
func Test_普通のSetGet(t *testing.T) {
	// set
	{
		statusCmd := client.Set(key, value,0)
		if statusCmd.Err() != nil {
			t.Fatalf("error!!! %s", statusCmd.Err())
		}
		t.Logf("set result: %#v", statusCmd)
	}
	// get
	{
		val, err := client.Get(key).Result()
		if err != nil {
			t.Logf("noting %s", err)
			return
		}
		t.Logf("get result: %s %s", key, val)
	}
}

func Test_TTL(t *testing.T) {
	// set
	{
		statusCmd := client.Set(key, value,1 *time.Second) // TTL 2秒j
		if statusCmd.Err() != nil {
			t.Fatalf("error!!! %s", statusCmd.Err())
		}
		t.Logf("set result: %#v", statusCmd)
	}

	// すぐ取得
	{
		val, err := client.Get(key).Result()
		if err != nil {
			t.Logf("noting %s", err)
			return
		}
		t.Logf("get result: %s %s", key, val)
	}

	// 2秒後に取得
	time.Sleep(2 * time.Second)
	{
		val, err := client.Get(key).Result()
		if err != nil {
			t.Logf("noting %s", err)
			return
		}
		t.Logf("get result: %s %s", key, val)
	}
}


type A struct {
	Id   string
	Name string
	Age  int64
	Tags []string
	Type []B
}

type B struct {
	No   int64
	Name string
}

func CreateAMock() A {
	return A{
		Id:   "ID-A",
		Name: "Name-A",
		Age:  60,
		Tags: []string{"TAG-A", "TAG-B"},
		Type: []B{
			{No: 1, Name: "Type-A"},
			{No: 2, Name: "Type-B"},
		},
	}
}

// 構造体のSET/GET
func Test_構造体をbyteで(t *testing.T) {
	// set
	{
		m := CreateAMock()
		buf := bytes.NewBuffer(nil)
		_ = gob.NewEncoder(buf).Encode(&m) // TODO errorハンドリング

		result := client.Set("key-b", buf.Bytes(), 0)
		if result.Err() != nil {
			t.Fatalf("%v", result.Err())
		}
		t.Logf("Set ok.")
	}

	// get
	{
		var m A
		b, _ := client.Get("key-b").Bytes() // TODO error
		buf := bytes.NewBuffer(b)
		_ = gob.NewDecoder(buf).Decode(&m) // TODO error

		t.Logf("result %#v", m)
	}

}
