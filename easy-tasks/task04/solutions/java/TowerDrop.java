import java.util.Scanner;

/**
 * Copyright (c) 2019 Valentin B. and contributors. All rights reserved.
 *
 * @author Valentin B.
 */
class Main {
    private static final double GRAVITY = 9.8d;

    /**
     * Gets the tower height from the user and returns it.
     */
    public static double getTowerHeight() {
        Scanner scanner = new Scanner(System.in);
        System.out.print("Please enter the height of the tower in metres: ");
        double towerHeight = scanner.nextDouble();
        scanner.close();

        return towerHeight;
    }

    /**
     * Returns height from the ground after "seconds" seconds.
     */
    private static double calculateHeight(double towerHeight, int seconds) {
        // Using formula: [ s = u * t + (a * t^2) / 2 ], here u(initial velocity) = 0
        double distanceFallen = (GRAVITY * (seconds * seconds)) / 2;
        double currentHeight = towerHeight - distanceFallen;
        return currentHeight;
    }

    /**
     * Prints the height every second until the ball has reached the ground.
     */
    private static void printHeight(double height, int seconds) {
        if (height > 0.0)
            System.out.println(String.format("After %d seconds, the ball is at %.1f metres.", seconds, height));
        else
            System.out.println(String.format("After %d seconds, the ball has already hit the ground.", seconds));
    }

    private static void calculateAndPrintHeight(double towerHeight, int seconds) {
        double height = calculateHeight(towerHeight, seconds);
        printHeight(height, seconds);
    }

    public static void main(String[] args) {
        double towerHeight = getTowerHeight();

        calculateAndPrintHeight(towerHeight, 0);
        calculateAndPrintHeight(towerHeight, 1);
        calculateAndPrintHeight(towerHeight, 2);
        calculateAndPrintHeight(towerHeight, 3);
        calculateAndPrintHeight(towerHeight, 4);
        calculateAndPrintHeight(towerHeight, 5);
    }
}