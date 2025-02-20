package uniqueNames

import (
	"testing"
)

// Convierte un slice en un map para facilitar la comparación
func sliceToMap(slice []string) map[string]bool {
	m := make(map[string]bool)
	for _, v := range slice {
		m[v] = true
	}
	return m
}

// Test para UniqueNames
func TestUniqueNames(t *testing.T) {
	trails := []struct {
		name1    []string
		name2    []string
		expected []string
	}{
		{
			name1:    []string{"Ava", "Emma", "Olivia"},
			name2:    []string{"Olivia", "Sophia", "Emma"},
			expected: []string{"Ava", "Emma", "Olivia", "Sophia"},
		},
	}

	for i, trail := range trails {
		t.Run(
			"TestCase#"+string(rune(i)),
			func(t *testing.T) {
				result := UniqueNames(trail.name1, trail.name2)

				// Convertimos ambos slices en maps
				resultMap := sliceToMap(result)
				expectedMap := sliceToMap(trail.expected)

				// Comparamos los maps
				if len(resultMap) != len(expectedMap) {
					t.Errorf("UniqueNames(%v, %v) = %v; expected %v",
						trail.name1, trail.name2, result, trail.expected)
					return
				}

				for key := range expectedMap {
					if !resultMap[key] {
						t.Errorf("UniqueNames(%v, %v) = %v; expected %v",
							trail.name1, trail.name2, result, trail.expected)
						return
					}
				}
			},
		)
	}
}
