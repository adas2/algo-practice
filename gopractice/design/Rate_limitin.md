Rate limter problem
----

Problem statement (EPI):

There are n web crawlers which is supposed to download ther web (millions of URLs). Each URL is mapped into using a hascode to a server which is responsible for downloading it. For example for URL x the server m = hash(x)%n is responsible for downlaoding it. Assuming the craw job eats up the bandwidth of the server it is downlaoding from, design a rate limitting system with at most b bandwidth downloaded per 1 sec.


