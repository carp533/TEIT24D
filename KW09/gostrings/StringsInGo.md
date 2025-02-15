# Strings

Strings sind in go byte-slices.
Ein Zeichen ist ein Unicode Code Point und kann aus mehreren bytes bestehen. Für lateinische Buchstaben reicht ein byte (ASCII-Codierung). 

**Wichtig**: nicht jeder byte-slice ist ein valider UTF-8-string. Beim iterieren über einen String darauf achten, dass man über die Zeichen (Code Points) iteriert und nicht über die bytes (siehe Beispielprogramm).

https://go.dev/blog/strings

```go 
const nihongo = "日本語"
for index, runeValue := range nihongo {
    fmt.Printf("%#U starts at byte position %d\n", runeValue, index)
}
```
```
U+65E5 '日' starts at byte position 0
U+672C '本' starts at byte position 3
U+8A9E '語' starts at byte position 6
```