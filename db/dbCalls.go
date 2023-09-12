package db

import (
	"log"
)

// Fetches all websites, no filter
func GetAllWebSites() []string { 
	// Regex to catch some mistakes in results : /(?<!http|https):\/\//
	websites := baseOneColQuery("SELECT `website` FROM `sources`")
	return websites
}

// Fetches websites, filtered by country code and source type.
func GetWebSites(countryCode string, sourceTypeId int) []string {
	websites := baseOneColQuery("SELECT `website` FROM `sources` WHERE `country_code` = ? AND `source_type_id` = ?", countryCode, sourceTypeId)
	return websites
}

// Fetches websites, filtered only by country code.
func GetWebsitesByCountry(countryCode string) []string {
	websites := baseOneColQuery("SELECT `website` FROM `sources` WHERE `country_code` = ?", countryCode)
	return websites
}

// Gets the available CC, meaning the CC for the countries that have url sources in the database (aka only the useful CC).
func GetAvailCountryCodes() []string {
	countryCodes := baseOneColQuery("SELECT lower(`country_code`) FROM `sources` GROUP BY lower(`country_code`);")
	return countryCodes
}

// Gets the available source types IDs, meaning the ones used in the db.
func GetAvailSourceTypes() []string {
	sourceTypeIds := baseOneColQuery("SELECT `source_type_id` FROM `sources` GROUP BY `source_type_id`")
	return sourceTypeIds
}



// Base logic for a SELECT query to fetch one column.
func baseOneColQuery(query string, replacements ...any) []string {
	var record string
	var recordList = []string{}

	rows, err := Db.Query(query, replacements...)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&record)
		if err != nil {
			log.Fatal(err)
		}	
		recordList = append(recordList, record)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err.Error())
	}

	return recordList
}
