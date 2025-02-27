# Gato CLI

HTTP Client on terminal, only for requests and responses in JSON

## Roadmap
- [x] Set values
- [x] Store values in yaml
- [x] GET request
- [x] Better format for response
- [x] Show values (check command)
- [x] Clear values
- [x] Read file for json body
- [ ] POST request
- [ ] Add Authentication
- [ ] Refactor commands
- [ ] Add command to get values directly from yaml file
- [ ] Add tests for every command

## Commands
### Set:
Set the values needed for the HTTP request, this values are:

- **Host (-H)** The host, default: localhost

- **Path (-D)** The path to use on host, default: None 

- **Port (-P)** The port for request, default: 4000

- **Method (-M)** The method for request, default: GET 

- **Body (-B)** File to use as json body, default None

### Check:
Displays the current values of the HTTP request.

If used without arguments, it lists all set values, if a specific parameter name is provided as an argument, it returns only the value for that parameter.

### Clear:
Restore the values of the HTTP request to the values of the default.yaml provided.

If used without arguments, it restore all set values, if a specific parameter name is provided as an argument, it restore only the value for that parameter.


### Send:
Send the HTTP request with the values that were set.

## Files

### gato.yaml
We'll use gato.yaml to save the values that were set for the HTTP requests.

``` yaml
host: 
path:
port: 
method:
body:
```

### default.yaml

We'll use default.yaml to have the default parameters set in case we need to clear data.

``` json
{
  "host":"http://127.0.0.1",
  "path":"",
  "port":8000,
  "method":"GET",
  "body":""
}
```

## Example:

```
gato set -H localhost.com -D api/todos -P 4000 -M GET -B body.json
gato send
```
<p align="center">
  <img src="https://github.com/user-attachments/assets/1cfcb045-b4da-4546-9ce9-4d9a638be93c", alt="Response" />
</p>

<h1></h1>

<p align="center">
  <img src="https://github.com/user-attachments/assets/5c1d6653-690d-47e5-9917-a60902738241", width="150", height="150", alt="Gato Logo" /> 
</p>
