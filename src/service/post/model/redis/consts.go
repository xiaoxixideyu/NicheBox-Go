package redis

const (
	KeyPrefixPost        = "post:"
	KeyPrefixBloomFilter = "bloomfilter:"

	KeyUserView = "userview:"
)

const (
	UpdateUserViewBloomFilterExpiration = 3 * 24 * 60 * 60
)
