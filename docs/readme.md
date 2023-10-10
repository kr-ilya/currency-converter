# Currency converter - Telegram Mini App

The project allows you to find out the current exchange rates. 

🔸 Made on the basis of Telegram Mini App (Web App). 

🔸 Currency exchange rates are updated once an hour. (Currency exchange rate provider: https://fixer.io/)

🔸 Localization: English, Russian 

![preview mini app][main_img]  


## Description of services  

* **[Frontend Mini App][miniapp_docs]** (WebApp)
* **[Backend Mini App][backend_docs]**
* **[Backend bot][bot_docs]**

## Project structure 

The project consists of three main services, each of which is located in the corresponding folder:  
* */backend* - backend Mini App - It is responsible for storing and receiving exchange rates from the supplier and implements an API for receiving rates by the Mini App frontend.
* */webapp* - frontend Mini App
* */bot* - backend bot: processes updates (user requests) of the bot.

The *config* folder contains the service configuration files.  
The *data* folder is used to store the redis database on disk.

## Technical features of the project

To store the exchange rates received from the supplier, [Redis][redis] is used.
[Nginx][nginx] is used as a web server for sending static Mini App files and [reverse proxy server][reverse_proxy].  

Services run in [docker][docker] containers. Docker is a platform for creating, deploying and managing containers. Containers allow you to package an application and all its dependencies into a single executable environment.  
Docker containers are managed using [Docker Compose][docker_compose]. Docker Compose makes it easy to set up and combine services together and launch an application with a single command. Docker Compose configuration file: docker-compose.yml ([documentation][compose_config]).  

To create a mini-application with Telegram, it must be hosted on a domain with an SSL certificate. This guide does not cover the steps to obtain a certificate. You can use the free letsencrypt certificate.

---

## Installation and launch  

Make sure you have the following installed:  
* Git ([install][git_download])
* Docker (v24.0.6)* ([install][install_docker])
* Docker Compose (v2.21.0)*

*version used during development

1. Clone the repository:

    ```shell
        git clone https://github.com/kr-ilya/currency-converter.git
    ```
2. Go to the project folder  

    ```shell
        cd currency-converter
    ```

3. Configuring the configuration

The main configuration of the project is set by editing files in the config folder.  
Folder contents:
* /nginx - contains nginx configuration file (reverse proxy)
* /redis - contains redis configuration file
* backend.env - backend Mini App configuration file 
* bot.env - bot's backend configuration file

In addition, in the file */webapp/src/js/app.js* it is necessary to specify the URL of the API server (backend Mini App) (link of the type: https://example.com/api).

### bot.env

*Parameters required for installation:*    

**BOT_TOKEN** - Telegram bot token 

**WEBHOOK_BASE** - The URL for the webhook server, the address should look like "scheme://host", for example: https://example.com  

**WEBAPP_URL** - Mini App URL (link to the main page, matches with *WEBHOOK_BASE*)

<details>
<summary>Parameters that can be left by default</summary> 

**LOGGER_TYPE** - type of logging (prod/dev)  

**LISTEN_ADDRESS** - the address with the webhook port of the bot server.  
When changing, also change the port in the /config/nginx/nginx.conf file. (in the *upstream bot* section) 

</details>

### backend.env  

*Parameters required for installation:*  

**FIXER_ACCESS_TOKEN** - API token of the service for obtaining exchange rates ([fixer][fixer])  

**REDIS_PASS** - password to the redis database (also specify in the /redis/redis.conf file, the *requirepass* parameter)

**BOT_TOKEN** - Telegram bot token

<details>
<summary>Parameters that can be left by default</summary> 

**LISTEN_ADDRESS** - the address with the backend port of the bot server.  
When changing, also change the port in the /config/nginx/nginx.conf file. (in the *upstream api* section)

**LOGGER_TYPE** - logging type (prod/dev)  

**REDIS_DB=0** - redis database number  

**REDIS_HOST** - redis host (corresponds to the name of the service from docker-compose.yml)

**REDIS_PORT** - redis port (also specify the *port* parameter in the /redis/redis.conf file)
</details>

### Redis Configuration  

Configuration file: /redis/redis.conf  
[Description of the Redis configuration file][redis_config]

It is required to specify:  
**requirepass** - password to the redis database

### nginx configuration

Configuration file: /config/nginx.conf  
[Nginx Documentation][nginx_docs]  

In the file, you must specify:
* server_name  
* ssl_certificate  
* ssl_certificate_key



4. Launch  

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
[miniapp_docs]: ./miniapp_en.md
[backend_docs]: ./backend_en.md
[bot_docs]: ./bot_en.md