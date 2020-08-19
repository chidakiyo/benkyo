package spanner

import (
	"cloud.google.com/go/spanner"
	"context"
	"google.golang.org/api/iterator"
	"os"
	"strings"
	"testing"
	"time"
)

var con *spanner.Client

func TestMain(m *testing.M) {
	println("before all...")
	ctx := context.Background()
	con = ConcreteNewClient(ctx)
	defer con.Close()
	code := m.Run()
	println("after all...")
	os.Exit(code)
}

func Test_アカウント10件作成する(t *testing.T) {
	ctx := context.Background()
	_, err := con.ReadWriteTransaction(ctx, func(ctx context.Context, transaction *spanner.ReadWriteTransaction) error {
		stmt := spanner.Statement{
			SQL: `INSERT User (ID, UserID, Email,  Password, CreatedAt, ModifiedAt) VALUES
					('74514912-1b78-4afb-8119-034f8e241e2b', 'U1' ,'U1@example.com' ,'password','2016-12-25 05:30:00+07', '2016-12-25 05:30:00+07'),
					('b8dfe98a-db5c-4177-83b8-ea0326a12032', 'U2' ,'U2@example.com' ,'password','2016-12-25 05:30:00+07', '2016-12-25 05:30:00+07'),
					('d684b2a5-55ac-4d7b-80be-e5d9157a2b7d', 'U3' ,'U3@example.com' ,'password','2016-12-25 05:30:00+07', '2016-12-25 05:30:00+07'),
					('7647dd05-7e8d-407b-8582-d8dc06defe43', 'U4' ,'U4@example.com' ,'password','2016-12-25 05:30:00+07', '2016-12-25 05:30:00+07'),
					('e6f36cb0-4ece-4618-9dde-9440dd247acb', 'U5' ,'U5@example.com' ,'password','2016-12-25 05:30:00+07', '2016-12-25 05:30:00+07'),
					('2bc48d3a-f709-4cfc-b46f-6642ed5aa598', 'U6' ,'U6@example.com' ,'password','2016-12-25 05:30:00+07', '2016-12-25 05:30:00+07'),
					('8d00e09b-1ffe-4393-be5f-9cb2e46d81c1', 'U7' ,'U7@example.com' ,'password','2016-12-25 05:30:00+07', '2016-12-25 05:30:00+07'),
					('7b0d2492-b2b2-4176-9cf5-47364e192f1a', 'U8' ,'U8@example.com' ,'password','2016-12-25 05:30:00+07', '2016-12-25 05:30:00+07'),
					('992c5a0b-f93e-4362-8d44-9060a6c2115e', 'U9' ,'U9@example.com' ,'password','2016-12-25 05:30:00+07', '2016-12-25 05:30:00+07'),
					('92f96ab5-400a-4211-aa93-5312096194b2', 'U10','U10@example.com','password','2016-12-25 05:30:00+07', '2016-12-25 05:30:00+07')`,
		}
		rowCount, err := transaction.Update(ctx, stmt)
		t.Logf("Rows %v", rowCount)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Errorf("%v", err)
	}
}

func Test_すでに存在するアカウントを追加する(t *testing.T) {
	ctx := context.Background()
	_, err := con.ReadWriteTransaction(ctx, func(ctx context.Context, transaction *spanner.ReadWriteTransaction) error {
		err := transaction.BufferWrite([]*spanner.Mutation{
			spanner.Insert("User", []string{"ID", "UserID", "Email", "Password", "CreatedAt", "ModifiedAt"}, []interface{}{
				"74514912-1b78-4afb-8119-034f8e241e2b", "U1", "U1@example.com", "password", time.Now(), spanner.CommitTimestamp,
			}),
		})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		// すでに存在するレコードを投入するテストなのでerrに入ってきて良い
		if strings.Index(err.Error(), "AlreadyExists") > 0 {
			t.Logf("OK %v", err)
		} else {
			t.Errorf("%v", err)
		}
	}
}

// Result : 単純にstructを渡すってことはできないようだ
func Test_Structを利用してinsertする(t *testing.T) {
	ctx := context.Background()
	_, err := con.ReadWriteTransaction(ctx, func(ctx context.Context, transaction *spanner.ReadWriteTransaction) error {
		i := Tweet{
			ID:         "a",
			Text:       "hello",
			CreatedAt:  time.Now(),
			ModifiedAt: spanner.CommitTimestamp,
			UserID:     "74514912-1b78-4afb-8119-034f8e241e2b",
		}
		m, err := spanner.InsertStruct("Tweet", i)
		if err != nil {
			panic(err)
		}
		err = transaction.BufferWrite([]*spanner.Mutation{m})
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		t.Errorf("%v", err) // spanner: code = "InvalidArgument", desc = "client doesn't support type spanner.Tweet"
	}
}

type Tweet struct {
	ID         string    `spanner:"ID"`
	Text       string    `spanner:"Text"`
	UserID     string    `spanner:"UserID"`
	CreatedAt  time.Time `spanner:"CreatedAt"`
	ModifiedAt time.Time `spanner:"ModifiedAt"`
}

type User struct {
	ID         string    `spanner:"ID"`
	Email      string    `spanner:"Email"`
	UserID     string    `spanner:"UserID"`
	Password   string    `spanner:"Password"`
	CreatedAt  time.Time `spanner:"CreatedAt"`
	ModifiedAt time.Time `spanner:"ModifiedAt"`
}

func Test_登録したユーザを検索する_Range(t *testing.T) {
	ctx := context.Background()
	rtx := con.ReadOnlyTransaction()
	defer rtx.Close()

	iter := rtx.Read(ctx, "User", spanner.AllKeys(), []string{"ID", "Email", "UserID", "Password", "CreatedAt", "ModifiedAt"})
	defer iter.Stop()

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		// Read()で指定したカラム順で取得できる
		var user User
		if err := row.ToStruct(&user); err != nil {
			panic(err)
		}
		t.Logf("%v\n", user)
	}
}

func Test_登録したユーザを検索する_key(t *testing.T) {
	ctx := context.Background()
	rtx := con.ReadOnlyTransaction()
	defer rtx.Close()

	iter := rtx.Read(ctx, "User", spanner.Key{"74514912-1b78-4afb-8119-034f8e241e2b"}, []string{"ID", "Email", "UserID", "Password", "CreatedAt", "ModifiedAt"})
	defer iter.Stop()

	for {
		row, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			panic(err)
		}
		// Read()で指定したカラム順で取得できる
		var user User
		if err := row.ToStruct(&user); err != nil {
			panic(err)
		}
		t.Logf("%v\n", user)
	}
}

func Test_トランザクション閉じたあとに再度実行(t *testing.T) {
	ctx := context.Background()
	rtx := con.ReadOnlyTransaction()

	{
		iter := rtx.Read(ctx, "User", spanner.Key{"74514912-1b78-4afb-8119-034f8e241e2b"}, []string{"ID", "Email", "UserID", "Password", "CreatedAt", "ModifiedAt"})
		for {
			row, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				panic(err)
			}
			// Read()で指定したカラム順で取得できる
			var user User
			if err := row.ToStruct(&user); err != nil {
				panic(err)
			}
			t.Logf("%v\n", user)
		}
		iter.Stop()
		rtx.Close()
	}

	{
		iter := rtx.Read(ctx, "User", spanner.Key{"74514912-1b78-4afb-8119-034f8e241e2b"}, []string{"ID", "Email", "UserID", "Password", "CreatedAt", "ModifiedAt"})
		for {
			_, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				t.Logf("トランザクションが閉じているのでここに入るのは期待値")
				break
			}
			t.Logf("ここには来ないはず")
		}
		iter.Stop()
		rtx.Close()
	}

	{
		rtx = con.ReadOnlyTransaction() // トランザクションを再取得する

		iter := rtx.Read(ctx, "User", spanner.Key{"74514912-1b78-4afb-8119-034f8e241e2b"}, []string{"ID", "Email", "UserID", "Password", "CreatedAt", "ModifiedAt"})
		for {
			row, err := iter.Next()
			if err == iterator.Done {
				break
			}
			if err != nil {
				panic(err)
			}
			// Read()で指定したカラム順で取得できる
			var user User
			if err := row.ToStruct(&user); err != nil {
				panic(err)
			}
			t.Logf("%v\n", user)
		}
		iter.Stop()
		rtx.Close()
	}
}
