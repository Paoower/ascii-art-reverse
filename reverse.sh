#!/bin/bash

# Function to print styled separator
print_separator() {
  printf "\n \033[36m%s\033[0m" "//////////////////////////////////////////////////////////"
}
# TESTS REVERSE
GO_PROGRAM="go run ."
printf "\n\n \033[31m%s\033[0m" "////////////////////////TESTS REVERSE///////////////////////////"
printf "\n\n \033[35m◼\033[0m : Input \033[32m◼\033[0m : Output \n"
FILES=("test/example00.txt" "test/example01.txt" "test/example02.txt" "test/example03.txt" "test/example04.txt" "test/example05.txt" "test/example06.txt" "test/example07.txt")

for file in "${FILES[@]}"; do
  printf "\033[35m\nTesting file: %s\n" "$file"
  echo "
--------Input-----------"
  cat "$file"
  printf "\033[32m"
  echo "
--------Output-----------"
  go run . --reverse="$file"
  printf "\033[0m"
  print_separator
done