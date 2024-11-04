<h2>GO STUDY</h2>
 A SIMPLE RESTFUL API FOR STUDENT & TEACHER MANAGEMENT

<h3> Motivation: </h3>
The GOSTUDY Project was created as an experiment to learn and gain some Hands-On on the Go Programming Language

<br>
<br>

<b>Tech Stack:</b>

- Programming Language: GO
- API Testing : Postman
- Type of API: RESTFUL
- Database: SQLITE3

<h3> Structure: </h3>

The Project Contains 3 Main Folders

- _cmd_
- _config_
- _internal_

Additionally the `go.mod` file contains the dependencies installed for this project using `go get package-name`

<h4> cmd folder:</h4>

The cmd folder contains the Entry Point to project. The `main.go` file contains code to run the server on `http:localhost:5001`

- The src folder wraps the procedures under the package name `router`.

  - defineRoutes.go: Contains all the endpoints defined and supported by this application

    <h5>Supported API Calls:</h5>

For Student:

`GET /api/student/id` : Gets Student by ID
`POST /api/create-student` : Provides ability to create student
`GET /api/students`: Gets all Students as a List
`PUT /api/update-student/id`: Updates Students based on ID
`DELETE /api/delete-student/id` : Deletes Student based on ID

  <br>

Similarly for Teachers:
`GET /api/teacher/id` : Gets Teacher by ID
`POST /api/create-teacher` : Provides ability to create Teacher
`GET /api/teachers`: Gets all Teachers as a List
`PUT /api/update-teacher/id`: Updates Teachers based on ID
`DELETE /api/delete-teacher/id` : Deletes Teacher based on ID

<h4> config folder:</h4>

The Config Folder contains the `local.yaml` folder containing the configuration for the project

<h4> internal folder:</h4>

This folder has several self explainatory sub folders disecting code to indivudial responsibiliteis such as models, interfaces, responses, validators, etc.

### Contributing:

To contribute to this project, kindly fork it and raise a pull request. valid pr's will be accepted.

## Note: This is just an exploratory project hence the learning is still in progress
