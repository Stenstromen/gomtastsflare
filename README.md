# GoMTA-STSFlare

Go binary for creating/updating MTA-STS records on Cloudflare, and create the accompanying Nginx configuration.

<br>

## Generate Cloudflare API Token
1. Visit [https://dash.cloudflare.com/profile/api-tokens](https://dash.cloudflare.com/profile/api-tokens)
2. Create Token
3. "Edit Zone DNS" Template
4. "Zone Resources" Include > Specific Zone > example.com

## Installation via Homebrew (MacOS/Linux - x86_64/arm64)
```
brew install stenstromen/tap/gomtastsflare
```
## Download and Run Binary
* For **MacOS** and **Linux**: Checkout and download the latest binary from [Releases page](https://github.com/Stenstromen/gomtastsflare/releases/latest/)
* For **Windows**: Build the binary yourself.

## Build and Run Binary
```
go build
./gomtastsflare
```

## Example Usage
```
- Create MTA-STS DNS Records and Nginx Configuration
export TOKEN="# Cloudflare API TOKEN"
./gomtastsflare create -d example.com -4 127.0.0.1 -6 2001:0db8:cafe:0001 -m email.example.com -r report@example.com

- Update MTA-STS DNS Records and/or Nginx Configuration
export TOKEN="# Cloudflare API TOKEN"
./gomtastsflare update -d example.com -4 127.0.0.2 -r new_report_user@example.com

- Verify MTA-STS Setup for domain
./gomtastsflare verify -d example.com

Go binary for creating/updating MTA-STS records on Cloudflare, and create the accompanying Nginx configuration.

Usage:
  gomtastsflare [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  create      Create required DNS records and Nginx configuration
  help        Help about any command
  update      Update DNS Records and/or Nginx Configuration
  verify      Verify DNS Records and Web server Configuration

Flags:
  -h, --help   help for gomtastsflare

Use "gomtastsflare [command] --help" for more information about a command.
```

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