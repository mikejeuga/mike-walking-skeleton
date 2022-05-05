test:
	go test -v ./...

c: commit
commit:
	@git commit -am "$m"
	@git pull --rebase
	git push

lf: lintfix
lintfix:
	@golangci-lint run ./... --fix