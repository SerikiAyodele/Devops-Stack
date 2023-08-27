# Understanding IPv4
- An important aspect of network communication id addressing, it's a way we can indetify devices on a network.
- An IPv4 address is made up of 32 bits, hence 2^32 possible addresses.
- Each address is broken into 4 sections referred to as octects because they are 8 bits each.
  Each octect can have a value from 0-255, again 2^8
- The IPv4 address itself represents two elements:
    - The network element and 
    - The host element
- network element tell what network the IP address resides on, the host element identifies the particular host on that network.
- In IPv4 we have three types of transmissions:
    - unicast: sends data from one sender to one specific recipient
    - multicast: from one sender to a specific group of recipients
    - broadcast: from one sender to all devices in a specific network segment

When planning your network, you're going to use one of the following addressing:
- Classful IPv4 addressing
- Classless IPv4 addressing

## Classful - Classless Network
- Classful network has a predefined number of bits allocated to the network element of the IP address, and therfore a pre-defined subnet mask.
| classes | range |purpose
| ----------- | ----------- | ----------- |
| A | 0.0.0.0 – 127.255.255.255 | |
| B | 128.0.0.0 – 192.255.255.255||
| C | 193.0.0.0 – 223.255.255.255||
| D | 224.0.0.0 – 239.255.255.255 |multicast transmission purposes|
| E | 240.0.0.0 – 255.255.255.255 | experimental purposes|

- Classless networks do not have boundaries of predefined network element.

## Understanding subnet masks
A subnet is a smaller network within a larger network, a subnet mask tells which which part of an IP 
address is anetwork element and which part is a host element
- To know if the source and destination Ip addresses are on the same network, their subnet values should be the same 

## CIDR
