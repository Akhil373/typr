APP=typr
BINDIR=bin

.PHONY: build build-all clean

build:
	go build -o $(BINDIR)/$(APP) .

build-all:
	GOOS=linux   GOARCH=amd64 go build -o $(BINDIR)/$(APP)-linux   .
	GOOS=darwin  GOARCH=arm64 go build -o $(BINDIR)/$(APP)-macos   .
	GOOS=windows GOARCH=amd64 go build -o $(BINDIR)/$(APP).exe    .

clean:
	rm -f $(BUILD)/$(APP) $(BUILD)/$(APP)-linux $(BUILD)/$(APP)-macos $(BUILD)/$(APP).exe
