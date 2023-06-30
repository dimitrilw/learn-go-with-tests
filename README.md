# learn-go-with-tests

My repo for working through the book: [Learn Go with Tests](https://quii.gitbook.io/learn-go-with-tests)

## todo

- learn the [dedent package](https://github.com/lithammer/dedent/blob/master/dedent.go)
- experiment with [fyne](https://fyne.io/) for UI
- get better with logging
    - learn [Uber's zap logging](https://github.com/uber-go/zap)
    - log contexts, like in [this blog post](https://blog.gopheracademy.com/advent-2016/context-logging/)
- research code profiling, like with `pprof` and other tools

## tips

- run go tool commands across all directories via `fd -td -x GO TOOL COMMAND`, such as: `fd -td -x go fmt`
- `go-staticcheck` = linting
- `go test -bench=.` = benchmark
- `errcheck` = error check
- `go test -race` = check for race conditions
- `go vet` = check for unsafe parallel routines (and maybe more?)
