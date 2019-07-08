package network

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/sirupsen/logrus"
	"reflect"
)

type Session interface {
	SendMsg(msg proto.Message)
	SendRawMsg(msgID uint32, data []byte)
}

type Client struct {
	conn Conn
	mgr  *Mgr
}

func NewClient(conn Conn, mgr *Mgr) *Client {
	p := &Client{
		conn: conn,
		mgr:  mgr,
	}
	return p
}

func (p *Client) ReadLoop() {
	for {
		data, err := p.conn.ReadMsg()
		if err != nil {
			fmt.Printf("read message: %v", err)
			break
		}

		msg, err := p.mgr.processor.Unmarshal(data)
		if err != nil {
			logrus.Debugf("unmarshal message error: %v", err)
			break
		}
		p.mgr.Post(func() {
			err = p.mgr.processor.Handle(msg, p)
		})
	}
}

func (p *Client) OnNew() {
	fmt.Println("client new")
	p.mgr.Post(func() {
		p.mgr.sesID2Client[p.conn.ID()] = p
		p.mgr.onNew(p)
	})
}

func (p *Client) OnClose() {
	fmt.Println("client close")
	p.mgr.Post(func() {
		delete(p.mgr.sesID2Client, p.conn.ID())
		p.mgr.onClose(p)
	})
}

func (p *Client) SendMsg(msg proto.Message) {
	data, err := p.mgr.processor.Marshal(msg)
	if err != nil {
		logrus.Errorf("marshal message %v error: %v", reflect.TypeOf(msg), err)
		return
	}
	err = p.conn.WriteMsg(data)
	if err != nil {
		logrus.Error("write message %v error: %v", reflect.TypeOf(msg), err)
	}
}

func (p *Client) SendRawMsg(msgID uint32, data []byte) {
	newData := p.mgr.processor.Encode(msgID, data)
	err := p.conn.WriteMsg(newData)
	if err != nil {
		logrus.Error("write message error: %v", err)
	}
}
