import Pristine from 'pristinejs';

import base from '@/layout/base';
import infoAside from '@/components/info-aside/info-aside';

import '@/styles/common/main-flow.scss';
import '@/styles/common/categories-list.scss';
import '@/styles/common/category.scss';
import '@/styles/common/tags-cloud.scss';
import '@/styles/common/tag.scss';
import '@/styles/common/post.scss';

import '@/pages/post/styles/post.scss';

import './styles/editor.scss';
import './styles/post-form.scss';
import './styles/preview.scss';
import './styles/switch-mode.scss';
import { md2html } from '@/scripts/api/helpers';

const root = document.querySelector('.editor-js');

/** @type {HTMLFormElement} */
const $form = root.querySelector('.post-form-js');

const $title = root.querySelector('.preview-title-js');
const $description = root.querySelector('.preview-description-js');
const $img = root.querySelector('.preview-img-js');
const $content = root.querySelector('.preview-content-js');

const $submit = root.querySelector('.submit-post-js');

const validator = new Pristine($form, {
   classTo: 'post-form__field',
   errorClass: 'post-form__field--error',
   errorTextParent: 'post-form__field',
   errorTextClass: 'post-form__error',
}, false);

$form.noValidate = false;

/** @type {Post} */
const post = {};

base();
infoAside();

initEvents();

function initEvents() {
   root.addEventListener('click', ({ target }) => {
      const btn = target.closest('.switch-btn-js');
      if (btn) changeMode(btn.dataset.mode);
   });

   $submit.addEventListener('click', () => {
      submit();
   });

   $form.addEventListener('submit', (e) => {
      e.preventDefault();
      $form.reportValidity();
   });
}

async function submit() {
   if (!validator.validate()) return;
   updatePostData();

   console.log(post);
   
}

/** @param {'preview'|'edit'} mode */
async function changeMode(mode) {
   if (mode === 'preview') {
      await updatePreview();
   }
   root.dataset.mode = mode;
}

async function updatePreview() {
   updatePostData();

   $title.innerHTML = post.title || 'No title';
   $description.innerHTML = post.description;
   $content.innerHTML = await getMd();

   $img.src = post.image || randomImg();
}

function updatePostData() {
   post.title = $form.title.value.trim();
   post.description = $form.description.value.trim();
   post.content = $form.content.value.trim();
   post.image = $form.image.value.trim();
   post.author = $form.author.value.trim();
   post.tags = parseTags($form.tags.value);
   post.category = $form.category.value.trim();
   post.password = $form.password.value.trim();
}

function parseTags(raw) {
   return raw.split(',').map(t => t.trim()).filter(t => t);
}

async function getMd() {
   try {
      return await md2html(post.content) || 'No content';
   } catch (e) {
      if (typeof e !== 'string') throw e;
      alert(e);
   }

   return 'No content';
}

function randomImg() {
   return `https://picsum.photos/1200/800?${Date.now()}`;
}