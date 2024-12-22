package application

import (
	"bufio"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"

	"calc/pkg/calculation"
)

type Config struct {
	Addr string
}

type BadResponse struct {
	Result string `json:"error"`
}

type Response struct {
	Result string `json:"result"`
}

func ConfigFromEnv() *Config {
	config := new(Config)
	config.Addr = os.Getenv("PORT")
	if config.Addr == "" {
		config.Addr = "8080"
	}
	return config
}

type Application struct {
	config *Config
}

func New() *Application {
	return &Application{
		config: ConfigFromEnv(),
	}
}

// Функция запуска приложения
// тут будем чиать введенную строку и после нажатия ENTER писать результат работы программы на экране
// если пользователь ввел exit - то останаваливаем приложение
func (a *Application) Run() error {
	for {
		// читаем выражение для вычисления из командной строки
		log.Println("input expression")
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Println("failed to read expression from console")
		}
		// убираем пробелы, чтобы оставить только вычислемое выражение
		text = strings.TrimSpace(text)
		// выходим, если ввели команду "exit"
		if text == "exit" {
			log.Println("aplication was successfully closed")
			return nil
		}
		//вычисляем выражение
		result, err := calculation.Calc(text)
		if err != nil {
			log.Println(text, " calculation failed wit error: ", err)
		} else {
			log.Println(text, "=", result)
		}
	}
}

type Request struct {
	Expression string `json:"expression"`
}

func CalcHandler(w http.ResponseWriter, r *http.Request) {
	request := new(Request)
	defer r.Body.Close()
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		response := BadResponse{
			Result: "Expression is not valid",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(response)
		return
	}
	request.Expression = strings.ReplaceAll(request.Expression, " ", "")
	re := regexp.MustCompile(`[^0-9\-+/*()]`)
	if re.MatchString(request.Expression) {
		response := BadResponse{
			Result: "Expression is not valid",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(response)
		return
	}

	result, err := calculation.Calc(request.Expression)
	if true == false {
		// как?
		response := BadResponse{
			Result: "Internal Server Error",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(500)
		json.NewEncoder(w).Encode(response)
	}
	if err != nil {
		response := BadResponse{
			Result: "Expression is not valid",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(422)
		json.NewEncoder(w).Encode(response)

	} else {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		response := Response{
			Result: strconv.FormatFloat(result, 'f', 6, 64),
		}
		json.NewEncoder(w).Encode(response)
	}
}

func (a *Application) RunServer() error {
	http.HandleFunc("/api/v1/calculate", CalcHandler)
	return http.ListenAndServe(":"+a.config.Addr, nil)
}
