package main

import (
	"encoding/json"
	flag "github.com/spf13/pflag"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func flatten(acc []map[string]interface{}, curr interface{}, key string, id, parentid int, idname, parentidname, parentrefname string) ([]map[string]interface{}, int) {
	var cid int
	switch t := curr.(type) {
	case map[string]interface{}:
		id++
		t[idname] = id
		t[parentidname] = parentid
		t[parentrefname] = key
		parentid = id

		acc = append(acc, t)

		for k, v := range t {
			if acc, cid = flatten(acc, v, k, id, parentid, idname, parentidname, parentrefname); cid != 0 {
				delete(t, k)
				id = cid
			}
		}

		return acc, id
	case []interface{}:
		for _, v := range t {
			if acc, cid = flatten(acc, v, key, id, parentid, idname, parentidname, parentrefname); cid != 0 {
				id = cid
			}
		}

		return acc, id

	default:
		return acc, 0
	}
}

func main() {
	fileIn := flag.StringP("file", "f", "", "input file")
	fileOut := flag.StringP("output", "o", "", "output file")
	firstId := flag.Int("id", 1, "begin id numbering with")
	nameId := flag.String("propid", "_id", "id property name")
	nameParentId := flag.String("parentid", "_parentid", "parent id property name")
	nameParentRef := flag.String("ref", "_parentref", "parent ref property name")

	flag.Parse()

	var data []byte
	var err error

	if *fileIn == "" {
		data, err = ioutil.ReadAll(os.Stdin)
		check(err)
	} else {
		data, err = ioutil.ReadFile(*fileIn)
		check(err)
	}

	var dat interface{}

	err = json.Unmarshal(data, &dat)
	check(err)

	flat, _ := flatten(make([]map[string]interface{}, 0), dat, "", *firstId, *firstId, *nameId, *nameParentId, *nameParentRef)

	var out []byte

	out, err = json.MarshalIndent(flat, "", "  ")
	check(err)

	if *fileOut == "" {
		_, err = os.Stdout.Write(out)
	} else {
		err = ioutil.WriteFile(*fileOut, out, 0644)
	}
	check(err)
}
