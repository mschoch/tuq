# tuq (Tool for Unstructured Querying)

A tool for querying unstructured datasources.  Run queries like:

```
tuq> SELECT {"style":doc.style, "count":count(doc.style), "avg_abv": avg(doc.abv)} 
FROM beersample WHERE doc.style.analyzed == "lager" && doc.abv > 0 GROUP BY doc.style
```

## Features
* UNQL-like query language
* Interactive query shell with readline-like support and command history
* HTTP mode
* DataSources
  * CSV file
  * ElasticSearch
  * Couchbase+ElasticSearch
  * MongoDB
* Baseline support for all operations in memory (allows working with databases that have limited query capability)
* Pluggable architecture (for easier experimentation)
  * Parser
  * Planner
  * Optimizer
  * Datasources

## Getting Started

See the wiki for full details:

https://github.com/mschoch/tuq/wiki
