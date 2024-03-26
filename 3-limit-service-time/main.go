//////////////////////////////////////////////////////////////////////
//
// Your video processing service has a freemium model. Everyone has 10
// sec of free processing time on your service. After that, the
// service will kill your process, unless you are a paid premium user.
//
// Beginner Level: 10s max per request
// Advanced Level: 10s max per user (accumulated)
//

package main

import (
	"sync"
	"time"
)

// User defines the UserModel. Use this to check whether a User is a
// Premium user or not
type User struct {
	ID        int
	IsPremium bool
	TimeUsed  int64 // in seconds
	mx        sync.Mutex
}

// HandleRequest runs the processes requested by users. Returns false
// if process had to be killed
func HandleRequest(process func(), u *User) bool {
	n := time.Now()

	process()
	d := time.Since(n).Seconds()
	u.mx.Lock()
	defer u.mx.Unlock()
	u.TimeUsed += int64(d)
	if !u.IsPremium && u.TimeUsed > 10 {
		return false
	}
	return true
}

func main() {
	RunMockServer()
}
