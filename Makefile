stable_marriage: stable_marriage.bonsai main.go .git/index
	git submodule update --init --recursive
	go install

clean:
	rm -f stable_marriage
