package tasks

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/VATUSA/api-v3/pkg/certificate"
	"github.com/VATUSA/api-v3/pkg/database"
	"github.com/VATUSA/api-v3/pkg/vatsim"
	"gorm.io/gorm"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
)

const (
	PageSize           = 100
	WorkerRoutines int = 10
)

type RosterPage struct {
	Items []vatsim.Member `json:"items"`
}

func SyncRosterFromVATSIM() {
	var wg sync.WaitGroup
	for i := 0; i < WorkerRoutines; i++ {
		wg.Add(1)
		go func(offset int) {
			defer wg.Done()
			SyncPageWorker(offset)
		}(i)
	}
	wg.Wait()
}

func SyncPageWorker(offset int) {
	for i := 0; true; i++ {
		page := (i * WorkerRoutines) + offset
		members, err := FetchDivisionRosterPage(page)
		if err != nil {
			println(fmt.Sprintf("Error in SyncPageWorker %d: %s", offset, err.Error()))
		} else if len(members) == 0 {
			break
		}
		ProcessMembers(members)
	}
}

func ProcessMembers(members []vatsim.Member) {
	for _, member := range members {
		err := ProcessMember(member)
		if err != nil {
			println(fmt.Sprintf("Error during ProcessMember for CID %d: %s", member.ID, err.Error()))
		}
	}
}

func ProcessMember(member vatsim.Member) error {
	cert, err := database.FetchCertificateByID(member.ID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	if cert != nil {
		err := certificate.UpdateCertificate(cert, &member)
		if err != nil {
			return err
		}
	} else {
		err := certificate.CreateCertificate(&member)
		if err != nil {
			return err
		}
	}
	// TODO
	return nil
}

func FetchDivisionRosterPage(page int) ([]vatsim.Member, error) {
	println(fmt.Sprintf("FetchDivisionRosterPage - Fetching Page %d", page))
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
