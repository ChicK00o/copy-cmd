package userFunctions

import (
	"encoding/json"
	"fmt"
	"github.com/otiai10/copy"
	"io/ioutil"
	"regexp"
)

func ExecuteCopy(source string, destination string, config string) error {

	data, err := ioutil.ReadFile(config)
	if err != nil {
		fmt.Print(err)
		return err
	}

	// define data structure
	type Data struct {
		Ignore []string `json:"ignore"`
	}

	// json data
	var obj Data

	// unmarshall it
	err = json.Unmarshal(data, &obj)
	if err != nil {
		fmt.Println("error:", err)
		return err
	}

	var regs = []*regexp.Regexp{}
	for _, data := range obj.Ignore {
		re, err := regexp.Compile(data)
		if err != nil {
			fmt.Println("error:", err)
			return err
		}
		regs = append(regs, re)
	}

	err = copy.Copy(source, destination, copy.Options{
		OnSymlink: func(_ string) copy.SymlinkAction {
			return copy.Shallow
		},
		Skip: func(src string) (skip bool, err error) {
			for _, data := range regs {
				if data.MatchString(src) {
					return true, nil
				}
			}
			return false, nil
		},
		Sync: true,
	})
	return err
}
