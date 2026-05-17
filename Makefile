run:
	@DEBUG=1 go run cmd/errgen/main.go errgen \
		-i ./data/errors.json \
		-c ./data/econst.go \
		-r ./data/error.response.json \
		-e "internal server error"

help:
	@go run main.go --help
