<h2>веб-сервис: пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.<h2\>
    <h2>Требования:<h2\>
        <ul dir="auto">
    <li>Нужен golang, можно скачать на https://go.dev/dl/</li>
    <li>(необязательно) git, можно скачать на https://go.dev/dl</li>
    </ul>
<h2>инструкция:<h2\>
<h4>шаг 1 скачать веб сервис через git init mod github.com/konodop/Project_go или просто скачать и распаковать .zip файл из гитхаба если не установлен git<h4\>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>go run ./cmd/main.go</pre><div class="zeroclipboard-container">
<h4>шаг 2 Запуск приложения в папке с помщью терминала либо git bash и команды:<h4\>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>go run ./cmd/main.go</pre><div class="zeroclipboard-container">
</div></div>
<h4>шаг 3 Отправка POST-запроса через curl: например:<h4\>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1+1\"}"</pre><div class="zeroclipboard-container">
Ответ:
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>{"result":"2.000000"}</pre><div class="zeroclipboard-container">
<h3>Можно подставлять другие значения в expression и проверять их<h3\>
<h2>Примеры запросов:<h2\>
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1/0\"}"</pre><div class="zeroclipboard-container">
    </clipboard-copy>
</div></div>
