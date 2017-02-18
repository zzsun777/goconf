## zzsun777/goconf

fork for `Terry-Mao/goconf`

## Requeriments
* Go 1.2+

## Installation

Just pull `zzsun777/goconf` from github using `go get`:

```sh
# download the code
$ go get -u github.com/zzsun777/goconf
```

## Usage

```go
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

```

## Examples

```sh
# configuration examples
# this is comment, goconf will ignore it.

id=1

col=777

# this is the section name
[core]

# a key-value config which key is test and value is 1
test 1

# one mb
test1 1mb

# one second
test2 1s

# boolean
test3 true

# arr
arr hello,the,world

# map
m 1=hello,2=the,3=world
```
