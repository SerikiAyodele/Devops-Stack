## logical port
- A logical port is one that is defined in software
- Both TCP and UDP headers contain a source and destination port number, 
  when the data is recived by the recipient device, the destination port
  is inspected and the device will pass the data to the relevant application
  or service
- There are 65,535 logical ports available

common ports and their numbers
| port number | service |
| ----------- | ----------- |
| HTTP | 80 |
| HTTPS | 443 |
| SSH | 22 |
| DNS | 53 |
|Telnet | 23 |

#### TCP 
- transmission user protocol is referred to as a connection oriented protocol
  beacsue before and data can be transferred between two devices, a connection
  has to be established.
- To establish this connection, TCP carries out a process known as the three way handshake
- TCP guarantees delivery of data through a process of using sequence numbers and acknowledgement
- This comes at a price, the TCP header adds 20-60 bytes of data per segment sent, this can make
  the transfer of data slow.


#### UDP
- Is connectionless
- The sending device literally sends the data out on a wire and hopes it is recived by
  the destination device
- It is faster because it has a smaller header and there's no acknowledgement to add to
  the header

## Layer3- Network Layer
- The network layer of the OSI model is responsible for the logical addressing of 
  devices through the use of IP addresses
- The most common protocol in the network layer is the IP(internet protocol)

## Layer2- Data-Link Layer
- The protocol data unit in layer 2 is the FRAME
- The responsibilities of the data link layer includes:
    - Placing the data unto a physical media
    - Error notifications and
    - Flow control
- This layer is split into:
    - LLC (logical link control) - Identifies what network layer protocol is being used
    - MAC (media access control) - MAC addresses are unique identifiers associated with network interface cards (NICs) of devices

## Layer1- Physical Layer
- The transmission of data in the form of bits takes place in this layer.
- At this layer, there are no protocols per second, but there are sets of standards and criteria
  that the cabling and network cards will need to adhere to. These standards include the
  following:
    - Voltages
    - Speeds
    - Wiring