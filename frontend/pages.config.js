const baseLayout = [
   layout('./layout/base.tmpl'),
   layout('./layout/header.tmpl'),
   layout('./layout/logo.tmpl'),
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
];

const allLayouts = [
   ...baseLayout,
];

module.exports = {
   pages: [...allPages, ...allLayouts],
};

function page(name, path ) {
   return {
      name,
      entry: `${path}/${name}.js`,
      tmpl: `${path}/${name}.tmpl`,
   };
}

function layout(tmpl) {
   return { tmpl };
}
