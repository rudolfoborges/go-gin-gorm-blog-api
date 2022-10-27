package utils

type Serializer interface {
	Serialize() map[string]interface{}
}
