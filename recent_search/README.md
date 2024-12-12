## Implement recent search
Capture search logs and make them queryable
From an HTTP request, extract all possible meta info
Ingest them in ES
Plot different graphs, segmenations, and gain insights using Kibana

When a user performs a search on a website or application, the system captures important metadata from the search request (e.g., search terms, timestamp, user agent). This data is then:

Ingested into Elasticsearch for efficient searching and analysis.
Stored temporarily in Redis for fast access.
Sent through Kafka to populate MongoDB for long-term storage.
Visualized using Kibana to create useful insights, such as search trends, popular queries, and segmentation analysis.


