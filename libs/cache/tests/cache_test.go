package tests

import (
	"route256/libs/cache"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type C struct {
	key   string
	value int32
}

func (c C) GetKey() string {
	return c.key
}

func TestCache(t *testing.T) {
	t.Run("correctly working cache", func(t *testing.T) {
		t.Parallel()

		c := cache.NewCache[C](time.Second)

		values := []C{
			{key: "1", value: 1},
			{key: "2", value: 2},
			{key: "3", value: 3},
			{key: "4", value: 4},
			{key: "5", value: 5},
			{key: "6", value: 6},
			{key: "22", value: 22},
			{key: "21", value: 21},
			{key: "20", value: 20},
		}

		for _, v := range values {
			c.AddToCache(v.key, v)
		}

		res, _ := c.GetFromCache("22")
		require.Equal(t, res.value, int32(22))
	})

	t.Run("missing key", func(t *testing.T) {
		t.Parallel()

		c := cache.NewCache[C](time.Second)

		values := []C{
			{key: "1", value: 1},
			{key: "2", value: 2},
			{key: "3", value: 3},
			{key: "4", value: 4},
			{key: "5", value: 5},
			{key: "6", value: 6},
			{key: "22", value: 22},
			{key: "21", value: 21},
			{key: "20", value: 20},
		}

		for _, v := range values {
			c.AddToCache(v.key, v)
		}

		var expected *C
		res, _ := c.GetFromCache("30")
		require.Equal(t, expected, res)
	})

	t.Run("least used is removed", func(t *testing.T) {
		t.Parallel()

		c := cache.NewCache[C](time.Second)

		values := []C{
			{key: "21", value: 21},
			{key: "23", value: 23},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
			{key: "22", value: 22},
		}

		for _, v := range values {
			c.AddToCache(v.key, v)
		}

		var expected *C

		first, _ := c.GetFromCache("21")
		second, _ := c.GetFromCache("23")
		require.Equal(t, expected, first)
		require.Equal(t, int32(23), second.value)
	})

	t.Run("latest value is returned", func(t *testing.T) {
		t.Parallel()

		c := cache.NewCache[C](time.Second)

		values := []C{
			{key: "21", value: 21},
			{key: "21", value: 22},
			{key: "21", value: 23},
		}

		for _, v := range values {
			c.AddToCache(v.key, v)
		}

		second, _ := c.GetFromCache("21")
		require.Equal(t, int32(23), second.value)
	})

}

func createCache(cnt int) *cache.Cache[C] {
	c := cache.NewCache[C](time.Second)

	for i := 0; i < cnt; i++ {
		c.AddToCache(strconv.Itoa(i), C{value: int32(i), key: strconv.Itoa(i)})
	}

	return c
}

func readCache(cache *cache.Cache[C], cnt int) {
	for i := 0; i < cnt; i++ {
		cache.GetFromCache(strconv.Itoa(i))
	}
}

func BenchmarkCacheLookup1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache := createCache(1000)
		readCache(cache, 1000)
	}
}

func BenchmarkCacheLookup10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache := createCache(10000)
		readCache(cache, 10000)
	}
}

func BenchmarkCacheLookup20000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		cache := createCache(20000)
		readCache(cache, 20000)
	}
}

func BenchmarkCacheMemory1000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createCache(1000)
	}
}

func BenchmarkCacheMemory10000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createCache(10000)
	}
}

func BenchmarkCacheMemory20000(b *testing.B) {
	for i := 0; i < b.N; i++ {
		createCache(20000)
	}
}
