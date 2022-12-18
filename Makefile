APP_BIN = build/web

lint:
	golangci-lint run

build: clean $(APP_BIN)

$(APP_BIN):
	go build -o $(APP_BIN) ./cmd/main.go

clean:
	rm -rf ./build || true

swagger:
	swag init -g ./app/cmd/main.go -o ./app/docs


git:
	git add .
	git commit -a -m "$m"
	git push -u origin main

update_contracts:
	go get -u github.com/i-b8o/read-only_contracts@$m

test:
	go test -p 1 ./...