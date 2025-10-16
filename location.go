package vnlocation

import (
	_ "embed"
	"encoding/json"
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

//go:embed data/vn-provinces.json
var vnProvincesData []byte

//go:embed data/vn-tree.json
var vnTreeData []byte

var (
	provinces     []Province
	provinceWards []ProvinceWard
)

func loadData() {
	if err := json.Unmarshal(vnProvincesData, &provinces); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(vnTreeData, &provinceWards); err != nil {
		panic(err)
	}
}

func init() {
	loadData()
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
