# Attendance application

- Build an application for attendance of the student so that the manual work of writing into the register can get replaced.

# plan

- To build the attendance application step by step in small chunks and follow MVC principle
- For now it is only for student who can make there attendance by this application from there account.
- In this iteration it is terminal based system.

# Step 1

- [X] show simple menu like

```
    1. Mark Attendance
    2. Exit
```

# step 2

- [X] remove the bug of storing duplicate value.

# step 3

- [] use JSON file to store the attendance.

- structure 
```json
{
    "date1": ["student1", "student2"],
    "date2": ["student1", "student2"]
}
```
- [X] read the data if available otherwise create the file.
- [X] add the data if needed.
- [X] write it back to the file.