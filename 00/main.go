package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/fatih/color"
	"github.com/howeyc/gopass"
)

func main() {

	silent, _ := gopass.GetPasswd()

	if string(silent) != "Meriadoc" {
		os.Exit(1)
	}

	red := color.New(color.FgRed).SprintFunc()
	green := color.New(color.FgGreen).SprintFunc()
	ok, err := ex00()
	if ok {
		fmt.Printf("ex00 %s\n", green("OK"))
	} else {
		fmt.Printf("ex00 %s %s\n", red("KO"), err.Error())
	}

	ok, err = ex01()
	if ok {
		fmt.Printf("ex01 %s\n", green("OK"))
	} else {
		fmt.Printf("ex01 %s %s\n", red("KO"), err.Error())
	}

	ok, err = ex02()
	if ok {
		fmt.Printf("ex02 %s\n", green("OK"))
	} else {
		fmt.Printf("ex02 %s %s\n", red("KO"), err.Error())
	}

	ok, err = ex04()
	if ok {
		fmt.Printf("ex04 %s\n", green("OK"))
	} else {
		fmt.Printf("ex04 %s %s\n", red("KO"), err.Error())
	}

}

func controllofile(numfiles int, file, dir string) (string, error) {
	_, err := os.Stat(dir + "/" + file)
	if err != nil {
		return "", fmt.Errorf("file %s assente", file)
	}

	files, err := ioutil.ReadDir(dir)
	if len(files) > numfiles {
		var listfiles []string
		for _, f := range files {
			if f.Name() == file {
				continue
			}
			listfiles = append(listfiles, f.Name())
		}
		return "", fmt.Errorf("troppi file nella directory %s: %s", dir, listfiles)
	}

	return "ok", nil
}
