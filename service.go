package main

import (
	"errors"
	"fmt"
	"net/http"
  "encoding/json"
)

type Service1 struct {
	Service1Config
}

type Service1Config struct {
	Host            string        `json:"host"`
	Port            int           `json:"port"`
}

type Hero struct {
	ID string `json:"id"`
	Name string `json:"name"`
}

func getHeros() []Hero{
	heros := []Hero {
    {"1", "fred"},
    {"2", "bert"},
  }
  return heros
}

func (s *Service1) validateRunConfig() error {
	if len(s.Host) == 0 {
		return errors.New("run requires a non-empty Host value")
	}
	if s.Port == 0 {
		return errors.New("run requires a non-zero Port value")
	}
	return nil
}

func handler(w http.ResponseWriter, r *http.Request) {
   hostname := r.URL.Query().Get("hostname")
   if len(hostname) != 0 {
      fmt.Fprintf(w, "Hi there, I love %s!", hostname)
   } else {
      fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
   }
}

func handler1(w http.ResponseWriter, r *http.Request) {
  heros := getHeros()
	jsonBlob, _ := json.Marshal(heros);

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*");
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT");
	w.Header().Set("Access-Control-Allow-Headers", "X-Requested-With, Content-Type, X-Codingpedia");

     w.Write(jsonBlob)
}

func (s *Service1) Run() error {
	if err := s.validateRunConfig(); err != nil {
		return fmt.Errorf("can't run: %s", err)
	}
	listenAddress := fmt.Sprintf("%s:%d", s.Host, s.Port)

  http.HandleFunc("/", handler)
  http.HandleFunc("/me/", handler1)
  http.ListenAndServe(listenAddress, nil)

	return nil
}
