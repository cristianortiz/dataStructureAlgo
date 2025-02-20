package findallmovies

// implement the findAllMoviFan function that takes a movie, and a map consisting of people's names 
// mapped to their respective favorites movies. The function should return a slice containing the names of the people, 
//in any order, who likes the movie.
//  for example the following code should display the name 'Chad'
movies:= map[string][]string{
	"steve":[]string{"StarWars","Rocky","TheLordsOfTheRings"},
	"patty":[]string{"Avengers","Ray","TheLastSamurai"},
	"chad":[]string{"Avengers","TheGoodFather","ReadyPlayerOne"},
	} 
	 fmt.println(findAllMovieFans("TheGoodFather",movies)
)

func findAllMovieFans(movie string, movies map[string][]string) []string {
	var result []string

	// Iterates over every person and their favorite movies list
	for person, movieList := range movies {
		for _, m := range movieList {
			if h == movie {
				result = append(result, person)
				break // if we find the movie in movie list, there is no reason to continue the search
			}
		}
	}

	return result
}
