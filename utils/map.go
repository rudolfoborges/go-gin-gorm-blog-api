package utils

func ToSerializedArray[T Serializer](data []T) []map[string]interface{} {
	maps := make([]map[string]interface{}, len(data))

	for i, serializer := range data {
		maps[i] = serializer.Serialize()
	}

	return maps
}
