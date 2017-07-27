# Ban Chá! 

Ban chá starts a simple HTTP server that log any HTTP method in the output. 


Everytime I want to test some webhook or see what my API is sending/receiving to a external client 
I have to write a simple dummy server. So I decide to create that. 

Maybe can be useful for others. 

 Before all: 
 [Install go!](https://golang.org/dl/) 
 

 Open your terminal and type one of these commands:


`go get github.com/xild/bancha`
  
  To see options: 
  

  `bancha --help`
  
  Default mode
  
  `bancha`
  
  
  # NOTE 
  The project's propose is to log and inspect HTTP requestion. Differen than (https://docs.python.org/2/library/simplehttpserver.html)[Python Simple HTTP Server] propose
