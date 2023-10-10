# Frontend мини-приложения 

Фронтенд мини-приложения (WebApp).

## Структура проекта  

Основные файлы и папки:  

* */src* - папка с исходным кодом сервиса
* *nginx.conf* - файл конфигурации nginx (веб-сервер для раздачи статического контента)
* *package.json* - файл, сожержащий информацию о проекте и его зависимостях
* *webpack.config.js* - конфигурационный файл для Webpack
* *Dockerfile* - файл содержит инструкции для создания Docker-образа 

## Технические особенности проекта  

Сервис написан на javascript. Используется инструмент сборки javascript приложений - [Webpack][webpack]. Он позволяет оптимизировать и упаковывать ресурсы (такие как js файлы, стили, изображения) в оптимизированный бандл (bundle) для развертывания на веб-сайтах. Webpack также позволяет использовать модульную систему, что облегчает организацию кода.

При сборке проекта js, css файлы минимизируются, объединяются в бандлы (несколько js файлов в один js файл, несколько css файлов - в один css).

При запуске проекта с помощью docker compose кроме сборки проекта выполняется установка веб-сервера nginx для отдачи статического контента.

## Сборка проекта  

При запуске проекта с помощью docker compose сервис собирается автоматически в контейнере (выполняется команда *npm run build*).

Сборка сервиса в среде разработки (dev) осуществляется командой npm run dev. 

## Внутреннее устройство сервиса  

Сервис написан следуя [MVP][mvp] паттерну проектирования приложений (Model-View-Presenter).  
В соответсвтии с паттерном в структуре сервиса есть три основных класса, находящихся в одноименных файлах:  
* *model.js* - Model - отвечает за данные приложения
* *view.js* - View - отвечает за отображение данных пользователю
* *presenter.js* - Presenter - обрабатывает события от пользователя и обновляет модель и представление.

Также в структуре проекта имеются следующие составляющие:  
* *event.js* - содержит класс, реализующий механизм событий, для взаимодействия View с Presenter.
* *i18n.js* - содержит локализацию сервиса
* *app.js* - точка входа  

### Интеграция с Telegram

[Документация][webapp_docs] Telegram Mini App.  

#### Начало работы  

Предисловие: далее по тексту могут встериться конструкции вида (файл:номер_строки), например (file.js:123) - это говорит о местонахождении обговоренного ранее функционала в файлах проекта.

В теге *head* в файле *index.html* необходимо подключить скрипт для интеграции с Telegram.

```html
<script src="https://telegram.org/js/telegram-web-app.js"></script>
```

Скрипты приложения, взаимодействующие с Telegram должны подключаться **после** подключения *telegram-web-app.js*.  

При инициализации приложения необходимо вызвать функцию *ready()* - Метод, который информирует приложение Telegram о том, что мини-приложение готово к отображению. (presenter.js:10)  

```js
Telegram.WebApp.ready();
```

#### Заглушка  

При открытии приложения вне Telegram Mini App (такое возможно, т.к это всё еще обычное веб приложение доступное по ссылке) показывается информационный блок и интерфейс приложения скрывается.  

В данном проекте это реализовано с помощью проверки нданных, получаемых от Telegram при запуске Mini App (). (presenter.js:12)  

```js
if ((!Telegram.WebApp.initDataUnsafe || !Telegram.WebApp.initDataUnsafe.query_id)) {
    this.view.showUnavailableBlock();
    this.view.showApp();
    return;
}
```  

#### Стилизация  

Мини приложение может получить информацию о установленной теме и основных цветах, используемых в telegram клиенте пользователя.  

[Подробннее в документации][theme_params]

Изменение стиля интерфейса осуществляется с помощью css переменных.   
Для изменения темы в мини-приложении при изменении темы пользователем в клиенте telegram Telegram.WebApp позволяет установить обработчк на соответствующее событие. (view.js:29)  

```js
bindTelegramThemeChangedAction(handler) {
    Telegram.WebApp.onEvent('themeChanged', handler)
}
``` 

[Подробннее о событиях в документации][events_list]

#### Облачное хранилище  

Пользователь мини-приложения может индивидуально настроить список валют для получения актуального курса.  
Список выбранных пользователем валют сохраняется в [облачном хранилище Telegram][cloud_storage].  
Функции работы с облачным хранилищем представлены в файле *model.js*  
Пример функции сохранения пользовательских данных: (model.js:189)  

```js
async saveUserSymbols() {
    let symbols = this._userSymbols.map(item => item.s);
    let data = JSON.stringify(symbols);
    Telegram.WebApp.CloudStorage.setItem('symbols', data, (err, ok) => {
        if (err) {
            console.log("Save user symbols error ", err);
            this.alertEvent.notify('Save user symbols error: ' + err);
        } else {
            if (ok) {
                console.log("User symbols saved in cloud storage!")
            }
        }
    });
}
```

#### Интернационализация  

Проект поддерживает два языка, которые выбираются автоматически в зависимости от значения `user.language_code`, получаемого от Telegram при старте мини-приложения. (model.js:33)


```js
if (this.initDataUnsafe.user !== undefined) {
    if (this.initDataUnsafe.user.language_code !== undefined) {
        if (this.langs.has(this.initDataUnsafe.user.language_code)) {
            this._userlang = this.initDataUnsafe.user.language_code
        }
    }
}
```


#### Валидация данных от мини-приложения 

Для подтверждения того, что запрос к API поступил именно от 
Telegram Mini App при инициализации мини-приложения передает `hash`. С его помощью можно получить подтверждение того, что запрос к API мини-приложения выполнен именно от мини-приложения.  
В проекте, при запросе к API на получении курсов валют передается значение `Telegram.WebApp.initData`. (model.js:119)  

```js
async apiRequest(path, data) {
    data.auth = this.initData
    try {
        let res = await fetch(this.apiUrl + path + "?" + new URLSearchParams(data), {
            method: 'get'
        });

        let d = await res.json();
        return d;
    } catch (err) {
        return { ok: false, data: err }
    }
}
```

[Валидация на серверной части][auth_on_back]

[Подробнее о валидации в документации][auth]

#### Кнопка "назад" 

В мини-приложении пользователь может открыть интерфейс выбора валюты - "экран" полностью перекрывающий главный экран.  
Возврат на главный экран реалзиован с помощью кнопки "назад" в интерфейсе Telegram Mini App.  
При переходе на экран выбора валюты вызывается фукция отображения кнопки "назад". (presenter.js:81)

```js
openSymbolsList() {
    this.view.openSymbolsList();
    Telegram.WebApp.BackButton.show();
}
```

При нажатии на кнопку вызывается обработчик (view.js:25)

```js
bindTelegramBackButtonAction(handler) {
    Telegram.WebApp.BackButton.onClick(handler);
}
```

[Подробнее о кнопке "назад"][back_button]


[//]: # (LINKS)
[webpack]: https://webpack.js.org/
[mvp]: https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93presenter
[webapp_docs]: https://core.telegram.org/bots/webapps
[theme_params]: https://core.telegram.org/bots/webapps#themeparams
[events_list]: https://core.telegram.org/bots/webapps#events-available-for-mini-apps
[cloud_storage]: https://core.telegram.org/bots/webapps#cloudstorage
[auth]:https://core.telegram.org/bots/webapps#validating-data-received-via-the-mini-app
[back_button]:https://core.telegram.org/bots/webapps#backbutton
[auth_on_back]:./backend_ru.md
