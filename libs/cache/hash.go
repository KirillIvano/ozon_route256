package cache

func hashFAQ6(key string) int32 {
	var hash int32

	for i := 0; i < len(key); i++ {
		hash += int32(key[i])
		hash += (hash << 10)
		hash ^= (hash >> 6)
	}

	hash += (hash << 3)
	hash ^= (hash >> 11)
	hash += (hash << 15)

	return hash
}
