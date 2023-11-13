package types

type genericSearchAnswer struct {
	Success      bool     `json:"success"`
	ResponseType string   `json:"responseType"`
	TotalEntries int      `json:"totalEntries"`
	MetaData     metaData `json:"metaData"`
}

type metaFields struct {
	Name   string `json:"name"`
	Header any    `json:"header"`
	Hidden bool   `json:"hidden"`
	Width  int    `json:"width,omitempty"`
	Flex   int    `json:"flex,omitempty"`
}
type metaData struct {
	TotalProperty string       `json:"totalProperty"`
	Root          string       `json:"root"`
	ID            string       `json:"id"`
	Fields        []metaFields `json:"fields"`
}

type genericSingleAnswer struct {
	Success      bool `json:"success"`
	ResponseType any  `json:"responseType"`
	Message      any  `json:"message"`
	Title        any  `json:"title"`
}

type SearchMemberAnswer struct {
	genericSearchAnswer
	Member Member `json:"data"`
}
type Kontoverbindung struct {
	ID                  int    `json:"id"`
	Institut            string `json:"institut"`
	Bankleitzahl        string `json:"bankleitzahl"`
	Kontonummer         string `json:"kontonummer"`
	Iban                string `json:"iban"`
	Bic                 string `json:"bic"`
	Kontoinhaber        string `json:"kontoinhaber"`
	MitgliedsNummer     int    `json:"mitgliedsNummer"`
	ZahlungsKonditionID int    `json:"zahlungsKonditionId"`
	ZahlungsKondition   string `json:"zahlungsKondition"`
}
type Member struct {
	Jungpfadfinder               any             `json:"jungpfadfinder"`
	MglType                      string          `json:"mglType"`
	Geschlecht                   string          `json:"geschlecht"`
	Staatsangehoerigkeit         string          `json:"staatsangehoerigkeit"`
	ErsteTaetigkeitID            any             `json:"ersteTaetigkeitId"`
	ErsteUntergliederung         string          `json:"ersteUntergliederung"`
	EmailVertretungsberechtigter string          `json:"emailVertretungsberechtigter"`
	LastUpdated                  string          `json:"lastUpdated"`
	ErsteTaetigkeit              any             `json:"ersteTaetigkeit"`
	NameZusatz                   string          `json:"nameZusatz"`
	ID                           int             `json:"id"`
	StaatsangehoerigkeitID       int             `json:"staatsangehoerigkeitId"`
	Version                      int             `json:"version"`
	Sonst01                      bool            `json:"sonst01"`
	Sonst02                      bool            `json:"sonst02"`
	Spitzname                    string          `json:"spitzname"`
	LandID                       int             `json:"landId"`
	StaatsangehoerigkeitText     string          `json:"staatsangehoerigkeitText"`
	GruppierungID                int             `json:"gruppierungId"`
	MglTypeID                    string          `json:"mglTypeId"`
	Beitragsart                  string          `json:"beitragsart"`
	Nachname                     string          `json:"nachname"`
	Eintrittsdatum               string          `json:"eintrittsdatum"`
	Rover                        any             `json:"rover"`
	Region                       string          `json:"region"`
	Status                       string          `json:"status"`
	Konfession                   string          `json:"konfession"`
	FixBeitrag                   any             `json:"fixBeitrag"`
	KonfessionID                 int             `json:"konfessionId"`
	Zeitschriftenversand         bool            `json:"zeitschriftenversand"`
	Pfadfinder                   any             `json:"pfadfinder"`
	Telefon3                     string          `json:"telefon3"`
	Kontoverbindung              Kontoverbindung `json:"kontoverbindung"`
	GeschlechtID                 int             `json:"geschlechtId"`
	Land                         string          `json:"land"`
	Email                        string          `json:"email"`
	Telefon1                     string          `json:"telefon1"`
	Woelfling                    any             `json:"woelfling"`
	Telefon2                     string          `json:"telefon2"`
	Strasse                      string          `json:"strasse"`
	Vorname                      string          `json:"vorname"`
	MitgliedsNummer              int             `json:"mitgliedsNummer"`
	Gruppierung                  string          `json:"gruppierung"`
	AustrittsDatum               string          `json:"austrittsDatum"`
	Ort                          string          `json:"ort"`
	ErsteUntergliederungID       any             `json:"ersteUntergliederungId"`
	WiederverwendenFlag          bool            `json:"wiederverwendenFlag"`
	RegionID                     int             `json:"regionId"`
	GeburtsDatum                 string          `json:"geburtsDatum"`
	Stufe                        string          `json:"stufe"`
	GenericField1                string          `json:"genericField1"`
	GenericField2                string          `json:"genericField2"`
	Telefax                      string          `json:"telefax"`
	BeitragsartID                int             `json:"beitragsartId"`
	Plz                          string          `json:"plz"`
}

type SearchAnswer struct {
	genericSearchAnswer
	Members []SearchMember `json:"data"`
}
type SearchMember struct {
	TaetigkeitID                 any    `json:"entries_ersteTaetigkeitId"`
	GenericField1                string `json:"entries_genericField1"`
	Version                      int    `json:"entries_version"`
	Telefon3                     string `json:"entries_telefon3"`
	Telefon2                     string `json:"entries_telefon2"`
	Telefon1                     string `json:"entries_telefon1"`
	Descriptor                   string `json:"descriptor"`
	EntriesID                    int    `json:"entries_id"`
	Staatsangehoerigkeit         string `json:"entries_staatsangehoerigkeit"`
	RepresentedClass             string `json:"representedClass"`
	Rover                        string `json:"entries_rover"`
	Pfadfinder                   string `json:"entries_pfadfinder"`
	MitgliedsNummer              int    `json:"entries_mitgliedsNummer"`
	WiederverwendenFlag          bool   `json:"entries_wiederverwendenFlag"`
	ErsteUntergliederungID       any    `json:"entries_ersteUntergliederungId"`
	RowCSSClass                  string `json:"entries_rowCssClass"`
	Vorname                      string `json:"entries_vorname"`
	ID                           int    `json:"id"`
	Woelfling                    string `json:"entries_woelfling"`
	Beitragsarten                string `json:"entries_beitragsarten"`
	Stufe                        string `json:"entries_stufe"`
	Email                        string `json:"entries_email"`
	Konfession                   string `json:"entries_konfession"`
	FixBeitrag                   string `json:"entries_fixBeitrag"`
	EmailVertretungsberechtigter string `json:"entries_emailVertretungsberechtigter"`
	LastUpdated                  string `json:"entries_lastUpdated"`
	Status                       string `json:"entries_status"`
	Jungpfadfinder               string `json:"entries_jungpfadfinder"`
	MglType                      string `json:"entries_mglType"`
	Kontoverbindung              string `json:"entries_kontoverbindung"`
	Geschlecht                   string `json:"entries_geschlecht"`
	Spitzname                    string `json:"entries_spitzname"`
	GeburtsDatum                 string `json:"entries_geburtsDatum"`
	StaatangehoerigkeitText      string `json:"entries_staatangehoerigkeitText"`
	Nachname                     string `json:"entries_nachname"`
	Eintrittsdatum               string `json:"entries_eintrittsdatum"`
	AustrittsDatum               string `json:"entries_austrittsDatum"`
	GenericField2                string `json:"entries_genericField2"`
	Telefax                      string `json:"entries_telefax"`
}

type SearchValues struct {
	Vorname              string           `json:"vorname,omitempty"`
	Nachname             string           `json:"nachname,omitempty"`
	MinAlter             int              `json:"alterVon,omitempty"`
	MaxAlter             int              `json:"alterBis,omitempty"`
	City                 string           `json:"mglWohnort,omitempty"`
	Mitgliedsnummer      string           `json:"mitgliedsNummber,omitempty"`
	MglStatusID          MITGLIEDSTATUSID `json:"mglStatusId,omitempty"`
	MglTypeID            MITGLIEDTYPEID   `json:"mglTypeId,omitempty"`
	TagID                string           `json:"tagId,omitempty"`
	Zeitschriftenversand bool             `json:"zeitschriftenversand,omitempty"`
	UntergliederungID    UNTERGLIEDERUNG  `json:"untergliederungId,omitempty"`
	TaetigkeitID         TAETIGKEIT       `json:"taetigkeitId,omitempty"`
	EbeneID              int              `json:"ebeneId,omitempty"`
	GroupID              string           `json:"grpNummer,omitempty"`
	GroupName            string           `json:"grpName,omitempty"`
	SearchType           string           `json:"searchType,omitempty"`
}

func (sv *SearchValues) Verify() bool {
	count := 0
	if sv.EbeneID != 0 {
		count++
	}
	if sv.GroupID != "" {
		count++
	}
	if sv.GroupName != "" {
		count++
	}
	if count > 1 {
		return false
	} else {
		return true
	}
}

type Config struct {
	Username    string
	Password    string
	Gruppierung string
}

type ActivitiesAnswer struct {
	genericSearchAnswer
	Activities []ActivityListItem `json:"data,omitempty"`
}
type ActivityListItem struct {
	AktivBis         string `json:"entries_aktivBis,omitempty"`
	BeitragsArt      string `json:"entries_beitragsArt,omitempty"`
	CaeaGroup        string `json:"entries_caeaGroup,omitempty"`
	AktivVon         string `json:"entries_aktivVon,omitempty"`
	Descriptor       string `json:"descriptor,omitempty"`
	RepresentedClass string `json:"representedClass,omitempty"`
	Anlagedatum      string `json:"entries_anlagedatum,omitempty"`
	GroupForGf       string `json:"entries_caeaGroupForGf,omitempty"`
	Untergliederung  string `json:"entries_untergliederung,omitempty"`
	Taetigkeit       string `json:"entries_taetigkeit,omitempty"`
	Gruppierung      string `json:"entries_gruppierung,omitempty"`
	ID               int    `json:"id,omitempty"`
	Mitglied         string `json:"entries_mitglied,omitempty"`
}

type Activity struct {
	ID            int        `json:"id,omitempty"`
	Gruppierung   string     `json:"gruppierung,omitempty"`
	GruppierungID int        `json:"gruppierungId,omitempty"`
	Taetigkeit    string     `json:"taetigkeit,omitempty"`
	TaetigkeitID  TAETIGKEIT `json:"taetigkeitId,omitempty"`
	CaeaGroup     string     `json:"caeaGroup,omitempty"`
	CaeaGroupID   int        `json:"caeaGroupId,omitempty"`
	AktivVon      string     `json:"aktivVon,omitempty"`
	AktivBis      string     `json:"aktivBis,omitempty"`
}

type ActivityAnswer struct {
	genericSingleAnswer
	Act Activity `json:"data,omitempty"`
}

type EducationsAnswer struct {
	genericSearchAnswer
	Educations []EducationListItem `json:"data,omitempty"`
}
type EducationListItem struct {
	VstgTag          string `json:"entries_vstgTag,omitempty"`
	Veranstalter     string `json:"entries_veranstalter,omitempty"`
	VstgName         string `json:"entries_vstgName,omitempty"`
	Baustein         string `json:"entries_baustein,omitempty"`
	ID               int    `json:"id,omitempty"`
	Descriptor       string `json:"descriptor,omitempty"`
	EntryID          int    `json:"entries_id,omitempty"`
	RepresentedClass string `json:"representedClass,omitempty"`
	Mitglied         string `json:"entries_mitglied,omitempty"`
}

type EducationAnswer struct {
	genericSingleAnswer
	Edu Education `json:"data,omitempty"`
}
type Education struct {
	ID               int    `json:"id,omitempty"`
	Baustein         string `json:"baustein,omitempty"`
	BausteinID       int    `json:"bausteinId,omitempty"`
	Mitglied         string `json:"mitglied,omitempty"`
	VstgTag          string `json:"vstgTag,omitempty"`
	VstgName         string `json:"vstgName,omitempty"`
	Veranstalter     string `json:"veranstalter,omitempty"`
	LastModifiedFrom string `json:"lastModifiedFrom,omitempty"`
}

type MemberTagsAnswer struct {
	genericSearchAnswer
	Tags []TagListEntry `json:"data,omitempty"`
}
type TagListEntry struct {
	Tag              string `json:"entries_tag,omitempty"`
	Identitaet       string `json:"entries_identitaet,omitempty"`
	ID               int    `json:"id,omitempty"`
	Descriptor       string `json:"descriptor,omitempty"`
	RepresentedClass string `json:"representedClass,omitempty"`
}

type MemberTagAnswer struct {
	genericSingleAnswer
	Tag Tag `json:"data,omitempty"`
}
type Tag struct {
	ID         int    `json:"id,omitempty"`
	Tag        string `json:"tag,omitempty"`
	TagID      int    `json:"tagId,omitempty"`
	Mitglied   string `json:"mitglied,omitempty"`
	MitgliedID int    `json:"mitgliedId,omitempty"`
}

type GroupTagsAnswer struct {
	genericSearchAnswer
	Tags []GroupTagListItem `json:"data,omitempty"`
}
type GroupTagListItem struct {
	Name             string `json:"entries_name,omitempty"`
	ID               int    `json:"id,omitempty"`
	Descriptor       string `json:"descriptor,omitempty"`
	RepresentedClass string `json:"representedClass,omitempty"`
}

type GroupTagAnswer struct {
	genericSingleAnswer
	Tag GroupTag `json:"data,omitempty"`
}
type GroupTag struct {
	ID                 int    `json:"id,omitempty"`
	Name               string `json:"name,omitempty"`
	OwnerGruppierung   string `json:"ownerGruppierung,omitempty"`
	OwnerGruppierungID int    `json:"ownerGruppierungId,omitempty"`
}

type GroupInfoAnswer struct {
	Success      bool        `json:"success,omitempty"`
	Group        []GroupInfo `json:"data,omitempty"`
	ResponseType string      `json:"responseType,omitempty"`
}
type GroupInfo struct {
	Descriptor       string `json:"descriptor,omitempty"`
	Name             string `json:"name,omitempty"`
	RepresentedClass string `json:"representedClass,omitempty"`
	ID               int    `json:"id,omitempty"`
}
