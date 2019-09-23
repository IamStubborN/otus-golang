//go:generate protoc event.proto --go_out=plugins=grpc:../delivery/gc/ --proto_path=../delivery/gc/
package domain

type Event struct {
	ID          uint64
	Name        string
	Description string
	Date        string
}
