path = pkg/repository
version = $(shell cat CHANGES.rst | awk '/^[0-9]+\.[0-9]+(\.[0-9]+)?/' | head -n1)
tag_name = $(path)/v$(version)

tag:
	@git rev-parse --abbrev-ref HEAD | grep '^main$$'
	@git tag -a $(tag_name) -m "New version of $(path) library."
	@git push origin $(tag_name)
	@echo "-- $(tag_name) pushed to origin. --"