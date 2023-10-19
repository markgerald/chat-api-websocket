package rediscache

func Persist(botMessage string) {
	err := Client().Set(ctx, "botcontent", botMessage, 6e+10).Err()
	if err != nil {
		return
	}
}
