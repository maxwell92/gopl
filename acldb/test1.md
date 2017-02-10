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
| userId     | int(10)      | NO   | UNI | NULL    |                |       1
| permId     | int(10)      | NO   | MUL | NULL    |                |       1
| status     | int(10)      | NO   |     | NULL    |                |
| createdAt  | varchar(256) | NO   |     | NULL    |                |
| modifiedAt | varchar(256) | NO   |     | NULL    |                |
| modifiedOp | int(10)      | NO   |     | NULL    |                |
| comment    | varchar(256) | YES  |     | NULL    |                |
+------------+--------------+------+-----+---------+----------------+
