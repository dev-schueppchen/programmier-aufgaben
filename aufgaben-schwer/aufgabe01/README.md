# Aufgabe 01
> Contributors: zekro#9131, Ron31#2338
> Letzte Aktualisierung:  2020/06/24

---

## Chat Room

Erstelle einen Web-Service, welcher z.b. über `localhost:8080` im Browser erreichbar ist, in dem man in einem Text-Feld seinen Nicknamen und in einem anderen eine Nachricht eintippen kann. Sobald man auf einen Senden-Button klick, wird diese Nachricht bei allen verbundenen Clients angezeigt mit dem Name des Senders, dem Text der Nchricht und einem Zeitstempel, wann die Nachricht versandt wurde.

![](https://i.zekro.de/firefox_2019-01-29_12-53-57.png)

Voraussetzungen:
- Eingehende Messages müssen erscheinen ohne dass man die Website neu laden muss
- Nickname, Text und der Zeitstempel müssen an alle verbundenen Clients gesendet werden
- Es dürfen externe Libraries *(z.B. express, gorilla/websocket, flask, websocket-client...)* verwendet werden

---

## Musterlösungen

- [Go](loesungen/go)
- [NodeJS](loesungen/js)
