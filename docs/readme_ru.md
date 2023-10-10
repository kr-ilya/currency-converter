# Конвертер валют - Telegram Mini App

Проект позволяет узнать актуальные курсы валют.  

🔸 Выполнен на базе Telegram Mini App (Web App).  

🔸 Курсы валют обновляются раз в час. (Поставщик курсов валют: https://fixer.io/)  

🔸 Локализация: Английский, Русский

![preview mini app][main_img]  


## Описание сервисов  

* **[Frontend мини-приложения][miniapp_docs]** (WebApp)
* **[Backend мини-приложения][backend_docs]**
* **[Backend бота][bot_docs]**

## Структура проекта  

Проект состоит из трех основных сервисов, каждый из которых находитится в соответствующей папке:  
* */backend* - backend мини-приложения - отвечает за хранение и получение курсов валют от поставщика и реализует API для получения курсов фронтендом мини-приложения.
* */webapp* - frontend мини-приложения
* */bot* - backend бота: обрабатывает обновления (запросы пользователя) бота. 

Папка *config* содержит файлы конфигурации сервисов.  
Папка *data* используется для хранения базы данных redis на диске.

## Технические особенности проекта

Для хранения курсов валют, полученных от поставщика используется [Redis][redis].
В качестве веб-сервера для отдачи статических файлов мини-приложения и [обратного прокси-сервера][reverse_proxy] используется [Nginx][nginx].  

Сервисы запускаются в [docker][docker] контейнерах. Docker - это платформа для создания, развертывания и управления контейнерами. Контейнеры позволяют упаковать приложение и все его зависимости в единое исполняемое окружение.  
Управление docker контейнерами осуществляется с помощью [Docker Compose][docker_compose]. Docker Compose позволяет легко настроить и объекдинить сервисы вместе и запустить приложение одной командой. Файл конфигурации Docker Compose: docker-compose.yml ([документация][compose_config]). 

Для интеграции мини-приложения с Telegram нужно, чтобы оно размещалось на домене с ssl сертификатом. В данном руководстве не затрагиваются этапы получения сертификата. Вы можете использовать бесплатный сертификат от letsencrypt.

---

## Установка и запуск

Убедитесь, что у вас установлено следующее:  
* Git ([install][git_download])
* Docker (v24.0.6)* ([install][install_docker])
* Docker Compose (v2.21.0)*

*версия, используемая при разработке  

1. Клонируйте репозиторий:

    ```shell
        git clone https://github.com/kr-ilya/currency-converter.git
    ```
2. Перейдите в папку с проектом  

    ```shell
        cd currency-converter
    ```

3. Настройка конфигурации

Основная конфигурация проекта задается путем редактирования файлов в папке config.  
Содержимое папки:  
* /nginx - содержит файл конфигурации nginx (reverse proxy)
* /redis - содержит файл конфигруации redis
* backend.env - файл конфигурации backend мини-приложения 
* bot.env - файл конфигурации backend бота

Кроме этого в файле */webapp/src/js/app.js* необходимо указать URL адрес API сервера (backend мини-приложения) (ссылка вида:  https://example.com/api).

### bot.env

*Параметры, обязательные к установке:*    

**BOT_TOKEN** - токен Telegram бота

**WEBHOOK_BASE** - URL адрес для webhook сервера, адрес должен иметь вид "scheme://host", например: https://example.com  

**WEBAPP_URL** - URL адрес мини-приложения (ссылка на главную, совпадает с *WEBHOOK_BASE*)

<details>
<summary>Параметры, которые можно оставить по умолчанию</summary> 

**LOGGER_TYPE** - тип логирования (prod/dev)  

**LISTEN_ADDRESS** - адрес с портом webhook сервера бота.  
При изменении также изменить порт в файле /config/nginx/nginx.conf. (в разделе *upstream bot*)  

</details>

### backend.env  

*Параметры, обязательные к установке:*  

**FIXER_ACCESS_TOKEN** - API токен сервиса для получения курсов валют ([fixer][fixer])  

**REDIS_PASS** - пароль к базе redis (также указать в файле /redis/redis.conf, параметр *requirepass*)

**BOT_TOKEN** - токен Telegram бота

<details>
<summary>Параметры, которые можно оставить по умолчанию</summary> 

**LISTEN_ADDRESS** - адрес с портом backend сервера бота.  
При изменении также изменить порт в файле /config/nginx/nginx.conf. (в разделе *upstream api*)

**LOGGER_TYPE** - тип логирования (prod/dev)  

**REDIS_DB=0** - номер базы данных redis  

**REDIS_HOST** - хост redis (соответствует названию сервиса из docker-compose.yml)

**REDIS_PORT** - порт redis (также указать в файле /redis/redis.conf, параметр *port*)
</details>

### Конфигруация Redis

Файл конфигурации: /redis/redis.conf  
[Описание конфигурационного файла Redis][redis_config]

Требуется указать:  
**requirepass** - пароль к базе redis

### Конфигурация nginx

Файл конфигурации: /config/nginx.conf  
[Документация Nginx][nginx_docs]  

В файле необходимо указать:  
* server_name  
* ssl_certificate  
* ssl_certificate_key



4. Запуск  

```shell
    docker compose up
```


[//]: # (LINKS)
[main_img]: ./assets/main.png
[redis]: https://redis.io/
[reverse_proxy]: https://en.wikipedia.org/wiki/Reverse_proxy
[nginx]: https://nginx.org/en/
[docker]: https://www.docker.com/get-started/
[docker_compose]: https://docs.docker.com/compose/
[git_download]: https://git-scm.com/downloads
[install_docker]: https://docs.docker.com/engine/install/
[fixer]: https://fixer.io/
[redis_config]: https://redis.io/docs/management/config-file/
[nginx_docs]: https://nginx.org/en/docs/
[compose_config]: https://docs.docker.com/compose/compose-file/03-compose-file/
[miniapp_docs]: ./miniapp_ru.md
[backend_docs]: ./backend_ru.md
[bot_docs]: ./bot_ru.md