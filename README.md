# 🚗 SwiftRide – Ride Sharing Backend Assignment

## 📌 Overview

SwiftRide is a backend system for a ride-sharing platform designed to connect riders with nearby drivers efficiently and safely.

This assignment focuses on building a scalable, concurrency-safe backend service that powers:

- Ride requests
- Driver matching
- Real-time ride tracking
- Fare calculation
- Ride lifecycle management

**Tagline:** *Move smarter. Move faster.*

---

## 🚨 Problem Statement

Modern ride-sharing platforms face the following challenges:

- Long driver allocation times
- Driver double allocation during peak traffic
- No real-time ride tracking
- Unclear surge pricing
- Unreliable payment processing
- Invalid ride state transitions

SwiftRide aims to solve these using clean architecture and robust backend design.

---

## 🎯 Business Goals

- 5-second ride matching SLA (Service Level Agreement)
- Real-time driver tracking
- Transparent fare calculation
- Concurrency-safe driver allocation
- Scalable architecture for high traffic

---

## 👥 Sample Data

### Riders

| RiderId   | FirstName | LastName |
|------------|------------|------------|
| rider101  | John       | Doe        |

### Drivers

| DriverId   | FirstName | LastName | Status     |
|------------|------------|------------|------------|
| driver101 | Alice      | Kumar      | AVAILABLE  |
| driver102 | Rahul      | Sharma     | BUSY       |

### Vehicles

| VehicleId  | Type   | DriverRefId |
|------------|--------|-------------|
| vehicle101 | Sedan  | driver101   |
| vehicle102 | SUV    | driver102   |

---

## 🌐 API Specification

Base URL:

http://localhost:8080


#### Request Body

```json
{
  "riderId": "rider101",
  "pickup": {
    "latitude": 12.9716,
    "longitude": 77.5946
  },
  "drop": {
    "latitude": 12.9352,
    "longitude": 77.6245
  },
  "vehicleType": "Sedan"
}

#### Response

{
  "rideId": "ride123",
  "status": "REQUESTED",
  "estimatedFare": 220.5,
  "estimatedArrivalTimeInMinutes": 6
}


POST /ride/accept

{
  "rideId": "ride123",
  "driverId": "driver101"
}
{
  "rideId": "ride123",
  "status": "ACCEPTED"
}



POST /ride/start
POST /ride/complete
{
  "rideId": "ride123",
  "status": "COMPLETED",
  "finalFare": 235.75,
  "paymentStatus": "SUCCESS"
}


GET /ride/track?rideId=ride123
{
  "rideId": "ride123",
  "driverLocation": {
    "latitude": 12.9701,
    "longitude": 77.5950
  },
  "etaMinutes": 4
}


POST /driver/location
POST /driver/status


Supported ride states:
REQUESTED → ACCEPTED → STARTED → COMPLETED
                    ↘
                   CANCELLED


Matching Strategy

Design a pluggable matching strategy interface.

Support:

Nearest driver

Highest rating

Lowest ETA

Surge-zone priority



Fare Strategy

Support:

Base fare

Per kilometer pricing

Surge pricing

Night pricing

Discount handling


Concurrency Handling (Critical)

If two riders request simultaneously and only one driver is available:

The driver must not be assigned twice.

Driver allocation must be atomic.

Demonstrate thread safety.

5️⃣ Scalability Considerations

Design for:

High concurrent ride requests

Real-time driver location updates

Low-latency driver matching



🛠 Suggested Development Commands
Command	Description
make run	Run application locally
make test	Run unit tests
make tidy	Clean dependencies


Bonus Challenges

Add ride cancellation fee

Add driver earnings wallet

Add rider/driver rating system

Add surge zones using geo-hashing

Add idempotent payment handling

Convert to event-driven architecture




Notes -
Missing payment 
Data is currently global -> ugh sucks!
When to take string/int and when to take pointer to string/int?
Need to add testcases
