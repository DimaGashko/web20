import Pristine from 'pristinejs';

import base from '@/layout/base';
import infoAside from '@/components/info-aside/info-aside';

import { readJsonData } from '@/scripts/utils/base';
import { md2html } from '@/scripts/api/helpers';
import { savePost } from '@/scripts/api/posts';

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


const root = document.querySelector('.editor-js');

/** @type {HTMLFormElement} */
const $form = root.querySelector('.post-form-js');

const $title = root.querySelector('.preview-title-js');
const $category = root.querySelector('.category-js');
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
const post = readJsonData('post') || {
   category: 'news',
   title: '',
   description: '',
   content: '',
   author: '',
   image: '',
   listed: true,
};

post.secret = '';


base();
infoAside();

initEvents();
updateForm();

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
   if (!validate()) return;
   updatePostData();

   try {
      console.log(post);
      
      const { slug } = await savePost(post);
      location.href = `/posts/${slug}`;
   } catch (e) {
      if (typeof e !== 'string') throw e;
      console.error(e);
   }
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
   $category.innerHTML = post.category;
   $description.innerHTML = post.description || 'No description';
   $content.innerHTML = await getMd();

   $img.src = post.image || randomImg();
}

function updatePostData() {
   post.category = $form.category.value;
   post.title = $form.title.value.trim();
   post.description = $form.description.value.trim();
   post.content = $form.content.value.trim();
   post.image = $form.image.value.trim();
   post.author = $form.author.value.trim();
   post.secret = $form.secret.value.trim();
   post.listed = $form.listed.checked;
}

function updateForm() {
   $form.category.value = post.category;
   $form.title.value = post.title;
   $form.description.value = post.description;
   $form.content.value = post.content;
   $form.image.value = post.image;
   $form.author.value = post.author;
   $form.listed.checked = post.listed;
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

function validate() {
   const isValid = validator.validate();
   
   if (!isValid) {
      root.scrollIntoView({ behavior: 'smooth' });
   }

   return isValid;
}