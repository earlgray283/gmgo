#!/bin/bash

for f in h7/psw/*; do 
    for path in $f; do
        go run tool/tsv2csv.go "${path}/in.csv"
        go run tool/tsv2csv.go "${path}/out.csv"
    done
done
