.PHONY: format vet tidy test clean
# -----------------------------------------------------------------------------
#  CONSTANTS
# -----------------------------------------------------------------------------

version = `cat VERSION`

src_dir = redash

build_dir = build

coverage_dir  = $(build_dir)/coverage
coverage_out  = $(coverage_dir)/coverage.out
coverage_html = $(coverage_dir)/coverage.html

# -----------------------------------------------------------------------------
#  FORMATTING
# -----------------------------------------------------------------------------

format:
	go fmt ./$(src_dir)
	gofmt -s -w ./$(src_dir)

vet:
	go vet ./$(src_dir)

tidy:
	go mod tidy 

# -----------------------------------------------------------------------------
#  TESTING
# -----------------------------------------------------------------------------

test:
	mkdir -p $(coverage_dir)
	go test ./$(src_dir) -tags test -v -covermode=count -coverprofile=$(coverage_out)
	go tool cover -html=$(coverage_out) -o $(coverage_html)

# -----------------------------------------------------------------------------
#  CLEANUP
# -----------------------------------------------------------------------------

clean:
	rm -rf $(build_dir)