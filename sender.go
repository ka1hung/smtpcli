package smtpcli

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

var InsecureMode bool

type Sender struct {
	Host     string
	Port     int
	User     string
	Password string
	auth     Auth
}

type Message struct {
	ContentType string
	To          []string
	CC          []string
	BCC         []string
	Subject     string
	Body        string
	Attachments map[string][]byte
}

func NewServer(addr string, port int, user, password string) *Sender {
	auth := PlainAuth("", user, password, addr)
	if user == "" && password == "" {
		auth = nil
	}

	return &Sender{
		Host:     addr + ":" + strconv.Itoa(port),
		Port:     port,
		User:     user,
		Password: password,
		auth:     auth}
}
func (s *Sender) Send(m *Message) error {
	return SendMail(s.Host, s.auth, s.User, m.To, m.toBytes())
}

// default set ContentType to text/plain; charset=UTF-8.
// body: use /r/n to make newline.
func NewMessage(subject, body string) *Message {
	return &Message{ContentType: "text/plain; charset=UTF-8", Subject: subject, Body: body, Attachments: make(map[string][]byte)}
}

// fs: set the files path
func (m *Message) AttachFiles(files []string) error {
	for _, f := range files {
		b, err := os.ReadFile(f)
		if err != nil {
			return err
		}

		_, fileName := filepath.Split(f)
		m.Attachments[fileName] = b
	}
	return nil
}

func (m *Message) toBytes() []byte {
	buf := bytes.NewBuffer(nil)
	withAttachments := len(m.Attachments) > 0
	buf.WriteString(fmt.Sprintf("Subject: %s\n", m.Subject))
	buf.WriteString(fmt.Sprintf("To: %s\n", strings.Join(m.To, ",")))
	if len(m.CC) > 0 {
		buf.WriteString(fmt.Sprintf("Cc: %s\n", strings.Join(m.CC, ",")))
	}

	if len(m.BCC) > 0 {
		buf.WriteString(fmt.Sprintf("Bcc: %s\n", strings.Join(m.BCC, ",")))
	}

	buf.WriteString("MIME-Version: 1.0\n")
	writer := multipart.NewWriter(buf)
	boundary := writer.Boundary()
	if withAttachments {
		buf.WriteString(fmt.Sprintf("Content-Type: multipart/mixed; boundary=%s\n", boundary))
		buf.WriteString(fmt.Sprintf("--%s\n", boundary))
	}
	buf.WriteString(fmt.Sprintf("Content-Type: %s\n", m.ContentType))

	buf.WriteString("\n" + m.Body)
	if withAttachments {
		for k, v := range m.Attachments {
			buf.WriteString(fmt.Sprintf("\n\n--%s\n", boundary))
			buf.WriteString(fmt.Sprintf("Content-Type: %s\n", http.DetectContentType(v)))
			buf.WriteString("Content-Transfer-Encoding: base64\n")
			buf.WriteString(fmt.Sprintf("Content-Disposition: attachment; filename=%s\n", k))

			b := make([]byte, base64.StdEncoding.EncodedLen(len(v)))
			base64.StdEncoding.Encode(b, v)
			buf.Write(b)
			buf.WriteString(fmt.Sprintf("\n--%s", boundary))
		}

		buf.WriteString("--")
	}

	return buf.Bytes()
}
