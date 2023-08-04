package hashtable_test

import (
	"fmt"
	"github.com/petermcneil/learning/hashtable"
	"testing"
)

func TestHashTable_Insert(t *testing.T) {
	t.Parallel()
	table := hashtable.New[bool](1000)

	for i := 0; i < 10; i++ {
		for j := 10; j > 0; j-- {
			key := fmt.Sprintf("%d:%d", i, j)
			table.Put(key, true)
		}
	}

	for i := 0; i < 10; i++ {
		for j := 10; j > 0; j-- {
			key := fmt.Sprintf("%d:%d", i, j)
			if !table.HasKey(key) {
				t.Errorf("table doesn't have key: %s", key)
			}
		}
	}
}

func TestHashTable_Get(t *testing.T) {
	t.Parallel()
	table := hashtable.New[string](1000)

	for i := 0; i < 10; i++ {
		for j := 10; j > 0; j-- {
			key := fmt.Sprintf("%d:%d", i, j)
			value := fmt.Sprintf("%d:%d", j, i)
			table.Put(key, value)
		}
	}

	for i := 0; i < 10; i++ {
		for j := 10; j > 0; j-- {
			key := fmt.Sprintf("%d:%d", i, j)
			value := fmt.Sprintf("%d:%d", j, i)
			got, ok := table.Get(key)
			if !ok {
				t.Errorf("table doesn't have key: %s", key)
			}

			if value != got {
				t.Errorf("key: %s\nexpected: '%s'\ngot: '%s'", key, value, got)
			}
		}
	}
}

func TestHashTable_GetNotInMap(t *testing.T) {
	t.Parallel()
	table := hashtable.New[int](20)

	for i := 0; i < 10; i++ {
		table.Put(fmt.Sprintf("%d", i), i)
	}

	for i := 10; i < 20; i++ {
		if val, ok := table.Get(fmt.Sprintf("%d", i)); ok {
			t.Errorf("Table returned a value: %d for the key: %d", val, i)
		}
	}
}

func TestHashTable_Capacity(t *testing.T) {
	t.Parallel()
	table := hashtable.New[int](20)

	if table.Capacity() != 23 {
		t.Errorf("table.Capacity() != 23")
	}

	for i := 0; i < 20; i++ {
		table.Put(fmt.Sprintf("%d", i), i)
	}

	if table.Capacity() == 23 {
		t.Errorf("table.Capacity() == 23, should have resized")
	}

	if table.LoadFactor() > 0.6 {
		t.Errorf("table.LoadFactor() > 0.6, should have resized")
	}
}
