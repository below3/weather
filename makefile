run:
	go run cmd/weatherapp/main.go $(prod)

benchmark:
	go test cmd/weatherapp/main_test.go  -bench=.

test:
	go test ./...

download:
	go run cmd/pulldata/pulldata.go

clean:
	rm result.json