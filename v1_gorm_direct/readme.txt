graphql:
https://gqlgen.com/getting-started/
https://github.com/99designs/gqlgen/issues/1483#issuecomment-949023482
sửa trong schema sau đó chạy lệnh :
go get -d github.com/99designs/gqlgen
go run github.com/99designs/gqlgen generate
https://www.youtube.com/watch?v=FkpCeXbXVhU
git config core.autocrlf true
https://www.soberkoder.com/go-graphql-api-mysql-gorm/



Query:
query {
  getAllUsers{
    id
    name
    email
  }
}
Mutation:
mutation {
  createNewUser(input:{
    name: "tun sugar"
    email: "tu333n@gmail.com"
    password: "123554"
  }) {
    email
    name
    password
  }
}

source: 
https://hasura.io/blog/building-a-graphql-api-with-golang-postgres-and-hasura/