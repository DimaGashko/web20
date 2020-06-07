const pages = [
   page('home', './pages/home'),
   page('post', './pages/post'),
   page('editor', './pages/editor'),
   page('posts-list', './pages/posts-list'),
   page('about', './pages/about'),
   page('contact-us', './pages/contact-us'),
   page('privacy', './pages/privacy'),
];

const errorPages = [
   page('err500', './pages/errors/err500'),
   page('err404', './pages/errors/err404'),
   page('err403', './pages/errors/err403'),
];

const baseLayout = [
   tmpl('header', './layout/components/header'),
   tmpl('logo', './layout/components/logo'),
   tmpl('menu', './layout/components/menu'),
   tmpl('footer', './layout/components/footer'),
   tmpl('base', './layout'),
];

const templates = [
   tmpl('info-aside', './components/info-aside'),
];

const allPages = [
   ...pages,
   ...errorPages,
];

const allLayouts = [
   ...baseLayout,
];

const allTemplates = [
   ...templates,
];

module.exports = {
   pages: [...allPages, ...allLayouts, ...allTemplates],
};

function page(name, path) {
   return {
      name,
      entry: `${path}/${name}.js`,
      tmpl: `${path}/${name}.tmpl`,
      tmplOut: `pages/${name}.tmpl`,
   };
}

function tmpl(name, path) {
   return {
      name,
      tmpl: `${path}/${name}.tmpl`,
      tmplOut: `layout/${name}.tmpl`,
   };
}