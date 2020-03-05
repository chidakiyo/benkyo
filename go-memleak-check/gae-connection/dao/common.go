package dao

import (
	cdatastore "cloud.google.com/go/datastore"
	"context"
	"go.mercari.io/datastore"
	"go.mercari.io/datastore/clouddatastore"
)

var connection datastore.Client

// コネクションを初期化する
func Init(project_id string) error {
	ctx := context.Background()
	datastoreclient, _ := cdatastore.NewClient(ctx, project_id)
	client, err := clouddatastore.FromClient(ctx, datastoreclient)
	if err != nil {
		return err
	}
	connection = client
	return nil
}

type (
	SpecializeXXX struct {
		datastore.Client
	}
	SpecializeYYY struct {
		datastore.Client
	}
)

func NewXXX() SpecializeXXX {
	return SpecializeXXX{connection}
}

func NewYYY() SpecializeYYY {
	return SpecializeYYY{connection}
}

type (
	SpecializeZZZ datastore.Client
)

func NewZZZ() SpecializeZZZ {
	return SpecializeZZZ(connection)
}

func(c *SpecializeXXX) Do() {
	c.Context() // mock
}

func(c *SpecializeYYY) Do() {
	c.Context() // mock
}

// type aliasだとinterfaceに対しては適用ができない
//func(c *SpecializeZZZ) Do() {
//	(*c).(datastore.Client).Context()
//}

