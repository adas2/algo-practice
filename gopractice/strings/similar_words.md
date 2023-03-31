
1. Given a string s and a dictionary d {s1, s2, s3....sN}, find the list strings that are similar to the given string. Similarity, is determined by k insertaions, deletion or substrituitios that are needed to transfrom one string to the other. E.g. "stars" is similar to { "tars", "stores", "scars" } when k <=2. To simplify assume k=2, also, length of dictionary is the order of N and length of the given string is bounded by lenght n. Describe the complexity of your solution.


2. Assume this functionality is converted into a service where multiple clients requests need to simultaneously; suggests some ways we can optimize this algorithm for performace and response time. 

```// sample logic
SpellCheckService(ServiceRequest req, ServiceResponse resp){
	word := processRequest(req)
	candidateList := findSimilarCandidates(word)
	resp.EncodeIntoResponse(candidateList)
}
```

Hint: Is single synchronous process good enough? Can we have concurrency? Can we have caching?
Bonus: Do we have any race conditions in our approach? How do we resolve them?

