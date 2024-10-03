build: cmd/o/main.go
	go build -o o cmd/o/main.go
	echo "Built Go binary successfully"

clean: o
	rm o
	echo "Cleaned successfully"
