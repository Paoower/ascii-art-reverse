GO_PROGRAM="go run ."

printf "\n \033[36m%s\033[0m" "////////////////////////TESTS///////////////////////////"
printf "\n \033[35m◼\033[0m : Input \033[32m◼\033[0m : Output \n"

# Predefined test cases
declare -a inputs=("banana" "hello" "hello world" "nice 2 meet you" "you & me" "123" "/(\")" "ABCDEFGHIJKLMNOPQRSTUVWXYZ" "\"#$%&/()*+,-./" "It's Working" "Train")
declare -a banners=("standard abc" "standard" "shadow" "thinkertoy" "standard" "shadow" "thinkertoy" "shadow" "thinkertoy" "thinkertoy" "train")

# Loop through predefined test cases
for i in "${!inputs[@]}"; do
    input="${inputs[$i]}"
    banner="${banners[$i]}"
    
    printf "\033[35m \n Test for go run . %s %s | cat -e : \n" "\"$input\"" "$banner"
    output="$($GO_PROGRAM "$input" $banner | cat -e)"
    printf "\033[32m%s\n" "$output"
    printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"
done
printf "\n////////////////////////More Tests ?///////////////////////////\n"
while true; do
    read -p "Enter a test input (or press Enter to finish): " input
    if [ -z "$input" ]; then
        break
    fi
    
    read -p "Enter a banner type for the input: " banner
    
    printf "\033[35m \n Test for go run . %s %s | cat -e : \n" "\"$input\"" "$banner"
    output="$($GO_PROGRAM "$input" $banner | cat -e)"
    printf "\033[32m%s\n" "$output"
    printf "\033[36m%s\033[0m\n" "//////////////////////////////////////////////////////////"
done
