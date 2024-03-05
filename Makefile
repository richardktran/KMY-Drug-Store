build:
	@ printf "Building the blog...\n"
	@ go build -o bin/kmy ./main.go

run: build
	@ ./bin/kmy