# HW3-APIserver

<img src="/home/stanislav/education/UpGolang/User.jpg" alt="User struct">

All users have access to all GET requests.
Requests PUT, POST, DELETE only for administrators.

Should be in the HEADER: 
* admin-name - value type string
* admin-password - value type string
(If you are not an administrator, these fields should be empty.)


"/users"
1. getUsers
  Get - "http://127.0.0.1:8000/users"
  <img src="/home/stanislav/education/UpGolang/getUsers.jpg" alt="getUsers">

2. addUser
  Post - "http://127.0.0.1:8000/users"
  (The new User must be sent in the body of the request)



"/users/user"
1 getUserById
  Get - "http://127.0.0.1:8000/users/user?id=" + id
  <img src="/home/stanislav/education/UpGolang/getUserById.png" alt="getUserById">
  
2 updateUserById
  Put - "http://127.0.0.1:8000/users/user?id=" + id
  (The body of the request must contain the User with the changed data)

3 deleteUserById
  Delete - "http://127.0.0.1:8000/users/user?id=" + id
