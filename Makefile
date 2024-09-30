add:
	@go run main.go add --name "TestServer" --ip "192.168.103.128" --username "hant" --password "1" --ssh-key "~/.ssh/id_rsa"
connect:
	@go run main.go connect --server "TestServer" --username "hant"
list:
	@go run main.go list
delete:
	@go run main.go delete --id 1
execute:
	@go run main.go execute --server "TestServer" --command "ip a"
build:
	@go build -o overfly main.go

