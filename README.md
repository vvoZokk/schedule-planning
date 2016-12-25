# The Critical Path Method
It's small educational project.

Implementation of CPM for scheduling a set of tasks. There are list of tasks and information how tasks are bound.
Each task have name, duration and as-soon-as-possible (ASAP) type or as-late-as-possible (ALAP) type.
### Example
Ten tasks from 0 to 9 with random duration and type. Links between tasks:
* after 0 - 1 and 4
* after 1 - 3 and 7
* after 2 - 4 and 7
* after 3 - 5
* after 4 - 9
* after 5 - 8
* after 6 - 7
* after	7 - 8

Compile and run `scheduling.go` to find critical path and schedule tasks (ALAP-tasks is priority).
Results ('#' marks CP-task):
```
________________________________________________________________________________
[Task_0######]
             [Task_1#######]
|---------------------[Task_2::::::::::]
                           [Task_3#######]
                 |---------------------[Task_4::::::::::]
                                         [Task_5#######]
[Task_6::::]---------------------------------|
                           |-----------[Task_7:::]-----|
                                                       [Task_8##########]
                                  |---------------------[Task_9:::::::::]
________________________________________________________________________________

```