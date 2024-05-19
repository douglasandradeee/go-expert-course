// package srp

// import (
// 	"io/fs"
// 	"net/url"
// 	"os"
// 	"strings"
// )

// func main() {
// 	journal := &Journal{}
// 	news := "news1"
// 	journal.AddNews(news)
// }

// type Journal struct {
// 	News []string
// }

// func (j *Journal) String() string {
// 	return strings.Join(j.News, "\n")
// }

// func (j *Journal) AddNews(news string) {
// 	//...
// }

// func (j *Journal) RemoveNews(news string) {
// 	//...
// }

// func (j *Journal) UpdateNews(news string) {
// 	//...
// }

// func (j *Journal) DeleteNews(news string) {
// 	//...
// }

// //métodos para printar/imprimir o jornal.
// /* Os métodos abaixo já fogem totalmente do escopo da nossa classe Journal, portanto estaríamos
// acrescentando à classe uma responsabilidade que vai além das de sua criação, seria assim adicionado
// o escopo de io(in/out) com metodos relacionados a forma como iremos montar o Journal e não a criação do Journal em si.
// Quebrando asssim o SRP- Single Responsibility Principle
// */
// /* Como resolução criaremos classe diferentes para trablhar com escopos diferente,
// cada classe com sua única responsibilidade. PS: em go como "não trablhamos com classes",
// iremos então separar por arquivos de acordo com o escopo.
// */

// func (j *Journal) PrintNews(filename string, news string) {
// 	_ = os.WriteFile(filename, []byte(j.String()), fs.FileMode(0644))
// }

// func (j *Journal) LoadNewsFromFile(filename string) {
// 	//...
// }

// func (j *Journal) LoadNewsFromWeb(website *url.URL) {
// 	//...
// }
