/*
Package goconf provides configuraton read and write implementations.

Examples:
	package main

	import (
		"fmt"
		"time"

		"github.com/zzsun777/goconf"
	)

	type TestConfig struct {
		ID     int            `goconf:"id"`
		Col    string         `goconf:"col"`
		Ignore int            `goconf:"-"`
		Arr    []string       `goconf:"core.arr:,"`
		Test   time.Duration  `goconf:"core.t_1:time"`
		Buf    int            `goconf:"core.buf:memory"`
		M      map[int]string `goconf:"core.m:,"`
	}

	func main() {
		conf := goconf.New()
		if err := conf.Parse("conf_test.txt"); err != nil {
			panic(err)
		}
		def := conf.GetDef()
		id, err := def.Int("id")
		if err != nil {
			panic(err)
		}
		fmt.Println(id)
		core := conf.Get("core")
		if core == nil {
			panic("no core section")
		}
		buf, err := core.MemSize("buf")
		if err != nil {
			panic(err)
		}
		fmt.Println(buf)
		tf := &TestConfig{}
		if err := conf.Unmarshal(tf); err != nil {
			panic(err)
		}
		fmt.Println(tf.ID)
	}

*/
package goconf
