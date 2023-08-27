# Understanding Name Resolution
We have four means of performing name resolution:
- DNS
- Host files
- Windows Internet Name Service (WINS)
- LMHOSTS files
Think of DHCP as issuing an IP address and DNS as resolving an IP address.
DNS provides an hierachical means of resolving an host name to an IP address - more
specifically, a fi=ully qualified domain name (FQDN).

## Fully Qualified Domain Name (FQDN)
- If we have `https://www.packtpub.com` everything after https here is the FQDN

WWW- host
packtpub- domain
com- TLD

The host is the hostname of the particular device; however, it may also be the Canonical Name (CNAME) linked to the host.

The domain element refers to the domain that the organization has registered with a domain name registrar. If the domain is restricted to internal use we don't have to register it.

The `.com` element is the top level domain.

For DNS to work, the servers need to be aware of the IP address that is linked to the FQDN.

## DNS records
DNS resolves an IP from a given FQDN. A DNS server can obtain records in one of three main ways:
- Manual entry
- Dynamic entry
- Zone transfer