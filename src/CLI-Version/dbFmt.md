The `database` folder is organised into multiple databases, each one storing the data for an individual school.
Each school database will contain sub folders known as `collections` (eg: data about the school, each year group can be a collection, etc.).

In all collections, JSON files exist called `documents`. The format of the document for a specific student could be:

```
{
  "id": int,
  "name": string, // name of student
  "YearGroup": int, //  year group that the student is in
  "age": int, // age of student
  "subjects": [ // subjects that the student is graded in
    {
      "SubjectName": string, // name of subject
      "StudentMark": float, // student's mark
      "StudentGrade": string, // student's grade
      "EnteredBy": int // teacher or staff member that entered the grade (give the ID)
    },
    .
    .
    .
  ]
}
```

For example, a 16 year old student called John Doe may be in Year 11 attending Hillside Highschool. Their Physics teacher (id=1234) gave them a grade of 75% - which is an A:
  `database/hillsideschool/year11/1.json`:
    ```
    {
      "id": 1,
      "name": "John Doe",
      "YearGroup": 11,
      "age": 16,
      "subjects": [
        {
          "SubjectName": "physics",
          "StudentMark": "75.0",
          "StudentGrade": "A",
          "EnteredBy": 1234
        },
        .
        .
        .
      ]
    }
    ```