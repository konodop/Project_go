<h1>веб-сервис<h1\>
    <h2>Описание:</h2>
    <h3>пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.
    строка-выражение состоит из односимвольных идентификаторов и знаков арифметических действий.
    Входящие данные - цифры(рациональные), операции +, -, *, /, операции приоритезации ( ). В случае ошибки записи выражения приложение выдает ошибку.
    </h3>
    <h1>Требования:<h1\>
    <h2>
    <ul>
    <li>golang, можно скачать на https://go.dev/dl/</li>
    <li>(необязательно) git, можно скачать на https://git-scm.com/downloads</li>
    </ul>
    </h2>
<hr><hr\>
<h1>инструкция для запуска проекта:<h1\>
<h4>шаг 1 скачать веб сервис через git clone https://github.com/konodop/Project_go.git или просто скачать и распаковать .zip файл из гитхаба если не установлен git<h4\>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>go run ./cmd/main.go</pre></div>
<h4>шаг 2 Запуск приложения в папке с помщью терминала либо git bash и команды:<h4\>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>go run ./cmd/main.go</pre></div>
<h4>шаг 3 Отправка POST-запроса через curl: например:<h4\>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1+1\"}"</pre></div>
    Ответ:
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>{"result":"2.000000"}</pre></div>
<h3>Можно подставлять другие значения в expression и проверять их<h3\>
<h3>Для того чтобы остановить приложение нужно написать exit <h3\>
<hr><hr\>
<h1>Примеры запросов:<h1\>
<h2>1</h2>
<h4>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1/0\"}"</pre></div>
Ответ:
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>{"error":"Expression is not valid"}</pre><div class="zeroclipboard-container"></div>
</h4>
<h2>2</h2>
<h4>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1-(1000*10)\"}"</pre></div>
Ответ:
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>{"result":"-9999.000000"}</pre><div class="zeroclipboard-container"></div>
</h4>
