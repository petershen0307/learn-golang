package openaddressedhashtable

import (
	"encoding/binary"
	"hash"
	"hash/fnv"
)

type hashPair[K comparable, V any] struct {
	key   K
	value V
}

type HashTable[K comparable, V any] struct {
	data   []*hashPair[K, V]
	hashFn hash.Hash32
	count  int
}

func New[K comparable, V any]() *HashTable[K, V] {
	return &HashTable[K, V]{
		data:   make([]*hashPair[K, V], 16),
		hashFn: fnv.New32(),
		count:  0,
	}
}

func keyToBytes[K comparable](key K) []byte {
	var b []byte
	switch v := any(key).(type) {
	case string:
		b = make([]byte, len(v))
		copy(b, v)
	case bool:
		b = make([]byte, binary.Size(v))
		if v {
			b[0] = 1
		} else {
			b[0] = 0
		}
	case int8:
		b = make([]byte, binary.Size(v))
		b[0] = byte(v)
	case uint8:
		b = make([]byte, binary.Size(v))
		b[0] = byte(v)
	case int16:
		b = make([]byte, binary.Size(v))
		binary.BigEndian.PutUint16(b, uint16(v))
	case uint16:
		b = make([]byte, binary.Size(v))
		binary.BigEndian.PutUint16(b, v)
	case int32:
		b = make([]byte, binary.Size(v))
		binary.BigEndian.PutUint32(b, uint32(v))
	case uint32:
		b = make([]byte, binary.Size(v))
		binary.BigEndian.PutUint32(b, v)
	case int:
		b = make([]byte, binary.Size(uint64(0)))
		binary.BigEndian.PutUint64(b, uint64(v))
	case uint:
		b = make([]byte, binary.Size(uint64(0)))
		binary.BigEndian.PutUint64(b, uint64(v))
	case int64:
		b = make([]byte, binary.Size(v))
		binary.BigEndian.PutUint64(b, uint64(v))
	case uint64:
		b = make([]byte, binary.Size(v))
		binary.BigEndian.PutUint64(b, v)
	}
	return b
}

func insert[K comparable, V any](data []*hashPair[K, V], hashFn hash.Hash32, key K, value V) {
	hashFn.Reset()
	keyBytes := keyToBytes(key)
	hashFn.Write(keyBytes)
	index := int(hashFn.Sum32()) % (len(data))
	if data[index] == nil {
		data[index] = &hashPair[K, V]{key: key, value: value}
		return
	}
	for i := (index + 1) % len(data); i != index; i = (i + 1) % len(data) {
		if data[i] == nil {
			data[i] = &hashPair[K, V]{key: key, value: value}
			return
		}
	}
}

func (ht *HashTable[K, V]) Insert(key K, value V) {
	if ht.count >= len(ht.data)*7/10 {
		// expand the table
		newData := make([]*hashPair[K, V], len(ht.data)*2)
		for _, pair := range ht.data {
			if pair != nil {
				insert(newData, ht.hashFn, pair.key, pair.value)
			}
		}
	}
	insert(ht.data, ht.hashFn, key, value)
	ht.count++
}
