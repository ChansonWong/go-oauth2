package utils

func ByteIsEmpty(value []byte) bool {
	if value != nil && len(value) > 0 {
		return false
	}

	return true
}

func ByteIsNotEmpty(value []byte) bool {
	return !ByteIsEmpty(value)
}
