# builds service executable
.PHONY: build
build:
	@go build -o ./bin/netr-service pkg/main.go

clean:
	@rm -rf bin build

debian: build
	@mkdir -p build/debian/netr/DEBIAN \
		build/debian/netr/usr/local/sbin \
		build/debian/netr/etc/netr/ \
		build/debian/netr/lib/systemd/system
	@cp config/netr.yaml.sample build/debian/netr/etc/netr
	@cp bin/netr-service build/debian/netr/usr/local/bin
	@cp netr.service build/debian/netr/lib/systemd/system
	@cp debian.control build/debian/netr/DEBIAN/control
	@dpkg-deb --build build/debian/netr

debian-install: debian
	sudo dpkg -i build/debian/netr-service.deb

# runs the service locally using the credentials provided by aws-vault for dev-00
.PHONY: run
run: build
	@./bin/netr-service -c config/netr.yaml
