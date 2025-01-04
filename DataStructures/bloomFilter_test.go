package DataStructures

import (
	"bytes"
	"fmt"
	"testing"
)

func TestBloomFilter(t *testing.T) {
	t.Run("Test_NewBloomFilter_BasicParameters", func(t *testing.T) {
		expectedElements := uint(100)
		falsePositiveRate := 0.01
		bf := NewBloomFilter(expectedElements, falsePositiveRate)

		if bf.Size() == 0 {
			t.Errorf("Expected non-zero size, got %d", bf.Size())
		}
		if bf.HashCount() == 0 {
			t.Errorf("Expected non-zero hash count, got %d", bf.HashCount())
		}
		if bf.ElementCount() != 0 {
			t.Errorf("Element count should start at 0, got %d", bf.ElementCount())
		}
	})

	t.Run("Test_Add_And_Contains_SingleItem", func(t *testing.T) {
		bf := NewBloomFilter(10, 0.05)
		item := []byte("test-item")
		bf.Add(item)

		if !bf.Contains(item) {
			t.Errorf("Expected to contain 'test-item', but got false")
		}
		if bf.ElementCount() != 1 {
			t.Errorf("Element count expected to be 1, got %d", bf.ElementCount())
		}
	})

	t.Run("Test_Add_MultipleItems_ContainsCheck", func(t *testing.T) {
		bf := NewBloomFilter(50, 0.01)
		items := [][]byte{
			[]byte("apple"),
			[]byte("banana"),
			[]byte("cherry"),
			[]byte("date"),
		}

		for _, it := range items {
			bf.Add(it)
		}

		// All items should be "possibly present" => Contains == true
		for _, it := range items {
			if !bf.Contains(it) {
				t.Errorf("Expected item %s to be present, but got false", it)
			}
		}
		if bf.ElementCount() != uint(len(items)) {
			t.Errorf("Element count expected %d, got %d", len(items), bf.ElementCount())
		}
	})

	t.Run("Test_Contains_UnknownItem", func(t *testing.T) {
		bf := NewBloomFilter(20, 0.01)
		bf.Add([]byte("known-item"))

		if bf.Contains([]byte("unknown-item")) && bf.Contains([]byte("unknown-item2")) {
			t.Error("False positive is possible, but having multiple in a small filter is suspicious; check logic!")
		}
	})

	t.Run("Test_FalsePositive_Rate_Calculation", func(t *testing.T) {
		bf := NewBloomFilter(100, 0.01)

		// Add exactly 100 items
		for i := 0; i < 100; i++ {
			str := fmt.Sprintf("item-%d", i)
			bf.Add([]byte(str))
		}

		// The actual fill ratio is approximate; just ensure the function runs and returns a sensible value
		fpRate := bf.CurrentFalsePositiveRate()
		if fpRate < 0.0 || fpRate > 1.0 {
			t.Errorf("CurrentFalsePositiveRate out of bounds: %f", fpRate)
		}
	})

	t.Run("Test_Add_LargeNumberOfItems", func(t *testing.T) {
		bf := NewBloomFilter(1000, 0.05)
		for i := 0; i < 1000; i++ {
			data := []byte(fmt.Sprintf("data-%d", i))
			bf.Add(data)
		}

		if bf.ElementCount() != 1000 {
			t.Errorf("Expected element count 1000, got %d", bf.ElementCount())
		}
	})

	t.Run("Test_DuplicateAdds", func(t *testing.T) {
		bf := NewBloomFilter(10, 0.01)
		item := []byte("duplicate-item")
		bf.Add(item)
		bf.Add(item)
		bf.Add(item)

		// Should not double- or triple-count
		if bf.ElementCount() != 3 {
			t.Errorf("ElementCount should track actual additions, expected 3, got %d", bf.ElementCount())
		}
		// Contains must still be true
		if !bf.Contains(item) {
			t.Errorf("Expected item to be present after multiple adds")
		}
	})

	t.Run("Test_BoundaryCases_SizeCalculation", func(t *testing.T) {
		// Very small bloom filter
		bf := NewBloomFilter(1, 0.1)
		if bf.Size() == 0 {
			t.Errorf("Size must be at least 1, got %d", bf.Size())
		}
		if bf.HashCount() == 0 {
			t.Errorf("HashCount must not be 0, got %d", bf.HashCount())
		}
	})

	t.Run("Test_BitArray_Internals", func(t *testing.T) {
		bf := NewBloomFilter(5, 0.01)
		item := []byte("bitarray-test")
		hashPositions := bf.getHashValues(item)

		bf.Add(item)

		for _, pos := range hashPositions {
			if !bf.bitArray[pos] {
				t.Errorf("bitArray at position %d should be true after adding the item", pos)
			}
		}
	})

	t.Run("Test_EmptyFilter_ContainsAlwaysFalse", func(t *testing.T) {
		bf := NewBloomFilter(10, 0.01)
		item := []byte("not-added")
		if bf.Contains(item) {
			t.Errorf("Empty filter should never contain an item, got Contains=true")
		}
	})

	t.Run("Test_BigData_Sanity", func(t *testing.T) {
		// Checking performance / not crashing for large expectedElements
		bf := NewBloomFilter(1000000, 0.01)
		if bf.Size() < 1000000 {
			t.Errorf("Size should be significantly larger than expectedElements, got %d", bf.Size())
		}
	})

	t.Run("Test_CollisionDemo", func(t *testing.T) {
		// We artificially check for collisions by using repeated item
		bf := NewBloomFilter(100, 0.05)
		itemA := []byte("abc")
		itemB := []byte("def")

		// Force same data, for demonstration
		if bytes.Equal(itemA, itemB) {
			t.Fatalf("Items are unexpectedly equal!") // just a safety net
		}

		bf.Add(itemA)
		if bf.Contains(itemB) {
			t.Logf("False positive possible (size=100, fpr=0.05) - might happen rarely")
		}
	})
}
