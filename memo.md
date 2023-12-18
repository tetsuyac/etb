
````
go build -buildmode=plugin -gcflags="all=-N -l" -o test/T001.so test/T001.go
go build -buildmode=plugin -gcflags="all=-N -l" -o code/C001.so code/C001.go 
````
#### no plugin debug feature in GoLand 2022.2.1 Build #GO-222.3345.146, built on August 4, 2022
 