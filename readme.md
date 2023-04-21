# IGlobal Backend Test

## Table of Contents
- [IGlobal Backend Test](#iglobal-backend-test)
  - [Table of Contents](#table-of-contents)
  - [List API Endpoint](#list-api-endpoint)
    - [Login](#login)
      - [Request Body](#request-body)
      - [Responses](#responses)
    - [Register](#register)
      - [Request Body](#request-body-1)
      - [Responses](#responses-1)
    - [Search Course](#search-course)
      - [Responses](#responses-2)
    - [Get Course By Id](#get-course-by-id)
      - [Parameters](#parameters)
      - [Responses](#responses-3)
    - [Enroll Course](#enroll-course)
      - [Parameters](#parameters-1)
      - [Responses](#responses-4)
    - [Create Course](#create-course)
      - [Request Body](#request-body-2)
      - [Responses](#responses-5)
    - [Get All Course](#get-all-course)
      - [Responses](#responses-6)
    - [Update Course By Id](#update-course-by-id)
      - [Parameters](#parameters-2)
      - [Responses](#responses-7)
    - [Delete Course By Id](#delete-course-by-id)
      - [Parameters](#parameters-3)
      - [Responses](#responses-8)
  - [How To Run This Project](#how-to-run-this-project)
      - [Run the Applications on Local Machine](#run-the-applications-on-local-machine)
  - [Tech Stack](#tech-stack)

## List API Endpoint
https://www.postman.com/adhityaf-learn/workspace/adhitya-febhiakbar/request/13223681-e2edfecc-634c-48cb-ba06-e19bdc238f52

Below is the list of features and API endpoints available in this project:
### Login
```http
POST /v1/login
```
#### Request Body
```java{
   "email"     : string,
   "password"  : string,
}
```
Example:
```javascript
{
   "email"     : "adit@mail.com",
   "password"  : "12341234"
}
```

#### Responses
```javascript
{
   "message"   : string,
   "token"     : string
}
```
Example: 
```javascript
{
   "message"   : "Login successful",
   "token"     : "eyJhbGciOJIUzI1NiIsIn5cCI6IkpXVCJ9.eyJleHBpcmUiOjE2ODA2NTU4DYsImV4cGlyZV9kYXRlIjoiMjAyMy0wNC0wNVQwNzo1MDowNi4xMj4NDUrMDc6MDAiLCJuYW1lIjoiYWRpdCB1c2VyIEiLCJyb2xlIjoidXNlciIsInVzZXJfaWQiOjF9.xyQYEXqoVnDbgrLDq8BXT9i4VIZX2LPCVZVTkxAk8"
}
```

### Register
```http
POST /v1/register
```
#### Request Body
```javascript
{
   "name"         : string,
   "email"        : string,
   "password"     : string,
}
```
Example:
```javascript
{
   "name"         : "adit",
   "email"        : "adit@mail.com",
   "password"     : "12341234"
}
```

#### Responses
```javascript
{
   "data"         : object
}
```
Example: 
```javascript
{
   "data"      : {
      "user_id"   : 1,
      "email"     : "adit@mail.com",
      "password"  : "12341234",
      "name"      : "Deleted User",
      "role"      : "user",
      "image_url" : "",
      "created_at": "2023-04-19T07:56:18.026+07:00",
      "updated_at": "2023-04-19T07:56:18.026+07:00",
      "deleted_at": null
   }
}
```

### Search Course

```http
GET /v1/courses/search?level=Beginner&language=Indonesia&query=java&size=5&page=1
```

| Parameter | Type     | Description           |
| :--- |:---------|:----------------------|
| `query` | `string` | for searching by title or description |
| `page` | `string` | current page number |
| `size` | `string` | limit size per page |
| `level` | `string` | filter by level: Beginner, Intermediate, Advance; Default: Beginner |
| `language` | `string` | filter by language: Indonesia, English; Default: Indonesia |


#### Responses
```javascript
{
   "items"                 : array of object
}
```
Example:
```javascript
{
   "items": [
      {
         "course_id"    : 2,
         "title"        : "Belajar Javascript",
         "slug"         : "belajar-javascript",
         "description"  : "Deskripsi Belajar Javascript",
         "price"        : 120000,
         "level"        : "Beginner",
         "language"     : "Indonesia",
         "UserCourses"  : null,
         "created_at"   : "2023-04-19T06:38:38.897+07:00",
         "updated_at"   : "2023-04-19T06:38:38.897+07:00"
      },
      {
         "course_id"    : 3,
         "title"        : "Belajar Java",
         "slug"         : "belajar-java",
         "description"  : "Deskripsi Belajar Java",
         "price"        : 120000,
         "level"        : "Beginner",
         "language"     : "Indonesia",
         "UserCourses"  : null,
         "created_at"   : "2023-04-19T06:38:45.86+07:00",
         "updated_at"   : "2023-04-19T06:38:45.86+07:00"
      }
   ]
}
```

### Get Course By Id
```http
GET /v1/courses/{courseId}
```
#### Parameters
```javascript
{
   "courseId"        : integer
}
```
Example:
```javascript
{
   "courseId"        : 1
}
```

#### Responses
```javascript
{
   "items"           : object
}
```
Example: 
```javascript
{
   "data": {
      "course_id"    : 5,
      "title"        : "Learn Golang",
      "slug"         : "learn-golang",
      "description"  : "Deskripsi Learn Golang",
      "price"        : 125000,
      "level"        : "Beginner",
      "language"     : "English",
      "created_at"   : "2023-04-19T06:39:28.509+07:00",
      "updated_at"   : "2023-04-19T06:39:28.509+07:00"
   }
}
```

### Enroll Course
```http
POST /v1/courses/enroll/{courseId}
```
#### Parameters
```javascript
{
   "courseId"           : integer
}
```
Example:
```javascript
{
   "courseId"           : 1
}
```

#### Responses
```javascript
{
   "data"               : object
}
```
Example: 
```javascript
{
   "data": {
      "user_id"         : 1,
      "course_id"       : 1,
      "score"           : 0,
      "created_at"      : "2023-04-19T08:57:53.276+07:00",
      "updated_at"      : "2023-04-19T08:57:53.276+07:00"
   }
}
```

### Create Course
```http
POST /v1/courses
```
#### Request Body
```javascript
{
   "title"              : string,
   "description"        : string,
   "price"              : integer,
   "level"              : string,
   "language"           : string
}
```
Example:
```javascript
{
   "title"              : "Learn Go",
    "description"       : "Deskripsi Learn Go",
    "price"             : 125000,
    "level"             : "Beginner",
    "language"          : "English"
}
```

#### Responses
```javascript
{
   "data"               : object
}
```
Example: 
```javascript
{
   "data": {
      "course_id"       : 1,
      "title"           : "Learn Go",
      "slug"            : "learn-go",
      "description"     : "Deskripsi Learn Go",
      "price"           : 125000,
      "level"           : "Beginner",
      "language"        : "English",
      "created_at"      : "2023-04-19T08:26:26.984+07:00",
      "updated_at"      : "2023-04-19T08:26:26.984+07:00"
   }
}
```

### Get All Course
```http
GET /v1/courses
```

#### Responses
```javascript
{
   "data"               : array of object
}
```
Example: 
```javascript
{
   "items": [
      {
         "course_id"    : 1,
         "title"        : "Learn Dart",
         "slug"         : "learn-dart",
         "description"  : "Deskripsi Learn Dart",
         "price"        : 75000,
         "level"        : "Advance",
         "language"     : "English",
         "created_at"   : "2023-04-19T06:36:57.826+07:00",
         "updated_at"   : "2023-04-19T06:56:30.44+07:00"
      },
      {
         "course_id"    : 2,
         "title"        : "Belajar Javascript",
         "slug"         : "belajar-javascript",
         "description"  : "Deskripsi Belajar Javascript",
         "price"        : 120000,
         "level"        : "Beginner",
         "language"     : "Indonesia",
         "created_at"   : "2023-04-19T06:38:38.897+07:00",
         "updated_at"   : "2023-04-19T06:38:38.897+07:00"
      }
   ]
}
```

### Update Course By Id
```http
PUT /v1/courses/{courseId}
```
#### Parameters
```javascript
{
   "courseId"     : integer
}
```
Example:
```javascript
{
   "courseId"     : 1
}
```

#### Responses
```javascript
{
   "message"      : string
}
```
Example: 
```javascript
{
   "message"      : "Success update data with id : 1"
}
```

### Delete Course By Id
```http
DELETE /v1/courses/{courseId}
```
#### Parameters
```javascript
{
   "courseId"     : integer
}
```
Example:
```javascript
{
   "courseId"     : 1
}
```

#### Responses
```javascript
{
   "message"      : string
}
```
Example: 
```javascript
{
   "message"      : "Success DELETE data with id : 1"
}
```

## How To Run This Project
#### Run the Applications on Local Machine

```bash
# Clone into your workspace
$ git clone https://github.com/adhityaf/iglobal-be-test.git
# Change directory to project
$ cd iglobal-be-test
# Download all dependencies
$ go mod tidy
# setup the environment variables by copy .env-example file and rename it to .env
# run the application
$ go run main.go
```

## Tech Stack
- Go 1.19
- Gin
- MySQL
- GORM