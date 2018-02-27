package xyz_test

import (
	"testing"

	"github.com/chengxiaoer/geomGo"
	"github.com/chengxiaoer/geomGo/xyz"
)

func TestDistance(t *testing.T) {
	for i, tc := range []struct {
		p1, p2 geom.Coord
		result float64
	}{
		{
			p1:     geom.Coord{0.05790978948805059, 0.7487914339893872, 0.3097809883247721},
			p2:     geom.Coord{0.5544543275873585, 0.886648581134798, 0.8226072117199881},
			result: 0.727015685344633,
		},
		{
			p1:     geom.Coord{0.9058291726317237, 0.8027685679074159, 0.688622556298369},
			p2:     geom.Coord{0.46032889533195054, 0.0914326830074278, 0.8023165442624004},
			result: 0.8469920667395824,
		},
		{
			p1:     geom.Coord{0.9821931669640064, 0.15876616854845904, 0.8667924896764038},
			p2:     geom.Coord{0.4323369560803155, 0.8695237851937236, 0.5416149195711873},
			result: 0.9556456949969271,
		},
		{
			p1:     geom.Coord{0.35061245555917053, 0.8960408311058482, 0.7746746016337248},
			p2:     geom.Coord{0.32001410072017566, 0.10172334453953935, 0.7859415388836344},
			result: 0.7949864606764474,
		},
		{
			p1:     geom.Coord{0.9853903683176634, 0.09446182616512078, 0.6628108344597603},
			p2:     geom.Coord{0.05653802094599025, 0.2917024120888081, 0.1664559286406906},
			result: 1.0714656898305444,
		},
	} {
		distance := xyz.Distance(tc.p1, tc.p2)

		if distance != tc.result {
			t.Errorf("Test %v failed: expected %v but was %v", i+1, tc.result, distance)
		}
	}
}

func TestDistanceLineToLine(t *testing.T) {

	for i, tc := range []struct {
		line1Start, line1End, line2Start, line2End geom.Coord
		result                                     float64
	}{
		{
			line1Start: geom.Coord{0.7741081331418179, 0.21173250439922287, 0.38051542958688733},
			line1End:   geom.Coord{0.24786497671483976, 0.9619118750989877, 0.5580971994370003},
			line2Start: geom.Coord{0.15681267365775853, 0.880199192600833, 0.7922662558231603},
			line2End:   geom.Coord{0.7011220236566095, 0.03865047746667638, 0.8725545997644786},
			result:     0.2642018761133755,
		},
		{
			line1Start: geom.Coord{0.6866849797685883, 0.7768264682122077, 0.6906860665173048},
			line1End:   geom.Coord{0.3725562200019461, 0.3273245439357202, 0.3994607184268433},
			line2Start: geom.Coord{0.11184492626414366, 0.5616613808469744, 0.502654956297042},
			line2End:   geom.Coord{0.5615007976601206, 0.8098111847087466, 0.8597931511661152},
			result:     0.2120619972390977,
		},
		{
			line1Start: geom.Coord{0.14324831422763928, 0.1976764146480534, 0.6232929645076098},
			line1End:   geom.Coord{0.44953958873649036, 3.5239563737987645E-4, 0.7712169838831762},
			line2Start: geom.Coord{0.7545925980004722, 0.2637482401207386, 0.2724556780759071},
			line2End:   geom.Coord{0.25710142520446666, 0.8181769277392215, 0.6125714339070055},
			result:     0.5272955676770279,
		},
		{
			line1Start: geom.Coord{0.368839171686284, 0.19495025186182324, 0.48535590338977685},
			line1End:   geom.Coord{0.43949536746898277, 0.9141491603204804, 0.06329907441509253},
			line2Start: geom.Coord{0.6108422671987475, 0.7906113775550576, 0.04431489641722197},
			line2End:   geom.Coord{0.08734478071906804, 0.5638793723135712, 0.4573656361895869},
			result:     0.008723033474112478,
		},
		{
			line1Start: geom.Coord{0.021673226025721415, 0.733857154166252, 0.8547874752704443},
			line1End:   geom.Coord{0.9102489826857886, 0.2059521582591355, 0.569799533506188},
			line2Start: geom.Coord{0.04719383832534252, 0.579365817650374, 0.45780393239823747},
			line2End:   geom.Coord{0.8653051887590829, 0.3768753311397194, 0.2575471442865518},
			result:     0.34405671586126135,
		},
	} {
		distance := xyz.DistanceLineToLine(tc.line1Start, tc.line1End, tc.line2Start, tc.line2End)

		if distance != tc.result {
			t.Errorf("Test %v failed: expected %v but was %v", i+1, tc.result, distance)
		}
	}
}

func TestDistancePointToLine(t *testing.T) {

	for i, tc := range []struct {
		p1, p2, p3 geom.Coord
		result     float64
	}{
		{
			p1:     geom.Coord{0.6653423449125663, 0.25829302403948784, 0.09776737959637871},
			p2:     geom.Coord{0.9446232000628368, 0.8518235691174553, 0.7861793928492881},
			p3:     geom.Coord{0.08328344949056798, 0.11359339931864088, 0.7238095172075693},
			result: 0.7076174111284124,
		},
		{
			p1:     geom.Coord{0.3056340253059928, 0.548272835230905, 0.9174205725375176},
			p2:     geom.Coord{0.9473664418937926, 0.5138296811763381, 0.8850559878558496},
			p3:     geom.Coord{0.6875714575802029, 0.34243488501137487, 0.43883146208982726},
			result: 0.5845239962606498,
		},
		{
			p1:     geom.Coord{0.17925466242492893, 0.25326178211346384, 0.43900958861438044},
			p2:     geom.Coord{0.3462415885096438, 0.6058198433952049, 0.04313179002608547},
			p3:     geom.Coord{0.6299539377475025, 0.6777591017854, 0.8777643731216357},
			result: 0.47331436935257626,
		},
		{
			p1:     geom.Coord{0.6452433791475372, 0.7508024881952754, 0.9118212667989819},
			p2:     geom.Coord{0.043285204500404406, 0.825429320166729, 0.16788006864300675},
			p3:     geom.Coord{0.08903335403896706, 0.9435616168756346, 0.24718902961638278},
			result: 0.8878410242961485,
		},
		{
			p1:     geom.Coord{0.48829408889602244, 0.628882048212391, 0.9413675943438186},
			p2:     geom.Coord{0.910739892900317, 0.8188321183330397, 0.9792667071608953},
			p3:     geom.Coord{0.10973730289482764, 0.6344963586547071, 0.09552384295243632},
			result: 0.3198693050691922,
		},
	} {
		distance := xyz.DistancePointToLine(tc.p1, tc.p2, tc.p3)

		if distance != tc.result {
			t.Errorf("Test %v failed: expected %v but was %v", i+1, tc.result, distance)
		}
	}
}

func TestEquals(t *testing.T) {
	for i, tc := range []struct {
		p1, p2 geom.Coord
		result bool
	}{
		{
			p1:     geom.Coord{0.9674898015414836, 0.02624230336884803, 0.34208659688222187},
			p2:     geom.Coord{0.35932783559039183, 0.5111822990523828, 0.3927940423315276},
			result: false,
		},
		{
			p1:     geom.Coord{0.6538035324502031, 0.1197867740640145, 0.9713753698395488},
			p2:     geom.Coord{0.8526516675217474, 0.13873098842131826, 0.6071218166426663},
			result: false,
		},
		{
			p1:     geom.Coord{0.7171640483537938, 0.9237446591375224, 0.153341134286061},
			p2:     geom.Coord{0.10067104615631928, 0.9025079507239742, 0.2625867527336143},
			result: false,
		},
		{
			p1:     geom.Coord{0.21503646116881991, 0.3273197157358999, 0.07232487899719986},
			p2:     geom.Coord{0.5914194401153642, 0.02735318122495156, 0.3163177142935619},
			result: false,
		},
		{
			p1:     geom.Coord{0.43841576728845855, 0.7257060784407421, 0.9389273557450912},
			p2:     geom.Coord{0.23129358543810752, 0.38457770238327604, 0.3875535459930316},
			result: false,
		},
	} {
		isEqual := xyz.Equals(tc.p1, tc.p2)

		if isEqual != tc.result {
			t.Errorf("Test %v failed: expected %v but was %v", i+1, tc.result, isEqual)
		}
	}
}
