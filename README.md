# learn-go-with-tests

My repo for working through the book: [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)

## todo

- learn the [dedent package](https://github.com/lithammer/dedent/blob/master/dedent.go)
## tips

- benchmark all packages across all sub-dirs via `go test -bench=.`, such as:
    `fd -td -x go test -bench=.`
- error check all packages across all sub-dirs via `errcheck`, such as:
    `fd -td -x errcheck .`
- check for race conditions via `go test -race`, such as:
    `fd -td -x go test -race`
- check for unsafe parallel routines (and maybe more?) via `go vet`, such as:
    `fd -td -x go vet`
