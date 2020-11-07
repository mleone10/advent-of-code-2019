package main

import (
	"os"
	"testing"
)

var smallCase = make(chemicals)
var mediumCase = make(chemicals)
var largeCase = make(chemicals)

func TestMain(m *testing.M) {
	smallCase["A"] = chemical{yield: 10, components: requirements{"ORE": 10}}
	smallCase["B"] = chemical{yield: 1, components: requirements{"ORE": 1}}
	smallCase["C"] = chemical{yield: 1, components: requirements{"A": 7, "B": 1}}
	smallCase["D"] = chemical{yield: 1, components: requirements{"A": 7, "C": 1}}
	smallCase["E"] = chemical{yield: 1, components: requirements{"A": 7, "D": 1}}
	smallCase["FUEL"] = chemical{yield: 1, components: requirements{"A": 7, "E": 1}}

	mediumCase["A"] = chemical{yield: 2, components: requirements{"ORE": 9}}
	mediumCase["B"] = chemical{yield: 3, components: requirements{"ORE": 8}}
	mediumCase["C"] = chemical{yield: 5, components: requirements{"ORE": 7}}
	mediumCase["AB"] = chemical{yield: 1, components: requirements{"A": 3, "B": 4}}
	mediumCase["BC"] = chemical{yield: 1, components: requirements{"B": 5, "C": 7}}
	mediumCase["CA"] = chemical{yield: 1, components: requirements{"A": 1, "C": 4}}
	mediumCase["FUEL"] = chemical{yield: 1, components: requirements{"AB": 2, "BC": 3, "CA": 4}}

	largeCase["NZVS"] = chemical{yield: 5, components: requirements{"ORE": 157}}
	largeCase["DCFZ"] = chemical{yield: 6, components: requirements{"ORE": 165}}
	largeCase["PSHF"] = chemical{yield: 7, components: requirements{"ORE": 179}}
	largeCase["HKGWZ"] = chemical{yield: 5, components: requirements{"ORE": 177}}
	largeCase["GPVTF"] = chemical{yield: 2, components: requirements{"ORE": 165}}
	largeCase["QDVJ"] = chemical{yield: 9, components: requirements{"HKGWZ": 12, "GPVTF": 1, "PSHF": 8}}
	largeCase["XJWVT"] = chemical{yield: 2, components: requirements{"DCFZ": 7, "PSHF": 7}}
	largeCase["KHKGT"] = chemical{yield: 8, components: requirements{"DCFZ": 3, "NZVS": 7, "HKGWZ": 5, "PSHF": 10}}
	largeCase["FUEL"] = chemical{yield: 1, components: requirements{"XJWVT": 44, "KHKGT": 5, "QDVJ": 1, "NZVS": 29, "GPVTF": 9, "HKGWZ": 48}}

	os.Exit(m.Run())
}

func TestSmallCaseOneFuel(t *testing.T) {
	want, got := 31, smallCase.simplify(fuel, 1, make(requirements))
	if want != got {
		t.Errorf("incorrect fuel calculation; wanted %d, got %d", want, got)
	}
}

func TestMediumCaseOneFuel(t *testing.T) {
	want, got := 165, mediumCase.simplify(fuel, 1, make(requirements))
	if want != got {
		t.Errorf("incorrect fuel calculation; wanted %d, got %d", want, got)
	}
}

func TestLargeCaseOneFuel(t *testing.T) {
	want, got := 13312, largeCase.simplify(fuel, 1, make(requirements))
	if want != got {
		t.Errorf("incorrect fuel calculation; wanted %d, got %d", want, got)
	}
}

func TestLargeCaseOneTrillionFuel(t *testing.T) {
	want, got := 82892753, largeCase.maxFuel(1000000000000)
	if want != got {
		t.Errorf("incorrect fuel calculation; wanted %d, got %d", want, got)
	}
}
