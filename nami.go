package namigo

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/thisni1s/nami-go/types"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
)

func errchck(err *error) {
	if *err != nil {
		log.Fatal(*err)
	}
}

var jar *cookiejar.Jar
var client *http.Client

func Login(username string, password string) error {
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

func Logout() (bool, error) {
	response, error := client.Get(types.LOGOUT_URL)
	if response.StatusCode != 204 {
		return false, error
	}
	defer response.Body.Close()
	return true, error
}

func GetMemberDetails(userid string, groupid string) (types.Member, error) {
	var answer types.SearchMemberAnswer
	response, error := client.Get(fmt.Sprintf("%s%s/%s", types.MEMBER_URL, groupid, userid))
	if error != nil {
		return answer.Member, error
	}
	defer response.Body.Close()
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer.Member, nil
}

func GetMemberActivities(userid string) ([]types.ActivityListItem, error) {
	var answer types.ActivitiesAnswer
	response, error := client.Get(fmt.Sprintf("%s%s/flist", types.MEMBER_ACTIVITIES_URL, userid))
	if error != nil {
		return answer.Activities, error
	}
	defer response.Body.Close()
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer.Activities, error
}

func GetActivityById(userid string, activityId string) (types.Activity, error) {
	var answer types.ActivityAnswer
	response, error := client.Get(fmt.Sprintf("%s%s/%s", types.MEMBER_ACTIVITIES_URL, userid, activityId))
	if error != nil {
		return answer.Act, error
	}
	defer response.Body.Close()
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer.Act, error
}

func GetMemberEducation(userid string) ([]types.EducationListItem, error) {
	var answer types.EducationsAnswer
	response, error := client.Get(fmt.Sprintf("%s/%s/flist", types.MEMBER_EDUCATION_URL, userid))
	if error != nil {
		return answer.Educations, error
	}
	defer response.Body.Close()
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer.Educations, error
}

func GetEducationById(userid string, educationId string) (types.Education, error) {
	var answer types.EducationAnswer
	response, error := client.Get(fmt.Sprintf("%s/%s/%s", types.MEMBER_EDUCATION_URL, userid, educationId))
	if error != nil {
		return answer.Edu, error
	}
	defer response.Body.Close()
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer.Edu, error
}

func GetTags(userid string) ([]types.TagListEntry, error) {
	var answer types.MemberTagsAnswer
	response, error := client.Get(fmt.Sprintf("%s/%s/flist", types.TAG_URL, userid))
	fmt.Println(response.Request.URL)
	if error != nil {
		return answer.Tags, error
	}
	defer response.Body.Close()
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer.Tags, error
}

func GetTagById(userid string, tagId string) (types.Tag, error) {
	var answer types.MemberTagAnswer
	response, error := client.Get(fmt.Sprintf("%s/%s/%s", types.TAG_URL, userid, tagId))
	if error != nil {
		return answer.Tag, error
	}
	defer response.Body.Close()
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer.Tag, error
}

func GetGroupTags() ([]types.GroupTagListItem, error) {
	var answer types.GroupTagsAnswer
	request, error := http.NewRequest("GET", fmt.Sprintf("%sflist", types.GROUP_TAG_URL), nil)
	request.URL.RawQuery = url.Values{
		"page":  {"1"},
		"start": {"0"},
		"limit": {"999999"},
	}.Encode()
	if error != nil {
		return answer.Tags, error
	}
	response, error := client.Do(request)
	if error != nil {
		return answer.Tags, error
	}
	defer response.Body.Close()
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer.Tags, error
}

func GetGroupTagById(tagid string) (types.GroupTag, error) {
	var answer types.GroupTagAnswer
	response, error := client.Get(fmt.Sprintf("%s%s", types.GROUP_TAG_URL, tagid))
	if error != nil {
		return answer.Tag, error
	}
	defer response.Body.Close()
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer.Tag, error
}

func getJsonStringFromBody(body io.ReadCloser) (string, error) {
	jason, error := io.ReadAll(body)
	return string(jason[:]), error
}

func Search(searchValues types.SearchValues) ([]types.SearchMember, error) {
	var answer types.SearchAnswer
	if !searchValues.Verify() {
		return answer.Members, errors.New("nami: EbeneID, GroupID and GroupName are mutually exclusive in SearchValues. You set more then one of them!")
	}
	request, error := http.NewRequest("GET", types.SEARCH_URL, nil)
	params, error := json.Marshal(searchValues)
	if error != nil {
		return answer.Members, error
	}
	request.URL.RawQuery = url.Values{
		"searchedValues": {string(params)},
		"page":           {"1"},
		"start":          {"0"},
		"limit":          {"999999"},
	}.Encode()
	response, error := client.Do(request)
	if error != nil {
		return answer.Members, error
	}
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer.Members, error
}

func GetGroupInfo() (types.GroupInfo, error) {
	var answer types.GroupInfoAnswer
	response, error := client.Get(types.GROUPINFO_URL)
	if error != nil {
		return answer.Group[0], error
	}
	defer response.Body.Close()
	error = json.NewDecoder(response.Body).Decode(&answer)
	return answer.Group[0], error
}

/*
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
    for _, member := range memlist {
        fmt.Println(member.Vorname, member.Nachname)
    }


    //acts, _ := getMemberActivities(cnf.Username)
    //act, _ := getActivityById(cnf.Username, "123456")
    //answer, err := getMemberEducation(cnf.Username)
    //edus, _ := getEducationById(cnf.Username, "12345")
    //usertags, _ := getTags(cnf.Username)
    //tag, _ := getTagById(cnf.Username, "123456")
    //tags, _ := getGroupTags()
	errchck(&err)

}

*/
