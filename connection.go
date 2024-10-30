package quickfix

import (
	"io"
	"log"
	"time"

	glog "git.5th.im/lb-public/gear/log"
)

func writeLoop(connection io.Writer, messageOut chan []byte, log Log) {
	for {
		msg, ok := <-messageOut
		if !ok {
			return
		}
		tmNow := time.Now()
		if _, err := connection.Write(msg); err != nil {
			log.OnEvent(err.Error())
		} else {
			log.OnOutgoing(msg)
		}
		glog.Infof("[timetest-out] messageout - writeLoop time:%+v", time.Since(tmNow))
	}
}

func readLoop(parser *parser, msgIn chan fixIn) {
	defer close(msgIn)

	for {
		tmNow := time.Now()
		msg, err := parser.ReadMessage()
		if err != nil {
			log.Println(`parser read message failed,conection readLoop just quit, error_info->`, err.Error())
			return
		}
		msgIn <- fixIn{msg, parser.lastRead}
		glog.Infof("[timetest-in] readLoop time:%+v", time.Since(tmNow))
	}
}
