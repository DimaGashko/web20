import { setClassIf } from '@/utils/base';

import './styles/menu.scss';

export default function () {
   const root = document.querySelector('.menu');
   const $showBtn = root.querySelector('.show-menu-js');

   /** @type {HTMLElement} */
   const $list = root.querySelector('.menu-list-js');

   let expanded = false;

   setExpanded(expanded);
   $showBtn.addEventListener('click', () => {
      setExpanded(!expanded);
   });

   function setExpanded(val) {
      expanded = val;
      $showBtn.setAttribute('aria-expanded', expanded);
      setClassIf($list, 'menu__list--expanded', expanded);
   }
}