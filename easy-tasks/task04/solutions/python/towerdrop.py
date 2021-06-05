# Copyright (c) 2019 Valentin B. and contributors. All rights reserved.

# There are no real constant definitions in Python.
# The regular convention is to write constants in all capital letters.
GRAVITY = 9.8


def get_tower_height():
    """Gets the tower height from the user and returns it."""
    tower_height = input('Please enter the height of the tower in metres: ')
    return float(tower_height)


def calculate_height(tower_height, seconds):
    """Returns height from the ground after "seconds" seconds."""
    # Using formula: [ s = u * t + (a * t^2) / 2 ], here u(initial velocity) = 0
    distance_fallen = (GRAVITY * (seconds * seconds)) / 2
    current_height = tower_height - distance_fallen

    return current_height


def print_height(height, seconds):
    """Prints the height every second until the ball has reached the ground."""
    if height > 0.0:
        print(f'After {seconds} seconds, the ball is at {height} metres.')
    else:
        print(f'After {seconds} seconds, the ball has already hit the ground.')


def calculate_and_print_height(tower_height, seconds):
    height = calculate_height(tower_height, seconds)
    print_height(height, seconds)


if __name__ == '__main__':
    tower_height = get_tower_height()

    calculate_and_print_height(tower_height, 0)
    calculate_and_print_height(tower_height, 1)
    calculate_and_print_height(tower_height, 2)
    calculate_and_print_height(tower_height, 3)
    calculate_and_print_height(tower_height, 4)
    calculate_and_print_height(tower_height, 5)
