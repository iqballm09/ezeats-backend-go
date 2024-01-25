build:
	go build -o server main.go

run: build
	./server

watch:
	reflex -s -r './usr/local/go/bin/go$$' make run