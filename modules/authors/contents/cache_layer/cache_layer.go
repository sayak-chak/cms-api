package cache_layer

import (
	"cms-api/config"
	models "cms-api/models/responses"
	"sync"
	"time"
)

var cacheMap = map[string]*cachedData{}
var mapLock = sync.RWMutex{}

type cachedData struct {
	lastCacheTime  int
	cachedResponse *[]models.TopContentsResponse
}

func GetCachedResponseIfPossible(tag string) *[]models.TopContentsResponse {
	mapLock.RLock()
	defer mapLock.RUnlock()
	if !canSendCachedResponse(tag) {
		return nil
	}
	if cacheEntry, isPresent := cacheMap[tag]; isPresent {
		return cacheEntry.cachedResponse
	}
	return nil

}

func CacheThis(topArticles *[]models.TopContentsResponse, tag string) {
	mapLock.Lock()
	defer mapLock.Unlock()
	cacheMap[tag] = &cachedData{
		lastCacheTime:  int(time.Now().Unix()),
		cachedResponse: topArticles,
	}
}

func canSendCachedResponse(tag string) bool {
	return cacheMap != nil && !cacheHasExpired(tag)
}

func cacheHasExpired(tag string) bool {
	cacheEntry, isPresent := cacheMap[tag]
	if !isPresent {
		return false
	}
	return time.Now().Unix()-int64(cacheEntry.lastCacheTime) > config.CachePeriod
}
