package main

import (
	"reflect"
	"strings"
	"testing"
)

func CountWords(text string) map[string]int {
	_ = strings.Fields(text) // Zerlegt den Text in Wörter

	/*
		Bemerkung: möchte man alles außer Buchstaben und Leerzeichen entfernen,
		kann man einen regulären Ausdruck verwenden:
		re := regexp.MustCompile(`[^\w\s]`)
		cleanedText := re.ReplaceAllString(text, "")
		fmt.Println(cleanedText)
	*/
	return nil
}

// ermittelt das häufigste Wort in einem String
// Bemerkung bei gleicher Häufigkeit reicht es, eines der häufigsten Wörter auszugeben
// Satzzeichen müssen nicht beachtet werden ("beliebt." <-> "beliebt")
func MostFrequentWord(text string) (string, int) {
	var mostFrequent string
	maxCount := 0
	return mostFrequent, maxCount
}

func TestMostFrequentWord(t *testing.T) {
	text := "Go ist eine einfache und schnelle Programmiersprache. Go ist beliebt. Go Go Go!"

	expectedF, expectedC := "Go", 4
	mostFrequent, maxCount := MostFrequentWord(text)

	// Prüfen, ob die Maps gleich sind
	if mostFrequent != expectedF {
		t.Errorf("Erwartet: %v, aber erhalten: %v", expectedF, mostFrequent)
	}
	if maxCount != expectedC {
		t.Errorf("Erwartet: %v, aber erhalten: %v", expectedC, maxCount)
	}
}

func TestCountWords(t *testing.T) {
	text := "Go ist eine einfache und schnelle Programmiersprache. Go ist beliebt."

	expected := map[string]int{
		"Go":                  2,
		"ist":                 2,
		"eine":                1,
		"einfache":            1,
		"und":                 1,
		"schnelle":            1,
		"Programmiersprache.": 1,
		"beliebt.":            1,
	}

	result := CountWords(text)

	// Prüfen, ob die Maps gleich sind
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("Erwartet: %v, aber erhalten: %v", expected, result)
	}
}
