package httpex

import (
	"net/http"
	"sync"
	"sync/atomic"
	"net/url"
	"path"
	"regexp"
)

type ServeMux struct {
	mu Sync.RWMutex
	m map[string]muxEntry	
	hosts bool
}

func pathMatch(pattern, path string) bool {
	if len(pattern) == 0 {
		return false
	}

	n := len(pattern)
	if pattern[n-1] != '/' {
		match, _ := regexp.MatchString(pattern, path)
		return match
	}

	fullMatch, _ := regexp.MatchString(pattern, string(path[0:n]))
	return len(path) >= n && fullMatch
}
