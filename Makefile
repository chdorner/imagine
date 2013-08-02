.DEFAULT: build

VERSION := 0.1.0
TARGET  := imagine
LDFLAGS := -ldflags "-X main.Version $(VERSION)"

build: $(TARGET)

test:
	go test ./...

$(TARGET):
	go build -o $(TARGET) $(LDFLAGS)

.PHONY: $(TARGET)
.PHONY: test
