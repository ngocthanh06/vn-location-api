package vnlocation

import (
	"encoding/json"
	"os"
)

type Province struct {
	Code     string `json:"code"`
	Fullname string `json:"fullname"`
	Name     string `json:"name"`
	Slug     string `json:"slug"`
	Type     string `json:"type"`
}

type Ward struct {
	Code     string `json:"code"`
	Name     string `json:"name"`
	Fullname string `json:"fullname"`
	Slug     string `json:"slug"`
	Type     string `json:"type"`
}

type ProvinceWard struct {
	Province
	Wards []Ward
}

var provinces []Province

// var wards []Ward
var provinceWards []ProvinceWard

func loadData() {
	loadJson("data/vn-provinces.json", &provinces)
	//loadJson("data/vn-wards.json", &wards)
	loadJson("data/vn-tree.json", &provinceWards)
}

func init() {
	loadData()
}

// LoadJson loads a JSON file from the given path into the target variable.
func loadJson(path string, target any) {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(data, target); err != nil {
		panic(err)
	}
}

// GetProvinces returns a list of all provinces.
func GetProvinces() []Province {
	return provinces
}

// GetWardsByProvinceCode returns a list of wards for a given province code.
func GetWardsByProvinceCode(provinceCode string) []Ward {
	var wards []Ward

	for _, ward := range provinceWards {
		if ward.Province.Code == provinceCode {
			wards = append(wards, ward.Wards...)
		}
	}

	return wards
}
