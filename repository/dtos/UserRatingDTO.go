package dtos

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	UrlUserRating = "user.rating?handle=Fefer_Ivan"
	Handle        = "handle"
)

/*
UserRatingResultDTO
```go
{
  "contestId": 1,
  "contestName": "Codeforces Beta Round #1",
  "handle": "Fefer_Ivan",
  "rank": 30,
  "ratingUpdateTimeSeconds": 1266588000,
  "oldRating": 0,
  "newRating": 1502
}
```
*/
type UserRatingResultDTO struct {
	ContestID               int    `json:"contestId"`
	ContestName             string `json:"contestName"`
	Handle                  string `json:"handle"`
	Rank                    int    `json:"rank"`
	RatingUpdateTimeSeconds int    `json:"ratingUpdateTimeSeconds"`
	OldRating               int    `json:"oldRating"`
	NewRating               int    `json:"newRating"`
}

// UserRatingDTO Sample https://codeforces.com/api/user.rating?handle=Fefer_Ivan
type UserRatingDTO struct {
	BaseDTO
	Result []UserRatingResultDTO `json:"result"`
}

// RatingChange Returns the change in rating
// @return int: i.e. NewRating - OldRating
func (d UserRatingResultDTO) RatingChange() int {
	return d.NewRating - d.OldRating
}

func GetUserRating(handle string) (*UserRatingDTO, error) {
	req, err := http.NewRequest(http.MethodGet, BaseUrl+UrlUserRating, nil)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	queryParams := req.URL.Query()
	queryParams.Add(Handle, handle)

	client := &http.Client{}
	httpResponse, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	response, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return nil, err
	}
	var result = new(UserRatingDTO)
	err = json.Unmarshal(response, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
