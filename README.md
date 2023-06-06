# README #

## API Contract
find the api documentation at request_example.postman_collection file
## Setup

1. Install Go version 1.19
2. Download dependencies with command `go mod tidy and go vendor`
3. Create `.env` file based on `.env.example` and fill with your configuration

## Run

Use this command to run API app from root directory:

```shell
go run app/main.go
```

---

## Coach Appointment Tech Spec

This README would normally document whatever steps are necessary to get your application up and running.

### Feature Description ###
there are fifteen endpoint that could use for appointment process, which are:
1. Login
2. Register
3. Create Admin
4. Read Admin by id
5. Read All Admin
6. Update Admin
7. Delete Admin
8. Create Customer
9. Read Customer by id
10. Read All Customer
11. Update Customer 
12. Delete Customer
13. Read All Registration Approval
14. Approve/Reject Registration Approval
15. Activate/Deactivate admin account

### Acceptance Criteria ###
1. Admin or Super Admin can log in to system
2. New registered admin need approval by super admin before they can log in to system
3. Super Admin can deactivate or activate deactivated admin account

### Architecture and Design ###
this service using onion architecture, there are 5 layers
from inner to outer which are entity, repository, use case,
controller, and request handler. the usage and responsibility of
each layer are follow:
1. **Entity**: this layer contains the domain model or entities
   of the system. These are the core objects that
   represent the business concepts and rules.
2. **Repository**: This layer provides an interface for the
   application to access and manipulate the entities.
   It encapsulates the data access logic and provides
   a way to abstract the database implementation details.
3. **Use case** : This layer contains the business logic
   or use cases of the system. It defines the operations
   that can be performed on the entities and orchestrates
   the interactions between the entities and the repository layer.
4. **Controller**: This layer handles the HTTP requests and
   responses. It maps the incoming requests to the appropriate
   use case and returns the response to the client.
5. **Request handler**: This layer is responsible for handling
   the incoming HTTP requests and passing them on to
   the controller layer.

[//]: # (### Service State Diagram ###)

[//]: # (this diagram will explain the state of appointment through)

[//]: # (th every process at system to achieve the output state.)

[//]: # ()
[//]: # ([//]: # &#40;![state diagram]&#41;)
[//]: # ()
[//]: # (As explain at state diagram. there are 3 action could be performed)

[//]: # (within the system which are create appointment, approval of appointment,)

[//]: # (and reschedule appointment. So to cover the business logic of each)

[//]: # (action. there I provide the activity diagram for them.)

### Create Admin Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/admin/create%20admin.png)

### Read Admin Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/admin/read%20admin.png)

### Update Admin Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/admin/update%20admin.png)

### Delete Admin Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/admin/delete%20admin.png)

### Activate Admin Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/admin/activate%20admin.png)

### Deactivate Admin Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/admin/deactivate%20admin.png)

### Login Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/authenticate/login.png)

### Register Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/authenticate/register.png)

### Create Customer Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/customer/create%20customer.png)

### Read Customer Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/customer/read%20customer.png)

### Update Customer Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/customer/update%20customer.png)

### Delete Customer Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/customer/delete%20customer.png)

### Read Registration Approval Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/register-approval/read%20register.png)

### Approve Registration Approval Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/register-approval/approve.png)

### Reject Registration Approval Activity Diagram ###

![create appointment activity](https://raw.githubusercontent.com/nashrul-be/crm/main/plantuml/register-approval/reject.png)

###  Data Flow Diagram ###

![data flow diagram](https://raw.githubusercontent.com/nashrul-be/mini-project/main/entity%20diagram.png)