test:
	go test -v ./...

commit:
	@git commit -am "$m"
	@git pull --rebase
	git push