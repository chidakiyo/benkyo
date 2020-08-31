package spanner

import (
	"cloud.google.com/go/spanner"
	"context"
	"fmt"
	"github.com/chidakiyo/benkyo/spanner/models"
	"github.com/google/uuid"
	"github.com/greymd/ojichat/generator"
	"log"
	"sync"
	"testing"
	"time"
)

func Test_Yoしてみる(t *testing.T) {
	ctx := context.Background()
	tx := con.ReadOnlyTransaction()
	defer tx.Close()

	user, err := models.ReadUser(ctx, tx, spanner.AllKeys())
	if err != nil {
		t.Errorf("%v", err)
	}

	for _, i := range user {
		t.Logf("%v", i)
	}
}

func Test_tweet投稿してみる(t *testing.T) {
	ctx := context.Background()
	start := time.Now()
	tt, err := con.ReadWriteTransaction(ctx, func(ctx context.Context, transaction *spanner.ReadWriteTransaction) error {
		log.Printf("■ %s", time.Now().Sub(start))
		text := message_gen()
		log.Printf("■ %s", time.Now().Sub(start))

		uuid, _ := uuid.NewRandom()
		log.Printf("■ %s", time.Now().Sub(start))

		tw := models.Tweet{
			ID:         "T_" + uuid.String(),
			UserID:     "U_74514912-1b78-4afb-8119-034f8e241e2b",
			Text:       text,
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
		}
		mut := tw.Insert(ctx)
		transaction.BufferWrite([]*spanner.Mutation{mut})
		return nil
	})
	log.Printf("■ %s", time.Now().Sub(start))
	t.Logf("%v, %v", tt, err)
}

func message_gen() string {
	text, _ := generator.Start(generator.Config{})
	return "[ojichat] " + text
}

func tweet_uuid_gen() string {
	uuid, _ := uuid.NewRandom()
	return "T_" + uuid.String()
}

func user_uuid_gen() string {
	uuid, _ := uuid.NewRandom()
	return "U_" + uuid.String()
}

func Test_データ大量に投入する(t *testing.T) {
	ctx := context.Background()

	tr := con.ReadOnlyTransaction()
	// アカウントを検索する
	users, err := models.ReadUser(ctx, tr, spanner.AllKeys())
	if err != nil {
		t.Errorf("%v", err)
	}
	tr.Close()

	var muChan = make(chan *spanner.Mutation, 10000)
	var wg sync.WaitGroup

	// FIXME この辺適当。なおす。
	go func() {
		for {
			var tmp []*spanner.Mutation
		BRK:
			for {
				select {
				case m := <-muChan:
					tmp = append(tmp, m)
					if len(tmp) > 100 {
						break BRK
					}
				}
			}
			_, err := con.ReadWriteTransaction(ctx, func(c context.Context, tr *spanner.ReadWriteTransaction) error {
				return tr.BufferWrite(tmp)
			})
			if err != nil {
				t.Logf("commit fail %v", err)
			} else {
				wg.Add(len(tmp) * -1)
				t.Logf("commit success count:%d", len(tmp))
			}
		}
	}()

	for _, u := range users {
		for i := 0; i < 10000; i++ {
			user := u
			ctx := ctx
			wg.Add(1)
			func() {
				//t.Logf("insert : %s", user.ID)
				start := time.Now()

				text := message_gen()
				uuid := tweet_uuid_gen()
				now := time.Now()
				tw := models.Tweet{
					ID:         uuid,
					UserID:     user.ID,
					Text:       text,
					CreatedAt:  now,
					ModifiedAt: now,
				}
				mut := tw.Insert(ctx)
				muChan <- mut
				t.Logf("time : %s [%P]", time.Now().Sub(start), err)
			}()
		}
	}
	close(muChan)
	wg.Wait()
}

func Benchmark_Spanner投入(b *testing.B) {
	ctx := context.Background()
	for i := 0; i < b.N; i++ {
		text := message_gen()
		uuid := tweet_uuid_gen()
		now := time.Now()
		tw := models.Tweet{
			ID:         uuid,
			UserID:     "U_74514912-1b78-4afb-8119-034f8e241e2b",
			Text:       text,
			CreatedAt:  now,
			ModifiedAt: now,
		}
		mut := tw.Insert(ctx)
		_, err := con.ReadWriteTransaction(ctx, func(c context.Context, tr *spanner.ReadWriteTransaction) error {
			b.Logf("Insert %v", tw)
			return tr.BufferWrite([]*spanner.Mutation{mut})
		})
		if err != nil {
			b.Logf("%v", err)
		}
	}
}

func Benchmark_文字列生成パフォーマンス(b *testing.B) {
	for i := 0; i < b.N; i++ {
		message_gen()
	}
}

func Benchmark_UUID生成パフォーマンス(b *testing.B) {
	for i := 0; i < b.N; i++ {
		tweet_uuid_gen()
	}
}

func insert_tweet(t *testing.T, ctx context.Context, user models.User) {
	t.Logf("insert : %s", user.ID)
	start := time.Now()

	text := message_gen()
	uuid := tweet_uuid_gen()
	now := time.Now()
	tw := models.Tweet{
		ID:         uuid,
		UserID:     user.ID,
		Text:       text,
		CreatedAt:  now,
		ModifiedAt: now,
	}
	mut := tw.Insert(ctx)
	_, err := con.ReadWriteTransaction(ctx, func(c context.Context, tr *spanner.ReadWriteTransaction) error {
		t.Logf("Insert %v", tw)
		return tr.BufferWrite([]*spanner.Mutation{mut})
	})
	t.Logf("time : %s [%P]", time.Now().Sub(start), err)
}

func Test_アカウントを大量に投入する(t *testing.T) {
	ctx := context.Background()
	var mutations []*spanner.Mutation
	start0 := time.Now()
	for i := 0; i < 10; i++ {
		start := time.Now()
		u := models.User{
			ID:         user_uuid_gen(),
			//UserID:     fmt.Sprintf("U%s", strings.Split(uuid, "-")[4]),
			UserID:     fmt.Sprintf("U%d", i),
			Email:      fmt.Sprintf("U%d@example.com", i),
			Password:   "password",
			CreatedAt:  time.Now(),
			ModifiedAt: time.Now(),
		}
		uu := u.Insert(ctx)
		mutations = append(mutations, uu)
		t.Logf("time : %s", time.Now().Sub(start))
	}
	_, err := con.ReadWriteTransaction(ctx, func(c context.Context, tr *spanner.ReadWriteTransaction) error {
		return tr.BufferWrite(mutations)
	})
	if err != nil {
		t.Errorf("%v", err)
	}
	t.Logf("time : %s", time.Now().Sub(start0))
}
