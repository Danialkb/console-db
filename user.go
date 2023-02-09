package main

import "console-db/models"

func main() {
	ss := &models.SearchService{}

	ss.SearchByName("sh")
	//models.Some_func()

}
