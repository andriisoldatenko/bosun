package annotate

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// SendAnnotation sends a annotation to an annotations server
// apiRoot should be "http://foo/api" where the annotation routes
// would be available at http://foo/api/annotation...
func SendAnnotation(apiRoot string, a Annotation) (Annotation, error) {
	b, err := json.Marshal(a)
	if err != nil {
		return a, err
	}

	res, err := http.Post(apiRoot+"/annotation", "application/json", bytes.NewReader(b))
	if err != nil {
		return a, err
	}
	defer res.Body.Close()
	err = json.NewDecoder(res.Body).Decode(&a)
	return a, err
}

//i.e.
// func main() {
// 	a := annotate.Annotation{
// 		CreationUser: "kbrandt",
// 		Owner:        "sre",
// 		Category:     "test",
// 		Message:      "lib send test",
// 		Source:       "go lib",
// 	}
// 	a, err := annotate.SendAnnotation("http://annotate", a)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	log.Println(a)
// }