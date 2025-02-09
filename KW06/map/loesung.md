func CountWords(text string) map[string]int {
	/*
		Bemerkung: möchte man alles außer Buchstaben und Leerzeichen entfernen,
		kann man einen regulären Ausdruck verwenden:
		re := regexp.MustCompile(`[^\w\s]`)
		cleanedText := re.ReplaceAllString(text, "")
		fmt.Println(cleanedText)
	*/
	wordCounts := make(map[string]int) // Map zur Speicherung der Wortanzahl
	words := strings.Fields(text)      // Zerlegt den Text in Wörter

	for _, word := range words {
		wordCounts[word]++ // Erhöht den Zähler für das Wort
	}

	return wordCounts
}

func MostFrequentWord(text string) (string, int) {
	var mostFrequent string
	maxCount := 0
	wordCounts := CountWords(text)
	for word, count := range wordCounts {
		if count > maxCount {
			mostFrequent = word
			maxCount = count
		}
	}

	return mostFrequent, maxCount
}