package resource

import (
	"context"
	"fmt"

	"io"
	"log"
	"net"
	"net/http"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func ResourceVerify(cmd *cobra.Command, args []string) error {
	var successCount int

	domain, err := cmd.Flags().GetString("domain")
	if err != nil {
		return err
	}

	resp, err := http.Get("https://mta-sts." + domain + "/.well-known/mta-sts.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	resptext, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	if cloudflareDNSLookupIP("mta-sts."+domain, false) != "" {
		fmt.Println("OK - mta-sts A record exists")
		if cloudflareDNSLookupIP("mta-sts."+domain, true) != "" {
			fmt.Println("OK - mta-sts AAAA record exists")
		} else {
			fmt.Println("INFO - mta-sts AAAA record does not exist")
		}
		successCount++
	} else {
		println("FAIL - No mta-sts A/AAAA record exists")
	}

	if strings.Contains(cloudflareDNSLookupTXT("_mta-sts."+domain), "v=STSv1;") {
		fmt.Println("OK - STSv1 TXT record exists")
		successCount++
	} else {
		fmt.Println("FAIL - No STSv1 TXT record exists")
	}

	if strings.Contains(cloudflareDNSLookupTXT("_smtp._tls."+domain), "v=TLSRPTv1;") {
		fmt.Println("OK - TLSRPTv1 TXT record exists")
		successCount++
	} else {
		fmt.Println("FAIL - No TLSRPTv1 TXT record exists")
	}

	if strings.Contains(string(resptext), "version: STSv1") {
		fmt.Println("OK - " + "mta-sts." + domain + "/.well-known/mta-sts.txt" + " exists")
		successCount++
	} else {
		fmt.Println("FAIL - No mta-sts.txt exists")
	}

	fmt.Println("Success Count: " + fmt.Sprint(successCount) + "/4")

	return nil
}

func cloudflareDNSLookupIP(domainrecord string, ipv6 bool) string {
	resolve := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			dial := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return dial.DialContext(ctx, "udp", "1.1.1.1:53")
		},
	}

	ipresult, _ := resolve.LookupHost(context.Background(), domainrecord)

	if ipv6 {
		return ipresult[1]
	} else {
		return ipresult[0]
	}
}

func cloudflareDNSLookupTXT(domainrecord string) string {
	resolve := &net.Resolver{
		PreferGo: true,
		Dial: func(ctx context.Context, network, address string) (net.Conn, error) {
			dial := net.Dialer{
				Timeout: time.Millisecond * time.Duration(10000),
			}
			return dial.DialContext(ctx, "udp", "1.1.1.1:53")
		},
	}

	txtresult, _ := resolve.LookupTXT(context.Background(), domainrecord)

	return txtresult[0]
}
