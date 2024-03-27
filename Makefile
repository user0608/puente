build:
	@go build -ldflags \
		"-X 'puente/cmd.version=$(shell go run scripts/version/version.go)' \
		-X 'puente/cmd.descripcion=$(shell cat description.txt)' \
		-X 'puente/cmd.branch=$(shell git rev-parse --abbrev-ref HEAD)' \
		-X 'puente/cmd.commit=$(shell git log -1 --format=%H)'" \
		-X 'puente/cmd.configExample=$(shell cat config.yml.example)'" \
		-o build/puente main.go
run:
	@go run main.go -c config.yml