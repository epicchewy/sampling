VERSION := v0.1.0
LDFLAGS := -X main.version=$(VERSION) -s -w

build: fmt vet rsampling ssampling wrsampling

rsampling: cmd/rsampling/main.go
	go build -ldflags "$(LDFLAGS)" -o bin/$@ $<

ssampling: cmd/ssampling/main.go
	go build -ldflags "$(LDFLAGS)" -o bin/$@ $<

wrsampling: cmd/wrsampling/main.go
	go build -ldflags "$(LDFLAGS)" -o bin/$@ $<

clean:
	rm -rf bin/

test:
	go test ./...

fmt:
	go fmt ./...

vet:
	go vet ./...
