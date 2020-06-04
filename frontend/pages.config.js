const pages = [
   page('home', './pages/home'),
];

const errorPages = [
   page('err500', './pages/errors/err500'),
   page('err404', './pages/errors/err404'),
   page('err403', './pages/errors/err403'),
];

const layoutPages = [
   ...[
      // 'header',
      // 'footer',
   ].map((name) => ({
      tmpl: `./pages/layout/${name}.tmpl`,
      tmplRes: `../templates/layout/${name}.tmpl`,
      name,
   })), {
      ...page('base', './pages/layout/base'),
      tmplRes: `../templates/layout/base.tmpl`,
   },
];

const allPages = [
   ...pages,
   ...errorPages,
   // ...layoutPages,
];

function page(name, path) {
   return {
      name: name,
      entry: `${path}/${name}.js`,
      tmpl: `${path}/${name}.tmpl`,
   };
}

function getPages(pagesNames = []) {
   return [
      ...layoutPages,
      ...allPages.filter(({ name }) => pagesNames.includes(name)),
   ];
}

module.exports = {
   allPages,
   getPages,
};