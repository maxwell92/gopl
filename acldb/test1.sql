Permisssion
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | int(10)      | NO   | PRI | NULL    | auto_increment |       1                  2
| name       | varchar(256) | NO   | UNI | NULL    |                |       S                  O
| status     | int(10)      | NO   |     | NULL    |                |       1                  1
| createdAt  | varchar(256) | NO   |     | NULL    |                |       
| modifiedAt | varchar(256) | NO   |     | NULL    |                |
| modifiedOp | int(10)      | NO   |     | NULL    |                |
| comment    | varchar(256) | YES  |     | NULL    |                |
| navList    | text         | NO   |     | NULL    |                |       ..SyncImage..    ..Operation..
+------------+--------------+------+-----+---------+----------------+

User
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | int(10)      | NO   | PRI | NULL    | auto_increment |       1
| name       | varchar(256) | NO   | UNI | NULL    |                |     maxwell92
| password   | varchar(256) | NO   |     | NULL    |                |
| orgId      | int(10)      | NO   | MUL | NULL    |                |
| status     | int(10)      | NO   |     | NULL    |                |
| createdAt  | varchar(256) | NO   |     | NULL    |                |
| modifiedAt | varchar(256) | NO   |     | NULL    |                |
| modifiedOp | int(10)      | NO   |     | NULL    |                |
| comment    | varchar(256) | YES  |     | NULL    |                |
| navList    | text         | NO   |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+

UserPermission
+------------+--------------+------+-----+---------+----------------+
| Field      | Type         | Null | Key | Default | Extra          |
+------------+--------------+------+-----+---------+----------------+
| id         | int(10)      | NO   | PRI | NULL    | auto_increment |       1
| userId     | int(10)      | NO   |     | NULL    |                |       1
| permId     | int(10)      | NO   |     | NULL    |                |       1
| status     | int(10)      | NO   |     | NULL    |                |
| createdAt  | varchar(256) | NO   |     | NULL    |                |
| modifiedAt | varchar(256) | NO   |     | NULL    |                |
| modifiedOp | int(10)      | NO   |     | NULL    |                |
| comment    | varchar(256) | YES  |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+

CREATE TABLE permission(id INT(10) PRIMARY KEY NOT NULL AUTO_INCREMENT, name VARCHAR(256) NOT NULL UNIQUE, status INT(10) NOT NULL, createdAt VARCHAR(256) NOT NULL, modifiedAt VARCHAR(256) NOT NULL, modifiedOp INT(10) NOT NULL, comment VARCHAR(256), navList VARCHAR(256) NOT NULL);

REATE TABLE userpermission ( id INT(10) NOT NULL PRIMARY KEY AUTO_INCREMENT, userId INT(10) NOT NULL, permId INT(10) NOT NULL, status INT(10) NOT NULL, createdAt VARCHAR(256) NOT NULL, modifiedAt VARCHAR(256) NOT NULL, modifiedOp INT(10) NOT NULL, comment VARCHAR(256) NOT NULL);
