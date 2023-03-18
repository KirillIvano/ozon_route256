package rate_limiter

import (
	"sync"
	"time"
)

type RateLimiter struct {
	mx            *sync.Mutex
	lastTime      *time.Time
	expectedDelay time.Duration
}

func (r *RateLimiter) updateLastTime(newTime time.Time) {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.lastTime = &newTime
}

func (r *RateLimiter) Take() time.Time {
	// первый запрос, не ждем ответа
	if r.lastTime == nil {
		r.updateLastTime(time.Now())
		return time.Time{}
	}

	currentTime := time.Now()
	diff := r.lastTime.Sub(currentTime)

	// не ждем запрос, если предыдущий закончился в прошлом
	if diff < 0 {
		diff = 0
	}

	// выставляем новое время для последующих запросов
	r.updateLastTime(currentTime.Add(diff + r.expectedDelay))
	// ждем, пока пройдет достаточно времени для нашего запроса
	time.Sleep(diff)

	return time.Now()
}

func New(rate int32) *RateLimiter {
	return &RateLimiter{
		expectedDelay: time.Second / time.Duration(rate),
		mx:            &sync.Mutex{},
		lastTime:      nil,
	}
}
