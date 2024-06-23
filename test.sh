./push-swap go run .

./push-swap/main.go "2 1 3 6 5 8".

./push-swap/main.go "0 1 2 3 4 5"

./push-swap/main.go "0 one 2 3"

./push-swap/main.go "1 2 2 3"

./checker

./checker "0 one 2 3"

echo -e "sa\npb\nrrr\n" | ./checker "0 9 1 8 2 7 3 6 4 5"

echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | ./checker "0 9 1 8 2"

ARG="4 67 3 87 23"; ./push-swap "$ARG" | ./checker "$ARG"

ARG="<100 random numbers>"; ./push-swap "$ARG"

ARG="<100 random numbers>"; ./push-swap "$ARG" | ./checker "$ARG"
