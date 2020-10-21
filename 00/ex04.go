package main

func ex04() (bool, error) {

	dir := "ex04"
	file := "midLS"

	_, err := controllofile(1, file, dir)
	if err != nil {
		return false, err
	}

	return true, nil
}
