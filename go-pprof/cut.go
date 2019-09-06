// Copyright 2019 Yoshi Yamaguchi
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// This application is Go port of `cut` command.
// This is just a sample application so it's not the perfect clone.
// This is used for the exercise part of https://connpass.com/event/144347/
//
// For testing, prepare large file where fields are splitted by same characters.
// * http://eforexcel.com/wp/downloads-18-sample-csv-files-data-sets-for-testing-sales/
// * https://www.transtats.bts.gov/DL_SelectFields.asp?Table_ID=236
package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"runtime/pprof"
)

var (
	delimiter byte
	field     int
)

func init() {
	var delimiterStr string
	flag.StringVar(&delimiterStr, "d", ",", "delimiter")
	delimiter = delimiterStr[0]
	flag.IntVar(&field, "f", 1, "field position")
}

// 使用方法: ./cut1 -f 3 -d ',' foo.csv
// フィールドに関しては範囲を表すハイフンは利用しない。
func main() {
	flag.Parse()
	f, err := os.Open(flag.Arg(0))
	if err != nil {
		log.Fatalf("Could not open file %q: %v", flag.Arg(0), err)
	}
	defer f.Close()

	pro,_ := os.Create("profile")
	defer pro.Close()
	pprof.StartCPUProfile(pro)
	defer pprof.StopCPUProfile()

	infield := false
	pos := field - 1
	s := []byte{}
	var buf [1]byte
	r := bufio.NewReader(f)
	w := bufio.NewWriter(os.Stdout)
	for {
		// ファイルからの読み込みが出来ない場合やファイル末尾の場合は終了する
		_, err := r.Read(buf[:])
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not read file %q properly: %v", flag.Arg(0), err)
		}

		// 1文字ずつ走査していく前提で状態として
		// * 行内/行末
		// * ターゲットのフィールド内外
		// * デリミタ
		// が考えられるので、その有限状態を管理する
		c := buf[0]
		if pos == 0 {
			infield = true
		}
		if c == delimiter {
			if pos > 0 && !infield {
				pos--
			} else if pos == 0 && infield {
				pos--
				infield = false
			}
			continue
		}
		if c == '\n' {
			infield = false
			pos = field - 1
			//fmt.Println(string(s))
			w.Write(s)
			s = []byte{}
		}
		if infield {
			s = append(s, c)
		}
	}
}