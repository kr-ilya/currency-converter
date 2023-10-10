# Frontend mini-app

Frontend mini-applications (WebApp).

## Project structure

Main files and folders:  

* */src* - folder with the source code of the service
* *nginx.conf* - nginx configuration file (web server for distributing static content)
* *package.json* - a file containing information about the project and its dependencies
* *webpack.config.js* - configuration file for Webpack
* *Dockerfile* - the file contains instructions for creating a Docker image

## Technical features of the project  

The service is written in javascript. The javascript application build tool is used - [Webpack][webpack]. It allows you to optimize and package resources (such as js files, styles, images) into an optimized bundle for deployment on websites. Webpack also allows you to use a modular system, which makes it easier to organize the code.

When building a js project, css files are minimized, combined into bundles (several js files into one js file, several css files into one css).

When starting a project using docker compose, in addition to building the project, the nginx web server is installed to return static content.

## Building a project  

When starting a project using docker compose, the service is assembled automatically in a container (the *npm run build* command is executed).

The service is built in the development environment (dev) by the npm run dev command. 

## Internal structure of the service  

The service is written following the [MVP][mvp] application design pattern (Model-View-Presenter).  
In accordance with the pattern, there are three main classes in the structure of the service, located in the files of the same name:  
* *model.js* - Model - responsible for application data
* *view.js* - View - responsible for displaying data to the user
* *presenter.js* - Presenter - processes events from the user and updates the model and view.

Also in the structure of the project there are the following components:
* *event.js* - contains a class that implements an event mechanism for interaction between the View and the Presenter.
* *i18n.js* - contains localization of the service
* *app.js* - entry point  

### Integration with Telegram

[Documentation][webapp_docs] Telegram Mini App.  

#### Getting started  

Preface: further along the text, constructions of the form (file:line_number), for example (file.js:123) - this indicates the location of the previously discussed functionality in the project files.

In the *head* tag in the file *index.html* it is necessary to connect the script for integration with Telegram.

```html
<script src="https://telegram.org/js/telegram-web-app.js"></script>
```

Application scripts interacting with Telegram should be connected **after** connection *telegram-web-app.js*.

When initializing the application, you need to call the *ready() function* - A method that informs the Telegram application that the mini-application is ready to display. (presenter.js:10)

```js
Telegram.WebApp.ready();
```

#### Plug  

When opening the application outside of the Telegram Mini App (this is possible, because it is still a regular web application available by link), an information block is shown and the application interface is hidden.  

In this project, this is implemented by checking the data received from Telegram when launching the Mini App. (presenter.js:12)


```js
if ((!Telegram.WebApp.initDataUnsafe || !Telegram.WebApp.initDataUnsafe.query_id)) {
    this.view.showUnavailableBlock();
    this.view.showApp();
    return;
}
```  

#### Stylization  

The mini application can get information about the installed theme and the primary colors used in the user's telegram client.  

[More details in the documentation][theme_params]

The interface style is changed using css variables.   
To change the theme in the mini-application when the user changes the theme in the telegram Telegram client.WebApp allows you to set a handler for the corresponding event. (view.js:29)  

```js
bindTelegramThemeChangedAction(handler) {
    Telegram.WebApp.onEvent('themeChanged', handler)
}
``` 

[More information about events in the documentation][events_list]

#### Cloud storage  

The user of the mini-application can individually configure the list of currencies to get the current exchange rate.  
The list of currencies selected by the user is saved in [Telegram cloud storage][cloud_storage].  
The functions of working with cloud storage are presented in the file *model.js*  
Example of user data saving function: (model.js:189)  

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

#### Internationalization  

The project supports two languages, which are selected automatically depending on the value of `user.language_code` received from Telegram at the start of the mini-application. (model.js:33)


```js
if (this.initDataUnsafe.user !== undefined) {
    if (this.initDataUnsafe.user.language_code !== undefined) {
        if (this.langs.has(this.initDataUnsafe.user.language_code)) {
            this._userlang = this.initDataUnsafe.user.language_code
        }
    }
}
```


#### Validation of data from the mini-application 

To confirm that the API request came from 
Telegram Mini App sends a `hash` when initializing the mini-application. With its help, you can get confirmation that the request to the API of the mini-application was made from the mini-application.  
In the project, when requesting the API to receive exchange rates, the value `Telegram' is passed.WebApp.initData`. (model.js:119)  

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

[Validation on the server side][auth_on_back]

[Read more about validation in the documentation][auth]

#### Back button 

In the mini-application, the user can open the currency selection interface - a "screen" that completely overlaps the main screen.  
The return to the main screen is implemented using the "back" button in the Telegram Mini App interface.  
When switching to the currency selection screen, the display of the "back" button is called. (presenter.js:81)

```js
openSymbolsList() {
    this.view.openSymbolsList();
    Telegram.WebApp.BackButton.show();
}
```

When the button is clicked, the handler is called (view.js:25)

```js
bindTelegramBackButtonAction(handler) {
    Telegram.WebApp.BackButton.onClick(handler);
}
```

[More about the "back" button][back_button]


[//]: # (LINKS)
[webpack]: https://webpack.js.org/
[mvp]: https://en.wikipedia.org/wiki/Model%E2%80%93view%E2%80%93presenter
[webapp_docs]: https://core.telegram.org/bots/webapps
[theme_params]: https://core.telegram.org/bots/webapps#themeparams
[events_list]: https://core.telegram.org/bots/webapps#events-available-for-mini-apps
[cloud_storage]: https://core.telegram.org/bots/webapps#cloudstorage
[auth]:https://core.telegram.org/bots/webapps#validating-data-received-via-the-mini-app
[back_button]:https://core.telegram.org/bots/webapps#backbutton
[auth_on_back]:./backend_en.md
