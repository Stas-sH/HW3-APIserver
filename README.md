


# HW3-APIserver

![User](https://github.com/Stas-sH/HW3-APIserver/assets/64601990/dd275a58-1171-44e8-bf97-b0974f5aa156)

All users have access to all GET requests.
Requests PUT, POST, DELETE only for administrators.

Should be in the HEADER: 
* admin-name - value type string
* admin-password - value type string
(If you are not an administrator, these fields should be empty.)


"/users"
* Get Users '\n'
  Get - "http://127.0.0.1:8000/users"
  ![getUsers](https://github.com/Stas-sH/HW3-APIserver/assets/64601990/4e345117-da32-4eac-88bb-56c9076cde3b)

* addUser
  Post - "http://127.0.0.1:8000/users"
  (The new User must be sent in the body of the request)



"/users/user"
* getUserById
  Get - "http://127.0.0.1:8000/users/user?id=" + id
  ![getUserById](https://github.com/Stas-sH/HW3-APIserver/assets/64601990/b55040b3-6e49-4234-b57f-4b7c4a7a9c5c)
  
* updateUserById
  Put - "http://127.0.0.1:8000/users/user?id=" + id
  (The body of the request must contain the User with the changed data)

* deleteUserById
  Delete - "http://127.0.0.1:8000/users/user?id=" + id
