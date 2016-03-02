# crawlfish

`POST {{url}}/api/v1/crawl`

Request payload

```json
{
    "url" : "http://0value.com",
    "words": "Doer \n struct"
}
```

User will receive immediate response with empty body and HTTP status 200 


For dev env crawling results will be pushed to https://crawlfish-dev.firebaseio.com