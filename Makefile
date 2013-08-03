.DEFAULT: build

VERSION := 0.1.0
TARGET  := imagine
REPO    := github.com/chdorner/imagine
LDFLAGS := -ldflags "-X $(REPO)/server.Version $(VERSION)"

build: $(TARGET)

test:
	go test ./...

$(TARGET):
	go build -o $(TARGET) $(LDFLAGS)

.PHONY: $(TARGET)
.PHONY: test
