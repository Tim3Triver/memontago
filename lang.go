package memontago

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
)

var Lang Language

type Language struct {
	Second  string `json:"second" yaml:"second"`
	Seconds string `json:"seconds" yaml:"seconds"`
	Minute  string `json:"minute" yaml:"minute"`
	Minutes string `json:"minutes" yaml:"minutes"`
	Hour    string `json:"hour" yaml:"hour"`
	Hours   string `json:"hours" yaml:"hours"`
	Day     string `json:"day" yaml:"day"`
	Days    string `json:"days" yaml:"days"`
	Week    string `json:"week" yaml:"week"`
	Weeks   string `json:"weeks" yaml:"weeks"`
	Month   string `json:"month" yaml:"month"`
	Months  string `json:"months" yaml:"months"`
	Year    string `json:"year" yaml:"year"`
	Years   string `json:"years" yaml:"years"`
	JustNow string `json:"just_now" yaml:"just_now"`
	Online  string `json:"online" yaml:"online"`
	Ago     string `json:"ago" yaml:"ago"`
	Later   string `json:"later" yaml:"later"`
}

func UnMarshalToStruct(lang string) error {
	data, err := ioutil.ReadFile("./language/" + lang + ".yaml")
	if err != nil {
		return fmt.Errorf("momentago/language/%s/.yaml is not exised, please check", lang)
	}
	err = yaml.Unmarshal(data, &Lang)
	if err != nil {
		return fmt.Errorf("momentago/language/%s/.yaml is not yaml style, please check", lang)
	}
	return nil
}
