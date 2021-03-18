# simplebank

## Description

### Transaction
* We have 4 property in transaction is Atomicity, Consistency, Isolation, Durability.

### Race condition
* It's mean one record will be updated by 2 tx at the same time. It will make the value of that's record wrong.

### Deadlock
* When we use lock for update database to avoid race condition. We can make a deadlock in DB. Deadlock is the table will be locked by tx1 when it's processing, and tx2 want to do someting in table 1 and it occur deadlock.
