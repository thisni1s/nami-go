package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
)

type Config struct {
	Username string
	Password string
	Gruppierung string
}

type SearchAnswer struct {
	sucess bool
	responseType string 
	message string 
	title string 
	data struct {
		id int
		vorname string
		nachname string
	}
}

func errchck(err *error) {
	if *err != nil {
		log.Fatal(err)
	}
}

func main() {
	conf, err := os.ReadFile("./config.json")
	errchck(&err)

	var cnf Config 
	err = json.Unmarshal(conf, &cnf)
	errchck(&err)

	jar, _ := cookiejar.New(nil)
	data := url.Values{
		"Login": {"API"},
		"username": {cnf.Username},
		"password": {cnf.Password},
	}


	client := &http.Client{
		Jar: jar,
	}
	url := "https://nami.dpsg.de/ica/rest/nami/auth/manual/sessionStartup"

	resp, err := client.PostForm(url, data)

	errchck(&err)

	defer resp.Body.Close()

	res := make(map[string]string)
	json.NewDecoder(resp.Body).Decode(&res)

	url2 := fmt.Sprintf("https://nami.dpsg.de/ica/rest/nami/mitglied/filtered-for-navigation/gruppierung/gruppierung/%s/%s", "350161", "182574")

	r, err := http.NewRequest("GET", url2, nil)
	for _, c := range resp.Cookies() {
		r.AddCookie(c)
	}
	resp, err = client.Do(r)
	errchck(&err)

	defer resp.Body.Close() 

	//var sa SearchAnswer
	res2 := make(map[string]interface{})
	json.NewDecoder(resp.Body).Decode(&res2)
	
	for k, v := range res2 {
		fmt.Println(k, v)
	}

	//fmt.Println(res2.data.id)
	//fmt.Println(res2.data.vorname)
	//fmt.Println(res2.data.nachname)

}
