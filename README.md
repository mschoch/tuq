# tuq (Tool for Unstructured Querying)

A tool for querying unstructured datasources.

## Features
* UNQL-like query language
* Interactive query shell with readline-like support and command history
* HTTP mode
* DataSources
  * CSV file
  * Couchbase+ElasticSearch
* Baseline support for all operations in memory (allows working with databases that have limitted query capability)
* Pluggable architecture (for easier experimentation)
  * Parser
  * Planner
  * Optimizer
  * Datasources

### Query Language



### Parser

The parser is responsible for converting a query string into a query tree.

The query language is based on UNQL.

### Planner

The planner converts the query tree into a plan.  The plan is a tree of steps which must be performed to answer the query.  Currently the planner builds the simplest plan possible to provide the correct answer.

### Optimizer

Currently there are 2 optimizer implementations

* null - The null optimizer does nothing.
* naive - The naive optimizer currently knows how to optimize 2 scenarios
  * Queries on a single datasource.  In this case the optimizer tries to get the datasource to do as much of the work as possible.
  * Queries joining 2 datasources.  In this case the optimizer tries to get each datasource to filter as much as possible prior to joining, then in some cases switch to a sort-merge JOIN.


### Datasources

Currently 3 datasources are supported, CSV and ElasticSearch, and Couchbase + ElasticSearch.

Adding support for new datasources is easy if you have a way to iterate all the items in the datasource.  Making the datasource work efficiently requires converting various query expressions into operations on the server-side.

CSV is an example of a bare-minimum datasource.  The only operation it supports is returning all values.
ElasticSearch is an example of a more complex datasource.  It has support for evaluating various expressions on the server-side.


## Things that do NOT work
* Sub-query is allowed in the syntax, but not supported.
* Array index access in expression (like anArray[index])

## Examples

Basic query pulling back 1 full document

    tuq> SELECT FROM beer-sample LIMIT 1
    [
        {
            "doc": {
                "abv": 0,
                "brewery_id": "110f213f0c",
                "category": "North American Ale",
                "description": "",
                "ibu": 0,
                "name": "Schlafly IPA",
                "srm": 0,
                "style": "American-Style India Pale Ale",
                "type": "beer",
                "upc": 0,
                "updated": "2010-07-22 20:00:20"
            },
            "meta": {
                "id": "110f9f1c9f",
                "rev": "1-000000057d1b8d340000000000000000",
                "expiration": 0,
                "flags": 0
            }
        }
    ]
    
Reformat the result row into new document, including an array.
    
    tuq> SELECT { "beer-name": doc.name, "abv_ibu_array": [doc.abv, doc.ibu] } FROM beer-sample LIMIT 1 OFFSET 5
    [
        {
            "abv_ibu_array": [
                0,
                5.8
            ],
            "beer-name": "Old BullDog Extra Special"
        }
    ]

Show a query with a WHERE clause.  This one not matching any rows, returning empty result set.

    tuq> SELECT FROM beer-sample WHERE doc.abv > 12 && doc.abv < 12.1
    []
    
Show a similar query that does match rows.
    
    tuq> SELECT FROM beer-sample WHERE doc.abv > 12 && doc.abv < 12.2
    [
        {
            "doc": {
                "abv": 12.1,
                "brewery_id": "110f1db022",
                "description": "Our fourth in the Firehouse Ales Series, Pompier means \"fireman\" in French and represents our continued commitment to celebrate and honor the men and women who respond to the call day after day.  Pompier is rich and  smooth with complexities that come from a huge grain bill comprised of premium imported specialty malts, French Strisselspalt aroma hops and a 3 month aging process in oak hogsheads where it is combined with toasted French oak wood chips and champagne yeast.  Pompier is intended to be a vintage quality English-Style Barleywine with a French twist.  Appreciate its fine character and 12.1%A(MISSING)BV when we release this single 10 barrel batch sometime in December or you may choose to cellar it for many years to come.  \r\n\r\nYou will find Pompier on retail shelves packaged in the same 1 Liter Swing-Top bottle that has become a signature for our specialty beers.",
                "ibu": 0,
                "name": "Pompier",
                "srm": 0,
                "type": "beer",
                "upc": 0,
                "updated": "2010-07-22 20:00:20"
            },
            "meta": {
                "id": "110fc6508b",
                "rev": "1-000000069767a9500000000000000000",
                "expiration": 0,
                "flags": 0
            }
        }
    ]
    
Show using ORDER BY to find the beer with the highest alcohol content.
    
    tuq> SELECT FROM beer-sample ORDER BY doc.abv DESC LIMIT 1
    [
        {
            "doc": {
                "abv": 99.99,
                "brewery_id": "110f2a25d2",
                "category": "British Ale",
                "description": "",
                "ibu": 0,
                "name": "Norfolk Nog Old Dark Ale",
                "srm": 0,
                "style": "Old Ale",
                "type": "beer",
                "upc": 0,
                "updated": "2010-07-22 20:00:20"
            },
            "meta": {
                "id": "110f645057",
                "rev": "1-0000000421986fb00000000000000000",
                "expiration": 0,
                "flags": 0
            }
        }
    ]
    
Aggregate query across whole data source.

    tuq> SELECT {"count": count(doc.abv), "minabv": min(doc.abv), "maxibu": max(doc.ibu)} FROM beer-sample
    [
        {
            "count": 5901,
            "maxibu": 93,
            "minabv": 0
        }
    ]

Group all the documents based on their type (beer or brewery) and show the count of each category.
    
    tuq> SELECT { "type": doc.type, "count": count(doc.type) } FROM beer-sample GROUP BY doc.type
    [
        {
            "count": 5901,
            "type": "beer"
        },
        {
            "count": 1414,
            "type": "brewery"
        }
    ]
    
A  more complex example, WHERE uses "doc.style.analyzed" to refer to the string field that was analyzed by ElasticSearch.  Essentially
this matches beers with a style containing the term "lager".  Also beers with 0 abv are exluded.  Then we group the beers by "doc.style".
This uses the not-analyzed field, so we group on the exact styles.  Finally, we added an aggregate function on another field.  This output
also returns the average ABV for each beer style.
    
    tuq> SELECT {"style":doc.style, "count":count(doc.style), "avg_abv": avg(doc.abv)} FROM beer-sample WHERE doc.style.analyzed == "lager" && doc.abv > 0 GROUP BY doc.style
    [
        {
            "avg_abv": 5.166702702702702,
            "count": 185,
            "style": "American-Style Lager"
        },
        {
            "avg_abv": 5.105970149253731,
            "count": 67,
            "style": "Light American Wheat Ale or Lager"
        },
        {
            "avg_abv": 4.180652173913043,
            "count": 46,
            "style": "American-Style Light Lager"
        },
        {
            "avg_abv": 6.540000000000001,
            "count": 10,
            "style": "American Rye Ale or Lager"
        },
        {
            "avg_abv": 5.557142857142857,
            "count": 7,
            "style": "American-Style Cream Ale or Lager"
        },
        {
            "avg_abv": 4.5,
            "count": 1,
            "style": "American-Style Dark Lager"
        }
   ]

## Building (to hack on the internals)

1.  Install Go (http://golang.org/)
2.  Install nex (https://github.com/blynn/nex)
3.  Install goyacc (distributed with go, but may not be installed by default)
4.  Clone this project
5.  Run the included build.sh (runs nex, goyacc and go build)

## Building (to use)

1.  go get github.com/mschoch/tuq

## Running (out of date)

1.  You must already have a Couchbase Server
2.  You must already have an ElasticSearch Server
3.  You must already have integrated the two following this (http://blog.couchbase.com/couchbase-and-full-text-search-couchbase-transport-elastic-search)
4.  You must install one additional ElasticSearch Index Template
    curl -XPOST http://localhost:9200/_template/couchbase_unql -d @couchbase_unql_template.json
5.  Run ./go-unql-couchbase (defaults to localhost for Couchbase and ES, override with command-line options)