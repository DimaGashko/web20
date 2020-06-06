import { setClassIf } from '@/utils/base';

import './styles/menu.scss';

export default function () {
   const root = document.querySelector('.menu');
   const $showBtn = root.querySelector('.show-menu-js');
   const $list = root.querySelector('.menu-list-js');

   $showBtn.addEventListener('click', () => {
      const state = $showBtn.getAttribute('aria-expanded') === 'true';
      $showBtn.setAttribute('aria-expanded', !state);
      setClassIf($list, 'menu__list--expanded', !state);
   });
}