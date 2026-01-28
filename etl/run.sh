# sud0 apt install notify-tools
# inotifywait --monitor --event modify **/*.go |
inotifywait --monitor --event modify main.go |
while read file event; do
    clear # clear the console
    rm -rf output
    mkdir output
    go run . # run the code
done

