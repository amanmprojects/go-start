echo "Benchmarking GET"
for i in $(seq 100)
do
    response=$(curl -s http://localhost:8080/todos)
    echo "$i: $response" >> response.txt
done
echo "Benchmarking GET completed"

echo "Benchmarking POST"
for i in $(seq 100)
do
    id=$i
    response=$(curl -s -X POST http://localhost:8080/todos \
        -H "Content-Type: application/json" \
        -d "{\"ID\": \"$id\", \"Task\": \"Benchmark Task: $id\", \"Done\": false}")
    echo "$i: $response" >> response.txt
done
echo "Benchmarking POST completed"
