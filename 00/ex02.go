package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func ex02() (bool, error) {

	type r struct {
		permessi string
		inode    string
		kb       string
		mese     string
		giorno   string
		orario   string
		nome     string
	}

	dir := "ex02"
	// rimuove file untarred quando ha finito.
	defer os.Remove(dir + "/" + "test0")
	defer os.Remove(dir + "/" + "test1")
	defer os.Remove(dir + "/" + "test2")
	defer os.Remove(dir + "/" + "test3")
	defer os.Remove(dir + "/" + "test4")
	defer os.Remove(dir + "/" + "test5")
	defer os.Remove(dir + "/" + "test6")

	file := "exo2.tar"
	_, err := controllofile(1, file, dir)
	if err != nil {
		return false, err
	}

	catCmd := exec.Command("bash", "-c", "tar -xvf exo2.tar")
	catCmd.Dir = dir

	_, err = catCmd.Output()
	if err != nil {
		return false, err
	}

	// nasconde momentaneamente il file tar
	nascondiTar := exec.Command("bash", "-c", "mv exo2.tar .exo2.tar")
	nascondiTar.Dir = dir
	_, err = nascondiTar.Output()
	if err != nil {
		return false, err
	}

	lsCMD := exec.Command("bash", "-c", "ls -l")
	lsCMD.Dir = dir
	lsOut, err := lsCMD.Output()
	output := string(lsOut)

	scanner := bufio.NewScanner(strings.NewReader(output))
	scanner.Split(bufio.ScanLines)

	var risultato []r
	for scanner.Scan() {
		campi := strings.Fields(scanner.Text())
		if len(campi) >= 7 {
			var c = r{
				permessi: campi[0],
				inode:    campi[1],
				kb:       campi[4],
				mese:     campi[5],
				giorno:   campi[6],
				orario:   campi[7],
				nome:     campi[8],
			}
			risultato = append(risultato, c)
		}
	}

	outputatteso := `
drwx--xr-x 2 XX XX XX Jun 1 20:47 test0
-rwx--xr-- 1 XX XX 4 Jun 1 21:46 test1
dr-x---r-- 2 XX XX XX Jun 1 22:45 test2
-r-----r-- 2 XX XX 1 Jun 1 23:44 test3
-rw-r----x 1 XX XX 2 Jun 1 23:43 test4
-r-----r-- 2 XX XX 1 Jun 1 23:44 test5
lrwxr-xr-x 1 XX XX 5 Jun 1 22:20 test6 -> test0`

	scanner = bufio.NewScanner(strings.NewReader(outputatteso))
	scanner.Split(bufio.ScanLines)

	var risultatoAtteso []r
	for scanner.Scan() {
		campi := strings.Fields(scanner.Text())
		if len(campi) >= 7 {
			var c = r{
				permessi: campi[0],
				inode:    campi[1],
				kb:       campi[4],
				mese:     campi[5],
				giorno:   campi[6],
				orario:   campi[7],
				nome:     campi[8],
			}
			risultatoAtteso = append(risultatoAtteso, c)
		}

	}

	for i := 0; i < len(risultatoAtteso); i++ {
		if i == 0 || i == 2 {
			risultato[i].kb = "XX"
		}
		switch {
		case risultatoAtteso[i].nome != risultato[i].nome:
			fmt.Printf("nome non corretto in riga %d %s %s\n", i, risultatoAtteso[i].nome, risultato[i].nome)
		case strings.TrimSpace(risultatoAtteso[i].kb) != strings.TrimSpace(risultato[i].kb):
			fmt.Printf("kb non corretti in riga %d %s %s\n", i, risultatoAtteso[i].kb, risultato[i].kb)

		case risultatoAtteso[i].inode != risultato[i].inode:
			fmt.Println("inode non corretti")
		case risultatoAtteso[i].mese != risultato[i].mese:
			fmt.Printf("mese non corretto in riga %d %s %s\n", i, risultatoAtteso[i].mese, risultato[i].mese)
		case risultatoAtteso[i].orario != risultato[i].orario:
			fmt.Println("orario non corretto")
		case risultatoAtteso[i].giorno != risultato[i].giorno:
			fmt.Println("giorno non corretto")

		case strings.TrimSpace(risultatoAtteso[i].permessi) != strings.TrimSpace(risultato[i].permessi):
			fmt.Printf("permessi non corretti in riga %d %s %s\n", i, risultatoAtteso[i].permessi, risultato[i].permessi)

		default:
		}

	}

	// mostra nuovamente il file tar
	mostraTar := exec.Command("bash", "-c", "mv .exo2.tar exo2.tar")
	mostraTar.Dir = dir
	_, err = mostraTar.Output()
	if err != nil {
		return false, err
	}

	return true, nil
}
