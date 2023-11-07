package types

const LOGIN_URL string = "https://nami.dpsg.de/ica/rest/nami/auth/manual/sessionStartup";
const MEMBER_URL string = "https://nami.dpsg.de/ica/rest/nami/mitglied/filtered-for-navigation/gruppierung/gruppierung/";
const SEARCH_URL string = "https://nami.dpsg.de/ica/rest/nami/search/result-list";
const GROUPINFO_URL string = "https://nami.dpsg.de/ica/rest/nami/gruppierungen/filtered-for-navigation/gruppierung/node/root"

func FillMemberUrl(userid string, groupid string) string {
    return (MEMBER_URL + groupid + "/" + userid)
}
