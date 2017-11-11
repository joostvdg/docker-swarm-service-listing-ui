package model

// Type interface for json marshalling
type JsonStruct interface {
	Unmarshal([]byte) (JsonStruct, error)
}
