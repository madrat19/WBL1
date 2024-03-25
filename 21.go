// Реализовать паттерн «адаптер» на любом примере

// Есть типы Twitter и Facebook со специфическими методами отправки сообщений
// Также есть тип Message котрый содержит единый метод send
// Необходимо реализовать отправку сообщений в любую из двух соцсетей общим образом
// Для этого используются объеты-адаптеры TwitterAdapter и  FacebookAdapter,
// согласующие интерфейсы вышеописанных типов

package main

import (
	"fmt"
)

// Адаптируемый объект
type Twitter struct {
	username string
}

// Cпецифический метод
func (t Twitter) twit(message string) {
	fmt.Printf("%s tweeted: %s\n", t.username, message)
}

// Адаптируемый объект
type Facebook struct {
	username string
}

// Cпецифический метод
func (f Facebook) post(message string) {
	fmt.Printf("%s posted: %s\n", f.username, message)
}

// Адаптер
type TwitterAdapter struct {
	twitter Twitter
}

// Общий метод
func (ta TwitterAdapter) send(message string) {
	ta.twitter.twit(message)
}

// Адаптор
type FacebookAdapter struct {
	facebook Facebook
}

// Общий метод
func (fa FacebookAdapter) send(message string) {
	fa.facebook.post(message)
}

type Sender interface {
	send(message string)
}

// Содержит встроенный адаптор одного издвух типов
type Message struct {
	webService Sender
	text       string
}

// Общий метод
func (m Message) send() {
	m.webService.send(m.text)
}

// Тест
func main() {
	twitter := Twitter{"Donald Trump"}
	facebook := Facebook{"Ayn Rand"}

	twitterAdapter := TwitterAdapter{twitter}
	facebookAdapter := FacebookAdapter{facebook}

	message1 := Message{twitterAdapter, "Make America great again!"}
	message2 := Message{facebookAdapter, "Work!"}

	message1.send()
	message2.send()
}
