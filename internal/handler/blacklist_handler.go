package handler

import (
	"sync"
	"time"
)

var (
	tokenBlacklist = make(map[string]time.Time)
	blacklistMutex sync.Mutex
)

// BlacklistToken adds the token to the blacklist until its expiration.
func BlacklistToken(token string) error {
	// For a real implementation, decode the token to get its expiration time.
	// Here we simply set it to expire in 72 hours for demonstration.
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()
	tokenBlacklist[token] = time.Now().Add(72 * time.Hour)
	return nil
}

// IsTokenBlacklisted checks if a token is blacklisted.
func IsTokenBlacklisted(token string) bool {
	blacklistMutex.Lock()
	defer blacklistMutex.Unlock()
	expiry, exists := tokenBlacklist[token]
	if !exists {
		return false
	}
	// Remove expired tokens.
	if time.Now().After(expiry) {
		delete(tokenBlacklist, token)
		return false
	}
	return true
}
