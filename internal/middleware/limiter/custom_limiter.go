package limiter

import (
	"github.com/gofiber/fiber/v2"
	"sync"
	"time"
)

type IPRateLimiter struct {
	ips map[string]*IPRateLimit
	mu  *sync.RWMutex
}

type IPRateLimit struct {
	count    int
	lastSeen time.Time
}

func NewIPRateLimiter() *IPRateLimiter {
	return &IPRateLimiter{
		ips: make(map[string]*IPRateLimit),
		mu:  &sync.RWMutex{},
	}
}

func (r *IPRateLimiter) EndpointLimiter(maxRequests int, duration time.Duration) fiber.Handler {
	return func(c *fiber.Ctx) error {
		ip := c.IP()
		r.mu.Lock()
		limiter, exists := r.ips[ip]

		if !exists {
			r.ips[ip] = &IPRateLimit{
				count:    1,
				lastSeen: time.Now(),
			}
			r.mu.Unlock()
			return c.Next()
		}

		if time.Since(limiter.lastSeen) > duration {
			limiter.count = 1
			limiter.lastSeen = time.Now()
			r.mu.Unlock()
			return c.Next()
		}

		if limiter.count >= maxRequests {
			r.mu.Unlock()
			return c.Status(fiber.StatusTooManyRequests).JSON(fiber.Map{
				"status":  "error",
				"message": "Too many requests. Please try again later",
			})
		}

		limiter.count++
		r.mu.Unlock()
		return c.Next()
	}
}
