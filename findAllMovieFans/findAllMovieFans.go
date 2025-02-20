package findallmoviefans

// implement the findAllMovies function that takes a movie, and a map consisting of people's names
// mapped to their respective favorites movies. The function should return a slice containing the names of the people,
//in any order, who likes the movie.
//  for example the following code should display the name 'Chad'
// hobbies:= map[string][]string{
// 	"steve":[]string{"StarWars","Rocky","TheLordsOfTheRings"},
// 	"patty":[]string{"Avengers","Ray","TheLastSamurai"},
// 	"chad":[]string{"Avengers","TheGoodFather","ReadyPlayerOne"},
// 	}
// 	 fmt.println(findAllHobbysts("TheGoodFather",hobbies)
// )

func findAllMovieFans(movie string, movies map[string][]string) []string {
	var result []string

	// Iterate over every person in the map and check their favorite movie list
	for person, movieList := range movies {
		for _, m := range movieList {
			if m == movie {
				result = append(result, person)
				break // if we found the movie in the movieList, there is no reason to continue the iteration
			}
		}
	}

	return result
}
