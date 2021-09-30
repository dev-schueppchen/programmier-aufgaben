//  Solution without extra tasks

import java.util.Scanner;
import java.util.concurrent.ThreadLocalRandom;

public class Main {

    public static void main(String[] args) {
        int number = ThreadLocalRandom.current().nextInt(1, 101);
        int attempts = 10;
        Scanner scanner = new Scanner(System.in);

        System.out.println("The number is between 1 and 100.");
        System.out.println("You now have 10 tries to guess the number.");

        while (attempts > 0) {
            System.out.println("Remaining attemps: " + attempts);
            System.out.print("Please enter the number you would like to guess: ");
            int input = scanner.nextInt();
            attempts--;

            if (input == number) {
                System.out.println("You guessed the number right!");
                System.exit(0);
            } else if (input < number) {
                System.out.println("The number you entered was too low.");
            } else {
                System.out.println("The number you entered was too high.");
            }
        }
        System.out.println("You have no more attempts left!");
        scanner.close();
    }
}

