
# Die Türme von Hanoi

## Ausgangssituation
- Zu Beginn sind alle Scheiben auf dem ersten Stab (A, links) gestapelt
- Die Scheiben sind der Größe nach geordnet
- Ziel ist es, den kompletten Stapel auf den dritten Stab (C, rechts) zu bringen

## Regeln
1. Es darf immer nur eine Scheibe bewegt werden
2. Eine größere Scheibe darf niemals auf einer kleineren Scheibe liegen

### Startzustand
<svg viewBox="0 0 400 200" xmlns="http://www.w3.org/2000/svg">
    <!-- Basis -->
    <rect x="20" y="160" width="360" height="20" fill="#8B4513"/>
    <!-- Stäbe -->
    <rect x="70" y="40" width="10" height="120" fill="#8B4513"/>
    <rect x="195" y="40" width="10" height="120" fill="#8B4513"/>
    <rect x="320" y="40" width="10" height="120" fill="#8B4513"/>
    <!-- Scheiben auf Stab 1 -->
    <rect x="20" y="140" width="110" height="20" rx="5" fill="#FF4444"/>
    <rect x="35" y="120" width="80" height="20" rx="5" fill="#44FF44"/>
    <rect x="50" y="100" width="50" height="20" rx="5" fill="#4444FF"/>
    <!-- Beschriftung -->
    <text x="70" y="190" text-anchor="middle" fill="black">A</text>
    <text x="195" y="190" text-anchor="middle" fill="black">B</text>
    <text x="320" y="190" text-anchor="middle" fill="black">C</text>
</svg>

### Endzustand
<svg viewBox="0 0 400 200" xmlns="http://www.w3.org/2000/svg">
    <!-- Basis -->
    <rect x="20" y="160" width="360" height="20" fill="#8B4513"/>
    <!-- Stäbe -->
    <rect x="70" y="40" width="10" height="120" fill="#8B4513"/>
    <rect x="195" y="40" width="10" height="120" fill="#8B4513"/>
    <rect x="320" y="40" width="10" height="120" fill="#8B4513"/>
    <!-- Scheiben auf Stab 3 -->
    <rect x="270" y="140" width="110" height="20" rx="5" fill="#FF4444"/>
    <rect x="285" y="120" width="80" height="20" rx="5" fill="#44FF44"/>
    <rect x="300" y="100" width="50" height="20" rx="5" fill="#4444FF"/>
    <!-- Beschriftung -->
    <text x="70" y="190" text-anchor="middle" fill="black">A</text>
    <text x="195" y="190" text-anchor="middle" fill="black">B</text>
    <text x="320" y="190" text-anchor="middle" fill="black">C</text>
</svg>

## Lösung für 3 Scheiben

1. Bewege die kleinste Scheibe auf den Zielstab (rechts)
2. Bewege die mittlere Scheibe auf den Hilfsstab (Mitte)
3. Bewege die kleinste Scheibe auf den Hilfsstab
4. Bewege die größte Scheibe auf den Zielstab
5. Bewege die kleinste Scheibe auf den Ausgangsstab (links)
6. Bewege die mittlere Scheibe auf den Zielstab
7. Bewege die kleinste Scheibe auf den Zielstab

> Die optimale Lösung für n Scheiben benötigt 2^n - 1 Züge.

### Rekursiver Lösungsansatz
Für n=1 ist die Lösung trivial (lege die Scheibe auf den Zielstab). Nun nimmt man an, dass man die Lösung für n-1 bereits kennt. Dann

1. Bewege n-1 Scheiben vom Ausgangsstab auf den Hilfsstab
2. Bewege die größte Scheibe auf den Zielstab
3. Bewege die n-1 Scheiben vom Hilfsstab auf den Zielstab

**Aufgabe**: schau dir die Funktion hanoi an und vergleiche diese mit der textuellen Beschreibung. Führe hanoi.go mit unterschiedlichen Turmgrößen aus. Wie oft wird jeweils die zweitgrößte Scheibe bewegt? 

