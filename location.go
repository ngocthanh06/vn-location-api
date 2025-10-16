package vnlocation

import (
	_ "embed"
	"encoding/json"
	"log"
	"strconv"
	"strings"
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

//go:embed data/vn-wards.json
var vnWardsData []byte

var (
	provinces     []Province
	wards         []Ward
	provinceWards []ProvinceWard
)

func loadData() {
	if err := json.Unmarshal(vnProvincesData, &provinces); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(vnTreeData, &provinceWards); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(vnWardsData, &wards); err != nil {
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
func GetWardsByProvinceCode(provinceCodeStr *string) []Ward {
	if provinceCodeStr == nil || strings.TrimSpace(*provinceCodeStr) == "" {
		return wards
	}

	provinceCode, err := strconv.Atoi(*provinceCodeStr)
	if err != nil {
		log.Printf("Invalid province code: %v", err)
		return nil
	}

	var result []Ward
	for _, ward := range provinceWards {
		code, convErr := strconv.Atoi(ward.Province.Code)
		if convErr != nil {
			log.Printf("Invalid province code in data: %v", err)
			continue
		}

		if code == provinceCode {
			result = append(result, ward.Wards...)
		}
	}

	return result
}
