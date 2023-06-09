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

func ResourceUpdate(cmd *cobra.Command, args []string) error {
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

	if ipv4 != "" {
		putToCloudflare("mta-sts."+domain, "mta-sts A", genCloudflareReq("A", "mta-sts", ipv4, "Updated"))
	}

	if ipv6 != "" {
		putToCloudflare("mta-sts."+domain, "mta-sts AAAA", genCloudflareReq("AAAA", "mta-sts", ipv6, "Updated"))
	}

	if rua != "" {
		putToCloudflare("_mta-sts._tls."+domain, "_mta-sts._tls", genCloudflareReq("TXT", "_mta-sts._tls", "v=TLSRPTv1; rua=mailto:"+rua, "Updated"))
	}

	if ipv4 != "" || ipv6 != "" || rua != "" || mx != "" {
		putToCloudflare("_mta-sts."+domain, "_mta-sts TXT", genCloudflareReq("TXT", "_mta-sts", "v=STSv1; "+"id="+currentTime.Format("20060102")+"0"+strconv.Itoa(rand.Intn(20)), "Updated"))
	}

	if mx != "" {
		genNginxConf(domain, mx)
	}

	return nil
}

func putToCloudflare(nameanddomain string, recordtype string, putBody string) {
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

	searchurl := "https://api.cloudflare.com/client/v4/zones/" + res.Result[0].ID + "/dns_records"
	req2, err2 := http.NewRequest("GET", searchurl, nil)
	if err2 != nil {
		log.Println(err2)
		os.Exit(1)
	}

	req2.Header.Add("Authorization", bearer)
	client2 := &http.Client{}
	resp2, err2 := client2.Do(req2)
	if err2 != nil {
		log.Println("Error on response.\n[ERROR] -", err)
	}
	defer resp2.Body.Close()
	body2, err2 := io.ReadAll(resp2.Body)
	if err2 != nil {
		log.Println("Error while reading the response bytes:", err)
	}

	var recordsres RecordsRes
	if err2 := json.Unmarshal(body2, &recordsres); err2 != nil {
		log.Println(err2)
		os.Exit(1)
	}

	var did string

	for i := range recordsres.Result {
		if nameanddomain == recordsres.Result[i].Name {
			did = recordsres.Result[i].ID
		}
	}

	puturl := "https://api.cloudflare.com/client/v4/zones/" + res.Result[0].ID + "/dns_records/" + did

	var jsonStr = []byte(putBody)
	req3, err3 := http.NewRequest("PATCH", puturl, bytes.NewBuffer(jsonStr))
	if err3 != nil {
		log.Println(err3)
		os.Exit(1)
	}

	req3.Header.Set("Content-Type", "application/json")
	req3.Header.Add("Authorization", bearer)

	client3 := &http.Client{}
	resp3, err3 := client3.Do(req3)
	if err3 != nil {
		log.Println(err3)
		os.Exit(1)
	}
	defer resp.Body.Close()

	log.Println("Cloudflare Response Status for: "+recordtype, resp3.Status)
}
