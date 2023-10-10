import Event from './event'

// default symbols (currencies)
const defaultUserSymbols = [
    {
        s: "EUR",
        hasImg: true
    },
    {
        s: "USD",
        hasImg: true
    },
];

class Model {
    constructor(apiUrl) {
        this.apiUrl = apiUrl;
        this.maxConvRows = 8; // max inputs rows
        this._userSymbols = defaultUserSymbols
        this._selectedRowId = -1;
        this._baseId = 0; //base symbol (currency) id
        this._lastUpdate = 0;
        this._rates = {};
        this._amount = 1;

        this.langs = new Set(['en', 'ru']);
        this._userlang = 'en';

        this.initData = Telegram.WebApp.initData || '';
        this.initDataUnsafe = Telegram.WebApp.initDataUnsafe || {};
        
        // set user language
        if (this.initDataUnsafe.user !== undefined) {
            if (this.initDataUnsafe.user.language_code !== undefined) {
                if (this.langs.has(this.initDataUnsafe.user.language_code)) {
                    this._userlang = this.initDataUnsafe.user.language_code
                }
            }
        }

        this.addInputRowEvent = new Event();
        this.alertEvent = new Event();
    }

    get userSymbolsLen() {
        return this._userSymbols.length;
    }

    get maxInputRows() {
        return this.maxConvRows
    }

    set selectedRowId(v) {
        this._selectedRowId = v;
    }

    get selectedRowId() {
        return this._selectedRowId;
    }

    get userlang() {
        return this._userlang;
    }

    get baseId() {
        return this._baseId;
    }
    
    set baseId(v) {
        this._baseId = v;
        this.saveBaseId();
    }

    get rates() {
        return this._rates;
    }

    get userSymbols() {
        return this._userSymbols;
    }

    get lastUpdate() {
        return this._lastUpdate;
    }

    get amount() {
        return this._amount;
    }

    set amount(v) {
        this._amount = v;
    }

    getSymbolTranslation(symbol) {
        return this._rates.translations[symbol][this._userlang];
    }

    async loadRates() {
        let reqData = {
            tr: 1 // with translations
        };

        // fetch rates
        let response = await this.apiRequest('/rates', reqData);
        if (response.ok) {
            this._lastUpdate = response.data.timestamp;

            this._rates.data = response.data.symbols;
            this._rates.translations = response.data.translations;

            return true;
        } else {
            console.log("Error on load rates", response.data);
            this.alertEvent.notify('Error on load rates: ' + response.data, true);
            return false;
        }
    }

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

    getUserData(callback, showAlert) {
        Telegram.WebApp.CloudStorage.getItems(['symbols', 'baseId'], (err, values) => {
            if (err) {
                console.log("get user data error ", err);
                this.alertEvent.notify('Get user data error: ' + err, true);
            } else {
                if (values['symbols'].length > 0) {
                    let data = JSON.parse(values['symbols']);
                    
                    let n = data.length;
                    if (n > 2) {
                        let i = 2;
                        while(i < n) {
                            this.addInputRowEvent.notify(i);
                            ++i;
                        }
                    }

                    this._userSymbols = new Array(n);
                    data.forEach((v, i) => {
                        this._userSymbols[i] = {
                            s: v,
                            hasImg: this._rates.data[v].hasImg
                        };
                    });
                } else {
                    console.log('Symbols not found in Tg storage')
                }

                if (values['baseId'].length > 0) {
                    let bid = parseInt(values['baseId']);
                    this.baseId = bid < this._userSymbols.length ? bid : 0;

                } else {
                    console.log('BaseId not found in Tg storage')
                }

                this.calcRates();
                callback();
            }
        });
    }

    async saveBaseId() {
        Telegram.WebApp.CloudStorage.setItem('baseId', this._baseId, (err, ok) => {
            if (err) {
                console.log("Save baseId error ", err);
                this.alertEvent.notify('Save baseId error: ' + err);
            } else {
                if (ok) {
                    console.log("BaseId saved in cloud storage!")
                }
            }
        });
    }

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

    calcRates() {
        let base = this._userSymbols[this.baseId].s;
        for (let i = 0; i < this._userSymbols.length; ++i) {
            let s = this._userSymbols[i].s;
            this._userSymbols[i].rate = this._rates.data[s].rate/this._rates.data[base].rate;
        }
    }

    getUserSymbol(id) {
        return this._userSymbols[id];
    }

    deleteSymbol(id) {
        this._userSymbols.splice(id, 1);
    }

    addUserSymbol(symbol) {
        this._userSymbols.push({
            s: symbol,
            hasImg: this._rates.data[symbol].hasImg,
            rate: this._rates.data[symbol].rate
        });
    }

    changeUserSymbol(symbol) {
        this._userSymbols[this._selectedRowId] = {
            s: symbol,
            hasImg: this._rates.data[symbol].hasImg,
            rate: this._rates.data[symbol].rate
        };
    }

    getRowAmount(rowId) {
        let x = this._userSymbols[rowId].rate * this._amount;
            
        let p = 2;
        if (x < 1) {
            p = 6;
        } else if (x > 1e6) {
            p = 0;
        }

        return parseFloat(x.toFixed(p)).toString().replace('.', ',');
    }
}

export default Model;