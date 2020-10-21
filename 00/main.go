package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

func main() {
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

func ex00() (bool, error) {
	dir := "ex00"
	file := "z"
	_, err := controllofile(1, file, dir)
	if err != nil {
		return false, err
	}
	catCmd := exec.Command("bash", "-c", "cat ex00/z")
	catOut, err := catCmd.Output()
	if err != nil {
		return false, err
	}
	if string(catOut) != "Z" {
		return false, errors.New("contenuto file z non corrisponde")
	}
	return true, nil
}

func ex01() (bool, error) {
	dir := "ex01"
	file := "testShell00.tar"
	_, err := controllofile(1, file, dir)
	if err != nil {
		return false, err
	}

	catCmd := exec.Command("bash", "-c", "tar -xvf testShell00.tar")
	catCmd.Dir = dir

	_, err = catCmd.Output()
	if err != nil {
		return false, err
	}
	// rimuove file untarred quando ha finito.
	defer os.Remove(dir + "/" + "testShell00")

	lsCMD := exec.Command("bash", "-c", "ls -l testShell00")
	lsCMD.Dir = dir
	lsOut, err := lsCMD.Output()
	output := string(lsOut)
	fields := strings.Split(output, " ")
	fields[2] = "XX"
	fields[3] = "XX"
	output = strings.Join(fields, " ")

	outputatteso := "-r--r-xr-x. 1 XX XX 40 Jun  1 23:42 testShell00"

	if !strings.Contains(output, outputatteso) {
		return false, fmt.Errorf("output non corretto\n%s%s", output, outputatteso)
	}

	return true, nil
}

func controllofile(numfiles int, file, dir string) (string, error) {
	_, err := os.Stat(dir + "/" + file)
	if err != nil {
		return "", fmt.Errorf("file %s assente", file)
	}

	files, err := ioutil.ReadDir(dir)
	if len(files) > numfiles {
		return "", fmt.Errorf("troppi file nella directory %s", dir)
	}

	return "ok", nil
}
