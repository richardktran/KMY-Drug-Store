build:
	@ printf "Building the blog...\n"
	@ go build -o bin/richardktranBlog ./bootstrap/

run: build
	@ ./bin/richardktranBlog