package main

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strings"
)

type Province struct {
	ProvinceID   string `json:"provinceId"`
	ProvinceName string `json:"provinceName"`
}

type Regency struct {
	ProvinceID   string `json:"provinceId"`
	ProvinceName string `json:"provinceName"`
	RegencyID    string `json:"regencyId"`
	RegencyName  string `json:"regencyName"`
}

type District struct {
	ProvinceID   string `json:"provinceId"`
	ProvinceName string `json:"provinceName"`
	RegencyID    string `json:"regencyId"`
	RegencyName  string `json:"regencyName"`
	DistrictID   string `json:"districtId"`
	DistrictName string `json:"districtName"`
}

type Village struct {
	ProvinceID   string `json:"provinceId"`
	ProvinceName string `json:"provinceName"`
	RegencyID    string `json:"regencyId"`
	RegencyName  string `json:"regencyName"`
	DistrictID   string `json:"districtId"`
	DistrictName string `json:"districtName"`
	VillageID    string `json:"villageId"`
	VillageName  string `json:"villageName"`
}

func main() {
	f, err := os.Open("./data/base/wilayah.csv")
	// f, err := os.Open("./data/sample.csv")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	var (
		province Province
		regency  Regency
		district District
		village  Village
	)

	for _, row := range data {
		if len(row[0]) == 2 {
			province = Province{
				ProvinceID:   row[0],
				ProvinceName: row[1],
			}
		}
		if len(row[0]) == 5 {
			regency = Regency{
				ProvinceID:   province.ProvinceID,
				ProvinceName: province.ProvinceName,
				RegencyID:    strings.ReplaceAll(row[0], ".", ""),
				RegencyName:  row[1],
			}
		}
		if len(row[0]) == 8 {
			district = District{
				ProvinceID:   province.ProvinceID,
				ProvinceName: province.ProvinceName,
				RegencyID:    regency.RegencyID,
				RegencyName:  regency.RegencyName,
				DistrictID:   strings.ReplaceAll(row[0], ".", ""),
				DistrictName: row[1],
			}
		}

		if len(row[0]) == 13 {
			village = Village{
				ProvinceID:   province.ProvinceID,
				ProvinceName: province.ProvinceName,
				RegencyID:    regency.RegencyID,
				RegencyName:  regency.RegencyName,
				DistrictID:   district.DistrictID,
				DistrictName: district.DistrictName,
				VillageID:    strings.ReplaceAll(row[0], ".", ""),
				VillageName:  row[1],
			}

			jsonString, err := json.Marshal(village)
			if err != nil {
				panic(err)
			}

			os.WriteFile("./data/village/"+village.VillageID+".json", jsonString, os.ModePerm)
		}
	}

}
