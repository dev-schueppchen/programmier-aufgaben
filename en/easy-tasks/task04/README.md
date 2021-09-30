# Task 04

> Contributors: Vale#5252  
> Last update: 2019/01/24

---

## Towerdrop

This task is maybe a little more challenging. This is about programming a simulation. First of all, the user sets a height for a tower. From this a tower, a ball is then thrown down. We assume normal gravitational acceleration (9.81 m/s²) and that the ball has no initial velocity. Your program should
then output the height of the ball above the ground after 0, 1, 2, 3, 4 and 5 seconds.

**Keep in mind:** The ball should never fall "into" the ground. The end is at height 0.

Define a constant that stores the value of the gravitational acceleration (9.81).
To calculate how many metres the ball has fallen after x seconds, you can use the formula
`fall_distance = gravitational_acceleration * x_seconds² / 2`.

An example of the output:

```
Please enter the height of the tower in metres: 100
After 0 seconds, the ball is at 100 metres.
After 1 second, the ball is at 95.1 metres.
After 2 seconds, the ball is at 80.4 metres.
After 3 seconds, the ball is at 55.9 metres.
After 4 seconds, the ball is at 21.6 metres.
After 5 seconds, the ball has already hit the ground.
```

**Hints:**

-   You must use simple functions in your programm.
-   Your ball may not reach the ground in 5 seconds. That's fine.
-   Loops are not a mandatory requirement.

---

## Sample solutions

-   [C++](solutions/cpp/towerdrop.cpp)
-   [Java](solutions/java/TowerDrop.java)
-   [Python](solutions/python/towerdrop.py)
