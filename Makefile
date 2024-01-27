# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Build target
build:
	$(GOBUILD) -o your_binary_name

# Clean target
clean:
	$(GOCLEAN)
	rm -f your_binary_name

# Test target
test:
	$(GOTEST) -v ./...

# Get dependencies
deps:
	$(GOGET) -v ./...

# Default target
default: build
