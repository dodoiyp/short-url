# URL Shortener
## Requirement
This project is for dacard pre-interview homework 
  
Design and implement (with unit tests) an URL shortener using Go programming language.  
Criteria:  
 1. URL shortener has 2 APIs, please follow API example to implement:  
   a. A RESTful API to upload a URL with its expired date and response with a shorten URL.  
   b. b.An API to serve shorten URLs responded by upload API, and redirect to original URL. If URL is expired,please response with status 404.
 1. Please feel free to use any external libs if needed.  
 1. It is also free to use following external storage including:  
   a. Relational database (MySQL, PostgreSQL, SQLite).  
   b. Cache storage (Redis, Memcached). 
 1. Please implement reasonable constrains and error handling of these 2 APIs.
 1. You do not need to consider auth.
 1. Many clients might access shorten URL simultaneously or try to access with non-existent shorten URL, please take. 
performance into account.  

API example :
```
# Upload URL API
curl -X POST -H "Content-Type:application/json" http://localhost/api/v1/urls -d '{
   "url": "<original_url>",
"expireAt": "2021-02-08T09:20:41Z"
}'
# Response
{
   "id": "<url_id>",
   "shortUrl": "http://localhost/<url_id>"
}
# ------------------
# Redirect URL API
curl -L -X GET http://localhost/<url_id> => REDIRECT to original URL
```
Constraints:  
1. QPS:  
  a. Write:500.  
  b. Read:100k. 
1. url max length:none. 
1. Is shorturl need to be analysis :no.
1. shorturl expire time limit : none.  
1. Availability>Consistency


