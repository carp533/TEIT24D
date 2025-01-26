# Bubblesort

Bubblesort ist ein einfacher Sortieralgorithmus, der durch wiederholtes Durchlaufen der Liste geht und benachbarte Elemente vertauscht, wenn sie in der falschen Reihenfolge sind. Dieser Vorgang wird solange wiederholt, bis die gesamte Liste sortiert ist.
Bubblesort hat eine quadratische Zeitkomplexität; in der Praxis werden oft bessere Sortierverfahren eingesetzt.

## Beispiel

Unsortierte Liste: `[5, 3, 8, 4, 2]`

1. **Erster Durchlauf**:  
   - Vergleiche `5` und `3` → Tausche → `[3, 5, 8, 4, 2]`  
   - Vergleiche `5` und `8` → Kein Tausch  
   - Vergleiche `8` und `4` → Tausche → `[3, 5, 4, 8, 2]`  
   - Vergleiche `8` und `2` → Tausche → `[3, 5, 4, 2, 8]`  

2. **Zweiter Durchlauf**:  
   - Vergleiche `3` und `4` → Kein Tausch  
   - Vergleiche `5` und `4` → Tausche → `[3, 4, 5, 2, 8]`  
   - Vergleiche `5` und `2` → Tausche → `[3, 4, 2, 5, 8]`  

3. **Dritter Durchlauf**:  
   - Vergleiche `3` und `4` → Kein Tausch  
   - Vergleiche `4` und `2` → Tausche → `[3, 2, 4, 5, 8]`  

4. **Vierter Durchlauf**:  
   - Vergleiche `3` und `2` → Tausche → `[2, 3, 4, 5, 8]`  

Sortierte Liste: `[2, 3, 4, 5, 8]`

## Eigenschaften:
- **Komplexität**:  
  - Best Case (schon sortiert): \(O(n)\)  
  - Worst Case (umgekehrt sortiert): \(O(n^2)\)  
- **Stabil**: Ja (Elemente mit gleichem Wert behalten ihre Reihenfolge)

# Binäre Suche
Die Binärsuche ist ein effizienter Suchalgorithmus, der auf einem sortierten
Array basiert.
Die Binärsuche hat eine logarithmische Laufzeitkomplexität von O(log n),
wodurch sie effizienter ist als lineare Suchalgorithmen, insbesondere für
große sortierte Datenmengen.

1. Initialisierung:
Die Suche beginnt mit einem sortierten Array. Es gibt zwei Indexe, low und high
(für Anfang und Ende des Bereichs).
Setze low auf den Index des ersten Elements und high auf den Index des letzten
Elements im Array.

2. Mittelpunkt bestimmen:
Berechne den Mittelpunkt des aktuellen Suchbereichs (mid = low + (high - low) / 2)

3. Vergleich:
Vergleiche das Element an der Position mid mit dem gesuchten Wert x.
Wenn a[mid] gleich x ist, wurde das Element gefunden, und die Suche ist erfolgreich.
	(return true)
Wenn a[mid] kleiner als x ist, liegt das gesuchte Element im rechten Teil des aktuellen
	Bereichs. Setze low = mid + 1 für den nächsten Durchlauf.
Wenn a[mid] größer als x ist, liegt das gesuchte Element im linken Teil des aktuellen
	Bereichs. Setze high = mid - 1 für den nächsten Durchlauf.

4. Wiederhole den Prozess:
Wiederhole die Schritte 2 und 3, bis low größer als high ist, in diesem Fall wurde das
Element nicht gefunden. (return false)