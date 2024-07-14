package docsv1

import "github.com/golang/protobuf/proto"

func (d *TDocument) MarshalBinary() ([]byte, error) {
	return proto.Marshal(d)
}

func (d *UserTDocument) MarshalBinary() ([]byte, error) {
	return proto.Marshal(d)
}

func (d *TDocument) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, d)
}

func (d *UserTDocument) UnmarshalBinary(data []byte) error {
	return proto.Unmarshal(data, d)
}
