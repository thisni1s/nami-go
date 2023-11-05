package types

type SearchAnswer struct {
	Success      bool `json:"success"`
	Data         Data `json:"data"`
	ResponseType any  `json:"responseType"`
	Message      any  `json:"message"`
	Title        any  `json:"title"`
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
type Data struct {
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
