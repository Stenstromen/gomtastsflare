package resource

import (
	"fmt"
)

func genNginxConf(domain string, mx string) {
	nginxConf := "location ^~ /.well-known/mta-sts.txt {\n" +
		"try_files $uri @mta-sts;\n" +
		"}\n" +
		"location @mta-sts {\n" +
		"return 200 \"version: STSv1\n" +
		"mode: enforce\n" +
		"max_age: 604800\n" +
		"mx: " + mx + "\n\";\n" +
		"}" + "\n"
	fmt.Print("\n" + "== Put this inside the server block for mta-sts." + domain + "\n\n" + nginxConf + "\n")
}
