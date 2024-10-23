## Airline checkin system - To assign passenger seats

1. Set your DB variables:

`set DBUSER = <your_database_user>` <br>
`set DBPASS = <your_database_password>`


2. Update the `set_sb.sql` file to inset data into the tables.
go run `insert_db_data.go`

3. Run the `set_db.sql` file to create DB `airline_checkin`.
Run: `mysql -u root -p < <path to set_db.sql>`

#### SIMPLE SELECT STATEMENT
- Retrieves the first available seat but does not lock it.
- Multiple transactions could select the same seat simultaneously.

```
SELECT seat_no from seats WHERE user_id IS null ORDER BY seat_no LIMIT 1;
```
```
Time taken to assign seats: 241.1619ms

Flight Seating Chart:
.........xx.........
.........xx.........
.........xx.........

.........x..........
.........x..........
.........x..........
```

#### FOR UPDATE
- Retrieves the first available seat and locks it for the current transaction.
- Other transactions cannot read or modify this seat until the lock is released.
- Assign seats in order, sequentially.

```
SELECT seat_no from seats WHERE user_id IS null ORDER BY seat_no LIMIT 1 FOR UPDATE;
```
```
Time taken to assign seats: 917.702ms

Flight Seating Chart:
xxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxx

xxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxx
```

#### FOR UPDATE SKIP LOCKED
- Retrieves and locks the first available seat, but skips rows locked by other transactions.
- Ensures that different transactions pick different available seats without waiting.
- May not guarantee sequential assignment.

```
SELECT seat_no from seats WHERE user_id IS null ORDER BY seat_no LIMIT 1 FOR UPDATE SKIP LOCKED;
```
```
Time taken to assign seats: 161.8957ms

Flight Seating Chart:
xxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxx

xxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxx
xxxxxxxxxxxxxxxxxxxx
```

Note : Run `UPDATE seats SET user_id = NULL;` to ensure seats are null, before re-assigning it again.