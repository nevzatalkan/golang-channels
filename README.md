# golang-channels
golang multiple channels example usage compared with for loop

# usage:
go test -bench=.

# output:
goos: windows
goarch: amd64
BenchmarkGoParalel-4       20000            267850 ns/op
BenchmarkForLoop-4           300           4526676 ns/op
PASS
ok      
