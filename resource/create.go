package resource

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

func ResourceCreate(cmd *cobra.Command, args []string) error {
	currentTime := time.Now()
	domain, err := cmd.Flags().GetString("domain")
	if err != nil {
		return err
	}

	mx, err := cmd.Flags().GetString("mx")
	if err != nil {
		return err
	}

	ipv4, err := cmd.Flags().GetString("ipv4")
	if err != nil {
		return err
	}
	ipv6, err := cmd.Flags().GetString("ipv6")
	if err != nil {
		return err
	}

	rua, err := cmd.Flags().GetString("rua")
	if err != nil {
		return err
	}

	postToCloudflare(domain, "mta-sts A", genCloudflareReq("A", "mta-sts", ipv4, "Created"))

	if ipv6 != "" {
		postToCloudflare(domain, "mta-sts AAAA", genCloudflareReq("AAAA", "mta-sts", ipv6, "Created"))
	}

	postToCloudflare(domain, "_mta-sts TXT", genCloudflareReq("TXT", "_mta-sts", "v=STSv1; "+"id="+currentTime.Format("20060102")+"0"+strconv.Itoa(rand.Intn(20)), "Created"))
	postToCloudflare(domain, "_mta-sts._tls", genCloudflareReq("TXT", "_mta-sts._tls", "v=TLSRPTv1; rua=mailto:"+rua, "Created"))
	genNginxConf(domain, mx)

	return nil
}

func postToCloudflare(nameanddomain string, recordtype string, postBody string) {
	url := "https://api.cloudflare.com/client/v4/zones"
	var bearer = "Bearer " + os.Getenv("TOKEN")
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	req.Header.Add("Authorization", bearer)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error while reading the response bytes:", err)
	}
	var res Res
	if err := json.Unmarshal(body, &res); err != nil {
		log.Println(err)
		os.Exit(1)
	}

	var did string

	for i := range res.Result {
		if nameanddomain == res.Result[i].Name {
			did = res.Result[i].ID
		}
	}

	posturl := "https://api.cloudflare.com/client/v4/zones/" + did + "/dns_records"
	var jsonStr = []byte(postBody)
	req2, err2 := http.NewRequest("POST", posturl, bytes.NewBuffer(jsonStr))
	if err2 != nil {
		log.Println(err2)
		os.Exit(1)
	}

	req2.Header.Set("Content-Type", "application/json")
	req2.Header.Add("Authorization", bearer)

	client2 := &http.Client{}
	resp2, err2 := client2.Do(req2)
	if resp2.StatusCode != 200 {
		log.Println("Error while posting to Cloudflare: Invalid TOKEN och Record Already Exists")
		os.Exit(1)
	}
	if err2 != nil {
		log.Println(err2)
		os.Exit(1)
	}

	defer resp.Body.Close()

	log.Println("Cloudflare Response Status for: "+recordtype, resp2.Status)
}
