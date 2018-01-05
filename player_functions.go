package main
import (
 "github.com/google/uuid"
)

func UUID() string {
  u, err := uuid.NewRandom()
  if err != nil {
    panic(err)
  }
  return u.String()
}

func CreatePlayer() *Player{
    return &Player{
     ID: UUID(),
   }
}
