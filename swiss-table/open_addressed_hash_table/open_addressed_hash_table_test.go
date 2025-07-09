package openaddressedhashtable

import (
	"hash/fnv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_KeyToBytes(t *testing.T) {
	assert.Equal(t, []byte("test"), keyToBytes("test"))
	assert.Equal(t, []byte{1}, keyToBytes(true))
	assert.Equal(t, []byte{0}, keyToBytes(false))
	assert.Equal(t, []byte{0x01}, keyToBytes(int8(1)))
	assert.Equal(t, []byte{0x01}, keyToBytes(uint8(1)))
	assert.Equal(t, []byte{0x00, 0x01}, keyToBytes(int16(1)))
	assert.Equal(t, []byte{0x00, 0x01}, keyToBytes(uint16(1)))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x01}, keyToBytes(int32(1)))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x01}, keyToBytes(uint32(1)))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, keyToBytes(int(1)))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, keyToBytes(uint(1)))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, keyToBytes(int64(1)))
	assert.Equal(t, []byte{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x01}, keyToBytes(uint64(1)))
}

func Test_insert_key_string_value_int(t *testing.T) {
	ht := make([]*hashPair[string, int], 3)
	var pair hashPair[string, int]
	pair = hashPair[string, int]{
		key:   "key1",
		value: 42,
	}
	insert(ht, fnv.New32(), pair.key, pair.value)
	testHashFn := fnv.New32()
	testHashFn.Write([]byte(pair.key))
	index := int(testHashFn.Sum32()) % len(ht)
	assert.Equal(t, pair, *ht[index])

	pair = hashPair[string, int]{
		key:   "key2",
		value: 100,
	}
	insert(ht, fnv.New32(), pair.key, pair.value)
	testHashFn.Reset()
	testHashFn.Write([]byte(pair.key))
	index = int(testHashFn.Sum32()) % len(ht)
	assert.Equal(t, pair, *ht[index+1]) // key2 and key1 map to the same index, so key2 is inserted at the next available index
}
