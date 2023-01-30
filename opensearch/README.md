# Opensearch
## URL
```
http://<opensearch_url>/<index>/_search?q=<field>:<value>&<field>:<value>
```

## Dashboard
### Sorting
```
GET search_sb_channel/_search
{
  "sort": [
    {
      "updated_at": {
        "order": "desc"
      }
    }
  ], 
  "_source": ["updated_at"]
}
```


### Counting
```
GET search_sb_channel/_count
```

