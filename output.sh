#!/bin/bash

# Function to print styled separator
print_separator() {
  printf "\n \033[36m%s\033[0m" "//////////////////////////////////////////////////////////"
}

# TEST OUTPUT

mkdir outputtest
printf "\n\n \033[31m%s\033[0m" "////////////////////////TESTS OUTPUT///////////////////////////"
printf "\n\n \033[35m◼\033[0m : Input \033[32m◼\033[0m : Output \n"

GO_PROGRAM="go run ."
output_file="outputtest/test00.txt"
input="banana"
banner="standard"

printf "\033[35m\n Test for %s --output %s %s %s : \n" "$GO_PROGRAM" "$output_file" "$input" "$banner"
$GO_PROGRAM --output "$output_file" "$input" "$banner"
print_separator

test_cases=(
  "outputtest/test00.txt|First\nTest|shadow"
  "outputtest/test01.txt|hello|standard"
  "outputtest/test02.txt|123 -> #$%|standard"
  "outputtest/test03.txt|432 -> #$%&@|shadow"
  "outputtest/test04.txt|There|shadow"
  "outputtest/test05.txt|123 -> \"#$%@|thinkertoy"
  "outputtest/test06.txt|2 you|thinkertoy"
  "outputtest/test07.txt|Testing long output!|standard"
  "outputtest/random01.txt|RandomStringWithMixedCase|standard"
  "outputtest/random02.txt|lowercase numbers and spaces 12345|standard"
  "outputtest/random03.txt|Special characters !@#$%^&*()|standard"
  "outputtest/random04.txt|Mixed CASE 1234 and spaces|standard"
)

for test_case in "${test_cases[@]}"; do
  IFS='|' read -r output_file input banner <<< "$test_case"

  printf "\033[35m\n Test for %s --output=%s %s %s : \n" "$GO_PROGRAM" "$output_file" "$input" "$banner"
  $GO_PROGRAM --output="$output_file" "$input" "$banner"
  
  if [ -f "$output_file" ]; then
    cat_output="$(cat -e "$output_file")"
    printf "\033[32m%s" "$cat_output"
  else
    printf "\033[31m%s" "Output file not found: $output_file"
  fi

  print_separator
done