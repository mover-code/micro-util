/***************************
@File        : email.go
@Time        : 2022/06/21 16:34:32
@AUTHOR      : small_ant
@Email       : xms.chnb@gmail.com
@Desc        : use amz send email
****************************/

package amz

import (
    "gopkg.in/gomail.v2"
)

type Conf struct {
    EndPoint string
    Sender   string
    Key      string
    Port     int
    From     string
}

func Init(endPoint, sender, key, from string, port int) *Conf {
    return &Conf{
        EndPoint: endPoint,
        Sender:   sender,
        Key:      key,
        Port:     port,
        From:     from,
    }
}

func (c *Conf) Send(reciver, subject, body string) error {

    m := gomail.NewMessage()
    m.SetHeader("From", c.From)
    m.SetHeader("To", reciver)
    //m.SetAddressHeader("Cc", "dan@example.com", "Dan")
    m.SetHeader("Subject", subject)
    m.SetBody("text/html", body)
    //m.Attach("/home/Alex/lolcat.jpg")
    d := gomail.NewDialer(c.EndPoint, c.Port, c.Sender, c.Key)
    // Send the email to Bob, Cora and Dan.
    return d.DialAndSend(m)
}
