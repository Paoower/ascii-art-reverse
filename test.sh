#!/bin/bash

FILES=("reversetest/example00.txt" "reversetest/example01.txt" "reversetest/example02.txt" "reversetest/example03.txt" "reversetest/example04.txt" "reversetest/example00.txt" "reversetest/example00.txt" "reversetest/example00.txt" "reversetest/example05.txt" "reversetest/example06.txt" "reversetest/example07.txt")

for file in "${FILES[@]}"
do
    echo "Testing file: $file"
    go run . --reverse="$file"
    echo "-------------------"
done
