package types

type SearchMemberAnswer struct {
	Success      bool   `json:"success"`
	Member       Member `json:"data"`
	ResponseType any    `json:"responseType"`
	Message      any    `json:"message"`
	Title        any    `json:"title"`
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
	Success      bool           `json:"success"`
	Members      []SearchMember `json:"data"`
	ResponseType string         `json:"responseType"`
	TotalEntries int            `json:"totalEntries"`
	MetaData     MetaData       `json:"metaData"`
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
type Fields struct {
	Name   string `json:"name"`
	Header any    `json:"header"`
	Hidden bool   `json:"hidden"`
	Width  int    `json:"width,omitempty"`
	Flex   int    `json:"flex,omitempty"`
}
type MetaData struct {
	TotalProperty string   `json:"totalProperty"`
	Root          string   `json:"root"`
	ID            string   `json:"id"`
	Fields        []Fields `json:"fields"`
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

type Config struct {
	Username    string
	Password    string
	Gruppierung string
}
