package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"

	"github.com/fatih/color"
)

const passHashed = "5e676818516b2521abd83fd6b8d038f3"

var red = color.New(color.FgRed).SprintFunc()
var green = color.New(color.FgGreen).SprintFunc()
var brown = color.New(color.FgCyan).SprintFunc()

func main() {

	// fmt.Println("Inserisci password: ")
	// silent, _ := gopass.GetPasswd()
	// h := md5.New()
	// io.WriteString(h, string(silent)+"\n")
	// test := fmt.Sprintf("%x", h.Sum(nil))
	// if test != passHashed {
	// 	fmt.Println("Passwod errata!")
	// 	os.Exit(1)
	// }

	type exam func() (string, error)

	var exams = []exam{
		ex00,
		ex01,
		ex02,
		ex03,
		ex04,
		ex05,
		ex06,
		ex07,
		ex08,
	}

	for _, e := range exams {
		check(e())
	}
}

func check(name string, err error) {

	if err == nil {
		fmt.Printf("%s %s\n", name, green("OK"))
	} else {
		fmt.Printf("%s %s %s\n", name, red("KO"), err.Error())
	}

}

func controllofile(numfiles int, file, dir string) error {
	fmt.Println()
	fmt.Println(brown("Controllo per ", dir))
	_, err := os.Stat(dir + "/" + file)
	if err != nil {
		return fmt.Errorf("file %s assente", file)
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
		return fmt.Errorf("troppi file nella directory %s: %s", dir, listfiles)
	}

	return nil
}

func stampaFile(file, dir string) error {
	catCmd := exec.Command("bash", "-c", "cat "+file)
	catCmd.Dir = dir

	output, err := catCmd.Output()
	if err != nil {
		return err
	}
	fmt.Println(string(output))

	inputReader := bufio.NewReader(strings.NewReader(string(output)))

	scanner := bufio.NewScanner(inputReader)
	// Count the words.
	var openCount = 0
	var closeCount = 0
	var linesOfCode = 0
	for scanner.Scan() {
		switch scanner.Text() {
		case "{":
			openCount++
		case "}":
			closeCount++
		}
		if closeCount != openCount {
			linesOfCode++
		}
		if closeCount == openCount {
			if linesOfCode > 0 {
				fmt.Println("Linee di codice: ", linesOfCode-1)
				if linesOfCode-1 > 25 {
					err = fmt.Errorf("troppe linee di codice: %d", linesOfCode-1)
				}
				linesOfCode = 0
				break
			}
		}
	}
	return err
}

func ex00() (string, error) {
	file := "ft_putchar.c"
	dir := "ex00"
	error := controllofile(1, file, dir)
	error = stampaFile(file, dir)

	return dir, error
}

func ex01() (string, error) {
	file := "ft_print_alphabet.c"
	dir := "ex01"
	error := controllofile(1, file, dir)
	error = stampaFile(file, dir)

	return dir, error
}

func ex02() (string, error) {
	file := "ft_print_reverse_alphabet.c"
	dir := "ex02"
	error := controllofile(1, file, dir)
	error = stampaFile(file, dir)

	return dir, error
}

func ex03() (string, error) {
	file := "ft_print_numbers.c"
	dir := "ex03"
	error := controllofile(1, file, dir)
	error = stampaFile(file, dir)

	return dir, error
}

func ex04() (string, error) {
	file := "ft_is_negative.c"
	dir := "ex04"
	error := controllofile(1, file, dir)
	error = stampaFile(file, dir)

	return dir, error
}

func ex05() (string, error) {
	file := "ft_print_comb.c"
	dir := "ex05"
	error := controllofile(1, file, dir)
	error = stampaFile(file, dir)

	return dir, error
}

func ex06() (string, error) {
	file := "ft_print_comb2.c"
	dir := "ex06"
	error := controllofile(1, file, dir)
	error = stampaFile(file, dir)

	return dir, error
}

func ex07() (string, error) {
	file := "ft_putnbr.c"
	dir := "ex07"
	error := controllofile(1, file, dir)
	error = stampaFile(file, dir)

	return dir, error
}

func ex08() (string, error) {
	file := "ft_print_combn.c"
	dir := "ex08"
	error := controllofile(1, file, dir)
	error = stampaFile(file, dir)

	return dir, error
}
