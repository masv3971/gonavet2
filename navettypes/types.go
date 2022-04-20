package navettypes

// HamtaRequest 5.1.2.1
type HamtaRequest struct {
	Bestallning Bestallning  `json:"bestallning" validate:"required"`
	Sokvillkor  []Sokvillkor `json:"sokvillkor" validate:"min=1"`
}

// OkHamtaResponse 5.1.2.2
type OkHamtaResponse struct {
	Fel                  []Felmeddelande           `json:"fel" validate:"omitempty"`
	Folkbokforingsposter []FolkbokforingspostHamta `json:"folkbokforingsposter" validate:"min=1"`
	Status               int                       `json:"status" validate:"required"`
	Tidpunkt             string                    `json:"tidpunkt" validate:"required"`
}

// BadRequest /FelResponse 5.1.2.3
type BadRequest struct {
	Fel      []Felmeddelande `json:"fel" validate:"min=1"`
	Status   int             `json:"status" validate:"required"`
	Tidpunkt string          `json:"tidpunkt" validate:"required"`
}

// SokRequest 5.1.2.4
type SokRequest struct {
	Bestallning Bestallning `json:"bestallning" validate:"required"`
	Sokvillkor  Sokvillkor  `json:"sokvillkor" validate:"required"`
}

// OkSokRequest 5.1.2.5
type OkSokResponse struct {
	Folkbokforingsposter []FolkbokforingsPostSok `json:"folkbokforingsposter" validate:"min=1"`
	Status               int                     `json:"status" validate:"required"`
	Tidpunkt             string                  `json:"tidpunkt" validate:"required"`
}

// Bestallning 9.1.1.1
type Bestallning struct {
	BestallningsIdentitet string `json:"bestallningsidentitet" validate:"required,max=18"`
	Organisationsnummer   string `json:"organisationsnummer" validate:"required, max=12"`
}

// Sokvillkorsidentitet 9.1.1.2
type Sokvillkorsidentitet struct {
	Identitetsbeteckning string `json:"identitetsbeteckning" validate:"required"`
}

// Felmeddelande 9.1.1.3
type Felmeddelande struct {
	ID                    string `json:"id" validate:"omitempty"`
	Orsakskod             int    `json:"orsakskod" validate:"required"`
	Orsakskodsbeskrivning string `json:"orsakskodsbeskrivning" validate:"required"`
}

// FolkbokforingspostHamta 9.1.1.4
type FolkbokforingspostHamta struct {
	Avregistrering                   Avregistrering                  `json:"avregistrering" validate:"omitempty"`
	Civiltilstand                    Civiltilstand                   `json:"civiltilstand" validate:"omitempty"`
	Fodelse                          Fodelse                         `json:"fodelse" validate:"omitempty"`
	Folkbokforing                    Folkbokforing                   `json:"folkbokforing" validate:"omitempty"`
	Hanvisningar                     []Hanvisningar                  `json:"hanvisningar" validate:"omitempty"`
	Identitet                        Identitet                       `json:"identitet" validate:"required"`
	Invandring                       Invandring                      `json:"invandring" validate:"omitempty"`
	Kontaktadress                    Kontaktadress                   `json:"kontaktadress" validate:"omitempty"`
	Medborgarskap                    []Medborgarskap                 `json:"medborgarskap" validate:"omitempty"`
	Namn                             Namn                            `json:"namn" validate:"omitempty"`
	RelationerTillAldrigFolkbokforda []RelationTillAldrigFolkbokford `json:"relationerTillAldrigFolkbokforda" validate:"omitempty"`
	RelationerTillFolkbokforda       []RelationTillFolkbokford       `json:"relationerTillFolkbokforda" validate:"omitempty"`
	Samordningsnummer                Samordningsnummeruppgifter      `json:"samordningsnummer" validate:"omitempty"`
	SenastAndrad                     SenastAndrad                    `json:"senastAndrad" validate:"required"`
	SkyddAvPersonuppgifter           string                          `json:"skyddAvPersonuppgifter" validate:"required,oneof=SAKNAS SEKRETESSMARKERING SKYDDAD_FOLKBOKFORING"`
}

// Avregistrering 9.1.1.5
type Avregistrering struct {
	Avregistreringsdatum  DatumAaaaMmDd `json:"avregistreringsdatum" validate:"required"`
	DatumAntraffadDod     DatumAaaaMmDd `json:"datumAntraffadDod" validate:"omitempty"`
	Orsakskod             string        `json:"orsakskod" validate:"required,oneof=AS AN AV FI GN GS OB TA UV"`
	Orsakskodsbeskrivning string        `json:"orsakskodsbeskrivning" validate:"required"`
}

// DatumAaaaMmDd 9.1.1.6
type DatumAaaaMmDd struct {
	RFC3339datum string `json:"rfc3339Datum" validate:"required,max=10"` // 2019-10-12
	Varde        string `json:"varde" validate:"required,max=8"`         // 20191012
}

// Civiltilstand 9.1.1.7
type Civiltilstand struct {
	Beskrivning string        `json:"beskrivning" validate:"omitempty"`
	Datum       DatumAaaaMmDd `json:"datum" validate:"omitempty"`
	Kod         string        `json:"kod" validate:"omitempty,oneof=OG G Ä S RP SP EP"`
}

// Fodelse 9.1.1.8
type Fodelse struct {
	Datum   DatumAaaaMmDd  `json:"datum" validate:"omitempty"`
	Sverige FodelseSverige `json:"sverige" validate:"omitempty"`
	Utland  FodelseUtland  `json:"utland" validate:"omitempty"`
}

// FodelseSverige 9.1.1.9
type FodelseSverige struct {
	Lanskod string `json:"lanskod" validate:"omitempty"`
	Ort     string `json:"ort" validate:"omitempty"`
}

// FodelseUtland 9.1.1.10
type FodelseUtland struct {
	Land string           `json:"land" validate:"omitempty"`
	Ort  FodelseUtlandOrt `json:"ort" validate:"omitempty"`
}

// FodelseUtlandOrt 9.1.1.11
type FodelseUtlandOrt struct {
	Namn   string `json:"namn" validate:"omitempty"`
	Strykt bool   `json:"strykt" validate:"omitempty"`
}

// Folkbokforing 9.1.1.12
type Folkbokforing struct {
	Adress         FolkbokforingAdress      `json:"adress" validate:"omitempty"`
	Datum          DatumAaaaMmDd            `json:"datum" validate:"omitempty"`
	Distriktskod   string                   `json:"distriktskod" validate:"omitempty"`
	Fastighet      FolkbokforingFastighet   `json:"fastighet" validate:"omitempty"`
	Forsamlingskod string                   `json:"forsamlingskod" validate:"omitempty"`
	Historik       []Folkbokforingshistorik `json:"historik" validate:"omitempty"`
	Kommunkod      string                   `json:"kommunkod" validate:"omitempty"`
	Lagenhet       FolkbokforingLagenhet    `json:"lagenhet" validate:"omitempty"`
	Lanskod        string                   `json:"lanskod" validate:"omitempty"`
}

// FolkbokforingAdress 9.1.1.13
type FolkbokforingAdress struct {
	Adressfortsattning string `json:"adressfortsattning" validate:"omitempty"`
	CareOf             string `json:"careOf" validate:"omitempty"`
	Gatuadress         string `json:"gatuadress" validate:"omitempty"`
	Nyckel             string `json:"nyckel" validate:"omitempty,uuid4"`
	Postnummer         string `json:"postnummer" validate:"required"`
	Postort            string `json:"postort" validate:"omitempty"`
}

// FolkbokforingFastighet 9.1.1.14
type FolkbokforingFastighet struct {
	Beteckning string `json:"beteckning" validate:"omitempty"`
	Nyckel     string `json:"nyckel" validate:"omitempty,uuid4"`
}

// Folkbokforingshistorik 9.1.1.15
type Folkbokforingshistorik struct {
	Datum             DatumAaaaMmDd                           `json:"datum" validate:"omitempty"`
	Fastighet         FolkbokforingFastighet                  `json:"fastighet" validate:"omitempty"`
	Folkbokforingstyp FolkbokforingshistorikFolkbokforingstyp `json:"folkbokforingstyp" validate:"omitempty"`
	Forsamlingskod    string                                  `json:"forsamlingskod" validate:"omitempty"`
	Kommun            string                                  `json:"kommun" validate:"omitempty"`
	Lanskod           string                                  `json:"lanskod" validate:"omitempty"`
}

// FolkbokforingLagenhet 9.1.1.16
type FolkbokforingLagenhet struct {
	Nyckel string `json:"nyckel" validate:"omitempty,uuid4"`
}

// FolkbokforingshistorikFolkbokforingstyp 9.1.1.17
type FolkbokforingshistorikFolkbokforingstyp struct {
	Beskrivning string `json:"beskrivning" validate:"omitempty"`
	Kod         string `json:"kod" validate:"omitempty,oneof=FB UV OB"`
}

// Hanvisningar 9.1.1.18
type Hanvisningar struct {
	Identitetsbeteckning string `json:"identitetsbeteckning" validate:"omitempty"`
}

// Identitet 9.1.1.19
type Identitet struct {
	Identitetsbeteckning string           `json:"identitetsbeteckning" validate:"required"`
	Status               Identitetsstatus `json:"status" validate:"omitempty"`
	Typ                  string           `json:"typ" validate:"required,oneof=PERSONNUMMER SAMORDNINGSNUMMER"`
}

// Identitetsstatus 9.1.1.20
type Identitetsstatus struct {
	Datum DatumAaaaMmDd `json:"datum" validate:"required"`
	Orsak string        `json:"orsak" validate:"omitempty"`
	Varde string        `json:"varde" validate:"required,oneof=AKTIVT VILANDEFÖRKLARAT VILANDEFÖRKLARAT STÄNGT AVREGISTRERAT"`
}

// Invandring 9.1.1.21
type Invandring struct {
	Datum               DatumAaaaMmDd              `json:"datum" validate:"omitempty"`
	NordiskaIdentiteter InvandringNordiskIdentitet `json:"nordiskaIdentiteter" validate:"omitempty"`
	Uppehallsratt       bool                       `json:"uppehallsratt" validate:"omitempty"`
}

// InvandringNordiskIdentitet 9.1.1.22
type InvandringNordiskIdentitet struct {
	Identitetsbeteckning string `json:"identitetsbeteckning" validate:"omitempty"`
	Land                 string `json:"land" validate:"omitempty"`
	Landskod             string `json:"landskod" validate:"omitempty"`
}

// Kontaktadress 9.1.1.23
type Kontaktadress struct {
	Sverige KontaktAdressSverige `json:"sverige" validate:"omitempty"`
	Utland  KontaktAdressUtland  `json:"utland" validate:"omitempty"`
}

// KontaktAdressSverige 9.1.1.24
type KontaktAdressSverige struct {
	Adressfortsattning string `json:"adressfortsattning" validate:"omitempty"`
	CareOf             string `json:"careOf" validate:"omitempty"`
	Gatuadress         string `json:"gatuadress" validate:"omitempty"`
	Postnummer         string `json:"postnummer" validate:"required"`
	Postort            string `json:"postort" validate:"omitempty"`
	Typ                string `json:"Typ" validate:"omitempty,oneof=FOLKBOKFORINGSADRESS SARSKILD_POSTADRESS"`
}

// KontaktAdressUtland 9.1.1.25
type KontaktAdressUtland struct {
	Adressrad1 string        `json:"adressrad1" validate:"omitempty"`
	Adressrad2 string        `json:"adressrad2" validate:"omitempty"`
	Adressrad3 string        `json:"adressrad3" validate:"omitempty"`
	FromDatum  DatumAaaaMmDd `json:"fromDatum" validate:"omitempty"`
	Land       string        `json:"land" validate:"omitempty"`
}

// Medborgarskap 9.1.1.26
type Medborgarskap struct {
	FromDatum DatumAaaaMmDd `json:"fromDatum" validate:"omitempty"`
	Landskod  string        `json:"landskod" validate:"omitempty"`
	Strykt    bool          `json:"strykt" validate:"omitempty"`
}

// Namn 9.1.1.27
type Namn struct {
	Aviseringsnamn string     `json:"aviseringsnamn" validate:"omitempty"`
	Efternamn      Efternamn  `json:"efternamn" validate:"omitempty"`
	Fornamn        Fornamn    `json:"fornamn" validate:"omitempty"`
	Mellannamn     Mellannamn `json:"mellannamn" validate:"omitempty"`
}

// Efternamn 9.1.1.28
type Efternamn struct {
	Namn   string `json:"namn" validate:"required"`
	Strykt bool   `json:"strykt" validate:"omitempty"`
}

// Fornamn 9.1.1.29
type Fornamn struct {
	Namn                   string `json:"namn" validate:"required"`
	Strykt                 bool   `json:"strykt" validate:"omitempty"`
	Tilltalsnamnsmarkering string `json:"tilltalsnamnsmarkering" validate:"required"`
}

// Mellannamn 9.1.1.30
type Mellannamn struct {
	Namn   string `json:"namn" validate:"required"`
	Strykt bool   `json:"strykt" validate:"omitempty"`
}

// RelationTillAldrigFolkbokford 9.1.1.31
type RelationTillAldrigFolkbokford struct {
	Fodelsetidsnummer string                            `json:"fodelsetidsnummer" validate:"required"`
	FromDatumVardnad  DatumAaaaMmDd                     `json:"fromDatumVardnad" validate:"omitempty"`
	Namn              RelationTillAldrigFolkbokfordNamn `json:"namn" validate:"omitempty"`
	Typ               RelationTillAldrigFolkbokfordTyp  `json:"typ" validate:"required"`
}

// RelationTillAldrigFolkbokfordNamn 9.1.1.32
type RelationTillAldrigFolkbokfordNamn struct {
	Efternamn  NamnRelationsperson `json:"efternamn" validate:"omitempty"`
	Fornamn    NamnRelationsperson `json:"fornamn" validate:"omitempty"`
	Mellannamn NamnRelationsperson `json:"mellannamn" validate:"omitempty"`
}

// NamnRelationsperson 9.1.1.33
type NamnRelationsperson struct {
	Namn string `json:"namn" validate:"required"`
}

// RelationTillAldrigFolkbokfordTyp 9.1.1.34
type RelationTillAldrigFolkbokfordTyp struct {
	Beskrivning string `json:"beskrivning" validate:"required"`
	Kod         string `json:"kod" validate:"required, oneof=B MO FA F V M P"`
}

// RelationTillFolkbokford 9.1.1.35
type RelationTillFolkbokford struct {
	Avregistrering       AvregistreringRelation     `json:"avregistrering" validate:"omitempty"`
	FromDatumVardnad     DatumAaaaMmDd              `json:"fromDatumVardnad" validate:"omitempty"`
	Identitetsbeteckning string                     `json:"identitetsbeteckning" validate:"required"`
	Typ                  RelationTillFolkbokfordTyp `json:"typ" validate:"required"`
}

// AvregistreringRelation 9.1.1.36
type AvregistreringRelation struct {
	Avregistreringsdatum  DatumAaaaMmDd `json:"avregistreringsdatum" validate:"omitempty"`
	Orsakskod             string        `json:"orsakskod" validate:"required, equal=AV"`
	Orsakskodsbeskrivning string        `json:"orsakskodsbeskrivning" validate:"required"`
}

// RelationTillFolkbokfordTyp 9.1.1.37
type RelationTillFolkbokfordTyp struct {
	Beskrivning string `json:"beskrivning" validate:"required"`
	Kod         string `json:"kod" validate:"required, oneof=B MO FA F V VF M P"`
}

// Samordningsnummeruppgifter 9.1.1.38
type Samordningsnummeruppgifter struct {
	DatumAvliden                       DatumAaaaMmDd `json:"datumAvliden" validate:"omitempty"`
	Fornyelsedatum                     DatumAaaaMmDd `json:"fornyelsedatum" validate:"omitempty"`
	PreliminärtVilandeforklaringsDatum DatumAaaaMmDd `json:"preliminärtVilandeforklaringsDatum" validate:"omitempty"`
	Tilldelningsdatum                  DatumAaaaMmDd `json:"tilldelningsdatum" validate:"omitempty"`
}

// SenastAndrad 9.1.1.39
type SenastAndrad struct {
	Tidpunkt string `json:"tidpunkt" validate:"required"`
}

// Sokvillkor 9.1.2.1
type Sokvillkor struct {
	Adress     SokvillkorAdress     `json:"adress" validate:"omitempty"`
	Fodelsetid SokvillkorFodelsetid `json:"fodelsetid" validate:"omitempty"`
	Kategori   string               `json:"kategori" validate:"omitempty, equal=FOLKBOKFÖRDA"`
	Kon        string               `json:"kon" validate:"omitempty, oneof=MAN KVINNA"`
	Namn       SokvillkorNamn       `json:"namn" validate:"omitempty"`
}

// SokvillkorAdress 9.1.2.2
type SokvillkorAdress struct {
	Gatuadress SokvillkorGatuadress `json:"gatuadress" validate:"omitempty"`
	Postnummer SokvillkorPostnummer `json:"postnummer" validate:"omitempty"`
	Postort    string               `json:"postort" validate:"omitempty"`
}

// SokvillkorGatuadress 9.1.2.3
type SokvillkorGatuadress struct {
	SammansattVarde string `json:"sammansattVarde" validate:"required"`
}

// SokvillkorPostnummer 9.1.2.4
type SokvillkorPostnummer struct {
	Fran int `json:"fran" validate:"required"`
	Till int `json:"till" validate:"required"`
}

// SokvillkorFodelsetid 9.1.2.5
type SokvillkorFodelsetid struct {
	Fran SokvillkorFodelsetidDatum `json:"fran" validate:"required"`
	Till SokvillkorFodelsetidDatum `json:"till" validate:"required"`
}

// SokvillkorNamn 9.1.2.6
type SokvillkorNamn struct {
	Fornamn              SokvillkorFornamn              `json:"fornamn" validate:"omitempty"`
	MellanEllerEfternamn SokvillkorMellanEllerEfternamn `json:"mellanEllerEfternamn" validate:"omitempty"`
}

// SokvillkorFornamn 9.1.2.7
type SokvillkorFornamn struct {
	Namndelar []string `json:"namndelar" validate:"omitempty"`
}

// SokvillkorMellanEllerEfternamn 9.1.2.8
type SokvillkorMellanEllerEfternamn struct {
	Namndelar []string `json:"namndelar" validate:"omitempty"`
}

// SokvillkorFodelsetidDatum 9.1.2.9
type SokvillkorFodelsetidDatum struct {
	Ar    string `json:"ar" validate:"required"`
	Manad string `json:"manad" validate:"omitempty"`
	Dag   string `json:"dag" validate:"omitempty"`
}

// FolkbokforingsPostSok 9.1.2.10
type FolkbokforingsPostSok struct {
	Avregistrering         AvregistreringsOrsak               `json:"avregistrering" validate:"omitempty"`
	Folkbokforing          FolkbokforingspostSokFolkbokforing `json:"folkbokforing" validate:"omitempty"`
	Identitet              IdentitetSokresultat               `json:"identitet" validate:"required"`
	Kontaktadress          FolkbokforingspostSokKontaktadress `json:"kontaktadress" validate:"omitempty"`
	Namn                   NamnSokresultat                    `json:"namn" validate:"omitempty"`
	SkyddAvPersonuppgifter string                             `json:"skyddAvPersonuppgifter" validate:"required,oneof=SAKNAS SEKRETESSMARKERING SKYDDAD_FOLKBOKFORING"`
}

// AvregistreringsOrsak 9.1.2.11
type AvregistreringsOrsak struct {
	Orsakskod             string `json:"orsakskod" validate:"required,oneof=AS AN AV FI GN GS OB TA UV"`
	Orsakskodsbeskrivning string `json:"orsakskodsbeskrivning" validate:"required"`
}

// FolkbokforingspostSokFolkbokforing 9.1.2.12
type FolkbokforingspostSokFolkbokforing struct {
	Address SvenskAdress `json:"adress" validate:"required"`
}

// SvenskAdress 9.1.2.13
type SvenskAdress struct {
	Adressfortsattning string `json:"adressfortsattning" validate:"omitempty,max=35"`
	CareOf             string `json:"careOf" validate:"omitempty,max=35"`
	Gatuadress         string `json:"gatuadress" validate:"omitempty,max=35"`
	Postnummer         string `json:"postnummer" validate:"required,max=5"`
	Postort            string `json:"postort" validate:"omitempty,max=27"`
}

// IdentitetSokresultat 9.1.2.14
type IdentitetSokresultat struct {
	Identitetsbeteckning string                      `json:"identitetsbeteckning" validate:"required"`
	Status               IdentitetsStatusSokresultat `json:"status" validate:"omitempty"`
	Typ                  string                      `json:"typ" validate:"required, oneof=PERSONNUMMER SAMORDNINGSNUMMER"`
}

// IdentitetsStatusSokresultat 9.1.2.15
type IdentitetsStatusSokresultat struct {
	Varde string `json:"varde" validate:"required, oneof=AKTIV VILANDEFÖRKLARAT VILANDEFÖRKLARAT STÄNGT AVREGISTRERAT"`
}

// FolkbokforingspostSokKontaktadress 9.1.2.16
type FolkbokforingspostSokKontaktadress struct {
	Svensk KontaktAdressSverige `json:"svensk" validate:"required"`
}

// FolkbokforingsPostSokKontaktadress 9.1.2.17 already exists

// NamnSokresultat 9.1.2.18
type NamnSokresultat struct {
	Fornamn    FornamnUtanStyrkt    `json:"fornamn" validate:"omitempty"`
	Mellannamn MellannamnUtanStyrkt `json:"mellannamn" validate:"omitempty"`
	Efternamn  EfternamnUtanStyrkt  `json:"efternamn" validate:"omitempty"`
}

// FornamnUtanStyrkt 9.1.2.19
type FornamnUtanStyrkt struct {
	Namn                   string `json:"namn" validate:"required,max=80"`
	TilltalsnamnsMarkering string `json:"tilltalsnamnsmarkering" validate:"omitempty,max=2"`
}

// MellannamnUtanStyrkt 9.1.2.20
type MellannamnUtanStyrkt struct {
	Namn string `json:"namn" validate:"required, max=40"`
}

// EfternamnUtanStyrkt 9.1.2.21
type EfternamnUtanStyrkt struct {
	Namn string `json:"namn" validate:"required,max=60"`
}
