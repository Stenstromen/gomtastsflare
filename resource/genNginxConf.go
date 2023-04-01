package resource

import (
	"fmt"
	"strings"
)

func genNginxConf(domain string, mx string) {
	var mxs string
	mxlist := strings.Split(mx, ",")

	for i := range mxlist {
		mxs = mxs + "mx: " + mxlist[i] + "\n"
	}

	nginxConf := "location ^~ /.well-known/mta-sts.txt {\n" +
		"try_files $uri @mta-sts;\n" +
		"}\n\n" +
		"location @mta-sts {\n" +
		"return 200 \"version: STSv1\n" +
		"mode: enforce\n" +
		"max_age: 604800\n" +
		mxs +
		"\";\n" +
		"}" + "\n"
	fmt.Print("\n" + "== Put this inside the server block for mta-sts." + domain + "\n\n" + nginxConf + "\n")
}
