# Stacks
Ein Stack ist eine Datenstruktur, die auf dem LIFO-Prinzip ("Last In, First Out")
basiert. 
Dies bedeutet, dass das zuletzt hinzugefügte Element als erstes wieder entfernt
wird. Man kann sich einen Stack wie einen Stapel Teller vorstellen: Der zuletzt 
aufgelegte Teller wird als erster wieder heruntergenommen.

Hier sind die wichtigsten Eigenschaften und Operationen eines Stacks:

 - **Push**: Das Hinzufügen eines Elements auf den Stack. 
 - **Pop**: Das Entfernen des obersten Elements vom Stack. 
 - **Top/Peek**: Diese Operation ermöglicht es, das oberste Element des Stacks anzusehen, 
ohne es zu entfernen.
 - **IsEmpty**: Eine Überprüfung, ob der Stack leer ist. Dies ist nützlich, um Fehler zu 
vermeiden, wie zum Beispiel das Versuchen, ein Element von einem leeren Stack zu 
entfernen.
 - **Size**: Gibt die Anzahl der Elemente im Stack an.

Beispiel: 
```
Initial   Push(1)   Push(2)   Push(3)   Pop()->3
--------  --------  --------  --------  --------
|   |     |   |     |   |     | 3 |     |   |
|   |     |   |     | 2 |     | 2 |     | 2 |
|   |     | 1 |     | 1 |     | 1 |     | 1 |
--------  --------  --------  --------  --------
```

In der Praxis wird ein Stack oft für Aufgaben verwendet, bei denen Dinge in 
umgekehrter Reihenfolge abgearbeitet werden müssen, wie zum Beispiel:

Rückgängig-Funktionen in Softwareanwendungen: Jede Aktion wird auf einen Stack 
gelegt, sodass man die Aktionen in umgekehrter Reihenfolge rückgängig machen kann.

Funktionsaufrufe: In den meisten Programmiersprachen wird ein Stack verwendet, 
um Informationen über aktive Funktionsaufrufe zu speichern (Call Stack).

Syntaxprüfung: In Compilern und Interpretern zur Überprüfung der Syntax von 
Programmiersprachen.

Ein Stack ist somit ein grundlegendes Werkzeug in der Informatik, das in vielen 
verschiedenen Bereichen zum Einsatz kommt.