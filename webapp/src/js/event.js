class Event {
    constructor() {
        this.listeners = [];
    }

    attach(listener) {
        if (typeof listener !== 'function') return;
        this.listeners.push(listener);
    }

    notify(...args) {
        let l = this.listeners.length;
        for (let i = 0; i < l; ++i) {
            this.listeners[i](args);
        }
    }
}

export default Event;