package resource

import (
	"encoding/json"
	"log"
	"os"
	"time"
)

func genCloudflareReq(recordtype string, subdomain string, content string, cu string) string {
	currentTime := time.Now()

	jsonRequest := JSONRequest{
		Type:     recordtype,
		Name:     subdomain,
		Content:  content,
		Ttl:      3600,
		Priority: 10,
		Proxied:  false,
		Comment:  cu + " by GoMTA-STSFlare - " + currentTime.Format("2006-01-02 15:04:05"),
	}

	byteArray, err := json.MarshalIndent(jsonRequest, "", "  ")

	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return string(byteArray)
}
