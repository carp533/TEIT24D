# Maps
Eine "map" ist eine Datenstruktur, die Schlüssel-Wert-Paare speichert. 
Eine Map ist eine ungeordnete Sammlung von Elementen, bei der jeder 
Schlüssel eindeutig ist und einem bestimmten Wert zugeordnet ist. 
In Go wird die Map mit dem eingebauten Typ map erstellt.
In der Informatik bezeichnet man die "map" auch als Hashtabelle.

Die Definition ist ```map[KeyType]ValueType```, also z.B.
```var m map[string]int```
Der KeyType einer map muss "comparable" sein, also einer der Datentypen
  boolean, numeric, string, pointer, channel, interface 
  oder Strukturen oder Arrays mit diesen Typen


Vor dem Verwenden der map muss diese mit dem eingebauten Befehl make 
initialisiert werden: ```m = make(map[string]int)```

Eindeutige Schlüssel: Jeder Schlüssel in einer Map ist eindeutig, es 
kann also nur ein Wert pro Schlüssel existieren.

Dynamische Größe: Die Größe einer Map kann sich dynamisch ändern, da 
Elemente hinzugefügt oder entfernt werden können.

Ungeordnete Sammlung: Die Reihenfolge der Elemente in einer Map ist 
nicht garantiert. Bei Iterationen über die Map gibt es keine 
festgelegte Reihenfolge.

Effiziente Schlüsselzugriffe: Der Zugriff auf Werte in einer Map erfolgt 
in O(1) Zeitkomplexität, was bedeutet, dass die Zugriffszeit konstant ist, 
unabhängig von der Größe der Map.