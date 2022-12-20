# qmgo

## Usage

```go
var zero1, one1 int
zero, one := &zero1, &one1
in := [][]int{
     {0, 0, 0},
     {0, 0, 1},
     {0, 1, 0},
     {0, 1, 1},
     {1, 0, 0},
     {1, 0, 1},
     {1, 1, 0},
     {1, 1, 1},
}
out := [][]*int{
     {one},
     {zero},
     {one},
     {zero},
     {one},
     {zero},
     {one},
     {zero},
}
significantGroupEachOutput, _ := QuineMccluskey(in, out)
for outputIndex, logicFunctionList := range significantGroupEachOutput {
     for logicFuntionIndex, logicFunction := range logicFunctionList {
          fmt.Println(logicFunction)
     }
}
```

## install cli application

```shell
$ go install github.com/earlgray283/quine-mccluskey/cmd/qmgo
```

## example

```shell
$ qmgo ./examples/input.csv ./examples/output.csv
     __         _
f = BCD + AC + AB
```