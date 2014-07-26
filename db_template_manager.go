package helper

import (
	"log"
)

var (
	/**
	  gobbs template -> timestamp
	  timestamp is get from db when user add or update the template in db.
	*/
	CacheMapGobbsDbTemplates map[string]string
)

func init() {
	CacheMapGobbsDbTemplates = make(map[string]string)
}

func GetTemplateTimestamp(template_file_path string) (timestamp string) {
	if timestamp, ok := CacheMapGobbsDbTemplates[template_file_path]; ok {
		if 0 == len(timestamp) {
			log.Printf("Error Error Error when GetTemplateTimestamp , can not find template timestamp")
		}
	} else {
		log.Printf("Error Error Error when GetTemplateTimestamp , can not find template timestamp")
	}
	return
}
func SetTemplateTimestamp(template_file_path, timestamp string) {
	CacheMapGobbsDbTemplates[template_file_path] = timestamp
}
