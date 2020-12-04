// Forces the user to use the builder and keeps the email object away
// function intializer as a parameter

package main

import (
	"fmt"
	"strings"
)

type email struct {
	from, to, subject, body string
}


type EmailBuilder struct {
	email email
}

func (b *EmailBuilder) From(from string) *EmailBuilder {
	if !strings.Contains(from, "@"){
		panic("email need @")
	}
	b.email.from = from;
	return b
}

func (b *EmailBuilder) To(to string) *EmailBuilder {
	
	b.email.to = to;
	return b
}

func (b *EmailBuilder) Subject(subject string) *EmailBuilder {
	
	b.email.subject = subject;
	return b
}

func (b *EmailBuilder) Body(body string) *EmailBuilder {
	
	b.email.body = body;
	return b
}

func sendMailImplementation(email *email){
	fmt.Println("email sent")
}

type build func(*EmailBuilder)
func SendEmail(action build) {
	build := EmailBuilder{}
	action(&build)
	sendMailImplementation(&build.email)
}

func main() {
	SendEmail(func(b *EmailBuilder){
		b.From("sgrumleydev@gmail.com").
			To("support@rec.com").
			Subject("Meeting").
			Body("Meeting at 3, pls confirm")
	})
}