const layoutPages = [
   page('base-layout', './pages/base-layout'),
];

const pages = [
   page('home', './pages/home'),
   page('add-article', './pages/add-article'),
];

const errorPages = [
   page('err500', './pages/errors/err500'),
   page('err404', './pages/errors/err404'),
   page('err403', './pages/errors/err403'),
];

const allPages = [
   ...pages,
   ...errorPages,
   ...layoutPages,
   // ...layoutPages,
];

function page(name, path) {
   return {
      name,
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