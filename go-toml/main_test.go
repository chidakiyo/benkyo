package main

import (
	"github.com/BurntSushi/toml"
	"testing"
)

type Config struct {
	Hoge Hoge `toml:"hoge"`
}

type Hoge struct {
	A string `toml:"a"`
	B int64  `toml:"b"`
}

func Test_Tomlのデコード(t *testing.T) {

	target := Config{}
	meta, err := toml.DecodeFile("test.toml", &target)
	if err != nil {
		t.Errorf("%v", err)
	}

	t.Logf("%+v", target)
	t.Logf("%+v", meta)
	// {mapping:map[hoge:map[A:a B:1] hogehoge:map[C:3 D:d]] types:map[hoge:Hash hoge.A:String hoge.B:Integer hogehoge:Hash hogehoge.C:Integer hogehoge.D:String] keys:[[hoge] [hoge A] [hoge B] [hogehoge] [hogehoge C] [hogehoge D]] decoded:map[hoge:true hoge.A:true hoge.B:true] context:[]}
}
