test:
	go test -v ./...

commit:
	git commit -am "$m"
	git push