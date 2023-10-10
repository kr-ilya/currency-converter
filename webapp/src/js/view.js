import { translations } from './i18n'

class View {
    constructor() {
        this.minInputRows = 2;
        this.appEl = document.querySelector(`.app`);
        this.addBtn = document.querySelector('.add-input-row');
        this.searchInputEl = document.querySelector('.search-input-value');
        this.searchClearEl = document.getElementById('clear-search');
        this.lastUpdateInfoEl = document.getElementById('last-update-info');
    }

    bindAddSymbol(handler) {
        this.addBtn.addEventListener('click', handler);
    }

    bindInputSearch(handler) {
        this.searchInputEl.addEventListener('input', handler);
    }
    
    bindClearSearch(handler) {
        document.querySelector('.clear-search-button').addEventListener('click', handler);
    }
    
    bindTelegramBackButtonAction(handler) {
        Telegram.WebApp.BackButton.onClick(handler);
    }
    
    bindTelegramThemeChangedAction(handler) {
        Telegram.WebApp.onEvent('themeChanged', handler)
    }

    bindChangeSymbol(handler) {
        document.querySelector('.input-label').addEventListener('click', handler);
    }
        
    bindSelectSymbol(handler) {
        document.querySelector(`.all-currencies`).addEventListener('click', handler);
    }

    bindInputEvents(id, handlers) {
        let row = document.querySelector(`.converter-input-row[data-row-id="${id}"]`);
        let inputEl = row.querySelector(`input`);
        inputEl.addEventListener('focus', handlers.hFocus);
        inputEl.addEventListener('keyup', (t) => {
            filterLetters(t);
            handlers.hEnterAmount(t.target);
        });
        inputEl.addEventListener('keydown', filterNum);
        inputEl.addEventListener('change', filterLetters);
        inputEl.addEventListener('input', filterLetters);

        if (id > this.minInputRows-1) {
            row.querySelector('.delete-input-row').addEventListener('click', handlers.hDel);
        }

        row.querySelector('.input-label').addEventListener('click', handlers.hChangeSymbol);
    }    

    addRow(id, empty = true, symbolData = null, amnt = 0) {
        let wrap = document.querySelector(`.inputs-wrap`);

        let removeBtn = '';
        if (id > this.minInputRows-1) {
            removeBtn = `<div class="delete-input-row"">
                <svg xmlns="http://www.w3.org/2000/svg" class="icon icon-tabler icon-tabler-trash" width="24" height="24" viewBox="0 0 24 24" stroke-width="2" stroke="currentColor" fill="none" stroke-linecap="round" stroke-linejoin="round">
                    <path stroke="none" d="M0 0h24v24H0z" fill="none"/>
                    <path d="M4 7l16 0" />
                    <path d="M10 11l0 6" />
                    <path d="M14 11l0 6" />
                    <path d="M5 7l1 12a2 2 0 0 0 2 2h8a2 2 0 0 0 2 -2l1 -12" />
                    <path d="M9 7v-3a1 1 0 0 1 1 -1h4a1 1 0 0 1 1 1v3" />
                </svg>
            </div>`
        } 

        let icon = '';
        let symbol = '';
        let amount = '';
        if (!empty) {
            symbol = symbolData.s;
            let hasImg = symbolData.hasImg;

            icon = `background-image: url(assets/currencies/${symbol}.svg);`;
            if (!hasImg) {
                icon = `background-image: url(assets/currencies/money.svg);`;
            }       
            
            amount = amnt;
        }

        let row = `<div class="converter-input-row" data-row-id="${id}">
            ${removeBtn}
            <div class="converter-input converter-input-filled">
                <div class="input-label pointer">
                    <div class="currency-label">
                        <span class="currency-icon">
                            <div class="currency-icon-img" style="${icon}"></div>
                        </span>
                        <span class="currency-code">${symbol}</span>
                    </div>
                </div>
                <div class="input-wrap">
                    <input value="${amount}" maxlength="11" type="text" autocomplete="off" class="input-field converter-input-value" inputmode="decimal">
                </div>
            </div>
        </div>`;
        
        wrap.insertAdjacentHTML('beforeend', row);

        this.updRowsMargin();
    }

    updateRow(rowId, symbolData, amount, baseId) {
        var row = document.querySelector(`.converter-input-row[data-row-id="${rowId}"]`);
        if (row === null){
            return;
        }

        let iconEl = row.querySelector('.currency-icon-img');
        let codeEl = row.querySelector('.currency-code');

        let symbol = symbolData.s;
        let hasImg = symbolData.hasImg;

        var icon = symbol;
        if (!hasImg) {
            icon = 'money';
        }

        iconEl.style.backgroundImage = `url(assets/currencies/${icon}.svg)`;
        codeEl.textContent = symbol;

        if (rowId == baseId) {
            row.className = 'converter-input-row base-input';
        }
        
        this.updateRowAmount(rowId, amount);
    }

    createSymbolsList(rates, userlang) {
        let wrap = document.querySelector(`.all-currencies`);

        let list = ``;
        for(let symbol in rates.data) {
            let title = rates.translations[symbol][userlang];
            let icon = symbol;

            let hasImg = rates.data[symbol].hasImg;
            if (!hasImg) {
                icon = 'money';
            }

            list += `<div class="currency-row" data-symbol="${symbol}">
                <span class="currency-icon">
                    <div lazy-style="background-image: url(assets/currencies/${icon}.svg);"></div>
                </span>
                <div class="currency-info">
                    <span class="currency-code">${symbol}</span>
                    <span class="currency-title">${title}</span>
                </div>
            </div>`;
        }

        wrap.innerHTML = list;
        
        this.lazyLoad();
    }

    createSelectedSymbolsList(userSymbols, rates, userlang) {
        let i = 0;
        while (i < userSymbols.length) {
            let symbol = userSymbols[i].s;
            let title = rates.translations[symbol][userlang];
            let hasImg = rates.data[symbol].hasImg;

            this.addSymbolToListOfSelected(symbol, title, hasImg);
            ++i;
        }

        this.stylizationFirstSymbolRow();
    }

    addSymbolToListOfSelected(symbol, title, hasImg) {
        let wrap = document.querySelector(`.selected-currencies`);

        let icon = symbol;
        if (!hasImg) {
            icon = 'money';
        }

        let el = `<div class="currency-row" data-symbol="${symbol}">
            <span class="currency-icon">
                <div class="currency-icon-img" style="background-image: url(assets/currencies/${icon}.svg);"></div>
            </span>
            <div class="currency-info">
                <span class="currency-code">${symbol}</span>
                <span class="currency-title">${title}</span>
            </div>
            <div class="selected-symbol-mark"></div>
        </div>`;

        // hide symbol from list of all symbols
        let listAll = document.querySelector(`.all-currencies`);
        listAll.querySelector(`[data-symbol="${symbol}"]`).className = 'currency-row dnone';

        wrap.insertAdjacentHTML('beforeend', el);
    }

    changeSymbolInListOfSelected(oldSymbol, symbol, title, hasImg) {
        let wrap = document.querySelector(`.selected-currencies`);
        let row = wrap.querySelector(`[data-symbol="${oldSymbol}"]`);
        let iconEl = row.querySelector('.currency-icon-img');
        let codeEl = row.querySelector('.currency-code');
        let titleEl = row.querySelector('.currency-title');

        let icon = symbol;
        if (!hasImg) {
            icon = 'money';
        }

        iconEl.style.backgroundImage = `url(assets/currencies/${icon}.svg)`;
        codeEl.textContent = symbol;
        titleEl.textContent = title;
        row.setAttribute('data-symbol', symbol);

        let listAll = document.querySelector(`.all-currencies`);
        listAll.querySelector(`[data-symbol="${oldSymbol}"]`).className = 'currency-row';
        listAll.querySelector(`[data-symbol="${symbol}"]`).className = 'currency-row dnone';
    }

    updRowsMargin() {
        let el = document.querySelector(`.converter-inputs`);
        el.setAttribute('style', `margin-bottom: calc(100vh - var(--top-padding) - ${el.offsetHeight}px);`);
    }

    openSymbolsList() {
        this.appEl.classList.add("selecting");

        let el = document.querySelector(`.selecting-page`);
        el.setAttribute('style', `display: block;`);

        let updInfoEl = document.getElementById('last-update-info');
        updInfoEl.className = "dnone";

        window.scrollTo(0,0);
    }

    closeSymbolsList() {
        this.appEl.classList.remove("selecting");
        this.lastUpdateInfoEl.className = "";

        setTimeout(function() {
            let el = document.querySelector(`.selecting-page`);
            el.setAttribute('style', `display: none;`);
        }, 300);
        
        Telegram.WebApp.BackButton.hide();
    }

    resetSymbolsLists() {
        // hide clear input button
        this.searchClearEl.className = "input-label dnone";

        let wrap = document.querySelector(`.currencies-list`)
        wrap.querySelector(`.selected-list-label`).setAttribute('style', `display: block;`);

        let rows = wrap.querySelectorAll('.onsearch-hide, .onsearch-first-row, .onsearch-last-row');
        let l = rows.length;
        let i = 0;
        while (i < l) {
            rows[i].className = rows[i].className.replace(' onsearch-hide', '')
                .replace(' onsearch-first-row', '')
                .replace(' onsearch-last-row', '');
            ++i;
        }
    }

    showSearchClearBtn() {
        this.searchClearEl.className = "input-label";
    }

    searchGetRow(symbol) {
        let wrap = document.querySelector(`.currencies-list`);
        return wrap.querySelector(`[data-symbol="${symbol}"]`);
    }

    searchShowRow(el) {
        el.className = el.className.replace(' onsearch-hide', '')
            .replace(' onsearch-first-row', '')
            .replace(' onsearch-last-row', '');
    }

    searchHideRow(el) {
        if (el.className.indexOf('onsearch-hide') === -1) {
            el.className += ' onsearch-hide';
        }
    }

    getInputRowId(t) {
        let rowEl = t.closest('.converter-input-row');
        return parseInt(rowEl.getAttribute('data-row-id'));
    }

    getInputRowIdFromParent(t) {
        return parseInt(t.parentElement.getAttribute('data-row-id'));
    }

    getClosestSymbolRowEl(t) {
        return t.closest('.currency-row');
    }

    stylizationSearchRows() {
        // correction of border and rounding rows
        let wrap = document.querySelector(`.currencies-list`);
        ['.all-currencies', '.selected-currencies'].forEach((e) => {
            let list = wrap.querySelector(e);
            let visibleEls = list.querySelectorAll('.currency-row:not(.onsearch-hide):not(.dnone)')
            let vl = visibleEls.length
            if (vl > 0) {
                if (e == '.selected-currencies') {
                    wrap.querySelector(`.selected-list-label`).setAttribute('style', `display: block;`);
                }

                if (visibleEls[0].className.indexOf('onsearch-first-row') === -1) {
                    visibleEls[0].className += ' onsearch-first-row';
                }
                
                if (visibleEls[vl-1].className.indexOf('onsearch-last-row') === -1) {
                    visibleEls[vl-1].className += ' onsearch-last-row';
                }
            } else {
                if (e == '.selected-currencies') {
                    wrap.querySelector(`.selected-list-label`).setAttribute('style', `display: none;`);
                }
            }
        });
    }

    stylizationFirstSymbolRow() {
        let allListEl  = document.querySelector('.all-currencies');
        let lastFirstEl = allListEl.querySelector('.first-currency-row');
        if (lastFirstEl) {
            lastFirstEl.className = lastFirstEl.className.replace(' first-currency-row', '');
        }

        let firstEl = allListEl.querySelector('.currency-row:not(.dnone)');
        if (firstEl.className.indexOf('first-currency-row') === -1) {
            firstEl.className += ' first-currency-row';
        }
    }

    clearInputSearch() {
        this.searchInputEl.value = '';
        this.searchInputEl.dispatchEvent(new Event('input'));
    }

    updateRowAmount(rowId, amount) {
        var row = document.querySelector(`.converter-input-row[data-row-id="${rowId}"]`);
        if (row !== null){
            const inputEl = row.querySelector('.converter-input-value');
            inputEl.value = amount;
        }
    }

    setLastUpdate(lastUpdate, userlang) {
        let d = new Date(lastUpdate * 1000);
        let day = d.getDate();
        let month = d.getMonth()+1;
        let year = d.getFullYear();
        let hours = d.getHours();
        let minutes = d.getMinutes();

        day = (day < 10 ? '0' : '') + day;       
        month = (month < 10 ? '0' : '') + month;       
        hours = (hours < 10 ? '0' : '') + hours;       
        minutes = (minutes < 10 ? '0' : '') + minutes;       

        document.getElementById('last-update-text').innerHTML = `${translations['updateTime'][userlang]}: ${day}.${month}.${year} ${hours}:${minutes}`;
    }

    hideAddInputRowBtn() {
        document.getElementById(`add-row-btn`).setAttribute('style', `display: none;`);
        this.updRowsMargin();
    }
    
    showAddInputRowBtn() {
        document.getElementById(`add-row-btn`).setAttribute('style', `display: block;`);
        this.updRowsMargin();
    }

    setBaseInput(rowId) {
        let baseInputEl = document.querySelector(`.converter-input-row[data-row-id="${rowId}"]`);
        baseInputEl.className = 'converter-input-row base-input';
    }
    
    resetBaseInput(rowId) {
        let baseInputEl = document.querySelector(`.converter-input-row[data-row-id="${rowId}"]`);
        baseInputEl.className = 'converter-input-row';
    }

    renumberInputRows(startRowId) {
        let wrap = document.querySelector(`.inputs-wrap`);
        let childrens = wrap.children;
        let i = startRowId-1;
        let l = childrens.length;
        while (i < l) {
            childrens[i].setAttribute("data-row-id", i);
            ++i;
        }
    }

    delSymbolFromListOfSelected(symbol) {
        let wrap = document.querySelector(`.selected-currencies`);
        wrap.querySelector(`[data-symbol="${symbol}"]`).remove();

        let listAll = document.querySelector(`.all-currencies`);
        listAll.querySelector(`[data-symbol="${symbol}"]`).className = 'currency-row';
    }

    getSearchInputVal() {
        let searchEl = document.querySelector(`.search-input-value`);
        return searchEl.value;
    }

    getSymbolFromEl(el) {
        return el.getAttribute('data-symbol');
    }

    onInputFocused(t) {
        setTimeout(function(){ t.selectionStart = t.selectionEnd = 10000; }, 0);
    }

    showApp() {
        this.appEl.removeAttribute('style');
    }

    showUnavailableBlock() {
        this.appEl.innerHTML = `<section class="unavailable-app">
            <div class="unavailable-content">
                <h1>WebApp is unavailable</h1>
                <span>Come back later</span>
            </div>
        </section>`;
    }

    disableApp() {
        this.appEl.classList.add('disabled-app');
        let inputs = this.appEl.querySelectorAll('input');
        inputs.forEach((el) => {
            el.setAttribute('disabled', 'disabled');
        })
    }

    i18n(userlang) {
        document.querySelector(`.add-input-text`).innerHTML = translations['addButton'][userlang];
        document.getElementById(`last-update-text`).innerHTML = translations['updateTime'][userlang] + ': --.--.---- --:--';
        document.querySelector(`.selected-list-label`).innerHTML = translations['selectedList'][userlang];
    }

    showAlert(message, closeApp = false) {
        if (!closeApp) {
            Telegram.WebApp.showAlert(message);
        } else {
            Telegram.WebApp.showAlert(message, () => {
                Telegram.WebApp.close();
            });
        }
        
    }

    lazyLoad() {
        let imgs;    
        if ("IntersectionObserver" in window) {
            imgs = document.querySelectorAll("[lazy-style]");
            var imageObserver = new IntersectionObserver((entries, observer) => {
                entries.forEach((entry) => {
                    if (entry.isIntersecting) {
                        let image = entry.target;
                        let style = image.getAttribute('lazy-style');
                        image.setAttribute('style', style);
                        image.removeAttribute('lazy-style');
                        imageObserver.unobserve(image);
                    }
                });
            });

            imgs.forEach((image) => {
                imageObserver.observe(image);
            });
        } else {  
            var lazyloadThrottleTimeout;
            imgs = document.querySelectorAll("[lazy-style]");
            
            function lazy () {
                if(lazyloadThrottleTimeout) {
                    clearTimeout(lazyloadThrottleTimeout);
                }    

                lazyloadThrottleTimeout = setTimeout(function() {
                    var scrollTop = document.body.scrollTop || document.documentElement.scrollTop;
                    imgs.forEach((img, i)=> {
                        if(img.offsetTop < (window.innerHeight + scrollTop)) {
                            let style = img.getAttribute('lazy-style');
                            if (style){
                                img.setAttribute('style', style);
                                img.removeAttribute('lazy-style');
                            }
                        }
                    });
                    imgs = document.querySelectorAll("[lazy-style]");
                    if(imgs.length == 0) { 
                        document.removeEventListener("scroll", lazy);
                        window.removeEventListener("resize", lazy);
                        window.removeEventListener("orientationChange", lazy);
                    }
                }, 20);
            }

            document.addEventListener("scroll", lazy);
            window.addEventListener("resize", lazy);
            window.addEventListener("orientationChange", lazy);
        }
    }

    setColorScheme(scheme) {
        if (scheme) {
            document.querySelector(`html`).className = scheme;
        } else {
            document.querySelector(`html`).className = 'light';
        }
        
    }
}

function filterNum(e) {
    var key = e.key
    if (key == 'ArrowLeft' || key == 'ArrowRight' || key == 'Delete' || key == 'Backspace') {
        return true;
    }

    if(key != ',' && key != '.' && !(key >= '0' && key <= '9')) {
        return false;
    }

    if ((key == ',' || key == '.') && e.target.value.indexOf(',') !== -1) {
        return false;
    }
}

function filterLetters(e) {
    e.target.value = e.target.value.replace(/\./g, ',');

    var v = e.target.value
    if (v === ',') {
        e.target.value = '0,';
        return
    }

    if (v.replace(/\s/g, '').match(/[^,\d]/) || (v.split(',').length - 1) > 1) {
        e.target.value = v.replace(/[^\,\d]+/g, '').replace(/^([^\,]*\,)|\,/g, '$1')
    }
    
    e.target.value = e.target.value.replace(/\s/g, '')
}

export default View;