```go
func CommonMultiples(m, n, max int) []int {
	result := []int{}
	for i := 1; i <= max; i++ {
		if i%n == 0 && i%m == 0 {
			result = append(result, i)
		}
	}
	return result
}

func ArraySums(list []int) []int {
	result := []int{}
	sum := 0
	for _, val := range list {
		sum += val
		result = append(result, sum)
	}
	return result
}

func Power2(x int) float64 {
	if x == 0 {
		return 1
	}
	if x < 0 {
		x *= -1
		return 1 / (2 * Power2(x-1))
	}
	return 2 * Power2(x-1)
}

func FindAll(list []int, x int) []int {
	result := []int{}
	for idx, val := range list {
		if val == x {
			result = append(result, idx)
		}
	}
	return result
}

func (dict *Dictionary) SetEntry(de, en string) {
	for i, entry := range dict.Entries {
		if entry.De == de {
			dict.Entries[i] = Entry{de, en}
			return
		}
	}
	dict.Entries = append(dict.Entries, Entry{de, en})
}

// Datentyp für Einträge eines Wörterbuchs.
type Eintrag struct {
	De string
	En []string
}

// Liefert einen neuen Eintrag.
func NewEntry(de string, en []string) Eintrag {
	return Eintrag{De: de, En: en}
}

// Liefert eine String-Repräsentation eines Eintrags.
func (entry Eintrag) String() string {
	return fmt.Sprintf("%s : %v", entry.De, entry.Translations())
}

// Liefert einen String mit allen englischen Wörtern aus entry.
// Die einzelnen Wörter sollen mit Kommata getrennt sein.
func (entry Eintrag) Translations() []string {
	return entry.En
}

// Fügt eine neue Übersetzung zu entry hinzu.
func (entry *Eintrag) AddTranslation(newEn string) {
	entry.En = append(entry.En, newEn)
}

func ExampleEntry_Translations() {
	e1 := NewEntry("Himmel", []string{"sky"})
	fmt.Println(e1)

	e1.AddTranslation("heaven")
	fmt.Println(e1)

	// Output:
	// Himmel : sky
	// Himmel : sky,heaven
}

func Div(x, y int) int {
	if x < 0 {
		return -Div(-x, y)
	}
	if y < 0 {
		return -Div(x, -y)
	}
	if x < y {
		return 0
	}
	return 1 + Div(x-y, y)
}
```