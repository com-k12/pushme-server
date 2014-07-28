pushme-server
=============

Демон для отправки сообщений на мобильное приложение ["Pusme-mobile"](https://github.com/com-k12/pushme-mobile).
Команды демон получает по http-протоколу по следующему синтаксису:

* _http://localhost/add/?level=0&message=all_is_ok_ - добавляет сообщение "all_is_ok" в очередь на отправку 
всем зарегестрированным пользователям.  Уровень сообщения - 0 (0-OK, 1-warning, 2-critical, 3-info)
* _http://localhost/add/?level=0&message=come_to_the_meeting&users=user1,user2_ - добавляет сообщение "come_to_the_meeting" в очередь на отправку 
указанным в users пользователям.  Уровень сообщения - 0 (0-OK, 1-warning, 2-critical, 3-info)
* _http://localhost/register/?user=user3_ - регистрирует пользователя в системе
* _http://localhost/?user=user1 - выдаёт все сообщения предназначенные пользователю user1 и уберает их из очереди сообщений.
