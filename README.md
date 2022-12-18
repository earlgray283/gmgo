# qmgo

## install

```shell
$ go install .
```

## execute

```shell
$ qmgo ./examples/input.csv ./examples/output.csv
0
0
0
0
1
0
0
0
1
-
1
1
1
0
-
1
===must-significant-list(1)===
 - | 1 | 0 | 0 (4, 12)

===new-significant-list(1)===
 1 | 0 | - | 0 (8, 10)
 1 | 0 | - | 1 (9, 11)
 1 | 0 | 1 | - (10, 11)
 1 | - | 1 | 1 (11, 15)
 1 | 0 | 0 | - (8, 9)
 1 | - | 0 | 0 (8, 12)
 1 | - | 1 | 0 (10, 14)
 1 | 1 | - | 0 (12, 14)
 1 | 1 | 1 | - (14, 15)
 - | 1 | 0 | 0 (4, 12)

    --     -
f: BCD+AC+AB
```