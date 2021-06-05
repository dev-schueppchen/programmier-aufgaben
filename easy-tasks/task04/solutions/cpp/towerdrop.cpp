// Copyright (c) 2019 Valentin B. and contributors. All rights reserved.

const float gravity = 9.8;

/// Gets the tower height from the user and returns it.
double getTowerHeight()
{
    std::cout << "Please enter the height of the tower in metres: ";
    double towerHeight;
    std::cin >> towerHeight;
    return towerHeight;
}

/// Returns height from the ground after "seconds" seconds.
double calculateHeight(double towerHeight, int seconds)
{
    // Using formula: [ s = u * t + (a * t^2) / 2 ], here u(initial velocity) = 0
    double distanceFallen = (gravity * (seconds * seconds)) / 2;
    double currentHeight = towerHeight - distanceFallen;

    return currentHeight;
}

/// Prints the height every second until the ball has reached the ground.
void printHeight(double height, int seconds)
{
    if (height > 0.0)
        std::cout << "After " << seconds << " seconds, the ball is at " << height << " metres.\n";
    else
        std::cout << "After " << seconds << " seconds, the ball has already hit the ground.\n";
}

void calculateAndPrintHeight(double towerHeight, int seconds)
{
    double height = calculateHeight(towerHeight, seconds);
    printHeight(height, seconds);
}

int main()
{
    const double towerHeight = getTowerHeight();

    calculateAndPrintHeight(towerHeight, 0);
    calculateAndPrintHeight(towerHeight, 1);
    calculateAndPrintHeight(towerHeight, 2);
    calculateAndPrintHeight(towerHeight, 3);
    calculateAndPrintHeight(towerHeight, 4);
    calculateAndPrintHeight(towerHeight, 5);

    return 0;
}