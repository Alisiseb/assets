package clock

import "fmt"

// Define the Clock type here.
type Clock struct {
	hour int
	min  int
}

func New(h, m int) Clock {
	for m < 0 {
		m = 60 + m
		h--
	}
	for h < 0 {
		h += 24
	}
	for m > 59 {
		m -= 60
		h++
	}
	for h > 23 {
		h = h - 24
	}
	return Clock{h, m}
}

func (c Clock) Add(m int) Clock {
	min := c.min + m
	hour := c.hour
	for min > 59 {
		min -= 60
		hour++
	}
	for hour > 23 {
		hour = hour - 24
	}
	return Clock{hour, min}
}

func (c Clock) Subtract(m int) Clock {
	min := c.min - m
	hour := c.hour
	for min < 0 {
		min += 60
		hour--
	}
	for hour < 0 {
		hour = hour + 24
	}
	return Clock{hour, min}
}

func (c Clock) String() string {
	return fmt.Sprintf("%02v:%02v", c.hour, c.min)
}
