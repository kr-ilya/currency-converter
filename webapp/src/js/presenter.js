class Presenter {
    constructor(view, model) {
        this.view = view;
        this.model = model;

        this.init();
    }

    async init() {
        Telegram.WebApp.ready();
        this.view.setColorScheme(Telegram.WebApp.colorScheme);
        if ((!Telegram.WebApp.initDataUnsafe || !Telegram.WebApp.initDataUnsafe.query_id)) {
            this.view.showUnavailableBlock();
            this.view.showApp();
            return;
        }

        this.view.addRow(0);
        this.view.addRow(1);

        this.inputRowHandlers = {
            'hFocus': this.onInputFocused.bind(this),
            'hEnterAmount': this.onEnterAmount.bind(this),
            'hDel': this.onDeleteInputRowClicked.bind(this),
            'hChangeSymbol': this.onChangeSymbolClicked.bind(this)
        }

        this.view.bindInputEvents(0, this.inputRowHandlers);
        this.view.bindInputEvents(1, this.inputRowHandlers);

        this.view.i18n(this.model.userlang);

        this.view.bindAddSymbol(() => {
            this.onAddSymbolClicked();
        });
        
        this.view.bindInputSearch((t) => {
            this.onInputSearch(t);
        });
        
        this.view.bindClearSearch(() => {
            this.onClearSearchClicked();
        });

        this.view.bindTelegramBackButtonAction(() => {
            this.onTgBackButtonClicked();
        });
        
        this.view.bindTelegramThemeChangedAction(() => {
            this.onTgThemeChanged();
        });

        this.view.bindSelectSymbol((e) => {
            this.onSelectSymbolClicked(e);
        });

        this.model.addInputRowEvent.attach(this.onAddInputRowEvent.bind(this));
        this.model.alertEvent.attach(this.onAlertEvent.bind(this));

        this.view.showApp();

        let r = await this.model.loadRates();
        if (!r) {
            this.view.disableApp();
            return
        }

        this.model.getUserData(() => {
            this.updateInputRows();
            this.view.createSymbolsList(this.model.rates, this.model.userlang);
            this.view.createSelectedSymbolsList(this.model.userSymbols, this.model.rates, this.model.userlang);

            if (this.model.userSymbolsLen >= this.model.maxInputRows) {
                this.view.hideAddInputRowBtn();
            }
        });

        this.view.setLastUpdate(this.model.lastUpdate, this.model.userlang);
    }

    openSymbolsList() {
        this.view.openSymbolsList();
        Telegram.WebApp.BackButton.show();
    }

    onAddSymbolClicked (e) {
        if (this.model.userSymbolsLen >= this.model.maxInputRows) {
            return
        }

        this.model.selectedRowId = -1;

        this.openSymbolsList();
    }

    onInputSearch(t) {
        let v = t.target.value.toLowerCase();

        // reset lists
        if (v === '') {
            this.view.resetSymbolsLists();
            return;
        }
        
        // show clear input button
        this.view.showSearchClearBtn();

        // searching algo
        for (let symbol in this.model.rates.data) {
            let sEl = this.view.searchGetRow(symbol);
            let symbolTranslation = this.model.getSymbolTranslation(symbol);

            if (!!~symbol.toLowerCase().indexOf(v) ||
                !!~symbolTranslation.toLowerCase().indexOf(v)){
                    this.view.searchShowRow(sEl);
            } else {
                this.view.searchHideRow(sEl);
            }
        }
        
        this.view.stylizationSearchRows();
    }

    onClearSearchClicked() {
        this.view.clearInputSearch();
    }

    onTgBackButtonClicked() {
        this.view.closeSymbolsList();
    }

    onAddInputRowEvent(id) {
        this.view.addRow(id);
        this.view.bindInputEvents(id, this.inputRowHandlers);
    }

    updateInputRows() {
        for (let i = 0; i < this.model.userSymbolsLen; ++i) {
            this.onUpdateInputRow(i);
        }
    }

    onUpdateInputRow(rowId) {
        let symbolData = this.model.getUserSymbol(rowId);
        let amnt = this.model.getRowAmount(rowId);

        this.view.updateRow(rowId, symbolData, amnt, this.model.baseId);
    }

    onChangeSymbolClicked(t) {
        let rowId = this.view.getInputRowId(t.target);
        this.model.selectedRowId = rowId;
        this.openSymbolsList();
    }

    onSelectSymbolClicked(e) {
        let t = this.view.getClosestSymbolRowEl(e.target);
        if (!t) return;

        let symbol = this.view.getSymbolFromEl(t);
        let title = this.model.rates.translations[symbol][this.model.userlang];
        let hasImg = this.model.rates.data[symbol].hasImg;

        // on add row
        if (this.model.selectedRowId == -1) {
            this.model.addUserSymbol(symbol);

            let sid = this.model.userSymbolsLen-1;
            let symbolData = this.model.getUserSymbol(sid);
            let amount = this.model.getRowAmount(sid);

            this.view.addRow(sid, false, symbolData, amount);
            this.view.bindInputEvents(sid, this.inputRowHandlers);
            this.view.addSymbolToListOfSelected(symbol, title, hasImg);
        // on change symbol
        } else {
            let oldSymbol = this.model.userSymbols[this.model.selectedRowId].s;
            
            this.model.changeUserSymbol(symbol);
            this.model.calcRates();

            let amount = this.model.getRowAmount(this.model.selectedRowId);

            let symbolData = this.model.getUserSymbol(this.model.selectedRowId);
            this.view.updateRow(this.model.selectedRowId, symbolData, amount, this.model.baseId);

            this.view.changeSymbolInListOfSelected(oldSymbol, symbol, title, hasImg);

            for (let i = 0; i < this.model.userSymbolsLen; ++i) {
                if (i != this.model.selectedRowId) {
                    let amount = this.model.getRowAmount(i);
                    this.view.updateRowAmount(i, amount);
                }
            }
        }

        this.model.saveUserSymbols();

        if (this.model.userSymbolsLen >= this.model.maxInputRows) {
            this.view.hideAddInputRowBtn();
        }

        if (this.view.getSearchInputVal() != '') {
            this.view.clearInputSearch();
        }

        this.view.closeSymbolsList();
        this.view.stylizationFirstSymbolRow();
    }

    onDeleteInputRowClicked(e) {
        let t = e.currentTarget;
        let rowId = this.view.getInputRowIdFromParent(t);
        let symbol = this.model.getUserSymbol(rowId).s;

        this.model.deleteSymbol(rowId);

        if (rowId == this.model.baseId) {
            this.model.baseId = 0;
            this.view.setBaseInput(0);
        } else if (rowId < this.model.baseId) {
            this.model.baseId = rowId - 1;
        }
        
        this.model.saveUserSymbols();

        if (this.model.userSymbolsLen < this.model.maxInputRows) {
            this.view.showAddInputRowBtn();
        }

        t.parentElement.remove();
        
        this.view.renumberInputRows(rowId);

        this.view.delSymbolFromListOfSelected(symbol);
        this.view.updRowsMargin();

        if (this.view.getSearchInputVal() != '') {
            this.view.stylizationSearchRows();
        }

        this.view.stylizationFirstSymbolRow();
    }

    onEnterAmount(t) {
        let rowId = this.view.getInputRowId(t);

        this.view.resetBaseInput(this.model.baseId);

        if (this.model.baseId != rowId) {
            this.model.baseId = rowId;
        }

        this.view.setBaseInput(this.model.baseId);

        this.model.amount = parseFloat(t.value.replace(",", ".")) || 0;
        this.model.calcRates();

        for (let i = 0; i < this.model.userSymbolsLen; ++i) {
            if (i !== rowId) {
                let amount = this.model.getRowAmount(i);
                this.view.updateRowAmount(i, amount);
            }
        }
    }

    onInputFocused(t) {
        this.view.onInputFocused(t.currentTarget);
    }

    onAlertEvent(args) {
        this.view.showAlert(...args);
    }

    onTgThemeChanged() {
        this.view.setColorScheme(Telegram.WebApp.colorScheme);
    }
}

export default Presenter;