package main

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
	"sort"
	"time"
)

type keyvalue struct {
	key   string
	value int
}

func sortMapByValue(items *map[string]int) (keyValues []keyvalue) {
	keyValues = make([]keyvalue, 0, len(*items))

	for k, v := range *items {
		keyValues = append(keyValues, keyvalue{k, v})
	}

	sort.Slice(keyValues, func(i, j int) bool {
		return keyValues[i].value > keyValues[j].value
	})

	return
}

func httpGetAndReadAll(url string) (xmldata []byte) {
	var (
		retries = 3
		resp    *http.Response
		err     error
	)

	for retries > 0 {
		resp, err = http.Get(url)
		if err != nil || resp.StatusCode == 202 {
			retries--
			time.Sleep(400 * time.Millisecond)
		} else {
			break
		}
	}

	if resp != nil {
		defer resp.Body.Close()

		xmldata, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalf("ERROR %s", err)
		}
	} else {
		log.Fatalf("ERROR %s", err)
	}

	return
}

func unmarshalOrDie(xmldata []byte, object interface{}) {
	err := xml.Unmarshal(xmldata, object)
	if err != nil {
		log.Fatalf("ERROR %s", err)
	}
	return
}
