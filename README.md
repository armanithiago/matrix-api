# League Backend Challenge
## Installing
#### Requirements
- Go already setup on the computer (otherwise [click here](https://golang.org/doc/install))
- Port 8080 available
##### Steps:
- Clone the repository
- Open the console
- Go to the project ``src`` folder 
- Enter the command below
```
go run .
````
## API
#### Requirements

``.csv`` file with same numbers of columns and rows

Example for all requests:
```
1,2,3
4,5,6
7,8,9
```
### /echo - POST
Return the matrix as a string in matrix format
**curl:**
```sh
curl --location --request POST 'http://localhost:8080/echo' --form 'file=@"<FILE_PATH>"'
```
**Return:**
```sh
"1,2,3
2,3,4
5,6,7"
```
### /invert - POST
Return the matrix as a string in matrix format where the columns and rows are inverted.
**curl:**
```sh
curl --location --request POST 'http://localhost:8080/invert' --form 'file=@"<FILE_PATH>"'
```
**Return:**
```sh
"1,4,7
2,5,8
3,6,9"
```
### /flatten - POST
Return the matrix as a 1 line string, with values separated by commas.
**curl:**
```sh
curl --location --request POST 'http://localhost:8080/flatten' --form 'file=@"<FILE_PATH>"'
```
**Return:**
```sh
"1,2,3,4,5,6,7,8,9"
```
### /sum - POST

Return the sum of the integers in the matrix.
*Especial Requirements: All elements in the matrix needs to be an integer*
**curl:**

```sh
curl --location --request POST 'http://localhost:8080/sum' --form 'file=@"<FILE_PATH>"'
```
**Return:**
```sh
45
```
### /multiply - POST
Return the product of the integers in the matrix.
*Especial Requirements: All elements in the matrix needs to be an integer*
**curl:**

```sh
curl --location --request POST 'http://localhost:8080/multiply' --form 'file=@"<FILE_PATH>"'
```
**Return:**
```sh
362880
```


## Testing

To execute the tests you can run the following command in the projects root folder.

```
go test -v ./...
```

This will run the tests recursively
