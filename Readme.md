Dating App
</br>
-------------------------------------------------------</br>
Technologies/tools Used:</br>
=================</br>
    Golang
    Postgres
    Gorm
    Gin

</br>

API Requests:</br>
=============</br>

1) GET    /dating/match/{id}                          => Matches and Returns the Users where User liked Other and Other liked User  </br>
2) GET    /dating/matchalpha/{match alphabet}         => Matches and Returns the name which contains the alphabet </br>
3) GET    /dating/range/{id}/{range}                         => Returns all the users withing given range of user with User ID </br>
4) GET    /dating/allmatches                           => Returns all the matches in the database


==================================</br>

Run locally 
```sh
go run cmd/server/main.go
```
Run Port 
```sh
localhost:7000
```