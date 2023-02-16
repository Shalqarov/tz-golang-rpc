package main

import (
	"crypto/md5"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"tz/models"
)

type Service struct {
	db map[string]*models.User
}

func (s *Service) Create(payload models.User, reply *models.User) error {
	if _, has := s.db[payload.Email]; has {
		return fmt.Errorf("user with %s email already exists", payload.Email)
	}

	hash := md5.New()
	hash.Write([]byte(payload.Password))
	hashedPassword := fmt.Sprintf("%x", hash.Sum([]byte(payload.Salt)))

	s.db[payload.Email] = &models.User{
		Email:    payload.Email,
		Salt:     payload.Salt,
		Password: hashedPassword,
	}
	log.Println(payload)
	*reply = payload
	log.Println(payload)

	return nil
}

func (s *Service) Get(payload string, reply *models.User) error {
	u, exist := s.db[payload]
	if !exist {
		return fmt.Errorf("user with %s email doesn't exists", payload)
	}

	*reply = *u
	return nil
}

func main() {
	s := &Service{db: make(map[string]*models.User)}
	rpc.Register(s)

	t, err := net.ResolveTCPAddr("tcp4", ":8080")
	if err != nil {
		log.Fatalln(err)
	}

	l, err := net.ListenTCP("tcp4", t)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", c.RemoteAddr())
		rpc.ServeConn(c)
	}
}
