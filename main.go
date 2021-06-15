package main

import (
	"fmt"
	"github.com/rishabhdeepsingh/codeforces_api/repository/dtos"
)

func main() {
	jsonResponse, err := dtos.GetUserRating("ads")
	if err != nil {
		return
	}
	fmt.Println(jsonResponse.Status)
	fmt.Println(jsonResponse.Result[1].ContestName)
	fmt.Println(jsonResponse.Result[1].Rank)
	fmt.Println(jsonResponse.Result[0].RatingChange())
}
