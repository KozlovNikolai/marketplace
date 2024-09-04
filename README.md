# marketplace
[Установка и запуск проекта.](#установка-и-запуск)

# Установка и запуск
1. установить утилиты: линтер и goose для миграций:
```
make install-golangci-lint
make install-goose
```
2. В папку `bin` сгенерировать и положить файлы сертификата и ключа для протокола  https:
```
openssl req -x509 -newkey rsa:4096 -keyout key.pem -out cert.pem -days 365 -nodes
```
3. Добавить переменные окружения - ключ шифрования для токена и пути нахождения файлов для http:
```
export TLS_CERT="/home/user/go/src/test-task/tasks/backend/GO/gremiha3/cert.pem"
export TLS_KEY="/home/user/go/src/test-task/tasks/backend/GO/gremiha3/key.pem"
export JWT_KEY="-my-256-bit-secret-"
```
4. собрать docker контейнеры:
    * postgres master
    * postgres replica
    * приложение
```
docker-compose up -d
```
5. Дальше заливаем миграции:
```
make local-migration-up
```
6. Открываем браузер:
```
https://localhost:8443/docs/index.html#/
```


## Примеры запросов:
### Регистрация:
```
http --verify=no POST https://localhost:8443/user/register login=qwerty@ya.com password=123456 username=Nikola role=super
```
### Авторизация:
```
http --verify=no POST https://localhost:8443/user/login login=qwerty@ya.com password=123456
```
### Вывести список всех пользователей:
```
http --verify=no GET https://localhost:8443/users Authorization:"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQwNDc2NTcsImxvZ2luIjoicXdlcnR5QHlhLmNvbSIsInJvbGUiOiJzdXBlciJ9.gLvjroKKz2FQrdDAt-SzaXMLTICl0s90VuRHC4wu6zo"
```
### Из браузера:
```
fetch(
  'https://localhost:8443/users',
  {
    method: 'GET',
    headers: { 'Content-Type': 'application/json','Authorization': 'Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MjQwNDY4NTYsImxvZ2luIjoiY21kQGNtZC5ydSIsInJvbGUiOiJzdXBlciJ9.Lz1tIHXDiSJQy6JspvFRSCCsGoNSFOg2S0SIzhTg_yk' }
  }
).then(resp => resp.text()).then(console.log)
```
[⬆️ Вернуться к оглавлению](#вопросы)