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

// Kontoverbindung represents all payment related information a member can has.
type Kontoverbindung struct {
	ID                  int    `json:"id"`
	Institut            string `json:"institut"`
	Bankleitzahl        string `json:"bankleitzahl"`
	Kontonummer         string `json:"kontonummer"`
	Iban                string `json:"iban"`
	Bic                 string `json:"bic"`
	Kontoinhaber        string `json:"kontoinhaber"`
	MitgliedsNummer     int    `json:"mitgliedsNummer"`
	ZahlungsKonditionID int    `json:"zahlungsKonditionId"` // ZahlungsKonditionID is either 1 or 2 depending on direct debit or sepa transaction. See also [thisni1s/nami-go/types/Zahlungskondition]
	ZahlungsKondition   string `json:"zahlungsKondition"`   // ZahlungsKondition is either "Std Lastschrift" or "Std Überweisung". See also [thisni1s/nami-go/types/Zahlungskondition]
}

// Member represents all the information NAMI has about one Member
type Member struct {
	ID                           int             `json:"id"`              // ID is the NAMI internal ID
	MitgliedsNummer              int             `json:"mitgliedsNummer"` // MitgliedsNummer is the NAMI Mitgliedsnummer
	Vorname                      string          `json:"vorname"`
	Spitzname                    string          `json:"spitzname"`
	Nachname                     string          `json:"nachname"`
	NameZusatz                   string          `json:"nameZusatz"`
	GeburtsDatum                 string          `json:"geburtsDatum"`
	Geschlecht                   string          `json:"geschlecht"`   // Geschlecht and GeschlechtID can together form a Geschlecht object defined in types. [thisni1s/nami-go/types/Geschlecht]
	GeschlechtID                 int             `json:"geschlechtId"` // GeschlechtID and Geschlecht can together form a Geschlecht object defined in types. [thisni1s/nami-go/types/Geschlecht]
	Stufe                        string          `json:"stufe"`        // String containing the "Stufenname"
	Strasse                      string          `json:"strasse"`
	Ort                          string          `json:"ort"`
	Plz                          string          `json:"plz"`
	Beitragsart                  string          `json:"beitragsart"`   // Beitragsart and BeitragsartID can together form a Beitragsart object defined in types. [thisni1s/nami-go/types/Beitragsart]
	BeitragsartID                int             `json:"beitragsartId"` // BeitragsartID and Beitragsart can together form a Beitragsart object defined in types. [thisni1s/nami-go/types/Beitragsart]
	FixBeitrag                   any             `json:"fixBeitrag"`
	Staatsangehoerigkeit         string          `json:"staatsangehoerigkeit"`
	StaatsangehoerigkeitID       int             `json:"staatsangehoerigkeitId"`
	StaatsangehoerigkeitText     string          `json:"staatsangehoerigkeitText"`
	Land                         string          `json:"land"`
	LandID                       int             `json:"landId"`
	MglType                      string          `json:"mglType"` // MglType is either "MITGLIED" or "NICHTMITGLIED"
	MglTypeID                    string          `json:"mglTypeId"`
	Email                        string          `json:"email"`
	EmailVertretungsberechtigter string          `json:"emailVertretungsberechtigter"`
	ErsteTaetigkeit              any             `json:"ersteTaetigkeit"`        // ErsteTaetigkeit is usually not set
	ErsteTaetigkeitID            any             `json:"ersteTaetigkeitId"`      // ErsteUntergliederungID is usually not set
	ErsteUntergliederung         string          `json:"ersteUntergliederung"`   // ErsteUntergliederung is the "Stufe"
	ErsteUntergliederungID       any             `json:"ersteUntergliederungId"` // ErsteUntergliederungID can have the values defined in [thisni1s/nami-go/types/UNTERGLIEDERUNG]
	Gruppierung                  string          `json:"gruppierung"`            // Gruppierung contains the name of the "Stammgruppierung"
	GruppierungID                int             `json:"gruppierungId"`          // GruppierungID is the internal ID of the "Stammgruppierung"
	Eintrittsdatum               string          `json:"eintrittsdatum"`
	AustrittsDatum               string          `json:"austrittsDatum"`
	Region                       string          `json:"region"`   // Region is the Federal State
	RegionID                     int             `json:"regionId"` // RegionID is the Federal State ID set by NAMI
	Status                       string          `json:"status"`   // Status usually is "AKTIV"
	Konfession                   string          `json:"konfession"`
	KonfessionID                 int             `json:"konfessionId"`
	Rover                        any             `json:"rover"`          // Rover is almost never set
	Pfadfinder                   any             `json:"pfadfinder"`     // Pfadfinder is almost never set
	Jungpfadfinder               any             `json:"jungpfadfinder"` // Jungpfadfinder is almost never set
	Woelfling                    any             `json:"woelfling"`      // Woelfling is almost never set
	Telefon1                     string          `json:"telefon1"`       // Telefon1 is the landline field in NAMI
	Telefon2                     string          `json:"telefon2"`       // Telefon2 is the mobile field in NAMI
	Telefon3                     string          `json:"telefon3"`       // Telefon3 is the business field in NAMI
	Telefax                      string          `json:"telefax"`
	Kontoverbindung              Kontoverbindung `json:"kontoverbindung"`
	Zeitschriftenversand         bool            `json:"zeitschriftenversand"`
	WiederverwendenFlag          bool            `json:"wiederverwendenFlag"`
	GenericField1                string          `json:"genericField1"`
	GenericField2                string          `json:"genericField2"`
	LastUpdated                  string          `json:"lastUpdated"`
	Version                      int             `json:"version"`
	Sonst01                      bool            `json:"sonst01"`
	Sonst02                      bool            `json:"sonst02"`
}

type SearchAnswer struct {
	genericSearchAnswer
	Members []SearchMember `json:"data"`
}
type SearchMember struct {
	ID                           int    `json:"id"`                      // ID is the NAMI internal ID
	EntriesID                    int    `json:"entries_id"`              // EntriesID is the NAMI internal ID
	MitgliedsNummer              int    `json:"entries_mitgliedsNummer"` // MitgliedsNummer is the NAMI Mitgliedsnummer
	Vorname                      string `json:"entries_vorname"`
	Nachname                     string `json:"entries_nachname"`
	Spitzname                    string `json:"entries_spitzname"`
	GeburtsDatum                 string `json:"entries_geburtsDatum"`
	Stufe                        string `json:"entries_stufe"`    // Stufe is the current "Stufe"
	Telefon1                     string `json:"entries_telefon1"` // Telefon1 is the landline field in NAMI
	Telefon2                     string `json:"entries_telefon2"` // Telefon2 is the mobile field in NAMI
	Telefon3                     string `json:"entries_telefon3"` // Telefon3 is the business field in NAMI
	Telefax                      string `json:"entries_telefax"`
	Staatsangehoerigkeit         string `json:"entries_staatsangehoerigkeit"`
	StaatangehoerigkeitText      string `json:"entries_staatangehoerigkeitText"` // StaatangehoerigkeitText is usually null
	TaetigkeitID                 any    `json:"entries_ersteTaetigkeitId"`       // TaetigkeitID is usually null
	ErsteUntergliederungID       any    `json:"entries_ersteUntergliederungId"`  // ErsteUntergliederungID is always null
	Beitragsarten                string `json:"entries_beitragsarten"`           // Beitragsarten is usually null
	Email                        string `json:"entries_email"`
	EmailVertretungsberechtigter string `json:"entries_emailVertretungsberechtigter"`
	Konfession                   string `json:"entries_konfession"`
	FixBeitrag                   string `json:"entries_fixBeitrag"`
	Status                       string `json:"entries_status"`  // Status is usually "AKTIV"
	MglType                      string `json:"entries_mglType"` // MglType is either "MITGLIED" or "NICHTMITGLIED"
	Geschlecht                   string `json:"entries_geschlecht"`
	Eintrittsdatum               string `json:"entries_eintrittsdatum"`
	AustrittsDatum               string `json:"entries_austrittsDatum"`
	Kontoverbindung              string `json:"entries_kontoverbindung"` // Kontoverbindung always null
	WiederverwendenFlag          bool   `json:"entries_wiederverwendenFlag"`
	Rover                        string `json:"entries_rover"`          // Rover is always null
	Pfadfinder                   string `json:"entries_pfadfinder"`     // Pfadfinder is always null
	Jungpfadfinder               string `json:"entries_jungpfadfinder"` // Jungpfadfinder is always null
	Woelfling                    string `json:"entries_woelfling"`      // Woelfling is always null
	LastUpdated                  string `json:"entries_lastUpdated"`
	Descriptor                   string `json:"descriptor"` // Descriptor usually is "Nachname, Vorname"
	GenericField1                string `json:"entries_genericField1"`
	GenericField2                string `json:"entries_genericField2"`
	RepresentedClass             string `json:"representedClass"`
	RowCSSClass                  string `json:"entries_rowCssClass"`
	Version                      int    `json:"entries_version"`
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
