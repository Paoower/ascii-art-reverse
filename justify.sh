#!/bin/bash

# Function to print styled separator
print_separator() {
  printf "\n \033[36m%s\033[0m" "//////////////////////////////////////////////////////////"
}
# TESTS Justify
GO_PROGRAM="go run ."
printf "\n\n \033[31m%s\033[0m" "////////////////////////TESTS JUSTIFY///////////////////////////"
printf "\n\n \033[35m◼\033[0m : Input \033[32m◼\033[0m : Output \n"

declare -a align_options=("left" "right" "center" "justify")
declare -a inputs=("left standard" "right standard" "center hello shadow" "justify \"1 Two 4\" shadow")
declare -a banners=("standard" "thinkertoy" "shadow" "thinkertoy")

for i in "${!inputs[@]}"; do
  align="${align_options[$i]}"
  input="${inputs[$i]}"
  banner="${banners[$i]}"
  
  printf "\033[35m \n Test for go run . --align=%s %s %s : \n" "$align" "\"$input\"" "$banner"
  output="$(go run . --align="$align" "$input" "$banner")"
  printf "\033[32m%s\n" "$output"
  print_separator
done
