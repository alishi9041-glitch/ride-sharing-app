package models

type DRIVER_STATUS string

const (
	BUSY      DRIVER_STATUS = "busy"
	AVAILABLE DRIVER_STATUS = "available"
)

type RIDE_STATUS string

const (
	REQUESTED RIDE_STATUS = "requested"
	ACCEPTED  RIDE_STATUS = "accepted"
	STARTED   RIDE_STATUS = "started"
	COMPLETED RIDE_STATUS = "completed"
	CANCELLED RIDE_STATUS = "cancelled"
)

type PAYMENT_STATUS string

const (
	PENDING PAYMENT_STATUS = "pending"
	SUCCESS PAYMENT_STATUS = "success"
)
