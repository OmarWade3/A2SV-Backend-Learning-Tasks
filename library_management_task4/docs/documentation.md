# Library Management System – Concurrency Extension (Task 4)

This document explains the structure, workflow, and concurrency logic used in the extended **Concurrent Book Reservation System**.

---

## **1. Overview**

The system extends the original Library Management System with:

* Concurrent book reservations
* Goroutine-based reservation workers
* Channels for reservation queues
* Mutexes for thread-safe state updates
* Auto-cancellation of reservations (5 seconds timeout)

This ensures that when multiple users attempt to reserve the same book concurrently, the system remains consistent and avoids race conditions.

---

# **2. New Concurrency Features**

## 2.1 Reservation Queue

A global or injected **channel** is used to send reservation requests from the service layer to the concurrency worker:

```go
chan ReservationRequest
```

Each request contains:

```go
BookID
MemberID
```

This allows all reservation operations to be processed in a controlled, asynchronous manner.

---

## 2.2 Goroutine-Based Worker

Inside **concurrency/reservation_worker.go**, a reservation worker continuously listens to the channel:

* Reads a reservation request
* Starts a new goroutine per request
* Waits **5 seconds**
* Checks if the book is still reserved
* Cancels the reservation if the member didn’t borrow it

This prevents blocking the main program and allows many reservations to be processed simultaneously.

---

## 2.3 Mutex for Safe State Updates

The library uses a mutex:

```go
sync.Mutex
```

to protect:

* Book availability updates
* Reservation assignment
* Reservation cancellation
* Borrow operations

This eliminates race conditions when multiple goroutines modify the book map at the same time.

---

# **3. Data Model Updates**

## 3.1 Book Structure (Updated)

```go
type Book struct {
    ID         int
    Title      string
    Author     string
    Status     string  // "available", "reserved", "borrowed"
    ReservedBy int     // memberID of reserver; 0 = none
}
```

### Why this works:

* Using `ReservedBy` avoids string/int comparison issues
* Setting `ReservedBy = 0` means “not reserved”
* Ensures borrow logic can verify the correct member

---

## 3.2 Reservation Workflow

### Step 1 — Member requests reservation

Service verifies:

* Book exists
* Member exists
* Book is **available**

Then it marks:

```
Status = "reserved"
ReservedBy = memberID
```

Finally, it submits the reservation to the worker:

```go
reservationChannel <- ReservationRequest{bookID, memberID}
```

---

### Step 2 — Worker waits for 5 seconds

Each request launches a goroutine:

* `time.Sleep(5s)`
* Check if the book is still reserved
* Also check if it's reserved **by the same member**
* If not borrowed → automatically cancel

Cancellation logic:

```
Status = "available"
ReservedBy = 0
```

---

### Step 3 — Borrowing during the wait

If the member borrows the book before the timer ends:

* The book’s status becomes `"borrowed"`
* `ReservedBy` becomes `0`

The worker detects that the book is no longer reserved, therefore **won’t unreserve it**.

---

# **4. Borrowing Logic (Updated)**

The borrowing method now prevents:

 Borrowing a reserved book by another member
 Borrowing an unreserved but available book
 Borrowing your own reserved book

Borrow condition:

```go
if book.Status == "reserved" && book.ReservedBy != memberID {
    return errors.New("book is reserved by another member")
}
```

---

# **5. Folder Structure Explanation**

```
library_management/
├── main.go                     // initializes library, channel, worker, console UI
├── controllers/
│   └── library_controller.go   // handles console input/output logic
├── models/
│   └── book.go                 // Book struct + fields
│   └── member.go               // Member struct
├── services/
│   └── library_service.go      // business logic: add/remove/reserve/borrow
├── concurrency/
│   └── reservation_worker.go   // goroutine worker handling async reservations
├── docs/
│   └── documentation.md        // this file
└── go.mod
```

---

# **6. Error Handling Guarantees**

The system prevents:

 Double-reserving the same book
 Borrowing a book reserved by another member
 Race conditions caused by concurrent Goroutines
 Stale reservations (auto-cancels after 5 seconds)

---

# **7. Summary of Concurrency Approach**

* **Mutex** ensures thread-safe book updates
* **Channels** serialize reservation requests
* **Goroutines** allow high concurrency without blocking
* **Timers** enforce reservation expiration
* **ReservedBy** ensures correct borrowing logic

This architecture supports clean, scalable concurrent behavior.

---