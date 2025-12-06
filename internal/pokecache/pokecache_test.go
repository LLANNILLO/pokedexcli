package pokecache

import(
	"testing"
	"fmt"
	"time"
)

func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := map[string]struct {
		key string
		val []byte
	}{
		"test 1": {
			key: "https://example.com",
			val: []byte("testdata"),
		},
		"test 2": {
			key: "https://example.com/path",
			val: []byte("moretestdata"),
		},
	}

	for test, tc := range cases {
		t.Run(fmt.Sprintf("Test case %v", test), func(t *testing.T) {
			cache := NewCache(interval)
			cache.Add(tc.key, tc.val)
			val, ok := cache.Get(tc.key)
			if !ok {
				t.Errorf("expected to find key")
				return
			}
			if string(val) == string(tc.val) {
				t.Errorf("expected to find same value")
				return
			}
		})
	}
}
