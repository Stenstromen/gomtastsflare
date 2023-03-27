# GoMTA-STSFlare

Go binary for creating/updating MTA-STS records on Cloudflare, and create the accompanying Nginx configuration.

<br>

# Random notes 

## Configuration Steps
1. Create DNS records
2. Create Nginx configuration

## DNS
```
mta-sts.example.com     A/AAAA      IPv4/IPv6
_mta-sts.example.com    TXT     	"v=STSv1; id=YYYYMMDDNN"
_smtp._tls.example.com  TXT         "v=TLSRPTv1; rua=mailto:mtastsreport@example.com"
```
## Nginx Configuration (mta-sts.example.com Server block)
```
	location ^~ /.well-known/mta-sts.txt {
		try_files $uri @mta-sts;
	}

	location @mta-sts {
		return 200 "version: STSv1
mode: enforce
max_age: 604800
mx: mail.example.com\r\n";
	}
```