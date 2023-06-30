# learn-go-with-tests

My repo for working through the book: [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)

## tips

- run tests inside VS-Code via the testing tab
- run go tool commands across all directories via `fd -td -x GO TOOL COMMAND`, such as: `fd -td -x go fmt`
- `go-staticcheck` = linting
- `go test -bench=.` = benchmark
- `errcheck` = error check
- `go test -race` = check for race conditions
- `go vet` = check for unsafe parallel routines (and maybe more?)
