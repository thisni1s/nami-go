package main

import (
	"encoding/json"
	"fmt"
	"github.com/thisni1s/nami-go/types"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

func errchck(err *error) {
	if *err != nil {
		log.Fatal(*err)
	}
}

var jar *cookiejar.Jar
var client *http.Client

func login(username string, password string) error {
	jar, _ = cookiejar.New(nil)
	client = &http.Client{
		Jar: jar,
	}
	data := url.Values{
		"Login":    {"API"},
		"username": {username},
		"password": {password},
	}
	response, error := client.PostForm(types.LOGIN_URL, data)
	if error != nil {
		return error
	}
	defer response.Body.Close()
	return nil
}

func getMemberDetails(userid string, groupid string) (types.SearchMemberAnswer, error) {
	var answer types.SearchMemberAnswer
	response, error := client.Get(types.FillMemberUrl(userid, groupid))
	if error != nil {
		return answer, error
	}
	defer response.Body.Close()
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer, nil
}

func search(searchValues types.SearchValues) (types.SearchAnswer, error) {
	var answer types.SearchAnswer
	request, error := http.NewRequest("GET", types.SEARCH_URL, nil)
	params, error := json.Marshal(searchValues)
	if error != nil {
		return answer, error
	}
	request.URL.RawQuery = url.Values{
		"searchedValues": {string(params)},
		"page":           {"1"},
		"start":          {"0"},
		"limit":          {"999999"},
	}.Encode()
	response, error := client.Do(request)
	if error != nil {
		return answer, error
	}
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer, error
}

func main() {
	conf, err := os.ReadFile("./config.json")
	errchck(&err)

	var cnf types.Config
	err = json.Unmarshal(conf, &cnf)
	errchck(&err)

	err = login(cnf.Username, cnf.Password)

	memlist, err := search(types.SearchValues{
		UntergliederungID: types.UG_ROVER,
		MglTypeID:         types.MITGLIED,
		MglStatusID:       types.AKTIV,
	})
	errchck(&err)
	for _, member := range memlist.Members {
		fmt.Println(member.Vorname, member.Nachname)
	}

}
