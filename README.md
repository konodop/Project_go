веб-сервис: пользователь отправляет арифметическое выражение по HTTP и получает в ответ его результат.

инструкция
1 Запуск приложения из терминала либо git bash с помощью команды 
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>go run ./cmd/main.go</pre><div class="zeroclipboard-container">
    <clipboard-copy aria-label="Copy" class="ClipboardButton btn btn-invisible js-clipboard-copy m-2 p-0 d-flex flex-justify-center flex-items-center" data-copy-feedback="Copied!" data-tooltip-direction="w" value="go run ./cmd/main.go" tabindex="0" role="button">
    </clipboard-copy>
</div></div>
2 Отправка POST-запроса через curl: например: 
<div class="highlight highlight-source-shell notranslate position-relative overflow-auto" dir="auto"><pre>curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1+1\"}"</pre><div class="zeroclipboard-container">
    <clipboard-copy aria-label="Copy" class="ClipboardButton btn btn-invisible js-clipboard-copy m-2 p-0 d-flex flex-justify-center flex-items-center" data-copy-feedback="Copied!" data-tooltip-direction="w" value="curl -X POST http://localhost:8080/api/v1/calculate -H "Content-Type: application/json" -d "{\"expression\": \"1+1\"}"" tabindex="0" role="button">
    </clipboard-copy>
</div></div>
