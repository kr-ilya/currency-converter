import '../css/app.css';

import Model from './model'
import View from './view'
import Presenter from './presenter'

const apiUrl = ''

const app = new Presenter(new View(), new Model(apiUrl))