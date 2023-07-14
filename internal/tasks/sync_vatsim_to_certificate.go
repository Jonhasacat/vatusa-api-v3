package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/VATUSA/api-v3/pkg/vatsim"
	"io/ioutil"
	"net/http"
	"os"
)

const (
	PageSize = 100
)

type RosterPage struct {
	Items []vatsim.Member `json:"items"`
}

func FetchDivisionRosterPage(page uint) ([]vatsim.Member, error) {
	client := http.Client{}
	url := fmt.Sprintf("%s/v2/orgs/division/USA?limit=%d&offset=%d", os.Getenv("VATSIM_API_URL"), PageSize, page*PageSize)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", fmt.Sprintf("Token %s", os.Getenv("VATSIM_API_TOKEN")))
	response, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var rosterPage RosterPage
	err = json.Unmarshal(responseData, &rosterPage)
	if err != nil {
		return nil, err
	}
	return rosterPage.Items, nil
}
