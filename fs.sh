#!/bin/bash

# Function to print styled separator
print_separator() {
  printf "\n \033[36m%s\033[0m" "//////////////////////////////////////////////////////////"
}
# TEST FS
GO_PROGRAM="go run ."
printf "\n\n \033[31m%s\033[0m" "////////////////////////TESTS FS///////////////////////////"
printf "\n\n \033[35m◼\033[0m : Input \033[32m◼\033[0m : Output \n"

declare -a inputs=("banana" "hello" "hello world" "nice 2 meet you" "you & me" "123" "/(\")" "ABCDEFGHIJKLMNOPQRSTUVWXYZ" "\"#$%&/()*+,-./" "It's Working" "Train")
declare -a banners=("standard abc" "standard" "shadow" "thinkertoy" "standard" "shadow" "thinkertoy" "shadow" "thinkertoy" "thinkertoy" "train")

for i in "${!inputs[@]}"; do
  input="${inputs[$i]}"
  banner="${banners[$i]}"
  
  printf "\033[35m \n Test for go run . %s %s | cat -e : \n" "\"$input\"" "$banner"
  output="$($GO_PROGRAM "$input" $banner | cat -e)"
  printf "\033[32m%s\n" "$output"
  print_separator
done