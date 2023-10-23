package main

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Quantity int    `json:"quantity"`
}

var books = []book{
	{ID: "1", Title: "Solo Leveling", Author: "Chu-Gong", Quantity: 5},
	{ID: "2", Title: "The Beginning after the end", Author: "TurtleMe", Quantity: 10},
	{ID: "3", Title: "Omnicient reader", Author: "Sing Shong", Quantity: 12},
}

func main() {

}
