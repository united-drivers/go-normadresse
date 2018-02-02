package normadresse

import (
	"fmt"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func ExampleNormalize() {
	fmt.Println(strings.ToUpper(Normalize("BOULEVARD DU MARECHAL JEAN MARIE DE LATTRE DE TASSIGNY")))
	// Output: BD MAL J M DE LATTRE DE TASSIGNY
}

func ExampleNormalizeLength_100() {
	fmt.Println(strings.ToUpper(NormalizeLength("BOULEVARD DU MARECHAL JEAN MARIE DE LATTRE DE TASSIGNY", 100)))
	// Output: BOULEVARD DU MARECHAL JEAN MARIE DE LATTRE DE TASSIGNY
}

func ExampleNormalizeLength_32() {
	fmt.Println(strings.ToUpper(NormalizeLength("BOULEVARD DU MARECHAL JEAN MARIE DE LATTRE DE TASSIGNY", 32)))
	// Output: BD MAL J M DE LATTRE DE TASSIGNY
}

func ExampleNormalizeLength_10() {
	fmt.Println(strings.ToUpper(NormalizeLength("BOULEVARD DU MARECHAL JEAN MARIE DE LATTRE DE TASSIGNY", 10)))
	// Output: BD MAL J M LATTRE TASSIGNY
}

func Test(t *testing.T) {
	Convey("Test package", t, FailureContinues, func() {
		tests := [][]string{
			{"BOULEVARD DU MARECHAL JEAN MARIE DE LATTRE DE TASSIGNY", "bd mal j m de LATTRE de TASSIGNY", "bd mal j m LATTRE TASSIGNY"},
			{"sentier de la Côte", "SENTIER DE LA COTE"},
			//{"SQUARE DES SOEURS DE SAINT VINCENT DE PAUL", "sq SOEURS DE st VINCENT DE PAUL"},
			//{"RUE LINO BORRINI DIT LINO VENTURA", "RUE l BORRINI DIT LINO VENTURA"},
			//{"RUE ETIENNE GEOFFROY DE SAINT HILAIRE", "RUE e GEOFFROY DE SAINT HILAIRE"},
			//{"RUE DU PRINCE FERDINAND DE BOURBON DES DEUX SICILES", "RUE P f de BOURBON DEUX SICILES"},
			//{"RUE DU MARECHAL PHILIPPE DE HAUTECLOCQUE DIT LECLERC", "RUE mal p de H DIT LECLERC"},
			//{"RUE DU LIEUTENANT DE VAISSEAU JEAN DOMINIQUE ANDRIEUX", "RUE DU ltdv j DOMINIQUE ANDRIEUX"},
			//{"RUE DU GROUPE SCOLAIRE LOUIS BRETON", "RUE DU grou SCOLAIRE l BRETON"},
			//{"RUE DU CORPS EXPEDITIONNAIRE FRANCAIS EN ITALIE 1943 1944", "RUE C E fr ITALIE 1943 1944"},
			{"RUE DES FRERES MARC ET JEAN MARIE GAMON", "RUE DES FRERES m ET j m GAMON"},
			//{"TRAVERSE NOTRE DAME DE BON SECOURS", "TRAVERSE nd DE BON SECOURS"},
			{"AVENUE GEORGES ET CLAUDE CAUSTIER", "av GEORGES ET CLAUDE CAUSTIER"},
			{"AVENUE DE LA DIVISION DU GENERAL LECLERC", "av DE LA DIVISION DU gal LECLERC"},
			//{"ALLÉE DE L'ABBAYE NOTRE DAME DU GRAND MARCHÉ", "all DE L ABBAYE nd DU gd MARCHE"},
			//{"ALLEE 1ER BATAILLON DU RGT DE BIGORRE FFI 1944 1945", "all 1ER btn R B FFI 1944 1945"},
			//{"AGGLOMERATION DE VILLENEUVE D ENTRAUNES", "aggl DE VILLENEUVE D ENTRAUNES"},
			//{"SENTIER DE LA RAVINE DES PONCEAUX", "sent DE LA RAVINE DES PONCEAUX"},
			//{"SENTIER DE LA FONTAINE DU PETIT DAMIETTE", "sent DE LA font DU pt DAMIETTE"},
			{"RUE EMMANUEL D ASTIER DE LA VIGERIE", "RUE e D ASTIER DE LA VIGERIE"},
			//{"PLACE CHARLES DE GAULLE ET DE LA RESISTANCE", "pl c de GAULLE DE LA RESISTANCE"},
			//{"ALLEE DES VILLAS DE LA GRANDE BASTIDE", "all DES VILLAS DE LA gde BASTIDE"},
			{"RUE L APPEL DU 18 JUIN DU GENERAL DE GAULLE", "RUE APPEL 18 JUIN gal de GAULLE"},
			//{"ANCIEN CHEMIN DE GRANGE DE PUZAIS AU SCEY", "anc chem GRANGE PUZAIS AU SCEY"},
			//{"BOURG DE SAINT MICHEL L ECLUSE ET LEPARON", "bour st MICHEL L eclu ET LEPARON"},
			//{"CARREFOUR DU GENERAL JACQUES PARIS DE BOLLARDIERE", "carr gal j PARIS DE BOLLARDIERE"},
			//{"CHEMIN DE L EGLISE DE SAINT CHRISTOPHE", "chem DE L egli DE st CHRISTOPHE"},
			{"CARREFOUR JEAN DE LATTRE DE TASSIGNY", "carr JEAN DE LATTRE DE TASSIGNY"},
			//{"CHEMIN DE LA FERME DES BOIS AU CHEMIN DES MOULINS", "chem ferm BOIS chem DES MOULINS"},
			//{"CHEMIN DE LA GRANDE RUE AU CHEMIN DES ROBERDES", "chem gde RUE chem DES ROBERDES"},
			//{"CHEMIN DE TRAVERSE DE LA FORET DE SENART", "chem DE trav DE LA for DE SENART"},
			//{"CHEMIN DE VILLENEUVE SAINT GEORGES A LA GRANGE", "chem de V st GEORGES GRANGE"},
			//{"CHEMIN DEPARTEMENTAL 116 D ARPAJON A AUNEAU", "chem dep 116 D ARPAJON A AUNEAU"},
			//{"CHEMIN DES DROITS DE L HOMME ET DE L ENFANT", "chem DROITS HOMME ET DE L ENFANT"},
			//{"CENTRE ADMINISTRATIF WALDECK L HUILIER", "ctre A WALDECK l HUILIER"},
			{"BOULEVARD SAINT JEAN BAPTISTE DE LA SALLE", "bd st JEAN BAPTISTE DE LA SALLE"},
			//{"BOURG DE SAINT ETIENNE DES LANDES", "BOURG DE st ETIENNE DES LANDES"},
			//{"BOURG SAINT JULIEN DES EGLANTIERS", "BOURG st JULIEN DES EGLANTIERS"},
			//{"CARREFOUR DE L EUROPE PRIX NOBEL DE LA PAIX 2012", "carr EUROPE PRIX NOBEL PAIX 2012"},
			//{"MONTEE DE L EGLISE SAINT PIERRE SAINT PAUL", "mont DE L egli st pierre st PAUL"},
			//{"VOIE DE LA DECLARATION UNIVERSELLE DES DROITS DE L HOMME", "VOIE D universelle droits HOMME"},
			//#CARREFOUR DES COMBATTANTS DE LA FRANCE LIBRE:
			//#RUE DE LA MADELEINE A L ETANG PROLONGEE:
			//#RUE DES SOUS LIEUTENANTS RUIBET ET GATTINEAU:
			//#AVENUE DU 1ER REGIMENT DE MARCHE DES SPAHIS MAROCAINS:
		}
		for _, test := range tests {
			long, short := test[0], test[1]
			So(Normalize(long), ShouldEqual, short)
		}
	})
}
