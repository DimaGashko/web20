const webpack = require('webpack');
const chokidar = require('chokidar');
const kill = require('tree-kill');
const path = require('path');
const chalk = require('chalk');

const DotenvPlugin = require('dotenv-webpack');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const MiniCssExtractPlugin = require('mini-css-extract-plugin');
const HtmlReplaceWebpackPlugin = require('html-replace-webpack-plugin');
const ScriptExtHtmlWebpackPlugin = require('script-ext-html-webpack-plugin');
const FriendlyErrorsWebpackPlugin = require('friendly-errors-webpack-plugin');
const WebpackNotifierPlugin = require('webpack-notifier');
const CopyPlugin = require('copy-webpack-plugin');
const { CleanWebpackPlugin } = require('clean-webpack-plugin');
const { BundleAnalyzerPlugin } = require('webpack-bundle-analyzer');
const { spawn } = require('child_process');

const autoprefixer = require('autoprefixer');
const normalizeCharset = require('postcss-normalize-charset');
const cssnano = require('cssnano');

const dotenv = require('dotenv').config().parsed;
const { allPages, getPages } = require('./pages.config');

const { GO_PORT, SERVE_PORT } = dotenv;

const PATHS = {
   src: path.resolve(__dirname, 'src/'),
   dist: path.resolve(__dirname, 'dist/'),
   logo: path.join(__dirname, 'src/img/logo.svg/'),
   node_modules: path.resolve(__dirname, 'node_modules/'),
};

const minifyOptions = {
   collapseWhitespace: true,
   conservativeCollapse: true,

   ignoreCustomFragments: [/\{\{[\s\S]*?\}\}/],

   caseSensitive: true,

   minifyJS: true,
   minifyCSS: true,

   removeRedundantAttributes: true,
   removeAttributeQuotes: true,
   removeComments: true,

   includeAutoGeneratedTags: false,
};

const htmlReplacements = [{
   search: '@blank',
   replace: 'target="_blank" rel="noopener noreferrer"',
   flags: 'g',
}];

let devServer = null;
let goProcess = null;

module.exports = ((env = {}) => {
   const isDev = (env.NODE_ENV === 'development');
   const hashType = (!isDev) ? '[contenthash]' : '';
   const pages = (env.pages) ? getPages(env.pages.split(',')) : allPages;

   const maps = true;
   const babel = ('babel' in env) ? env.babel !== 'false' : !isDev;
   const compact = ('compact' in env) ? env.compact !== 'false' : isDev;
   const manageGo = ('go' in env) ? env.go !== 'false' : isDev;
   const isServe = 'WEBPACK_DEV_SERVER' in process.env;

   /** @type {import('webpack/declarations/WebpackOptions').WebpackOptions} */
   const config = {
      context: PATHS.src,
      mode: (isDev) ? 'development' : 'production',
      stats: (compact) ? 'none' : 'normal',
      devtool: (maps) ? 'source-map' : false,
      entry: {
         ...getPagesEntries(pages),
      },
      output: {
         filename: `scripts/[name].js?${hashType}`,
         path: path.resolve(__dirname, 'dist/static'),
         publicPath: '/static/',
      },
      optimization: {
         splitChunks: false,
         removeAvailableModules: !isDev,
         removeEmptyChunks: !isDev,
      },
      resolve: {
         symlinks: false,
         extensions: ['.js'],
         alias: {
            '@': PATHS.src,
         },
      },
      devServer: {
         proxy: {
            '/': `http://127.0.0.1:${GO_PORT}`,
         },
         writeToDisk: (path) => {
            return /\.tmpl$/i.test(path);
         },
         port: SERVE_PORT,
         quiet: compact,
         stats: { colors: true },
         overlay: true,
         hot: true,
         watchOptions: {
            ignored: [PATHS.node_modules, PATHS.dist],
         },

         // To allow BrowserStack's bs-local on Safari 
         disableHostCheck: true,

         before(_, server) {
            devServer = server;

            watchTemplateFiles();

            if (manageGo) {
               watchGoFiles();
            }
         },
      },
      module: {
         rules: [{
            enforce: 'pre',
            test: /\.js$/,
            exclude: PATHS.node_modules,
            loader: 'eslint-loader',
         }, {
            test: /\.js$/i,
            use: [
               ...(babel ? [{
                  loader: 'babel-loader',
                  options: {
                     cacheDirectory: true,
                     sourceMaps: maps,
                     presets: [
                        ['@babel/preset-env', {
                           'targets': {
                              esmodules: true,
                           },
                        }],
                     ],
                  },
               }] : []),
            ],
            // NOTE: do not exclude node_modules, it can cause error like
            // "Cannot declare a const variable twice: 't'" in Safari 10.1
         }, {
            test: /\.(scss|css)$/i,
            sideEffects: true,
            use: [{
               loader: MiniCssExtractPlugin.loader,
               options: {
                  hmr: isDev,
                  reloadAll: false,
                  esModule: true,
               },
            },
            `css-loader?sourceMap=${maps}`,
            ...(!isDev ? [{
               loader: 'postcss-loader',
               options: {
                  plugins: [
                     normalizeCharset(),
                     autoprefixer(),
                     cssnano(),
                  ],
                  sourceMap: maps,
               },
            }] : []), {
               loader: 'sass-loader',
               options: {
                  sourceMap: maps,
                  prependData: `@import "~@/styles/disappearing";`,
               },
            }],
         }, {
            test: /\.tmpl$/i,
            include: PATHS.src,
            use: [
               'ejs-loader',
               'extract-loader',
               {
                  loader: 'html-loader',
                  options: {
                     attrs: [
                        ':src',
                        ':data-src',
                        ':srcset',
                        ':data-srcset',
                        'link:href',
                     ],
                  },
               }, {
                  loader: 'string-replace-loader',
                  options: {
                     multiple: htmlReplacements,
                  },
               },
            ],
         }, {
            test: /\.html$/i,
            include: PATHS.src,
            use: [{
               loader: 'html-loader',
               options: {
                  minimize: !isDev ? minifyOptions : false,
                  attrs: [
                     ':src',
                     ':data-src',
                     'link:href',
                  ],
               },
            }, {
               loader: 'string-replace-loader',
               options: {
                  multiple: htmlReplacements,
               },
            }],
         }, {
            test: /\.(png|jpe?g|webp|svg|gif|ico)$/i,
            loader: 'file-loader?name=[path][name].[ext]',
            include: path.resolve(__dirname, 'src/img'),
         }, {
            test: /\.(woff|woff2)$/i,
            loader: `file-loader?name=fonts/icons/[name]_v2.[ext]?${hashType}`,
            include: path.resolve(__dirname, 'src/fonts/icons'),
         }, {
            test: /\.(woff|woff2)$/i,
            loader: 'file-loader?name=fonts/[name].[ext]',
            include: path.resolve(__dirname, 'src/fonts'),
            exclude: path.resolve(__dirname, 'src/fonts/icons'),
         }],
      },
      plugins: [
         new webpack.NoEmitOnErrorsPlugin(),

         new webpack.ProgressPlugin({
            profile: false,
         }),

         new DotenvPlugin(),

         new CleanWebpackPlugin({
            cleanStaleWebpackAssets: false,

            cleanOnceBeforeBuildPatterns: ['**/*', '../templates/**/*'],
            dangerouslyAllowCleanPatternsOutsideProject: true,
            dry: false,
         }),

         ...(!env.noNotify ? [
            new WebpackNotifierPlugin({
               title: 'Giggster Webpack',
               contentImage: PATHS.logo,
               excludeWarnings: isDev,
            }),
         ] : []),

         new HtmlReplaceWebpackPlugin(
            htmlReplacements.map(({ search, replace }) => ({
               pattern: search,
               replacement: replace,
            })),
         ),

         new MiniCssExtractPlugin({
            filename: `styles/[name].css?${hashType}`,
         }),

         new ScriptExtHtmlWebpackPlugin({
            defaultAttribute: 'defer',
         }),

         ...(compact ? [
            new FriendlyErrorsWebpackPlugin({
               compilationSuccessInfo: {
                  messages: [chalk.green('Giggster (SSR Frontend)')],
                  notes: [
                     chalk`To create a production build run {blue npm run build}`,
                     (isServe) ? chalk`Project is running at {blue http://localhost:${SERVE_PORT}}` : '',
                     (manageGo) ? chalk`To reload go server run: {blue kill -SIGUSR1 ${process.pid}}\n` : '',
                  ].filter(n => n),
               },
            }),
         ] : []),

         ...(env.bundleAnalyzer ? [
            new BundleAnalyzerPlugin(),
         ] : []),

         ...getHtmlWebpackPlugins(pages, !isDev),

         ({ hooks }) => {
            if (!manageGo) return;
            let done = false;

            hooks.afterCompile.tap('ManageGoPlugin', () => {
               if (done) return;
               done = true;

               runGo();
            });

            process.on('SIGUSR1', runGo);
         },
      ],
   };

   return config;
});

function runGo() {
   if (goProcess) kill(goProcess.pid);

   goProcess = spawn('go', ['run', 'main.go'], { cwd: '../' });

   const { stdout, stderr } = goProcess;
   stdout.on('data', (data) => process.stdout.write(chalk.cyan(data)));
   stderr.on('data', (data) => process.stderr.write(chalk.cyan(data)));

   return goProcess;
}

function watchTemplateFiles() {
   chokidar.watch(`${PATHS.src}/**/*.{tmpl,html}`, {
      ignored: [PATHS.node_modules, PATHS.dist],
      awaitWriteFinish: true,
   }).on('change', reload);
}

function watchGoFiles() {
   const goPath = path.resolve(__dirname, '../**/*.go');

   chokidar.watch(goPath, {
      ignored: [path.resolve(__dirname, 'frontend')],
   }).on('change', () => {
      runGo();
      setTimeout(reload, 2000);
   });
}

function reload() {
   if (!devServer) return;
   devServer.sockWrite(devServer.sockets, 'content-changed');
}

function getPagesEntries(pages) {
   const entries = {};

   pages.forEach(({ name, entry }) => {
      if (!entry) return;
      entries[name] = entry;
   });

   return entries;
}

function getHtmlWebpackPlugins(pages, minify) {
   const defHtmlWebpackPluginOptions = {
      inject: false,
      minify: (minify) ? minifyOptions : false,
   };

   return pages.map(({ name, tmpl, tmplRes }) => {
      if (!tmpl) return null;

      return new HtmlWebpackPlugin({
         ...defHtmlWebpackPluginOptions,
         filename: tmplRes || `../templates/${name}.tmpl`,
         template: tmpl,
         chunks: [name],
      });
   }).filter(plugin => plugin);
}
