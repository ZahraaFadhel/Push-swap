# go build -o push_swap main.go

echo 1- Try to run ./push-swap. Does it display nothing?
go run ./push-swap

echo -e

echo '2- Try to run ./push-swap "2 1 3 6 5 8". does it display a valid solution and less than 9 instructions?'
go run ./push-swap "2 1 3 6 5 8"
echo -e "Total number of instructions: $(go run ./push-swap "2 1 3 6 5 8" | wc -l)"


echo -e

echo 3- Try to run ./push-swap "0 1 2 3 4 5". Does it display nothing?
go run ./push-swap "0 1 2 3 4 5"

echo -e


echo 4- Try to run ./push-swap "0 one 2 3". Does it display an error message?
go run ./push-swap "0 one 2 3"

echo -e

echo 5- Try to run ./push-swap "1 2 2 3". Does it display an error message?
go run ./push-swap "1 2 2 3"

echo -e

echo '6- Try to run ./push-swap "<5 random numbers>" with 5 random numbers instead of the tag. Does it display a valid solution and less than 12 instructions?'
echo 'Enter 5 numbers seperated by spaces:' 
read -p "> " Input
go run ./push-swap "$Input"
echo -e "Total number of instructions: $(go run ./push-swap "$Input" | wc -l)"


echo -e 

echo '7- Again, Try to run ./push-swap "<5 random numbers>" with another 5 random numbers instead of the tag. Does it display a valid solution and less than 12 instructions?'
echo 'Enter 5 numbers seperated by spaces:' 
read -p "> " Input
go run ./push-swap "$Input"
echo -e "Total number of instructions: $(go run ./push-swap "$Input" | wc -l)"

echo -e 

# go build -o checker main.go

echo 8- Try to run ./checker and input nothing. Does it display nothing?
go run ./checker

echo -e 

echo 9- Try to run ./checker "0 one 2 3" . Does it display an error message? 
go run ./checker "0 one 2 3"

echo -e 

echo "10- Try to run echo -e \"sa\\npb\\nrrr\\n\" | ./checker \"0 9 1 8 2 7 3 6 4 5\". Does it display \"KO\"?"
echo -e "sa\npb\nrrr\n" | go run ./checker "0 9 1 8 2 7 3 6 4 5"


echo -e 

echo "11- Try to run echo -e "pb\\nra\\npb\\nra\\nsa\\nra\\npa\\npa\\n" | ./checker "0 9 1 8 2". Does it display "OK"?"
echo -e "pb\nra\npb\nra\nsa\nra\npa\npa\n" | go run ./checker "0 9 1 8 2"

echo -e 

echo "12- Try to run ARG="4 67 3 87 23"; ./push-swap "'$ARG'" | ./checker "'$ARG'". Does it display "OK"?"
ARG="4 67 3 87 23"; go run ./push-swap "$ARG" | go run ./checker "$ARG"

echo -e

echo "13- Try to run ARG="100 random numbers"; ./push-swap "'$ARG'" with 100 random different numbers instead of the tag. Does it display less than 700 commands? \nWe will test 100 random numbers for you"
ARG=$(seq -100 100 | shuf -n 100 | tr '\n' ' ' | sed 's/[[:space:]]*$//')
echo Numbers are: [$ARG]
go run ./push-swap "$ARG"
echo -e "Total number of instructions: $(go run ./push-swap "$ARG" | wc -l)"


echo -e 

echo "Try to run ARG="100 random numbers"; ./push-swap "'$ARG'" | ./checker "'$ARG'" with the same 100 random different numbers as before instead of the tag." Does it display OK?
ARG=$(seq -100 100 | shuf -n 100 | tr '\n' ' ' | sed 's/[[:space:]]*$//')
go run ./push-swap "$ARG" | go run ./checker "$ARG"