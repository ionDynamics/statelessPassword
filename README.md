# statelessPassword
Derived Billemont's algorithm in Go

```
go get go.iondynamics.net/statelessPassword
```
[Documentation](https://go.iondynamics.net/statelessPassword)

### You just want the cli tool?
```
go get go.iondynamics.net/statelessPassword/cmd/slpw
```

### slpw
You may set the following environment variables to circumvent some repeated typing  
ID_SLPW_FULLNAME = {Your Full Name}  
ID_SLPW_DEFAULTVARIANT = {The algorithm parameter variant} //defaults to 3  

### You don't know which variant to choose?
```
go get go.iondynamics.net/statelessPassword/cmd/slpw_perf
```
Execute slpw_perf and choose the variant code with a runtime between 400ms and 1sec
