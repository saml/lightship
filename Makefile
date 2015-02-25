lightship: main.go app/*.go
	go build 

run: lightship
	./lightship

doc:
	godoc -http=:8080

clean:
	rm lightship
