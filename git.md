# Git Grundlagen für Programmierer

## Was ist Git?
Git ist ein Versionskontrollsystem, das Entwicklern ermöglicht, Änderungen im Quellcode zu verfolgen, zusammenzuarbeiten und verschiedene Versionen eines Projekts zu verwalten.

## Grundlegende Konzepte

### Repository
Ein Repository (kurz: Repo) ist ein Projektordner, der die gesamte Projekt-History und Versionsinformationen enthält.

### Arbeitsablauf

#### Repository Klonen
```bash
git clone https://github.com/carp533/TEIT24D.git
```

#### Status prüfen
```bash
git status
```
Zeigt alle aktuellen Änderungen und den Status der Dateien.

#### Dateien zum Commit Vorbereiten
```bash
# Einzelne Datei hinzufügen
git add dateiname.txt

# Alle Änderungen hinzufügen
git add .
```

#### Änderungen Committen
```bash
git commit -m "Beschreibende Commit-Nachricht"
```

#### Änderungen Hochladen
```bash
git push origin main
```

## Branching

### Branch Erstellen
```bash
# Neuen Branch erstellen
git branch neuer-branch

# Zu einem Branch wechseln
git checkout neuer-branch

# Branch erstellen und direkt wechseln
git checkout -b neuer-branch
```

### Tipps
- Committen Sie kleine, logische Änderungen
- Schreiben Sie aussagekräftige Commit-Nachrichten
- Verwenden Sie Branches für neue Features

## Visual Studio Code Integration

### Source Control Ansicht
Die Empfehlung ist git über die Source Control Ansicht in Visual Studio Code zu bedienen
- Tastenkombination: `Strg+Shift+G`
- Änderungen committen und synchronisieren
- Merge-Konflikte auflösen

## Umgang mit Merge-Konflikten
- Konflikte manuell in VS Code bearbeiten
- Entscheiden, welche Änderungen übernommen werden
- Änderungen zusammenführen

## Häufige Befehle
- `git log`: Commit-Verlauf anzeigen
- `git pull`: Änderungen vom Remote-Repository herunterladen
- `git branch`: Verfügbare Branches auflisten
