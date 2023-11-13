# HW3-APIserver

![User](https://github.com/Stas-sH/HW3-APIserver/assets/64601990/dd275a58-1171-44e8-bf97-b0974f5aa156)

All users have access to all GET requests.
Requests PUT, POST, DELETE only for administrators.

Should be in the HEADER: 
* admin-name - value type string
* admin-password - value type string
(If you are not an administrator, these fields should be empty.)


## "/users"
### Get Users
  Get - "http://127.0.0.1:8000/users"
  ![getUsers](https://github.com/Stas-sH/HW3-APIserver/assets/64601990/4e345117-da32-4eac-88bb-56c9076cde3b)

### Add User
  Post - "http://127.0.0.1:8000/users"
  (The new User must be sent in the body of the request)



## "/users/user"
### Get User by id
  Get - "http://127.0.0.1:8000/users/user?id=" + id
  ![getUserById](https://github.com/Stas-sH/HW3-APIserver/assets/64601990/b55040b3-6e49-4234-b57f-4b7c4a7a9c5c)
  
### Update User by id
  Put - "http://127.0.0.1:8000/users/user?id=" + id
  (The body of the request must contain the User with the changed data)

### Delete User by id
  Delete - "http://127.0.0.1:8000/users/user?id=" + id

## PostgreSQL
### Open
* sudo -i -u postgres
### Create your user for DB
* createuser -P username
### Create DB for user
* createdb testclientserver -O username
### Go to the database
* psql
* \c testclientserver
### Сreating an admins table
* create table admins ( id serial not null unique, adminname varchar(255) not null, password varchar(255) not null );
### Add data into the admins table
* insert into admins (adminname, password) values ('your_name', 'your_password');
### Сreating an users table
* create table users ( id serial not null unique, username varchar(255) not null, mail varchar(255) not null, phone varchar(255) not null, password varchar(255) not null );
### Add data into the users table
* insert into users (username, mail, phone, password ) values ('Stanislav', 'stas@mail.com', '+380970403381', 'qwerty1');
* insert into users (username, mail, phone, password ) values ('Amina', 'amina@mail.com', '+123456789', 'qwerty2');

## Server configuration
* in "server" floder - find the file "dbserver.go"
* in all functions "sql.Open("postgres", "host=127.0.0.1 port=5432 user=stas dbname=testclientserver sslmode=disable password=1234")" - you need to change the data to your own:
1. user = your_user_name
2. password = your_user_password
