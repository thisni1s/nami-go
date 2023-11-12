package types

const LOGIN_URL string = "https://nami.dpsg.de/ica/rest/nami/auth/manual/sessionStartup"
const MEMBER_URL string = "https://nami.dpsg.de/ica/rest/nami/mitglied/filtered-for-navigation/gruppierung/gruppierung/"
const SEARCH_URL string = "https://nami.dpsg.de/ica/rest/nami/search/result-list"
const GROUPINFO_URL string = "https://nami.dpsg.de/ica/rest/nami/gruppierungen/filtered-for-navigation/gruppierung/node/root"
const MEMBER_HISTORY_URL string = "https://nami.dpsg.de/ica/rest/nami/mitglied-history/filtered-for-navigation/mitglied/mitglied/"
const MEMBER_ACTIVITIES_URL string = "https://nami.dpsg.de/ica/rest/nami/zugeordnete-taetigkeiten/filtered-for-navigation/gruppierung-mitglied/mitglied/"
const MEMBER_EDUCATION_URL string = "https://nami.dpsg.de/ica/rest/nami/mitglied-ausbildung/filtered-for-navigation/mitglied/mitglied/"
const TAG_URL string = "https://nami.dpsg.de/ica/rest/nami/mitglied-tagged-item/filtered-for-navigation/identitaet/identitaet/"
const GROUP_TAG_URL string = "https://nami.dpsg.de/ica/rest/tagging/"

func FillMemberUrl(userid string, groupid string) string {
	return (MEMBER_URL + groupid + "/" + userid)
}
