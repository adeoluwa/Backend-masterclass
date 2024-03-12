# Backend Masterclass Project

This repository documents my learning expeirence in progressive complex topics as I create a Go backend appliction using technologies like Docker, K8, AWS, a lot of unit testing and other efforts.

## Requirements

BAsic Understanding of Golang

Go version 20 > installed

Docker (check the installation process for windows in the offical docs)

Postgres

sqlc

## Progress Report

Day 1:
Used [dbdiagrams](https://dbdiagram.io) to design a database schema, installed docker and table plus for database visualization

Day 2:
unit testing, Databse CRUD Operations, installed postgres driver.

Day 3:
Golang DB Transaction

# DB Transaction: Single Unit of work that is often made up of multiple db operations

# Example: Transfering 100 USD from Account A to Account B involves 5 operations

  1.  Create a transfer record with amount 100
  2.  Create an account entry for Account A with the amount - 100
  3.  Create an account entry for Account B with the amount + 100
  4.  Subtract 100 from the balance of Account A
  5.  Add 100 to the balance of Account B

# Why do use Db transaction?
  1.  To provide a reliable and consistent unit of work, even in case of system failure
  2.  To provide isolation between programs that access the database concurrently

# ACID PROPERTY
  1.  Atomicity: either all operations complete successfully or the transcation fails and the db is unchange.
  2.  Consistency: The db state must be valid after the trancation. All constraints must be met.
  3.  Isolation: Concurrent transaction must not affect each other.
  3. Durability: Data written by a successful transaction must be recorded in persistent storage.

  # How To Run SQL TX?
    BEGIN;                                    BEGIN;

    ...                                       ...

    COMMIT; (If transaction is successful)     ROLLBACK;(iF transaction is unsuccessful)