run:
	go build -o gyt main.go
	./gyt $(in)

cat-file:
	go build -o gyt main.go
	./gyt cat-file $(in)

hash-object:
	go build -o gyt main.go
	./gyt hash-object $(in)
