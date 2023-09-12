package ressources

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

var SourceTypes = map[string]int{
	"Journaux":                1,
	"Agences de presse":       2,
	"Gouvernement":            3,
	"Partis politiques":       4,
	"Banque nationale":        5,
	"Ambassades":              6,
	"Sécurité nationale":      7,
	"Think Tanks":             8,
	"Universités":             9,
	"Statistiques nationales": 10,
}

// Source types in english and snake_case
var SourceTypesEn = map[string]int{
	"journals":            1,
	"press_agencies":      2,
	"government":          3,
	"political_parties":   4,
	"national_banks":      5,
	"ambassies":           6,
	"national_security":   7,
	"think_tanks":         8,
	"universities":        9,
	"national_statistics": 10,
}

// Source types in english and snake_case, AND REVERSED (id to name)
var SourceTypesReversedEn = map[int]string{
	1:  "journals",
	2:  "press_agencies",
	3:  "government",
	4:  "political_parties",
	5:  "national_banks",
	6:  "ambassies",
	7:  "national_security",
	8:  "think_tanks",
	9:  "universities",
	10: "national_statistics",
}

// gets all the possible Countries and Country Codes from a JSON file, and uncapitalizes them.
func GetAllCountries(basePath string) map[string]string {
	type Country struct {
		Name string `json:"name"`
		Code string `json:"code"`
	}
	type Countries []Country

	var countries Countries

	jsonFile, err := os.Open(filepath.Join(basePath, "/ressources", "countries.json"))
	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	bytes, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(bytes, &countries)

	countryCodes := map[string]string{}
	for _, v := range countries {
		countryCodes[strings.ToLower(v.Name)] = strings.ToLower(v.Code)
	}

	return countryCodes
}
