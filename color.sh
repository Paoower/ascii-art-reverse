#!/bin/bash

GO_PROGRAM="go run ."

printf "\n\n \033[31m%s\033[0m" "////////////////////////TESTS COLOR///////////////////////////"
printf "\n\n \033[35m◼\033[0m : Input "

# Test case 0
color="red"
input="banana"
printf "\033[35m\n Test for %s --color %s %s %s : \033[0m\n" "$GO_PROGRAM" "$color"  "\"$input\""
$GO_PROGRAM --color "$color" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Test case 1
color="red"
input="hello world"
printf "\033[35m\n Test for %s --color=%s %s : \033[0m\n" "$GO_PROGRAM" "$color" "\"$input\""
$GO_PROGRAM --color="$color" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Test case 2
color="green"
input="1 + 1 = 2"
printf "\033[35m\n Test for %s --color=%s %s : \033[0m\n" "$GO_PROGRAM" "$color" "\"$input\""
$GO_PROGRAM --color="$color" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Test case 3
color="yellow"
input="(%&) ??"
printf "\033[35m\n Test for %s --color=%s %s : \033[0m\n" "$GO_PROGRAM" "$color" "\"$input\""
$GO_PROGRAM --color="$color" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Test case 4
color="blue"
substring="ello"
input="hello world"
printf "\033[35m\n Test for %s --color=%s %s %s : \033[0m\n" "$GO_PROGRAM" "$color" "$substring" "\"$input\""
$GO_PROGRAM --color="$color" "$substring" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Test case 5
color="purple"
substring="e"
input="hello"
printf "\033[35m\n Test for %s --color=%s %s %s : \033[0m\n" "$GO_PROGRAM" "$color" "$substring" "\"$input\""
$GO_PROGRAM --color="$color" "$substring" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Test case 6
color="cyan"
substring="ll"
input="hello"
printf "\033[35m\n Test for %s --color=%s %s %s : \033[0m\n" "$GO_PROGRAM" "$color" "$substring" "\"$input\""s
$GO_PROGRAM --color="$color" "$substring" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Test case 7
color="orange"
substring="GuYs"
input="HeY GuYs"
printf "\033[35m\n Test for %s --color=%s %s %s : \033[0m\n" "$GO_PROGRAM" "$color" "$substring" "\"$input\""
$GO_PROGRAM --color="$color" "$substring" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Test case 8
color="blue"
substring="B"
input='RGB()'
printf "\033[35m\n Test for %s --color=%s %s %s : \033[0m\n" "$GO_PROGRAM" "$color" "$substring" "'$input'"
$GO_PROGRAM --color="$color" "$substring" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Random Test case 1
color="grey"
input="RandomStringWithMixedCase"
printf "\033[35m\n Test for %s --color=%s %s : \033[0m\n" "$GO_PROGRAM" "$color" "\"$input\""
$GO_PROGRAM --color="$color" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Random Test case 2
color="green"
input="lowercase numbers and spaces 12345"
printf "\033[35m\n Test for %s --color=%s %s : \033[0m\n" "$GO_PROGRAM" "$color" "\"$input\""
$GO_PROGRAM --color="$color" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Random Test case 3
color="#ff0000"
substring="c"
input="Special characters !@#$%^&*()"
printf "\033[35m\n Test for %s --color=%s %s %s : \033[0m\n" "$GO_PROGRAM" "$color" "$substring" "\"$input\""
$GO_PROGRAM --color="$color" "$substring" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Random Test case 4
color="rgb(255, 0, 0)"
substring="CASE 1234"
input="Mixed CASE 1234 and spaces"
printf "\033[35m\n Test for %s --color=%s %s %s : \033[0m\n" "$GO_PROGRAM" "$color" "$substring" "\"$input\""
$GO_PROGRAM --color="$color" "$substring" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"

# Random Test case 5
color="hsl(0, 100%, 50%)"
substring="CASE"
input="Testing different color notations CASE"
printf "\033[35m\n Test for %s --color=%s %s %s : \033[0m\n" "$GO_PROGRAM" "$color" "$substring" "\"$input\""
$GO_PROGRAM --color="$color" "$substring" "$input"
printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"
