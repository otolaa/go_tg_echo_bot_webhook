# 👾 simple echo bot on webhook

You must have a token from the bot.
And write the token to .env

## 🍐 webhook start pinggy.io

```
ssh -p 443 -R0:127.0.0.1:8080 -L4300:localhost:4300 free.pinggy.io
```

🍎 add url webhook in .env

## 🌵 Kill Port

```
killall -9 go
```

## 🚀 start

```
go run .
```