package json_rpc

import (
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
	"golang.org/x/time/rate"
)

// RateLimiter rate limits requests.
type RateLimiter struct {
	rate           float64
	burst          int
	clientLimiters map[string]*clientLimiter
	mu             sync.Mutex
}

type clientLimiter struct {
	lastReservation time.Time
	*rate.Limiter
}

// NewRateLimiter creates a new rate limiter.
func NewRateLimiter(rateN float64, burst int) *RateLimiter {
	if rateN <= 0 {
		log.Panic("non-positive rate limiter arg")
	}
	rl := &RateLimiter{
		rate:           rateN,
		burst:          burst,
		clientLimiters: make(map[string]*clientLimiter),
	}
	go rl.autoCleanup()
	return rl
}

// ExceedsLimit tries adding a request to the limiting channel and returns boolean to signal
// if we hit the rate limit.
func (rl *RateLimiter) ExceedsLimit(clientID string) bool {
	rl.mu.Lock()
	defer rl.mu.Unlock()
	limiter := rl.clientLimiters[clientID]
	if limiter == nil {
		limiter = &clientLimiter{Limiter: rate.NewLimiter(rate.Limit(rl.rate), rl.burst)}
		rl.clientLimiters[clientID] = limiter
	}
	limiter.lastReservation = time.Now()
	return !limiter.Allow()
}

// deallocate inactive limiters
func (rl *RateLimiter) autoCleanup() {
	ticker := time.NewTicker(time.Hour)
	for range ticker.C {
		rl.mu.Lock()
		for clientID, limiter := range rl.clientLimiters {
			if time.Since(limiter.lastReservation) > time.Minute*10 {
				delete(rl.clientLimiters, clientID)
			}
		}
		rl.mu.Unlock()
	}
}
