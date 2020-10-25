package main

import (
	"errors"
	"io/ioutil"
	"regexp"
)

func ex04() (bool, error) {

	dir := "ex04"
	file := "midLS"

	_, err := controllofile(1, file, dir)
	if err != nil {
		return false, err
	}
	body, err := ioutil.ReadFile(dir + "/" + file)
	if err != nil {
		return false, err
	}

	filecontent := string(body)

	var re = regexp.MustCompile(`(?m)^ls\s((-(p|m|tU){1,})\s){1,}`)

	if !re.MatchString(filecontent) {
		return false, errors.New("mancano flag")
	}

	return true, nil
}
