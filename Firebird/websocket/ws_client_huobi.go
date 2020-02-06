package websocket

import (
	"Firebird/db"
	"fmt"
	"math/rand"
	"net"
	"time"

	"Firebird/data"
	"Firebird/service"
	"encoding/json"
	"github.com/gorilla/websocket"
	"strings"
)

func (cli *Client) RunClient() {
	AddClientNum()
	dialer := websocket.DefaultDialer
	dialer.NetDial = func(network, addr string) (net.Conn, error) {
		addrs := []string{string(localIP)}
		fmt.Println("IP Address = ", addrs)
		localAddr := &net.TCPAddr{IP: net.ParseIP(addrs[rand.Int()%len(addrs)]), Port: 0}
		d := net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
			LocalAddr: localAddr,
		}
		c, err := d.Dial(network, addr)
		return c, err
	}
	c, _, err := dialer.Dial(cli.Addr, nil)
	if err != nil {
		log.Error("Dial Erro:", err)
		SubClientNum()
		return
	}

	log.Info(c.LocalAddr().String())
	defer func() {
		c.Close()
		SubClientNum()
	}()

	symbolMap := db.GetAllSymbolFromCache()
	for k := range symbolMap {
		topic := fmt.Sprintf("{\"sub\":\"market.%s.kline.1day\", \"id\":\"fb1\"}", k)
		message := []byte(topic)
		err = c.WriteMessage(websocket.TextMessage, message)
		if err != nil {
			log.Error("write err :", err)
		}
	}

	go func() {
		pangTicker := time.NewTicker(time.Second * 5)

		for {
			select {
			case <-pangTicker.C:
				message := []byte(fmt.Sprintf("{\"pong\":%d}", time.Now().Unix()))
				err = c.WriteMessage(websocket.TextMessage, message)
				if err != nil {
					log.Error("send msg err:", err)
					return
				}
			}
		}
	}()

	for {
		_, zipmsg, err := c.ReadMessage()
		if err != nil {
			log.Error("Read Error : ", err, cli.Name)
			c.Close()
			return
		}

		msg, err := parseGzip(zipmsg)
		if err != nil {
			log.Error("gzip Error : ", err)
			return
		}

		klineData := data.KlineData{}
		json.Unmarshal(msg, &klineData)
		tick := klineData.Tick
		if tick.Id > 0 {
			symbols := strings.Split(klineData.Ch, ".")
			log.Debug(symbols[1], tick.Close)

			// update symbol price
			service.NotifySymbolPrice(symbols[1], tick.Close)

			// notify schedule
			service.NotifySchedulePrice(symbols[1], tick.Close)
		}
		//log.Println(string(msg[:]))
	}
}
