# Aufgabe 1: 
programmiere die Multiplikation zweier Zahlen
mit hilfe einer rekursiven Funktion und Addition

```go
func mult(a,b int) int
Hinweis: nutze aus, dass (y+1)*x = y*x+x, also
	mult(0,b)   = 0
	mult(a+1,b) = mult(a,b)+b
```

# Aufgabe 2: 
programmiere die Ackermann Funktion.
```
  A(m, n) =
  n+1                              if m = 0
  A(m-1, 1)                        if m > 0 and n = 0
  A(m-1, A(m, n-1))                if m > 0 and n > 0
  Bonusaufgabe: zähle die Anzahl der rekursiven Aufrufe
  Hinweis erweitere die Funktion um einen int-Pointer, der
  bei jedem Aufruf inkrementiert wird.
```

# Aufgabe 3:
Programmiere die Funktionen isEven und isOdd. Die Funktionen sollen nur die Abfrage `n == 0` haben und sich gegenseitig aufrufen.

# Aufgabe 4: Fibonacci II
Go kann mit *großen* Zahlen rechnen (module "math/big")
Schreibe die Funktion ```func fibBig(n uint64) *big.Int ```, die die n-te Fibonacci Zahl in einer Schleife berechnet. mit ```big.NewInt(int64(n))``` kannst du die Variablen mit *großen* Zahlen initialisieren und mit ```x.Add(x, y)``` addieren.

# Aufgabe 5: Fibonacci III
Schreibe die Funktion ```func fibList(n int) []int```. Diese soll die Fibonacci Zahlen in einer Liste zurückgeben. 
