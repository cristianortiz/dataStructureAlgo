package uniqueNames

// Implement the uniqueNames function. when passed two slices ofnames it will return a slice containing
// the names that appear in either or both slices.  the returned slice should have no duplicates
// Example:

// calling uniqueNames([]string{"Ava","Emma","Olivia"}, []string{"Olivia","Sophia","Emma"})
//  should return a slice containing Ava, Emma, Olivia, and Sophia in any order

// using a map[string]bool allows to store unique names without the need for comparisions, because
// in maps key/values must be unique
func UniqueNames(names1, names2 []string) []string {
	//note how we use the names as key for the map, this way the keys must be unique
	uniqueMap := make(map[string]bool)

	//add the names from slice names1 to the map
	for _, name := range names1 {
		uniqueMap[name] = true

	}
	//add the names from slice names2 to the map, this is the problem key:veremos 
	// the map will ignore automatically any name has a key wich is already stored in the map,
	for _, name := range names2 {
		uniqueMap[name] = true

	}
	//now , just convert the map into a slice and return it
	var result []string
	for name := range uniqueMap {
		result = append(result, name)
	}
	return result
}
