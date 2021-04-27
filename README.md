## Quick Setup

1. copy & rename .env.example to .env
2. add db configuration in .env

## Stack Libraries
1. gin gonic
2. sqlx
3. viper env


## Package Structure
Pada bahasa pemrograman Golang, satu folder atau *directory* itu dikatakan satu *package* sehingga berikut ini adalah strukur *package* / *directory* dengan menerapkan konsep *clean architecture*

```
cmd 
    - database
config
model
repository
    - mysql
transport
    - http
usecase
```