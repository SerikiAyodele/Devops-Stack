IPv6 addresses are 16 bytes long (128 bits) and are made up of hexadecimal numbers.

## The Format Of an IPv6 Address
- IPv6 addresses are formed out of 8 groups of hexaecimal numbers, with each
   group seperated by a column

- i.e `2001:0034:09FA:F3B2:20E4:1030:0001:4BC2`

- CIDR stands for classless inter-domain routing, which is a method used in IP      addressing to allocate and manage IP addresses more effiicently.

- IPv6 introduces:
    - Anycast: sends data to the nearest device pout of a group of devices

## Understanding Address Types
- Global-unicast Address: Global unicast addresses all begin with `2000::/3`. They
  are similar to public Ip addresses in IPv4, in that it is routable across the internet.

- Link-local addresses: It is similar to an IPv4 APIPA(automatic private IP         addressing) address, they are not designed to be routable accross the internet.
    - They all begin with  `FE80::/10`

- Unique local addresses: Unique local addresses are the IPv6 equivalent of private IPv4 addresses and begin with `FC00::/7`. Like their IPv4 counterparts, these are not routable across the internet.

- Multicast addresses: Like their IPv4 counterpart, IPv6 multicast addresses are used to transmit data to devices within a specified group. Multicast IPv6 addresses begin with `FF00::/8`.

## Assigning IPv6 Addresses
- Like IPv4, IPv6 addresses can be assigned manually or dynamically. However, to obtain an IP address automatically, we can use either stateless auto-configuration or DHCPv6.
#### ICMPv6
ICMP is the underlying protocol that provides us with toools such as ping, tracert
and pathping.

## Understanding Interoperability With IPv4
 Because IPv6 has not yet fully replaced IPv4, there needs to be a means of using the two protocols together to allow communication to take place.

**Dual stack:** is a means for a network interface card to support and process IPv4 and IPv6 traffic simultaneously.

**6to4:** Allows devices on seperate IPv6 networks to communicate with each other 
over and IPv4 network.
    -  6to4 addresses always begin with the prefix 2002::/16.
    -  It does not work for devices that are sitting behind a
       network using NAT.

**Teredo tunneling:** Teredo tunneling is very similar to 6to4 but allows IPv6 communication between devices sitting on a network that use NAT.