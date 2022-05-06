package memontago

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
	"os"
)

func fileExists(filePath string) (bool, error) {
	_, err := os.Stat(filePath)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}

	return false, nil
}
func parseJsonIntoLang(filePath, lang string) interface{} {
	context := getContent(filePath)
	switch lang {
	case "ch":
		var langch langCh
		err := json.Unmarshal(context, &langch)
		if err != nil {
			log.Fatalln(err)
		}
		return langch
	case "en":
		var langen langEn
		err := json.Unmarshal(context, &langen)
		if err != nil {
			log.Fatalln(err)
		}
		return langen
	}
	return nil
}
func getContent(filePath string) []byte {
	context, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatalln(err)
	}
	return context
}
