Football players info
=====================

The program extracts all the players from the given teams and render to 
stdout the information about players alphabetically ordered by name.
Each player entry contains the following information: 
```
full name; age; list of teams. 
```

## Usage

You can specify team names as arguments or at standard input:
```
./football-players-info England Germany "Manchester Utd"
```

or

```
echo "England" | ./football-players-info
```

If you don't specify any teams, then default team list is:
```
Germany
England
France
Spain
Manchester Utd
Arsenal
Chelsea
Barcelona
Real Madrid
FC Bayern Munich
```

## Configuration

You can configure number of parallel requests to api by setting `API_THREADS`
environment variable, default is 10. 

You can also configure maximum id to traverse to by setting `MAX_ID_LIMIT`
environment variable, default is 1000. 

## Exit code

If one or more teams is not found then script do not output anything and 
exit with code 1

## Testing

To run tests you need "testify" package installed:

```
go get github.com/stretchr/testify
```

To run tests:

```
go test ./...
```

