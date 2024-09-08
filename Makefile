build: main.go
	go build -o o main.go
	echo "Built Go binary successfully"

install: o
	install "o" "${HOME}/.local/bin"
	chmod +x  "${HOME}/.local/bin/o"
	echo "Installed 'o' successfully. Ensure it is in your PATH"

uninstall:
	rm ${HOME}/.local/bin/o
	echo "Uninstalled `o` successfully"

clean:
	rm ${HOME}/.local/bin/o o
	echo "Cleaned successfully"
