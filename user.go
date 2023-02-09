package main

import "console-db/models"

func main() {
	ss := &models.SearchService{}

	ss.SearchByName("dre")

}
