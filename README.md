gogen
=====

gogen is a struct base [go generate](https://golang.org/doc/go1.4#gogenerate) library.

# Usage

```
$ go get github.com/kyokomi/gogen
```

# Example

## generator install

```
$ go install github.com/kyokomi/gogen/example/genexample
```

## generate

```
$ cd ./example
$ go generate
```

## output

[main_gen.go](https://github.com/kyokomi/gogen/blob/master/example/main_gen.go)

# TODO

- [ ] paser部分で[msgp/paser](https://github.com/philhofer/msgp)を使うのをやめる
- [ ] gen構造体で[msgp/gen](https://github.com/philhofer/msgp)を使うのをやめる
- [ ] `go generate --debug`でログがでるようにする（今は全部でてる）
- [ ] testコードを書く
- [ ] documentを書く

# LICENSE

[MIT](https://github.com/kyokomi/gogen/blob/master/LICENSE)
