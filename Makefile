build:
	@ printf "Building the blog...\n"
	@ go build -o bin/richardktranBlog ./

run: build
	@ ./bin/richardktranBlog