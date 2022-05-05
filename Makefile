test:
	go test -v ./...


git-commit:
	git commit -am "$m"
	git push