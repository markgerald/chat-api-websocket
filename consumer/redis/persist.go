package rediscache

import "log"

func Persist(botMessage string) {
	log.Printf("CHEGUEI NO PERSIST DO REDIS")
	err := Client().Set(ctx, "botcontent", botMessage, 6e+10).Err()
	if err != nil {
		log.Printf("MAS CAÍ NO ERR!!!")
		return
	}
}
