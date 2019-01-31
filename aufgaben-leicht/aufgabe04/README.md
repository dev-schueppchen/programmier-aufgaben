# Aufgabe 02

> Contributors: Vale#5252  
> Letzte Aktualisierung: 2019/01/24

---

## Towerdrop

Diese Aufgabe ist vielleicht etwas anspruchsvoller. Hier geht es darum, eine Simulation
zu programmieren. Und zwar legt der Nutzer zuerst eine Höhe für einen Turm fest. Von diesem
Turm wird dann ein Ball hinuntergeworfen. Wir gehen von normaler Fallbeschleunigung (9.8 m/s²),
sowie davon, dass der Ball keine Anfangsgeschwindigkeit besitzt, aus. Dein Programm sollte
anschließend die Höhe des Balles über dem Boden nach 0, 1, 2, 3, 4 und 5 Sekunden ausgeben.

**Beachte:** Der Ball sollte niemals "in" den Boden fallen. Bei Höhe 0 ist Schluss.

Definiere eine Konstante, welchen den Wert der Fallbeschleunigung (9.8) speichert.
Um zu berechnen, wie viele Meter der Ball nach x Sekunden gefallen ist, kannst du die Formel
`gefallene_distanz = fallbeschleunigung * x_sekunden² / 2` verwenden.

Ein Beispiel für die Ausgabe:

```
Bitte gib die Höhe des Turmes in Metern ein: 100
Nach 0 Sekunden ist der Ball auf 100 Metern Höhe.
Nach 1 Sekunden ist der Ball auf 95.1 Metern Höhe.
Nach 2 Sekunden ist der Ball auf 80.4 Metern Höhe.
Nach 3 Sekunden ist der Ball auf 55.9 Metern Höhe.
Nach 4 Sekunden ist der Ball auf 21.6 Metern Höhe.
Nach 5 Sekunden ist der Ball bereits auf dem Boden aufgeschlagen.
```

Der Einfachheit halber musst du dich nicht um grammatische Richtigkeit in den Ausgaben
deines Programmes kümmern. ;)

**Hinweise:**
- Du musst in deinem Programm einfache Methoden verwenden.
- Dein Ball erreicht vielleicht in 5 Sekunden nicht den Boden. Das ist in Ordnung.
- Schleifen werden nicht zwingend vorrausgesetzt.