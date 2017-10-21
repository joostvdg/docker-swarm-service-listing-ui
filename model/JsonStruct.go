package model

type JsonStruct interface {
	Unmarshal([]byte) (JsonStruct, error)
}
