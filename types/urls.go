package types

const LOGIN_URL string = "https://nami.dpsg.de/ica/rest/nami/auth/manual/sessionStartup";
const MEMBER_URL string = "https://nami.dpsg.de/ica/rest/nami/mitglied/filtered-for-navigation/gruppierung/gruppierung/";
const SEARCH_URL string = "https://nami.dpsg.de/ica/rest/nami/search/result-list";

func FillMemberUrl(userid string, groupid string) string {
    return (MEMBER_URL + groupid + "/" + userid)
}
