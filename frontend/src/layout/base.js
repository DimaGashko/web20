import 'normalize.scss/normalize.scss';
import '@/icons/icomoon/icomoon.scss';
import '@/fonts/josefin/josefin.scss';

import header from './components/header/header.js';
import logo from './components/logo/logo.js';
import menu from './components/menu/menu.js';
import footer from './components/footer/footer.js';

import '@/styles/base.scss';

export default function() {
   header();
   logo();
   menu();
   footer();
}
