# Candy Store

This program will read data from a web page, and analyse the data to return information about the customers, or return an error if the data cannot be read.

This program uses go modules. Please use the command `go mod vendor` to get all the necessary dependencies.

## Install and run
```
git clone https://github.com/achenet/candy-store
go mod vendor
go build
./candy-store
```

## Alt mode
The output can be generated two different ways. The default is to parse the top customer summary table of the webpage. The alternative is to parse the top customer details table of the webpage, and then calculate the favourites from that.

To use the alternative method, simply use the `alt` flag when running the program.
```
./candy-store -alt
```
