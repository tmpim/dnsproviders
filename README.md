DNS providers for Casket v1
==========================

**⚠️ For Casket 1 only, which is obsoleted by Casket 2.**

These providers can be used to help solve the ACME DNS challenge by plugging them into Casket 0.9 and newer:

```go
import _ "github.com/casketserver/dnsproviders/<provider>"
```

You can then use this in your Casketfile with the `tls` directive like so:

```plain
tls {
	dns <provider>
}
```

Credentials for your DNS provider should be set in environment variables. This information is in the [Automatic HTTPS](https://casketserver.com/docs/automatic-https#dns-challenge) page of the Casket documentation. For more information about using your DNS provider, see [the docs for your provider](https://godoc.org/github.com/go-acme/lego/providers/dns) directly.

If you specify a DNS provider, the DNS challenge will be used exclusively; other challenge types will be disabled. Be aware that some DNS providers may be slow in applying changes.


## About these packages

Casket 0.9 and newer supports solving the ACME DNS challenge. This challenge is unique because the server that is requesting a TLS certificate does not need to start a listener and be accessible from external networks. This quality is essential when behind load balancers or in other advanced networking scenarios.

The DNS challenge sets a DNS record and the ACME server verifies its correctness in order to issue the certificate. Casket can do this for you automatically, but it needs credentials to your DNS provider to do so. Since every DNS provider is different, we have these adapters you can plug into Casket in order to complete this challenge.

The underlying logic that actually solves the challenge is implemented in a different package not far away from here. Casket uses [go-acme/lego](https://github.com/go-acme/lego), a library originally written for use in Casket, to solve ACME challenges. If you wish to add a new provider, see the documentation for that library and write your own provider implementation. Then writing the adapter for Casket is very easy: just copy+paste any of these existing ones, replace the names and tweak a few things, and submit a pull request. Done!
