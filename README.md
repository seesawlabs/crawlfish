# crawlfish

`POST {{url}}/api/v1/crawl`

Request payload

```json
{
    "url" : "http://0value.com",
    "words": "Doer \n struct"
}
```

Response example
```json
{
  "words_found": {
    "Doer": [
      {
        "ulr": "http://0value.com",
        "time_elapsed": 0.00000331800000000000,
        "word_count": 13
      },
      {
        "ulr": "http://0value.com/A-PEG-parser-generator-for-Go",
        "time_elapsed": 0.00000249000000000000,
        "word_count": 4
      },
      {
        "ulr": "http://0value.com/Let-the-Doer-Do-it",
        "time_elapsed": 0.000004889,
        "word_count": 13
      }
    ],
    "struct": [
      {
        "ulr": "http://0value.com",
        "time_elapsed": 0.000023339,
        "word_count": 1
      },
      {
        "ulr": "http://0value.com/A-PEG-parser-generator-for-Go",
        "time_elapsed": 0.000025394,
        "word_count": 2
      },
      {
        "ulr": "http://0value.com/Let-the-Doer-Do-it",
        "time_elapsed": 0.000032913,
        "word_count": 1
      }
    ]
  },
  "pages_searched": [
    {
      "ulr": "http://0value.com",
      "time_elapsed": 0.00005308300000000001
    },
    {
      "ulr": "http://0value.com/A-PEG-parser-generator-for-Go",
      "time_elapsed": 0.000061708
    },
    {
      "ulr": "http://0value.com/Let-the-Doer-Do-it",
      "time_elapsed": 0.000101397
    }
  ],
  "pages_found": [
    {
      "ulr": "http://0value.com",
      "time_elapsed": 0.00005308300000000001
    },
    {
      "ulr": "http://0value.com/A-PEG-parser-generator-for-Go",
      "time_elapsed": 0.000061708
    },
    {
      "ulr": "http://0value.com/Let-the-Doer-Do-it",
      "time_elapsed": 0.000101397
    }
  ],
  "total_time_taken": 3.734993063,
  "percentage_of_found": 100
}

```