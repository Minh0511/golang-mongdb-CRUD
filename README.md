# golang-mongdb-CRUD
CRUD API using golang and mongodb

* Clone the project, then open terminal and run

```
go mod tidy
```

* Run the project, then use postman to test the APIs
```
http://localhost:8080/api/getAllUser
```
```
http://localhost:8080/api/getUserByID/{input_id_here}
```
```
http://localhost:8080/api/createUser
```
```
http://localhost:8080/api/updateUser/{input_id_here}
```
```
http://localhost:8080/api/deleteUser/{input_id_here}
```


* To test the create API, input the following into Body-raw-JSON type
```
{
    "studentID": your_input,
    "email": "your_input",
    "password": "your_input",
    "fullName": "your_input",
    "class": "your_input",
    "dateOfBirth": "YYYY-MM-DD",
    "phoneNumber": your_input
}
```
* To test the update API, input the following into Body-raw-JSON type
```
{
    "any field you want to update": your_input
}
```

