package containsDuplicate

import "testing"

//testing different aproach for the solutions, just for clarity
func TestContainsDuplicate(t *testing.T) {
	t.Run("Map[]_Aproach", func(t *testing.T) {
		if !ContainDuplicateWithMap([]int{1, 2, 3, 1}) {
			t.Error("Solution is incorrect, TRUE was expected, received: FALSE")

		}
		if ContainDuplicateWithMap([]int{3, 2, 4, 7, 8}) {
			t.Error("Solution is incorrect, FALSE was expected, received: TRUE")

		}
		if !ContainDuplicateWithMap([]int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}) {
			t.Error("Solution is incorrect TRUE was expected, received:FALSE")

		}

	})

}
