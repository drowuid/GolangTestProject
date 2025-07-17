package data

type Joke struct {

	ID int `json:"id"`
	Setup string `json:"setup"`
	Punchline string `json:"punchline"`
	Category string `json:"category"`
}

var Jokes = []Joke{

	{ID: 1, Setup: "Why don't scientists trust atoms?", Punchline: "They make up everything!", Category: "science"},
    {ID: 2, Setup: "How does a penguin build its house?", Punchline: "Igloos it together.", Category: "animals"},
    {ID: 3, Setup: "Why did the computer go to therapy?", Punchline: "It had too many processing issues.", Category: "tech"},
    {ID: 4, Setup: "What do you call fake spaghetti?", Punchline: "An impasta.", Category: "food"},
    {ID: 5, Setup: "Why was the math book sad?", Punchline: "It had too many problems.", Category: "math"},
}
