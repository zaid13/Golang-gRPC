**Run client**

`go run client.go`

**Run server**

`go run server.go  `

**Generate Proto file**

`protoc --proto_path=src --go_out=out --go_opt=paths=source_relative foo.proto bar/baz.proto`