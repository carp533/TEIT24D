# TEIT24D Informatik I - Labor 

Hier finden Sie die Programme zur Vorlesung informatik 1. Dieses GIT Repository in das workspace Verzeichnis kopieren/auschecken.
(über GIT Zip Datei herunter laden oder git clone https://github.com/carp533/TEIT24D)

## Software
- Go Compiler: https://go.dev/ (aktuelle Version)
- Visual Studio Code: https://code.visualstudio.com/
  - hier muss noch das Go Plugin installiert werden (Go for Visual Studio Code)
- optional git: https://git-scm.com/

## Hello World

Um mit VS Code und go mit mehreren Modulen arbeiten zu können, müssen diese einem workspace zugeordnet werden. Der workspace ist nicht Bestandteil des GIT-Repositories und muss im Laufe der Vorlesung angepasst werden.

in workspace Verzeichnis wechseln
```bash
    $ mkdir myfirstmodule
    $ cd myfirstmodule
    $ go mod init myfirstmodule
```
eine go Datei erzeugen myfirst.go
```go
    package main

    import (
        "fmt"
    )

    func main() {
        fmt.Println("Hello")
    }
```
datei starten
```bash
    $ go run myfirst.go
```

zurück in das workspace Verzeichnis (TEIT24D) wechseln und das modul zum workspace hinzufügen
```bash
    $ cd ..
    $ go work init
    $ go work use .\myfirstmodule\
```
das `go work init` muss nur einmalig gesartet werden. Um die Module aus KW49 einzubinden `go work use ./KW49/hello` verwenden.
Falls Testdateien vorhanden sind (*_test.go) können die Tests mit "go test" oder "go test -v" ausgeführt werden. oder über VS Code in der Testdatei (Zeile 1) run package tests oder run file tests. Die Datei go.work sollte dann etwa so aussehen:
```go
go 1.23.3

use (
	./KW49/hello
	./KW49/ifelse
    ./myfirstmodule
)
```
## nützliche Links:
- Go Playground: https://go.dev/play/
- A Tour of Go: https://go.dev/tour/welcome/1

## git

Siehe [Git-Dokumentation](git.md).