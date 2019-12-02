package redis

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/go-redis/cache/v7"
	"github.com/go-redis/redis/v7"
	"github.com/vmihailenco/msgpack/v4"
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
	client = NewClient() // setup
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
		statusCmd := client.Set(key, value, 0)
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
		statusCmd := client.Set(key, value, 1*time.Second) // TTL 2秒j
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

// 構造体のSET/GET gobを利用する
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

func partialCodec() func(*redis.Client) *cache.Codec {
	return func(client *redis.Client) *cache.Codec {
		codec := &cache.Codec{
			Redis: client,
			Marshal: func(v interface{}) ([]byte, error) {
				return msgpack.Marshal(v)
			},
			Unmarshal: func(b []byte, v interface{}) error {
				return msgpack.Unmarshal(b, v)
			},
		}
		return codec
	}
}

func Test_構造体を入れる(t *testing.T) {
	client.Get("hoge").String()

	codec := &cache.Codec{
		Redis: client,
		Marshal: func(v interface{}) ([]byte, error) {
			return msgpack.Marshal(v)
		},
		Unmarshal: func(b []byte, v interface{}) error {
			return msgpack.Unmarshal(b, v)
		},
	}

	m := CreateAMock()

	codec.Set(&cache.Item{
		Key:        key,
		Object:     m,
		Expiration: time.Hour,
	})

	var mm A
	if err := codec.Get(key, &mm); err == nil {
		t.Logf("result: %#v", mm)
	}
}

func Test_onceの振る舞い確認(t *testing.T) {
	codec := partialCodec()(client)

	var mm A
	err := codec.Once(&cache.Item{
		Key:    "key-once",
		Object: &mm,
		Func: func() (i interface{}, e error) {
			t.Logf("ココ実行されたよ")
			return CreateAMock(), nil
		},
		Expiration: 10 * time.Second,
		//Ctx:
	})

	if err != nil {
		t.Fatalf("fail %s", err.Error())
	}

	t.Logf("success: %#v", mm)
}
