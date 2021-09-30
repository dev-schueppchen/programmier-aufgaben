import java.util.Map;
import java.util.Scanner;

public class Main {

    public static void main(String[] args) {
        Map<String, String> userpws = Map.of("root", "rootpw");
        Scanner scanner = new Scanner(System.in);

        System.out.print("User: ");
        String user = scanner.nextLine();

        System.out.print("Password: ");
        if (userpws.get(user).equals(scanner.nextLine())) {
            System.out.println("Du bist nun eingeloggt");
        } else {
            System.out.println("Falsches Passwort");
        }
    }
}

