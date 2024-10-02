package service

type Cache struct {
	origin  string
	port    string
	headers map[string]interface{}
	body    map[string]interface{}
}

func CreateNewCache(origin string, port string, headers map[string]interface{}) {

}
