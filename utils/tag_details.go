package utils

import (
	models "cms-api/models/database"

	"strings"
	"sync"
)

var tagNamesMap map[string]string = map[string]string{"action": "actions", "adventure": "adventures", "horror": "horrors", "drama": "dramas"}

var lock = sync.RWMutex{}

func GetTagDBNameIfValid(inputTag string) (string, bool) {
	tag := strings.ToLower(inputTag)
	if tagDBName, isValid := tagNamesMap[tag]; !isValid {
		return "", false
	} else {
		return tagDBName, true
	}
}

func GetTagTableModelFor(tagTable string, contentId int) interface{} {

	if strings.EqualFold("actions", tagTable) {
		return &models.Action{
			ContentId: contentId,
		}
	}

	if strings.EqualFold("adventures", tagTable) {
		return &models.Adventure{
			ContentId: contentId,
		}
	}

	if strings.EqualFold("horrors", tagTable) {
		return &models.Horror{
			ContentId: contentId,
		}
	}

	if strings.EqualFold("dramas", tagTable) {
		return &models.Drama{
			ContentId: contentId,
		}
	}
	return nil
}
