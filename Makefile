stable_marriage: rules.go main.go
	go install

rules.go: stable_marriage.bonsai .git/index
	git submodule update --init --recursive
	cd bonsai && ruby compile.rb ../stable_marriage.bonsai > ../rules.go

clean:
	rm -f stable_marriage rules.go
