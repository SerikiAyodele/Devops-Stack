Vertical scaling vs horizontal scaling
--------------------------------------

 Vertical scaling, referred to as “scale up”, means the process of adding more power (CPU, RAM, etc.) to your servers. Horizontal scaling, referred to as “scale-out”, allows you to scale
 by adding more servers into your pool of resources. 
When traffic is low, vertical scaling is a great option, and the simplicity of vertical scaling is its main advantage. Unfortunately, it comes with serious limitations.

  • Vertical scaling has a hard limit. It is impossible to add unlimited CPU and memory to a single server. 
  • Vertical scaling does not have failover and redundancy. If one server goes down, the website/app goes down with it completely. 

Horizontal scaling is more desirable for large scale applications due to the limitations of vertical scaling.
In the previous design, users are connected to the web server directly. Users will unable to access the website
if the web server is offline. In another scenario, if many users access the web server simultaneously and it
reaches the web server’s load limit, users generally experience slower response or fail to connect to the server.
A load balancer is the best technique to address these problems.
