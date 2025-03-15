# Hangman (Виселица) на Go

## Описание

Этот проект представляет собой классическую игру «Виселица», написанную на языке Go. Игроку необходимо угадать слово, вводя буквы по одной. Если буква есть в слове, она открывается; если нет — добавляется часть виселицы. Игра заканчивается победой при полном угадывании слова или поражением при превышении допустимого количества ошибок.

## Функционал

- Генерация случайного слова для угадывания
- Отображение текущего состояния слова (со скрытыми буквами) и виселицы.
- Подсчет количества ошибок
- Проверка введенных пользователем букв
- Конец игры при выигрыше или проигрыше

## Установка и запуск

1. Убедитесь, что у вас установлен Go.
2. Склонируйте репозиторий:

   ```sh
   git clone https://github.com/lemon4ikk/Gallows.git
   cd Gallows
   ```

3. Запустите игру:

   ```sh
   go run cmd/main.go
   ```

## Использование

1. При запуске игры предлагается `Начать новую игру` или `Выйти`.
2. В зависимоти от выбора игра начинается, либо программа завершает свою работу.
3. Если было выбрано начать игру, пользователь может начать вводит буквы.
4. Если буква есть в слове, она открывается; иначе увеличивается счетчик ошибок.
5. Игра продолжается до победы (все буквы открыты) или поражения (достигнут лимит ошибок).
