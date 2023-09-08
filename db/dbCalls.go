package db

import (
	"log"
)

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

func GetAllCountryCodes() []string {
	countryCodes := baseOneColQuery("SELECT `country_code` FROM `sources` GROUP BY `country_code`")
	return countryCodes
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
