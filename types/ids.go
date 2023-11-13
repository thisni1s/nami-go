package types

type UNTERGLIEDERUNG int

const (
	UG_WOE   UNTERGLIEDERUNG = 1
	UG_JUFFI UNTERGLIEDERUNG = 2
	UG_PFADI UNTERGLIEDERUNG = 3
	UG_ROVER UNTERGLIEDERUNG = 4
	UG_STAVO UNTERGLIEDERUNG = 5
	UG_SONST UNTERGLIEDERUNG = 48
)

type TAETIGKEIT int

const (
	TG_MITGLIED         TAETIGKEIT = 1
	TG_ELTERNVERTRETUNG TAETIGKEIT = 4
	TG_LEITER           TAETIGKEIT = 6
	TG_DELEGIERT        TAETIGKEIT = 7
	TG_BEOBACHTER       TAETIGKEIT = 8
	TG_KURAT            TAETIGKEIT = 11
	TG_VORSITZ          TAETIGKEIT = 13
	TG_ADMIN            TAETIGKEIT = 14
	TG_SONSTMITARBEITER TAETIGKEIT = 16
	TG_MITGLIEDAK       TAETIGKEIT = 18
	TG_GESCHÄFTSFÜHRER  TAETIGKEIT = 19
	TG_KASSIERER        TAETIGKEIT = 20
	TG_KASSENPRÜFER     TAETIGKEIT = 21
	TG_MATWART          TAETIGKEIT = 23
	TG_SCHNUPPER        TAETIGKEIT = 35
	TG_PASSIV           TAETIGKEIT = 39
	TG_SONSTMITGLIED    TAETIGKEIT = 40
	TG_SONSTEXT         TAETIGKEIT = 41
)

type MITGLIEDTYPEID string

const (
	MITGLIED      MITGLIEDTYPEID = "MITGLIED"
	NICHTMITGLIED MITGLIEDTYPEID = "Nicht-Mitglied"
)

type MITGLIEDSTATUSID string

const (
	AKTIV MITGLIEDSTATUSID = "AKTIV"
)

type Zahlungskondition struct {
	Name string `json:"descriptor"`
	ID   int    `json:"id"`
}

var ZahlKondLastschrift Zahlungskondition = Zahlungskondition{Name: "Std Lastschrift", ID: 1}
var ZahlKondÜberweisung Zahlungskondition = Zahlungskondition{Name: "Std Überweisung", ID: 2}

type BeitragsartMgl struct {
	Name string `json:"descriptor"`
	ID   int    `json:"id"`
}

var BeitragMglVoll BeitragsartMgl = BeitragsartMgl{Name: "DPSG Bundesverband 000000 (Voller Beitrag - VERBANDSBEITRAG)", ID: 1}
var BeitragMglFamErmäßigt BeitragsartMgl = BeitragsartMgl{Name: "DPSG Bundesverband 000000 (Familienermäßigt - VERBANDSBEITRAG)", ID: 2}
var BeitragMglErmäßigt BeitragsartMgl = BeitragsartMgl{Name: "DPSG Bundesverband 00000…ßigt - VERBANDSBEITRAG)", ID: 3}
var BeitragMglVollStiftung BeitragsartMgl = BeitragsartMgl{Name: "DPSG Bundesverband 000000 (Voller Beitrag - Stiftungseuro - VERBANDSBEITRAG)", ID: 4}
var BeitragMglFamErmäßigtStiftung BeitragsartMgl = BeitragsartMgl{Name: "DPSG Bundesverband 000000 (Familienermäßigt - Stiftungseuro - VERBANDSBEITRAG)", ID: 5}
var BeitragMglSozialErmäßigtStiftung BeitragsartMgl = BeitragsartMgl{Name: "DPSG Bundesverband 000000 (Sozialermäßigt - Stiftungseuro - VERBANDSBEITRAG)", ID: 6}
var BeitragMglÜbernahme BeitragsartMgl = BeitragsartMgl{Name: "DPSG Bundesverband 000000 (Übernahme - VERBANDSBEITRAG)", ID: 7}

type Beitragsart struct {
	Name string `json:"descriptor"`
	ID   int    `json:"id"`
}

var BeitragVoll Beitragsart = Beitragsart{Name: "Voller Beitrag", ID: 1}
var BeitragFamErmäßigt Beitragsart = Beitragsart{Name: "Familienermäßigt", ID: 2}
var BeitragErmäßigt Beitragsart = Beitragsart{Name: "Sozialermäßigt", ID: 3}
var BeitragVollStiftung Beitragsart = Beitragsart{Name: "Voller Beitrag - Stiftungseuro", ID: 4}
var BeitragFamErmäßigtStiftung Beitragsart = Beitragsart{Name: "Familienermäßigt - Stiftungseuro", ID: 5}
var BeitragSozialErmäßigtStiftung Beitragsart = Beitragsart{Name: "Sozialermäßigt - Stiftungseuro", ID: 6}
var BeitragÜbernahme Beitragsart = Beitragsart{Name: "Übernahme", ID: 7}

type Geschlecht struct {
	Name string `json:"descriptor"`
	ID   int    `json:"id"`
}

var GeschlechtDivers Geschlecht = Geschlecht{Name: "divers", ID: 24}
var GeschlechtKeinAngabe Geschlecht = Geschlecht{Name: "keine Angabe", ID: 23}
var GeschlechtMännlich Geschlecht = Geschlecht{Name: "männlich", ID: 19}
var GeschlechtWeiblich Geschlecht = Geschlecht{Name: "weiblich", ID: 20}
