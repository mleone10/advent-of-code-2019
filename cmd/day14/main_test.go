package main

import "testing"

func TestSmallCase(t *testing.T) {
	cs := make(chemicals)

	cs["A"] = chemical{yield: 10, components: requirements{"ORE": 10}}
	cs["B"] = chemical{yield: 1, components: requirements{"ORE": 1}}
	cs["C"] = chemical{yield: 1, components: requirements{"A": 7, "B": 1}}
	cs["D"] = chemical{yield: 1, components: requirements{"A": 7, "C": 1}}
	cs["E"] = chemical{yield: 1, components: requirements{"A": 7, "D": 1}}
	cs["FUEL"] = chemical{yield: 1, components: requirements{"A": 7, "E": 1}}

	want, got := 31, cs.simplify(fuel, 1, make(requirements))
	if want != got {
		t.Errorf("incorrect fuel calculation; wanted %d, got %d", want, got)
	}
}

func TestMediumCase(t *testing.T) {
	cs := make(chemicals)

	cs["A"] = chemical{yield: 2, components: requirements{"ORE": 9}}
	cs["B"] = chemical{yield: 3, components: requirements{"ORE": 8}}
	cs["C"] = chemical{yield: 5, components: requirements{"ORE": 7}}
	cs["AB"] = chemical{yield: 1, components: requirements{"A": 3, "B": 4}}
	cs["BC"] = chemical{yield: 1, components: requirements{"B": 5, "C": 7}}
	cs["CA"] = chemical{yield: 1, components: requirements{"A": 1, "C": 4}}
	cs["FUEL"] = chemical{yield: 1, components: requirements{"AB": 2, "BC": 3, "CA": 4}}

	want, got := 165, cs.simplify(fuel, 1, make(requirements))
	if want != got {
		t.Errorf("incorrect fuel calculation; wanted %d, got %d", want, got)
	}
}

func TestLargeCase(t *testing.T) {
	cs := make(chemicals)

	cs["NZVS"] = chemical{yield: 5, components: requirements{"ORE": 157}}
	cs["DCFZ"] = chemical{yield: 6, components: requirements{"ORE": 165}}
	cs["PSHF"] = chemical{yield: 7, components: requirements{"ORE": 179}}
	cs["HKGWZ"] = chemical{yield: 5, components: requirements{"ORE": 177}}
	cs["GPVTF"] = chemical{yield: 2, components: requirements{"ORE": 165}}
	cs["QDVJ"] = chemical{yield: 9, components: requirements{"HKGWZ": 12, "GPVTF": 1, "PSHF": 8}}
	cs["XJWVT"] = chemical{yield: 2, components: requirements{"DCFZ": 7, "PSHF": 7}}
	cs["KHKGT"] = chemical{yield: 8, components: requirements{"DCFZ": 3, "NZVS": 7, "HKGWZ": 5, "PSHF": 10}}
	cs["FUEL"] = chemical{yield: 1, components: requirements{"XJWVT": 44, "KHKGT": 5, "QDVJ": 1, "NZVS": 29, "GPVTF": 9, "HKGWZ": 48}}

	want, got := 13312, cs.simplify(fuel, 1, make(requirements))
	if want != got {
		t.Errorf("incorrect fuel calculation; wanted %d, got %d", want, got)
	}
}
