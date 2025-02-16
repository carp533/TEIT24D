package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Partei struct {
	name    string
	stimmen int
	sitze   int
}

type Koalition struct {
	Parteien []Partei
	Summe    int
}

func main() {

	gesamtSitze := 630

	// Datei lesen
	parteien, err := wahldatenLesen("wahldaten.txt")
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		return
	}

	// Gesamtzahl Stimmen
	gesamtStimmen := 0
	for _, p := range parteien {
		gesamtStimmen += p.stimmen
	}

	// 5%-Klasuel
	grenze5Proz := int(float64(gesamtStimmen) * 0.05)
	parteienGr5Proz := []Partei{}
	fmt.Println("=================================")
	for _, p := range parteien {
		fmt.Printf("%15s, %15d,\n", p.name, p.stimmen)
		if p.stimmen >= grenze5Proz {
			parteienGr5Proz = append(parteienGr5Proz, p)
		}
	}
	fmt.Println("=================================")
	fmt.Printf("Gesamtzahl Simmen:%15d\n", gesamtStimmen)
	fmt.Printf("5%%:              %16d\n", grenze5Proz)
	sitzVerteilung(parteienGr5Proz, gesamtSitze)

	// Print results
	fmt.Println("\nSitzverteilung:")
	fmt.Println("=================================")
	for _, p := range parteienGr5Proz {
		fmt.Printf("%15s: %15d Stimmen, %3d Sitze\n", p.name, p.stimmen, p.sitze)
	}

	koalitionen := KalkuliereKoalitionen(parteienGr5Proz, gesamtSitze/2+1)
	// sucheKoalitionen(parteienGr5Proz, gesamtSitze/2)
	fmt.Printf("\n%d Mögliche Koalitionen:\n", len(koalitionen))
	fmt.Println("=================================")
	for _, k := range koalitionen {
		fmt.Println(k)
	}
}

func wahldatenLesen(filename string) ([]Partei, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var parteien []Partei
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, ",")
		if len(parts) != 2 {
			continue
		}

		votes, err := strconv.Atoi(strings.TrimSpace(parts[1]))
		if err != nil {
			continue
		}

		parteien = append(parteien, Partei{
			name:    strings.TrimSpace(parts[0]),
			stimmen: votes,
			sitze:   0,
		})
	}

	return parteien, scanner.Err()
}

func (k Koalition) String() string {
	sort.Slice(k.Parteien, func(i, j int) bool {
		return k.Parteien[i].name < k.Parteien[j].name
	})
	return fmt.Sprintf("%d Sitze, %v", k.Summe, k.Parteien)
}

func (p Partei) String() string {
	return fmt.Sprintf("%v", p.name)
}

func sitzVerteilung(parteien []Partei, gesamtSitze int) {
	type quotient struct {
		parteiIndex int
		wert        float64
	}

	var quotients []quotient
	for i, p := range parteien {
		for seat := 1; seat <= gesamtSitze; seat++ {
			quotients = append(quotients, quotient{
				parteiIndex: i,
				// wert:        float64(p.stimmen) / float64(2*seat-1),
				wert: float64(p.stimmen) / (float64(seat) - float64(0.5)),
			})
		}
	}

	sort.Slice(quotients, func(i, j int) bool {
		return quotients[i].wert > quotients[j].wert
	})

	for i := 0; i < gesamtSitze; i++ {
		parteien[quotients[i].parteiIndex].sitze++
	}
	sort.Slice(parteien, func(i, j int) bool {
		return parteien[i].sitze > parteien[j].sitze
	})
}

func KalkuliereKoalitionen(parteien []Partei, minSitze int) []Koalition {
	var reslist [][]Partei
	var helper func(index int, currentCombination []Partei, sitzeakt int, sitzeuebr int)

	helper = func(index int, kombAkt []Partei, sitzeakt int, sitzeuebr int) {
		if sitzeakt >= minSitze {
			reslist = append(reslist, append([]Partei{}, kombAkt...))
			return
		}

		if sitzeakt+sitzeuebr < minSitze {
			return
		}

		for i := index; i < len(parteien); i++ {
			helper(i+1, append(kombAkt, parteien[i]), sitzeakt+parteien[i].sitze, sitzeuebr-parteien[i].sitze)
		}
	}

	gesamtSitze := 0
	for _, p := range parteien {
		gesamtSitze += p.sitze
	}

	helper(0, []Partei{}, 0, gesamtSitze)
	ergebnis := make([]Koalition, 0)
	for _, v := range reslist {
		sum := 0
		for _, p := range v {
			sum += p.sitze
		}
		ergebnis = append(ergebnis, Koalition{Parteien: v, Summe: sum})
	}
	// Sortiere Ergebnisse nach Sitzen
	sort.Slice(ergebnis, func(i, j int) bool {
		return ergebnis[i].Summe < ergebnis[j].Summe
	})
	return ergebnis
}
