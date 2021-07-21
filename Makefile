
.PHONY: gen
gen:
	go generate ./...


.PHONY: build-domain
build-domain:
	echo "build domain"
	go build -o snailgirl ./cmd/domain/main.go
