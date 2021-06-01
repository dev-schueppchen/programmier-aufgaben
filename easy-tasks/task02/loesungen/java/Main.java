//Lösung ohne Zusatzaufgaben

import java.util.Scanner;
import java.util.concurrent.ThreadLocalRandom;

public class Main {

    public static void main(String[] args) {
        int zahl = ThreadLocalRandom.current().nextInt(1, 101);
        int versuche = 10;
        Scanner scanner = new Scanner(System.in);

        System.out.println("Die Zahl befindet sich zwischen einschließlich 1 und 100.");
        System.out.println("Du hast nun 10 Versuche, die Zahl zu raten.");

        while (versuche > 0) {
            System.out.println("Verbleibende Versuche: " + versuche);
            System.out.print("Bitte gib die Zahl ein, die du erraten möchtest: ");
            int eingabe = scanner.nextInt();
            versuche--;

            if (eingabe == zahl) {
                System.out.println("Du hast die Zahl richtig erraten!");
                System.exit(0);
            } else if (eingabe < zahl) {
                System.out.println("Deine eingegebene Zahl war zu niedrig.");
            } else {
                System.out.println("Deine eingegebene Zahl war zu hoch.");
            }
        }
        System.out.println("Du hast keine Versuche mehr übrig!");
    }
}

