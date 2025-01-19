# Rekursion

Rekursion bedeutet, dass eine Funktion sich selbst aufruft, bis eine bestimmte Abbruchbedingung erreicht ist. Die Aufrufkette der Funktionen nennt man **Call-Stack**.

## Beispiele:
```
Fakultät:
n! = {
    1                   für n = 0,1
    n * (n-1)!          für n > 1  
}
7! = 7 * 6 * 5 * 4 * 3 * 2 * 1

Fibonacci Folge:
fₙ = {  
     0                   für n = 0  
     1                   für n = 1  
     fₙ₋₁ + fₙ₋₂          für n > 1  
}
die ersten 10 Folgenglieder: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55
```

```go
// Beispiel 1: Fakultät berechnen
func factorial(n uint) uint {
    // Abbruchbedingung
    if n <= 1 {
        return 1
    }
    // Rekursiver Aufruf
    return n * factorial(n-1)
}

// Beispiel 2: Fibonacci-Zahlen berechnen
func fibonacci(n uint) uint {
    // Abbruchbedingungen
    if n < 2 {
        print(1)
        return n
    }
    // Rekursive Aufrufe
    return fibonacci(n-1) + fibonacci(n-2)
}
```

Bei der Fakultät-Berechnung ruft sich die Funktion immer wieder selbst auf und multipliziert die Zahlen, bis sie bei 1 ankommt. Zum Beispiel:
```
factorial(4) -->: 4 * factorial(3)
             -->: 4 * 3 * factorial(2)
             -->: 4 * 3 * 2 * factorial(1)
             -->: 4 * 3 * 2 * 1 
Das Endergebnis ist dann: 4 * 3 * 2 * 1 = 24
```

### Aufgabe: 
schreibe dir Schritt für Schritt auf, wie fibonacci(4) ausgewertet wird. Wie oft wird fibonacci(1) berechnet?
