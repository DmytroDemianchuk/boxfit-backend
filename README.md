# BOX-FIT 
### BOX-FIT backend app

## API:

POST /auth/sign-up

Example input:
```
{
    "firstname" : "Dmytro",
    "secondname": "Demianchuk",
	"email": "dmytro@gmail.com",
	"password": "dmytro123"
} 
```

POST /auth/sign-in

Example input:
```
{
	"email": "dmytro@gmail.com",
	"password": "dmytro@gmail.com"
} 
```

Example Response:

```
{
	"token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NzEwMzgyMjQuNzQ0MzI0MiwidXNlciI6eyJJRCI6IjAwMDAwMDAwMDAwMDAwMDAwMDAwMDAwMCIsIlVzZXJuYW1lIjoiemhhc2hrZXZ5Y2giLCJQYXNzd29yZCI6IjQyODYwMTc5ZmFiMTQ2YzZiZDAyNjlkMDViZTM0ZWNmYmY5Zjk3YjUifX0.3dsyKJQ-HZJxdvBMui0Mzgw6yb6If9aB8imGhxMOjsk"
} 
```

## Start and run

make build - to build app
make run - to run app

