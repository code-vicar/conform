# Conform Manual
## Command
conform
## Flags
* -o Output file path
* -p Environment variable prefix
* -f Output file format
## Examples

### Simple object
ENV
* COUCHDB_HTTPD{}BIND_ADDRESS 0.0.0.0

#### ini
Execute
````
$ conform -p "COUCHDB_" -f "ini"
````
Output:
````
[httpd]
bind_address 0.0.0.0
````

#### yml

Execute
````
$ conform -p "COUCHDB_" -f "yml"
````
Output:
````
httpd:
  bind_address: "0.0.0.0"
````

#### json

Execute
````
$ conform -p "COUCHDB_" -f "json"
````
Output:
````
{
    httpd: {
        bind_address: "0.0.0.0"
    }
}
````

### Arrays
ENV
* COUCHDB_TEST{}ARR[] one
* COUCHDB_TEST{}ARR[] two

#### ini
Execute
````
$ conform -p "COUCHDB_" -f "ini"
````
Output:
````
[test]
arr "[one, two]"
````

#### yml
Execute
````
$ conform -p "COUCHDB_" -f "yml"
````
Output:
````
test:
    arr:
      - "one"
      - "two"
````

#### json
Execute
````
$ conform -p "COUCHDB_" -f "json"
````
Output:
````
{
    test: {
       arr: [
        "one",
        "two"
       ]
    }
}
````
