gogen
=====

gogen is a struct base [go generate](https://golang.org/doc/go1.4#gogenerate) library.

## Usage

```
$ go get github.com/kyokomi/gogen
```

## Example

### main.go

[main.go](https://github.com/kyokomi/gogen/blob/master/example/main.go)

```go
package main

//go:generate genexample
type Hoge struct {
	Name    string
	Num     int
	Message string
}

func main() {

	h := Hoge{
		Name: "hoge",
		Num: 1,
		Message: "test",
	}
	h.Sample()
}
```


### generator install

```
$ go install github.com/kyokomi/gogen/example/genexample
```

[genexample source](https://github.com/kyokomi/gogen/blob/master/example/genexample/main.go)

### generate

```
$ cd ./example
$ go generate
```

### output

[main_gen.go](https://github.com/kyokomi/gogen/blob/master/example/main_gen.go)

```
package main

import (
	"fmt"
)


// Sample sample code
func (z *Hoge) Sample() {

	fmt.Println(z.Name)

	fmt.Println(z.Num)

	fmt.Println(z.Message)

}
```

### Run Sample

```
$ go run main.go main_gen.go
hoge
1
test
```

## TODO

- [ ] paser部分で[msgp/parse](https://github.com/philhofer/msgp/tree/master/parse)を使うのをやめる
- [ ] gen構造体で[msgp/gen](https://github.com/philhofer/msgp/tree/master/gen)を使うのをやめる
- [ ] `go generate --debug`でログがでるようにする（今は全部でてる）
- [ ] testコードを書く
- [ ] documentを書く

## LICENSE

[MIT](https://github.com/kyokomi/gogen/blob/master/LICENSE)
