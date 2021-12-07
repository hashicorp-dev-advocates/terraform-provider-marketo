name = marketo
organization = hashicorp
version = 0.1.0

build:
	go build -o bin/terraform-provider-$(name)_v$(version)

install: build 
	mkdir -p ~/.terraform.d/plugins/local/$(organization)/$(name)/$(version)/linux_amd64
	mv bin/terraform-provider-$(name)_v$(version) ~/.terraform.d/plugins/local/$(organization)/$(name)/$(version)/linux_amd64/