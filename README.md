# Gato CLI

HTTP Client on terminal, only for requests and responses in JSON

## Roadmap
- [x] Set values
- [x] Store values in yaml
- [x] GET request
- [x] Better format for response
- [x] Show values (check command)
- [x] Clear values
- [ ] Read file for json body
- [ ] POST request
- [ ] Open values directly from yaml

## Commands
### Set:
Set the values needed for the request, this are:

- **Host (-H)** The host, default: localhost

- **Path (-D)** The path to use on host, default: None 

- **Port (-P)** The port for request, default: 4000

- **Method (-M)** The method for request, default: GET 

- **Body (-B)** File to use as json body, default None

### Check:
Show the values that were setted

### Clear:
Clear values that were setted and return all to the default settings.

### Send:
Send request

## Example:

```
gato set -H localhost.com -P 8080 -M GET -B body.json
gato send
```
