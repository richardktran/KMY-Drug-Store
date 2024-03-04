build:
	@ printf "Building the blog...\n"
	@ go build -o bin/kmy ./

run: build
	@ ./bin/kmy