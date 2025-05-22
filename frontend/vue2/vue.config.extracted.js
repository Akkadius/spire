{
  mode: 'development',
  context: '/home/eqemu/server/spire/frontend/vue2',
  devtool: 'cheap-module-eval-source-map',
  node: {
    setImmediate: false,
    dgram: 'empty',
    fs: 'empty',
    net: 'empty',
    tls: 'empty',
    child_process: 'empty'
  },
  output: {
    path: '/home/eqemu/server/spire/frontend/vue2/dist',
    filename: '[name].[hash].js',
    publicPath: '/',
    globalObject: '(typeof self !== \'undefined\' ? self : this)'
  },
  resolve: {
    alias: {
      '@': '/home/eqemu/server/spire/frontend/vue2/src',
      vue$: 'vue/dist/vue.esm.js'
    },
    extensions: [
      '.mjs',
      '.js',
      '.jsx',
      '.vue',
      '.json',
      '.wasm',
      '.ts',
      '.tsx'
    ],
    modules: [
      'node_modules',
      '/home/eqemu/server/spire/frontend/vue2/node_modules',
      '/home/eqemu/server/spire/frontend/vue2/node_modules/@vue/cli-service/node_modules'
    ]
  },
  resolveLoader: {
    modules: [
      '/home/eqemu/server/spire/frontend/vue2/node_modules/@vue/cli-plugin-typescript/node_modules',
      '/home/eqemu/server/spire/frontend/vue2/node_modules/@vue/cli-plugin-babel/node_modules',
      'node_modules',
      '/home/eqemu/server/spire/frontend/vue2/node_modules',
      '/home/eqemu/server/spire/frontend/vue2/node_modules/@vue/cli-service/node_modules'
    ]
  },
  module: {
    noParse: /^(vue|vue-router|vuex|vuex-router-sync)$/,
    rules: [
      /* config.module.rule('vue') */
      {
        test: /\.vue$/,
        use: [
          /* config.module.rule('vue').use('cache-loader') */
          {
            loader: 'cache-loader',
            options: {
              cacheDirectory: '/home/eqemu/server/spire/frontend/vue2/node_modules/.cache/vue-loader',
              cacheIdentifier: '0ef69236'
            }
          },
          /* config.module.rule('vue').use('vue-loader') */
          {
            loader: 'vue-loader',
            options: {
              compilerOptions: {
                preserveWhitespace: false
              },
              cacheDirectory: '/home/eqemu/server/spire/frontend/vue2/node_modules/.cache/vue-loader',
              cacheIdentifier: '0ef69236'
            }
          }
        ]
      },
      /* config.module.rule('images') */
      {
        test: /\.(png|jpe?g|gif|webp)(\?.*)?$/,
        use: [
          /* config.module.rule('images').use('url-loader') */
          {
            loader: 'url-loader',
            options: {
              limit: 4096,
              fallback: {
                loader: 'file-loader',
                options: {
                  name: 'img/[name].[hash:8].[ext]'
                }
              }
            }
          }
        ]
      },
      /* config.module.rule('svg') */
      {
        test: /\.(svg)(\?.*)?$/,
        use: [
          /* config.module.rule('svg').use('file-loader') */
          {
            loader: 'file-loader',
            options: {
              name: 'img/[name].[hash:8].[ext]'
            }
          }
        ]
      },
      /* config.module.rule('media') */
      {
        test: /\.(mp4|webm|ogg|mp3|wav|flac|aac)(\?.*)?$/,
        use: [
          /* config.module.rule('media').use('url-loader') */
          {
            loader: 'url-loader',
            options: {
              limit: 4096,
              fallback: {
                loader: 'file-loader',
                options: {
                  name: 'media/[name].[hash:8].[ext]'
                }
              }
            }
          }
        ]
      },
      /* config.module.rule('fonts') */
      {
        test: /\.(woff2?|eot|ttf|otf)(\?.*)?$/i,
        use: [
          /* config.module.rule('fonts').use('url-loader') */
          {
            loader: 'url-loader',
            options: {
              limit: 4096,
              fallback: {
                loader: 'file-loader',
                options: {
                  name: 'fonts/[name].[hash:8].[ext]'
                }
              }
            }
          }
        ]
      },
      /* config.module.rule('pug') */
      {
        test: /\.pug$/,
        oneOf: [
          /* config.module.rule('pug').oneOf('pug-vue') */
          {
            resourceQuery: /vue/,
            use: [
              /* config.module.rule('pug').oneOf('pug-vue').use('pug-plain-loader') */
              {
                loader: 'pug-plain-loader'
              }
            ]
          },
          /* config.module.rule('pug').oneOf('pug-template') */
          {
            use: [
              /* config.module.rule('pug').oneOf('pug-template').use('raw') */
              {
                loader: 'raw-loader'
              },
              /* config.module.rule('pug').oneOf('pug-template').use('pug-plain') */
              {
                loader: 'pug-plain-loader'
              }
            ]
          }
        ]
      },
      /* config.module.rule('css') */
      {
        test: /\.css$/,
        oneOf: [
          /* config.module.rule('css').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('css').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('css').oneOf('vue-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('css').oneOf('vue-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          },
          /* config.module.rule('css').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('css').oneOf('vue').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('css').oneOf('vue').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('css').oneOf('vue').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          },
          /* config.module.rule('css').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('css').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('css').oneOf('normal-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('css').oneOf('normal-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          },
          /* config.module.rule('css').oneOf('normal') */
          {
            use: [
              /* config.module.rule('css').oneOf('normal').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('css').oneOf('normal').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('css').oneOf('normal').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          }
        ]
      },
      /* config.module.rule('postcss') */
      {
        test: /\.p(ost)?css$/,
        oneOf: [
          /* config.module.rule('postcss').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('postcss').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('postcss').oneOf('vue-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('postcss').oneOf('vue-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          },
          /* config.module.rule('postcss').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('postcss').oneOf('vue').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('postcss').oneOf('vue').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('postcss').oneOf('vue').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          },
          /* config.module.rule('postcss').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('postcss').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('postcss').oneOf('normal-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('postcss').oneOf('normal-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          },
          /* config.module.rule('postcss').oneOf('normal') */
          {
            use: [
              /* config.module.rule('postcss').oneOf('normal').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('postcss').oneOf('normal').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('postcss').oneOf('normal').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          }
        ]
      },
      /* config.module.rule('scss') */
      {
        test: /\.scss$/,
        oneOf: [
          /* config.module.rule('scss').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('scss').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('scss').oneOf('vue-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('scss').oneOf('vue-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('scss').oneOf('vue-modules').use('sass-loader') */
              {
                loader: 'sass-loader',
                options: {
                  sourceMap: false,
                  implementation: {
                    load: function () { /* omitted long function */ },
                    compile: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileString: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileStringAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initAsyncCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    Compiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    AsyncCompiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    Value: function Value0() {
                        },
                    SassBoolean: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassArgumentList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassCalculation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationOperation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationInterpolation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassColor: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassFunction: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMixin: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMap: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassNumber: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassString: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    sassNull: {},
                    sassTrue: {
                      value: true
                    },
                    sassFalse: {
                      value: false
                    },
                    Exception: function () { /* omitted long function */ },
                    Logger: {
                      silent: {
                        warn: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                },
                        debug: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                }
                      }
                    },
                    NodePackageImporter: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    deprecations: {
                      'call-string': {
                        id: 'call-string',
                        status: 'active',
                        description: 'Passing a string directly to meta.call().',
                        deprecatedIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        },
                        obsoleteIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        }
                      },
                      elseif: {
                        id: 'elseif',
                        status: 'active',
                        description: '@elseif.',
                        deprecatedIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        }
                      },
                      'moz-document': {
                        id: 'moz-document',
                        status: 'active',
                        description: '@-moz-document.',
                        deprecatedIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        }
                      },
                      'relative-canonical': {
                        id: 'relative-canonical',
                        status: 'active',
                        description: 'Imports using relative canonical URLs.',
                        deprecatedIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        }
                      },
                      'new-global': {
                        id: 'new-global',
                        status: 'active',
                        description: 'Declaring new variables with !global.',
                        deprecatedIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        }
                      },
                      'color-module-compat': {
                        id: 'color-module-compat',
                        status: 'active',
                        description: 'Using color module functions in place of plain CSS functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        }
                      },
                      'slash-div': {
                        id: 'slash-div',
                        status: 'active',
                        description: '/ operator for division.',
                        deprecatedIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        }
                      },
                      'bogus-combinators': {
                        id: 'bogus-combinators',
                        status: 'active',
                        description: 'Leading, trailing, and repeated combinators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        }
                      },
                      'strict-unary': {
                        id: 'strict-unary',
                        status: 'active',
                        description: 'Ambiguous + and - operators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        }
                      },
                      'function-units': {
                        id: 'function-units',
                        status: 'active',
                        description: 'Passing invalid units to built-in functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        }
                      },
                      'duplicate-var-flags': {
                        id: 'duplicate-var-flags',
                        status: 'active',
                        description: 'Using !default or !global multiple times for one variable.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        }
                      },
                      'null-alpha': {
                        id: 'null-alpha',
                        status: 'active',
                        description: 'Passing null as alpha in the JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        }
                      },
                      'abs-percent': {
                        id: 'abs-percent',
                        status: 'active',
                        description: 'Passing percentages to the Sass abs() function.',
                        deprecatedIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        }
                      },
                      'fs-importer-cwd': {
                        id: 'fs-importer-cwd',
                        status: 'active',
                        description: 'Using the current working directory as an implicit load path.',
                        deprecatedIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        }
                      },
                      'css-function-mixin': {
                        id: 'css-function-mixin',
                        status: 'active',
                        description: 'Function and mixin names beginning with --.',
                        deprecatedIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        }
                      },
                      'mixed-decls': {
                        id: 'mixed-decls',
                        status: 'active',
                        description: 'Declarations after or between nested rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        }
                      },
                      'feature-exists': {
                        id: 'feature-exists',
                        status: 'active',
                        description: 'meta.feature-exists',
                        deprecatedIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        }
                      },
                      'color-4-api': {
                        id: 'color-4-api',
                        status: 'active',
                        description: 'Certain uses of built-in sass:color functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'color-functions': {
                        id: 'color-functions',
                        status: 'active',
                        description: 'Using global color functions instead of sass:color.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'legacy-js-api': {
                        id: 'legacy-js-api',
                        status: 'active',
                        description: 'Legacy JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'import': {
                        id: 'import',
                        status: 'active',
                        description: '@import rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'global-builtin': {
                        id: 'global-builtin',
                        status: 'active',
                        description: 'Global built-in functions that are available in sass: modules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'user-authored': {
                        id: 'user-authored',
                        status: 'user',
                        description: null,
                        deprecatedIn: null,
                        obsoleteIn: null
                      }
                    },
                    Version: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    loadParserExports_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    info: 'dart-sass\t1.83.1\t(Sass Compiler)\t[Dart]\ndart2js\t3.6.0\t(Dart Compiler)\t[Dart]',
                    render: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    renderSync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    types: {
                      Boolean: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Color: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      List: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Map: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Null: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Number: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      String: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Error: function Error() { [native code] }
                    },
                    NULL: {},
                    TRUE: {
                      value: true
                    },
                    FALSE: {
                      value: false
                    },
                    cli_pkg_main_0_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            }
                  }
                }
              }
            ]
          },
          /* config.module.rule('scss').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('scss').oneOf('vue').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('scss').oneOf('vue').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('scss').oneOf('vue').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('scss').oneOf('vue').use('sass-loader') */
              {
                loader: 'sass-loader',
                options: {
                  sourceMap: false,
                  implementation: {
                    load: function () { /* omitted long function */ },
                    compile: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileString: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileStringAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initAsyncCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    Compiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    AsyncCompiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    Value: function Value0() {
                        },
                    SassBoolean: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassArgumentList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassCalculation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationOperation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationInterpolation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassColor: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassFunction: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMixin: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMap: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassNumber: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassString: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    sassNull: {},
                    sassTrue: {
                      value: true
                    },
                    sassFalse: {
                      value: false
                    },
                    Exception: function () { /* omitted long function */ },
                    Logger: {
                      silent: {
                        warn: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                },
                        debug: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                }
                      }
                    },
                    NodePackageImporter: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    deprecations: {
                      'call-string': {
                        id: 'call-string',
                        status: 'active',
                        description: 'Passing a string directly to meta.call().',
                        deprecatedIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        },
                        obsoleteIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        }
                      },
                      elseif: {
                        id: 'elseif',
                        status: 'active',
                        description: '@elseif.',
                        deprecatedIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        }
                      },
                      'moz-document': {
                        id: 'moz-document',
                        status: 'active',
                        description: '@-moz-document.',
                        deprecatedIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        }
                      },
                      'relative-canonical': {
                        id: 'relative-canonical',
                        status: 'active',
                        description: 'Imports using relative canonical URLs.',
                        deprecatedIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        }
                      },
                      'new-global': {
                        id: 'new-global',
                        status: 'active',
                        description: 'Declaring new variables with !global.',
                        deprecatedIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        }
                      },
                      'color-module-compat': {
                        id: 'color-module-compat',
                        status: 'active',
                        description: 'Using color module functions in place of plain CSS functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        }
                      },
                      'slash-div': {
                        id: 'slash-div',
                        status: 'active',
                        description: '/ operator for division.',
                        deprecatedIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        }
                      },
                      'bogus-combinators': {
                        id: 'bogus-combinators',
                        status: 'active',
                        description: 'Leading, trailing, and repeated combinators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        }
                      },
                      'strict-unary': {
                        id: 'strict-unary',
                        status: 'active',
                        description: 'Ambiguous + and - operators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        }
                      },
                      'function-units': {
                        id: 'function-units',
                        status: 'active',
                        description: 'Passing invalid units to built-in functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        }
                      },
                      'duplicate-var-flags': {
                        id: 'duplicate-var-flags',
                        status: 'active',
                        description: 'Using !default or !global multiple times for one variable.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        }
                      },
                      'null-alpha': {
                        id: 'null-alpha',
                        status: 'active',
                        description: 'Passing null as alpha in the JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        }
                      },
                      'abs-percent': {
                        id: 'abs-percent',
                        status: 'active',
                        description: 'Passing percentages to the Sass abs() function.',
                        deprecatedIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        }
                      },
                      'fs-importer-cwd': {
                        id: 'fs-importer-cwd',
                        status: 'active',
                        description: 'Using the current working directory as an implicit load path.',
                        deprecatedIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        }
                      },
                      'css-function-mixin': {
                        id: 'css-function-mixin',
                        status: 'active',
                        description: 'Function and mixin names beginning with --.',
                        deprecatedIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        }
                      },
                      'mixed-decls': {
                        id: 'mixed-decls',
                        status: 'active',
                        description: 'Declarations after or between nested rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        }
                      },
                      'feature-exists': {
                        id: 'feature-exists',
                        status: 'active',
                        description: 'meta.feature-exists',
                        deprecatedIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        }
                      },
                      'color-4-api': {
                        id: 'color-4-api',
                        status: 'active',
                        description: 'Certain uses of built-in sass:color functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'color-functions': {
                        id: 'color-functions',
                        status: 'active',
                        description: 'Using global color functions instead of sass:color.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'legacy-js-api': {
                        id: 'legacy-js-api',
                        status: 'active',
                        description: 'Legacy JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'import': {
                        id: 'import',
                        status: 'active',
                        description: '@import rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'global-builtin': {
                        id: 'global-builtin',
                        status: 'active',
                        description: 'Global built-in functions that are available in sass: modules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'user-authored': {
                        id: 'user-authored',
                        status: 'user',
                        description: null,
                        deprecatedIn: null,
                        obsoleteIn: null
                      }
                    },
                    Version: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    loadParserExports_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    info: 'dart-sass\t1.83.1\t(Sass Compiler)\t[Dart]\ndart2js\t3.6.0\t(Dart Compiler)\t[Dart]',
                    render: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    renderSync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    types: {
                      Boolean: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Color: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      List: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Map: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Null: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Number: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      String: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Error: function Error() { [native code] }
                    },
                    NULL: {},
                    TRUE: {
                      value: true
                    },
                    FALSE: {
                      value: false
                    },
                    cli_pkg_main_0_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            }
                  }
                }
              }
            ]
          },
          /* config.module.rule('scss').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('scss').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('scss').oneOf('normal-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('scss').oneOf('normal-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('scss').oneOf('normal-modules').use('sass-loader') */
              {
                loader: 'sass-loader',
                options: {
                  sourceMap: false,
                  implementation: {
                    load: function () { /* omitted long function */ },
                    compile: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileString: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileStringAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initAsyncCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    Compiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    AsyncCompiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    Value: function Value0() {
                        },
                    SassBoolean: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassArgumentList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassCalculation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationOperation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationInterpolation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassColor: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassFunction: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMixin: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMap: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassNumber: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassString: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    sassNull: {},
                    sassTrue: {
                      value: true
                    },
                    sassFalse: {
                      value: false
                    },
                    Exception: function () { /* omitted long function */ },
                    Logger: {
                      silent: {
                        warn: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                },
                        debug: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                }
                      }
                    },
                    NodePackageImporter: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    deprecations: {
                      'call-string': {
                        id: 'call-string',
                        status: 'active',
                        description: 'Passing a string directly to meta.call().',
                        deprecatedIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        },
                        obsoleteIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        }
                      },
                      elseif: {
                        id: 'elseif',
                        status: 'active',
                        description: '@elseif.',
                        deprecatedIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        }
                      },
                      'moz-document': {
                        id: 'moz-document',
                        status: 'active',
                        description: '@-moz-document.',
                        deprecatedIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        }
                      },
                      'relative-canonical': {
                        id: 'relative-canonical',
                        status: 'active',
                        description: 'Imports using relative canonical URLs.',
                        deprecatedIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        }
                      },
                      'new-global': {
                        id: 'new-global',
                        status: 'active',
                        description: 'Declaring new variables with !global.',
                        deprecatedIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        }
                      },
                      'color-module-compat': {
                        id: 'color-module-compat',
                        status: 'active',
                        description: 'Using color module functions in place of plain CSS functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        }
                      },
                      'slash-div': {
                        id: 'slash-div',
                        status: 'active',
                        description: '/ operator for division.',
                        deprecatedIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        }
                      },
                      'bogus-combinators': {
                        id: 'bogus-combinators',
                        status: 'active',
                        description: 'Leading, trailing, and repeated combinators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        }
                      },
                      'strict-unary': {
                        id: 'strict-unary',
                        status: 'active',
                        description: 'Ambiguous + and - operators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        }
                      },
                      'function-units': {
                        id: 'function-units',
                        status: 'active',
                        description: 'Passing invalid units to built-in functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        }
                      },
                      'duplicate-var-flags': {
                        id: 'duplicate-var-flags',
                        status: 'active',
                        description: 'Using !default or !global multiple times for one variable.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        }
                      },
                      'null-alpha': {
                        id: 'null-alpha',
                        status: 'active',
                        description: 'Passing null as alpha in the JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        }
                      },
                      'abs-percent': {
                        id: 'abs-percent',
                        status: 'active',
                        description: 'Passing percentages to the Sass abs() function.',
                        deprecatedIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        }
                      },
                      'fs-importer-cwd': {
                        id: 'fs-importer-cwd',
                        status: 'active',
                        description: 'Using the current working directory as an implicit load path.',
                        deprecatedIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        }
                      },
                      'css-function-mixin': {
                        id: 'css-function-mixin',
                        status: 'active',
                        description: 'Function and mixin names beginning with --.',
                        deprecatedIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        }
                      },
                      'mixed-decls': {
                        id: 'mixed-decls',
                        status: 'active',
                        description: 'Declarations after or between nested rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        }
                      },
                      'feature-exists': {
                        id: 'feature-exists',
                        status: 'active',
                        description: 'meta.feature-exists',
                        deprecatedIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        }
                      },
                      'color-4-api': {
                        id: 'color-4-api',
                        status: 'active',
                        description: 'Certain uses of built-in sass:color functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'color-functions': {
                        id: 'color-functions',
                        status: 'active',
                        description: 'Using global color functions instead of sass:color.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'legacy-js-api': {
                        id: 'legacy-js-api',
                        status: 'active',
                        description: 'Legacy JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'import': {
                        id: 'import',
                        status: 'active',
                        description: '@import rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'global-builtin': {
                        id: 'global-builtin',
                        status: 'active',
                        description: 'Global built-in functions that are available in sass: modules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'user-authored': {
                        id: 'user-authored',
                        status: 'user',
                        description: null,
                        deprecatedIn: null,
                        obsoleteIn: null
                      }
                    },
                    Version: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    loadParserExports_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    info: 'dart-sass\t1.83.1\t(Sass Compiler)\t[Dart]\ndart2js\t3.6.0\t(Dart Compiler)\t[Dart]',
                    render: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    renderSync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    types: {
                      Boolean: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Color: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      List: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Map: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Null: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Number: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      String: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Error: function Error() { [native code] }
                    },
                    NULL: {},
                    TRUE: {
                      value: true
                    },
                    FALSE: {
                      value: false
                    },
                    cli_pkg_main_0_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            }
                  }
                }
              }
            ]
          },
          /* config.module.rule('scss').oneOf('normal') */
          {
            use: [
              /* config.module.rule('scss').oneOf('normal').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('scss').oneOf('normal').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('scss').oneOf('normal').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('scss').oneOf('normal').use('sass-loader') */
              {
                loader: 'sass-loader',
                options: {
                  sourceMap: false,
                  implementation: {
                    load: function () { /* omitted long function */ },
                    compile: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileString: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileStringAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initAsyncCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    Compiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    AsyncCompiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    Value: function Value0() {
                        },
                    SassBoolean: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassArgumentList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassCalculation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationOperation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationInterpolation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassColor: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassFunction: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMixin: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMap: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassNumber: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassString: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    sassNull: {},
                    sassTrue: {
                      value: true
                    },
                    sassFalse: {
                      value: false
                    },
                    Exception: function () { /* omitted long function */ },
                    Logger: {
                      silent: {
                        warn: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                },
                        debug: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                }
                      }
                    },
                    NodePackageImporter: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    deprecations: {
                      'call-string': {
                        id: 'call-string',
                        status: 'active',
                        description: 'Passing a string directly to meta.call().',
                        deprecatedIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        },
                        obsoleteIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        }
                      },
                      elseif: {
                        id: 'elseif',
                        status: 'active',
                        description: '@elseif.',
                        deprecatedIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        }
                      },
                      'moz-document': {
                        id: 'moz-document',
                        status: 'active',
                        description: '@-moz-document.',
                        deprecatedIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        }
                      },
                      'relative-canonical': {
                        id: 'relative-canonical',
                        status: 'active',
                        description: 'Imports using relative canonical URLs.',
                        deprecatedIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        }
                      },
                      'new-global': {
                        id: 'new-global',
                        status: 'active',
                        description: 'Declaring new variables with !global.',
                        deprecatedIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        }
                      },
                      'color-module-compat': {
                        id: 'color-module-compat',
                        status: 'active',
                        description: 'Using color module functions in place of plain CSS functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        }
                      },
                      'slash-div': {
                        id: 'slash-div',
                        status: 'active',
                        description: '/ operator for division.',
                        deprecatedIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        }
                      },
                      'bogus-combinators': {
                        id: 'bogus-combinators',
                        status: 'active',
                        description: 'Leading, trailing, and repeated combinators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        }
                      },
                      'strict-unary': {
                        id: 'strict-unary',
                        status: 'active',
                        description: 'Ambiguous + and - operators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        }
                      },
                      'function-units': {
                        id: 'function-units',
                        status: 'active',
                        description: 'Passing invalid units to built-in functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        }
                      },
                      'duplicate-var-flags': {
                        id: 'duplicate-var-flags',
                        status: 'active',
                        description: 'Using !default or !global multiple times for one variable.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        }
                      },
                      'null-alpha': {
                        id: 'null-alpha',
                        status: 'active',
                        description: 'Passing null as alpha in the JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        }
                      },
                      'abs-percent': {
                        id: 'abs-percent',
                        status: 'active',
                        description: 'Passing percentages to the Sass abs() function.',
                        deprecatedIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        }
                      },
                      'fs-importer-cwd': {
                        id: 'fs-importer-cwd',
                        status: 'active',
                        description: 'Using the current working directory as an implicit load path.',
                        deprecatedIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        }
                      },
                      'css-function-mixin': {
                        id: 'css-function-mixin',
                        status: 'active',
                        description: 'Function and mixin names beginning with --.',
                        deprecatedIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        }
                      },
                      'mixed-decls': {
                        id: 'mixed-decls',
                        status: 'active',
                        description: 'Declarations after or between nested rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        }
                      },
                      'feature-exists': {
                        id: 'feature-exists',
                        status: 'active',
                        description: 'meta.feature-exists',
                        deprecatedIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        }
                      },
                      'color-4-api': {
                        id: 'color-4-api',
                        status: 'active',
                        description: 'Certain uses of built-in sass:color functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'color-functions': {
                        id: 'color-functions',
                        status: 'active',
                        description: 'Using global color functions instead of sass:color.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'legacy-js-api': {
                        id: 'legacy-js-api',
                        status: 'active',
                        description: 'Legacy JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'import': {
                        id: 'import',
                        status: 'active',
                        description: '@import rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'global-builtin': {
                        id: 'global-builtin',
                        status: 'active',
                        description: 'Global built-in functions that are available in sass: modules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'user-authored': {
                        id: 'user-authored',
                        status: 'user',
                        description: null,
                        deprecatedIn: null,
                        obsoleteIn: null
                      }
                    },
                    Version: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    loadParserExports_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    info: 'dart-sass\t1.83.1\t(Sass Compiler)\t[Dart]\ndart2js\t3.6.0\t(Dart Compiler)\t[Dart]',
                    render: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    renderSync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    types: {
                      Boolean: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Color: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      List: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Map: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Null: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Number: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      String: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Error: function Error() { [native code] }
                    },
                    NULL: {},
                    TRUE: {
                      value: true
                    },
                    FALSE: {
                      value: false
                    },
                    cli_pkg_main_0_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            }
                  }
                }
              }
            ]
          }
        ]
      },
      /* config.module.rule('sass') */
      {
        test: /\.sass$/,
        oneOf: [
          /* config.module.rule('sass').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('sass').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('sass').oneOf('vue-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('sass').oneOf('vue-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('sass').oneOf('vue-modules').use('sass-loader') */
              {
                loader: 'sass-loader',
                options: {
                  sourceMap: false,
                  implementation: {
                    load: function () { /* omitted long function */ },
                    compile: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileString: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileStringAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initAsyncCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    Compiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    AsyncCompiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    Value: function Value0() {
                        },
                    SassBoolean: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassArgumentList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassCalculation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationOperation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationInterpolation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassColor: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassFunction: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMixin: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMap: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassNumber: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassString: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    sassNull: {},
                    sassTrue: {
                      value: true
                    },
                    sassFalse: {
                      value: false
                    },
                    Exception: function () { /* omitted long function */ },
                    Logger: {
                      silent: {
                        warn: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                },
                        debug: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                }
                      }
                    },
                    NodePackageImporter: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    deprecations: {
                      'call-string': {
                        id: 'call-string',
                        status: 'active',
                        description: 'Passing a string directly to meta.call().',
                        deprecatedIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        },
                        obsoleteIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        }
                      },
                      elseif: {
                        id: 'elseif',
                        status: 'active',
                        description: '@elseif.',
                        deprecatedIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        }
                      },
                      'moz-document': {
                        id: 'moz-document',
                        status: 'active',
                        description: '@-moz-document.',
                        deprecatedIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        }
                      },
                      'relative-canonical': {
                        id: 'relative-canonical',
                        status: 'active',
                        description: 'Imports using relative canonical URLs.',
                        deprecatedIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        }
                      },
                      'new-global': {
                        id: 'new-global',
                        status: 'active',
                        description: 'Declaring new variables with !global.',
                        deprecatedIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        }
                      },
                      'color-module-compat': {
                        id: 'color-module-compat',
                        status: 'active',
                        description: 'Using color module functions in place of plain CSS functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        }
                      },
                      'slash-div': {
                        id: 'slash-div',
                        status: 'active',
                        description: '/ operator for division.',
                        deprecatedIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        }
                      },
                      'bogus-combinators': {
                        id: 'bogus-combinators',
                        status: 'active',
                        description: 'Leading, trailing, and repeated combinators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        }
                      },
                      'strict-unary': {
                        id: 'strict-unary',
                        status: 'active',
                        description: 'Ambiguous + and - operators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        }
                      },
                      'function-units': {
                        id: 'function-units',
                        status: 'active',
                        description: 'Passing invalid units to built-in functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        }
                      },
                      'duplicate-var-flags': {
                        id: 'duplicate-var-flags',
                        status: 'active',
                        description: 'Using !default or !global multiple times for one variable.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        }
                      },
                      'null-alpha': {
                        id: 'null-alpha',
                        status: 'active',
                        description: 'Passing null as alpha in the JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        }
                      },
                      'abs-percent': {
                        id: 'abs-percent',
                        status: 'active',
                        description: 'Passing percentages to the Sass abs() function.',
                        deprecatedIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        }
                      },
                      'fs-importer-cwd': {
                        id: 'fs-importer-cwd',
                        status: 'active',
                        description: 'Using the current working directory as an implicit load path.',
                        deprecatedIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        }
                      },
                      'css-function-mixin': {
                        id: 'css-function-mixin',
                        status: 'active',
                        description: 'Function and mixin names beginning with --.',
                        deprecatedIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        }
                      },
                      'mixed-decls': {
                        id: 'mixed-decls',
                        status: 'active',
                        description: 'Declarations after or between nested rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        }
                      },
                      'feature-exists': {
                        id: 'feature-exists',
                        status: 'active',
                        description: 'meta.feature-exists',
                        deprecatedIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        }
                      },
                      'color-4-api': {
                        id: 'color-4-api',
                        status: 'active',
                        description: 'Certain uses of built-in sass:color functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'color-functions': {
                        id: 'color-functions',
                        status: 'active',
                        description: 'Using global color functions instead of sass:color.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'legacy-js-api': {
                        id: 'legacy-js-api',
                        status: 'active',
                        description: 'Legacy JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'import': {
                        id: 'import',
                        status: 'active',
                        description: '@import rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'global-builtin': {
                        id: 'global-builtin',
                        status: 'active',
                        description: 'Global built-in functions that are available in sass: modules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'user-authored': {
                        id: 'user-authored',
                        status: 'user',
                        description: null,
                        deprecatedIn: null,
                        obsoleteIn: null
                      }
                    },
                    Version: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    loadParserExports_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    info: 'dart-sass\t1.83.1\t(Sass Compiler)\t[Dart]\ndart2js\t3.6.0\t(Dart Compiler)\t[Dart]',
                    render: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    renderSync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    types: {
                      Boolean: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Color: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      List: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Map: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Null: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Number: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      String: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Error: function Error() { [native code] }
                    },
                    NULL: {},
                    TRUE: {
                      value: true
                    },
                    FALSE: {
                      value: false
                    },
                    cli_pkg_main_0_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            }
                  },
                  indentedSyntax: true
                }
              }
            ]
          },
          /* config.module.rule('sass').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('sass').oneOf('vue').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('sass').oneOf('vue').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('sass').oneOf('vue').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('sass').oneOf('vue').use('sass-loader') */
              {
                loader: 'sass-loader',
                options: {
                  sourceMap: false,
                  implementation: {
                    load: function () { /* omitted long function */ },
                    compile: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileString: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileStringAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initAsyncCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    Compiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    AsyncCompiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    Value: function Value0() {
                        },
                    SassBoolean: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassArgumentList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassCalculation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationOperation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationInterpolation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassColor: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassFunction: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMixin: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMap: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassNumber: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassString: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    sassNull: {},
                    sassTrue: {
                      value: true
                    },
                    sassFalse: {
                      value: false
                    },
                    Exception: function () { /* omitted long function */ },
                    Logger: {
                      silent: {
                        warn: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                },
                        debug: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                }
                      }
                    },
                    NodePackageImporter: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    deprecations: {
                      'call-string': {
                        id: 'call-string',
                        status: 'active',
                        description: 'Passing a string directly to meta.call().',
                        deprecatedIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        },
                        obsoleteIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        }
                      },
                      elseif: {
                        id: 'elseif',
                        status: 'active',
                        description: '@elseif.',
                        deprecatedIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        }
                      },
                      'moz-document': {
                        id: 'moz-document',
                        status: 'active',
                        description: '@-moz-document.',
                        deprecatedIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        }
                      },
                      'relative-canonical': {
                        id: 'relative-canonical',
                        status: 'active',
                        description: 'Imports using relative canonical URLs.',
                        deprecatedIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        }
                      },
                      'new-global': {
                        id: 'new-global',
                        status: 'active',
                        description: 'Declaring new variables with !global.',
                        deprecatedIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        }
                      },
                      'color-module-compat': {
                        id: 'color-module-compat',
                        status: 'active',
                        description: 'Using color module functions in place of plain CSS functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        }
                      },
                      'slash-div': {
                        id: 'slash-div',
                        status: 'active',
                        description: '/ operator for division.',
                        deprecatedIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        }
                      },
                      'bogus-combinators': {
                        id: 'bogus-combinators',
                        status: 'active',
                        description: 'Leading, trailing, and repeated combinators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        }
                      },
                      'strict-unary': {
                        id: 'strict-unary',
                        status: 'active',
                        description: 'Ambiguous + and - operators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        }
                      },
                      'function-units': {
                        id: 'function-units',
                        status: 'active',
                        description: 'Passing invalid units to built-in functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        }
                      },
                      'duplicate-var-flags': {
                        id: 'duplicate-var-flags',
                        status: 'active',
                        description: 'Using !default or !global multiple times for one variable.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        }
                      },
                      'null-alpha': {
                        id: 'null-alpha',
                        status: 'active',
                        description: 'Passing null as alpha in the JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        }
                      },
                      'abs-percent': {
                        id: 'abs-percent',
                        status: 'active',
                        description: 'Passing percentages to the Sass abs() function.',
                        deprecatedIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        }
                      },
                      'fs-importer-cwd': {
                        id: 'fs-importer-cwd',
                        status: 'active',
                        description: 'Using the current working directory as an implicit load path.',
                        deprecatedIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        }
                      },
                      'css-function-mixin': {
                        id: 'css-function-mixin',
                        status: 'active',
                        description: 'Function and mixin names beginning with --.',
                        deprecatedIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        }
                      },
                      'mixed-decls': {
                        id: 'mixed-decls',
                        status: 'active',
                        description: 'Declarations after or between nested rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        }
                      },
                      'feature-exists': {
                        id: 'feature-exists',
                        status: 'active',
                        description: 'meta.feature-exists',
                        deprecatedIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        }
                      },
                      'color-4-api': {
                        id: 'color-4-api',
                        status: 'active',
                        description: 'Certain uses of built-in sass:color functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'color-functions': {
                        id: 'color-functions',
                        status: 'active',
                        description: 'Using global color functions instead of sass:color.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'legacy-js-api': {
                        id: 'legacy-js-api',
                        status: 'active',
                        description: 'Legacy JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'import': {
                        id: 'import',
                        status: 'active',
                        description: '@import rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'global-builtin': {
                        id: 'global-builtin',
                        status: 'active',
                        description: 'Global built-in functions that are available in sass: modules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'user-authored': {
                        id: 'user-authored',
                        status: 'user',
                        description: null,
                        deprecatedIn: null,
                        obsoleteIn: null
                      }
                    },
                    Version: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    loadParserExports_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    info: 'dart-sass\t1.83.1\t(Sass Compiler)\t[Dart]\ndart2js\t3.6.0\t(Dart Compiler)\t[Dart]',
                    render: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    renderSync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    types: {
                      Boolean: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Color: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      List: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Map: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Null: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Number: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      String: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Error: function Error() { [native code] }
                    },
                    NULL: {},
                    TRUE: {
                      value: true
                    },
                    FALSE: {
                      value: false
                    },
                    cli_pkg_main_0_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            }
                  },
                  indentedSyntax: true
                }
              }
            ]
          },
          /* config.module.rule('sass').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('sass').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('sass').oneOf('normal-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('sass').oneOf('normal-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('sass').oneOf('normal-modules').use('sass-loader') */
              {
                loader: 'sass-loader',
                options: {
                  sourceMap: false,
                  implementation: {
                    load: function () { /* omitted long function */ },
                    compile: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileString: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileStringAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initAsyncCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    Compiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    AsyncCompiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    Value: function Value0() {
                        },
                    SassBoolean: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassArgumentList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassCalculation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationOperation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationInterpolation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassColor: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassFunction: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMixin: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMap: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassNumber: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassString: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    sassNull: {},
                    sassTrue: {
                      value: true
                    },
                    sassFalse: {
                      value: false
                    },
                    Exception: function () { /* omitted long function */ },
                    Logger: {
                      silent: {
                        warn: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                },
                        debug: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                }
                      }
                    },
                    NodePackageImporter: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    deprecations: {
                      'call-string': {
                        id: 'call-string',
                        status: 'active',
                        description: 'Passing a string directly to meta.call().',
                        deprecatedIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        },
                        obsoleteIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        }
                      },
                      elseif: {
                        id: 'elseif',
                        status: 'active',
                        description: '@elseif.',
                        deprecatedIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        }
                      },
                      'moz-document': {
                        id: 'moz-document',
                        status: 'active',
                        description: '@-moz-document.',
                        deprecatedIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        }
                      },
                      'relative-canonical': {
                        id: 'relative-canonical',
                        status: 'active',
                        description: 'Imports using relative canonical URLs.',
                        deprecatedIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        }
                      },
                      'new-global': {
                        id: 'new-global',
                        status: 'active',
                        description: 'Declaring new variables with !global.',
                        deprecatedIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        }
                      },
                      'color-module-compat': {
                        id: 'color-module-compat',
                        status: 'active',
                        description: 'Using color module functions in place of plain CSS functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        }
                      },
                      'slash-div': {
                        id: 'slash-div',
                        status: 'active',
                        description: '/ operator for division.',
                        deprecatedIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        }
                      },
                      'bogus-combinators': {
                        id: 'bogus-combinators',
                        status: 'active',
                        description: 'Leading, trailing, and repeated combinators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        }
                      },
                      'strict-unary': {
                        id: 'strict-unary',
                        status: 'active',
                        description: 'Ambiguous + and - operators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        }
                      },
                      'function-units': {
                        id: 'function-units',
                        status: 'active',
                        description: 'Passing invalid units to built-in functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        }
                      },
                      'duplicate-var-flags': {
                        id: 'duplicate-var-flags',
                        status: 'active',
                        description: 'Using !default or !global multiple times for one variable.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        }
                      },
                      'null-alpha': {
                        id: 'null-alpha',
                        status: 'active',
                        description: 'Passing null as alpha in the JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        }
                      },
                      'abs-percent': {
                        id: 'abs-percent',
                        status: 'active',
                        description: 'Passing percentages to the Sass abs() function.',
                        deprecatedIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        }
                      },
                      'fs-importer-cwd': {
                        id: 'fs-importer-cwd',
                        status: 'active',
                        description: 'Using the current working directory as an implicit load path.',
                        deprecatedIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        }
                      },
                      'css-function-mixin': {
                        id: 'css-function-mixin',
                        status: 'active',
                        description: 'Function and mixin names beginning with --.',
                        deprecatedIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        }
                      },
                      'mixed-decls': {
                        id: 'mixed-decls',
                        status: 'active',
                        description: 'Declarations after or between nested rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        }
                      },
                      'feature-exists': {
                        id: 'feature-exists',
                        status: 'active',
                        description: 'meta.feature-exists',
                        deprecatedIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        }
                      },
                      'color-4-api': {
                        id: 'color-4-api',
                        status: 'active',
                        description: 'Certain uses of built-in sass:color functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'color-functions': {
                        id: 'color-functions',
                        status: 'active',
                        description: 'Using global color functions instead of sass:color.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'legacy-js-api': {
                        id: 'legacy-js-api',
                        status: 'active',
                        description: 'Legacy JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'import': {
                        id: 'import',
                        status: 'active',
                        description: '@import rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'global-builtin': {
                        id: 'global-builtin',
                        status: 'active',
                        description: 'Global built-in functions that are available in sass: modules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'user-authored': {
                        id: 'user-authored',
                        status: 'user',
                        description: null,
                        deprecatedIn: null,
                        obsoleteIn: null
                      }
                    },
                    Version: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    loadParserExports_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    info: 'dart-sass\t1.83.1\t(Sass Compiler)\t[Dart]\ndart2js\t3.6.0\t(Dart Compiler)\t[Dart]',
                    render: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    renderSync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    types: {
                      Boolean: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Color: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      List: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Map: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Null: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Number: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      String: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Error: function Error() { [native code] }
                    },
                    NULL: {},
                    TRUE: {
                      value: true
                    },
                    FALSE: {
                      value: false
                    },
                    cli_pkg_main_0_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            }
                  },
                  indentedSyntax: true
                }
              }
            ]
          },
          /* config.module.rule('sass').oneOf('normal') */
          {
            use: [
              /* config.module.rule('sass').oneOf('normal').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('sass').oneOf('normal').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('sass').oneOf('normal').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('sass').oneOf('normal').use('sass-loader') */
              {
                loader: 'sass-loader',
                options: {
                  sourceMap: false,
                  implementation: {
                    load: function () { /* omitted long function */ },
                    compile: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileString: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    compileStringAsync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    initAsyncCompiler: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    Compiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    AsyncCompiler: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    Value: function Value0() {
                        },
                    SassBoolean: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassArgumentList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassCalculation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationOperation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    CalculationInterpolation: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassColor: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassFunction: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMixin: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassList: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassMap: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassNumber: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    SassString: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    sassNull: {},
                    sassTrue: {
                      value: true
                    },
                    sassFalse: {
                      value: false
                    },
                    Exception: function () { /* omitted long function */ },
                    Logger: {
                      silent: {
                        warn: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                },
                        debug: function() {
                                  return _call(f, Array.prototype.slice.apply(arguments));
                                }
                      }
                    },
                    NodePackageImporter: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    deprecations: {
                      'call-string': {
                        id: 'call-string',
                        status: 'active',
                        description: 'Passing a string directly to meta.call().',
                        deprecatedIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        },
                        obsoleteIn: {
                          major: 0,
                          minor: 0,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '0.0.0'
                        }
                      },
                      elseif: {
                        id: 'elseif',
                        status: 'active',
                        description: '@elseif.',
                        deprecatedIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 3,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.3.2'
                        }
                      },
                      'moz-document': {
                        id: 'moz-document',
                        status: 'active',
                        description: '@-moz-document.',
                        deprecatedIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 7,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.7.2'
                        }
                      },
                      'relative-canonical': {
                        id: 'relative-canonical',
                        status: 'active',
                        description: 'Imports using relative canonical URLs.',
                        deprecatedIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 14,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.14.2'
                        }
                      },
                      'new-global': {
                        id: 'new-global',
                        status: 'active',
                        description: 'Declaring new variables with !global.',
                        deprecatedIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 17,
                          patch: 2,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.17.2'
                        }
                      },
                      'color-module-compat': {
                        id: 'color-module-compat',
                        status: 'active',
                        description: 'Using color module functions in place of plain CSS functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 23,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.23.0'
                        }
                      },
                      'slash-div': {
                        id: 'slash-div',
                        status: 'active',
                        description: '/ operator for division.',
                        deprecatedIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 33,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.33.0'
                        }
                      },
                      'bogus-combinators': {
                        id: 'bogus-combinators',
                        status: 'active',
                        description: 'Leading, trailing, and repeated combinators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 54,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.54.0'
                        }
                      },
                      'strict-unary': {
                        id: 'strict-unary',
                        status: 'active',
                        description: 'Ambiguous + and - operators.',
                        deprecatedIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 55,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.55.0'
                        }
                      },
                      'function-units': {
                        id: 'function-units',
                        status: 'active',
                        description: 'Passing invalid units to built-in functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 56,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.56.0'
                        }
                      },
                      'duplicate-var-flags': {
                        id: 'duplicate-var-flags',
                        status: 'active',
                        description: 'Using !default or !global multiple times for one variable.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.0'
                        }
                      },
                      'null-alpha': {
                        id: 'null-alpha',
                        status: 'active',
                        description: 'Passing null as alpha in the JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 62,
                          patch: 3,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.62.3'
                        }
                      },
                      'abs-percent': {
                        id: 'abs-percent',
                        status: 'active',
                        description: 'Passing percentages to the Sass abs() function.',
                        deprecatedIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 65,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.65.0'
                        }
                      },
                      'fs-importer-cwd': {
                        id: 'fs-importer-cwd',
                        status: 'active',
                        description: 'Using the current working directory as an implicit load path.',
                        deprecatedIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 73,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.73.0'
                        }
                      },
                      'css-function-mixin': {
                        id: 'css-function-mixin',
                        status: 'active',
                        description: 'Function and mixin names beginning with --.',
                        deprecatedIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 76,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.76.0'
                        }
                      },
                      'mixed-decls': {
                        id: 'mixed-decls',
                        status: 'active',
                        description: 'Declarations after or between nested rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 77,
                          patch: 7,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.77.7'
                        }
                      },
                      'feature-exists': {
                        id: 'feature-exists',
                        status: 'active',
                        description: 'meta.feature-exists',
                        deprecatedIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 78,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.78.0'
                        }
                      },
                      'color-4-api': {
                        id: 'color-4-api',
                        status: 'active',
                        description: 'Certain uses of built-in sass:color functions.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'color-functions': {
                        id: 'color-functions',
                        status: 'active',
                        description: 'Using global color functions instead of sass:color.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'legacy-js-api': {
                        id: 'legacy-js-api',
                        status: 'active',
                        description: 'Legacy JS API.',
                        deprecatedIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 79,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.79.0'
                        }
                      },
                      'import': {
                        id: 'import',
                        status: 'active',
                        description: '@import rules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'global-builtin': {
                        id: 'global-builtin',
                        status: 'active',
                        description: 'Global built-in functions that are available in sass: modules.',
                        deprecatedIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        },
                        obsoleteIn: {
                          major: 1,
                          minor: 80,
                          patch: 0,
                          preRelease: [],
                          build: [],
                          _version$_text: '1.80.0'
                        }
                      },
                      'user-authored': {
                        id: 'user-authored',
                        status: 'user',
                        description: null,
                        deprecatedIn: null,
                        obsoleteIn: null
                      }
                    },
                    Version: function() {
                              return _call(f, this, Array.prototype.slice.apply(arguments));
                            },
                    loadParserExports_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    info: 'dart-sass\t1.83.1\t(Sass Compiler)\t[Dart]\ndart2js\t3.6.0\t(Dart Compiler)\t[Dart]',
                    render: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    renderSync: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            },
                    types: {
                      Boolean: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Color: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      List: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Map: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Null: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Number: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      String: function() {
                                return _call(f, this, Array.prototype.slice.apply(arguments));
                              },
                      Error: function Error() { [native code] }
                    },
                    NULL: {},
                    TRUE: {
                      value: true
                    },
                    FALSE: {
                      value: false
                    },
                    cli_pkg_main_0_: function() {
                              return _call(f, Array.prototype.slice.apply(arguments));
                            }
                  },
                  indentedSyntax: true
                }
              }
            ]
          }
        ]
      },
      /* config.module.rule('less') */
      {
        test: /\.less$/,
        oneOf: [
          /* config.module.rule('less').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('less').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('less').oneOf('vue-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('less').oneOf('vue-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('less').oneOf('vue-modules').use('less-loader') */
              {
                loader: 'less-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          },
          /* config.module.rule('less').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('less').oneOf('vue').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('less').oneOf('vue').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('less').oneOf('vue').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('less').oneOf('vue').use('less-loader') */
              {
                loader: 'less-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          },
          /* config.module.rule('less').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('less').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('less').oneOf('normal-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('less').oneOf('normal-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('less').oneOf('normal-modules').use('less-loader') */
              {
                loader: 'less-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          },
          /* config.module.rule('less').oneOf('normal') */
          {
            use: [
              /* config.module.rule('less').oneOf('normal').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('less').oneOf('normal').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('less').oneOf('normal').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('less').oneOf('normal').use('less-loader') */
              {
                loader: 'less-loader',
                options: {
                  sourceMap: false
                }
              }
            ]
          }
        ]
      },
      /* config.module.rule('stylus') */
      {
        test: /\.styl(us)?$/,
        oneOf: [
          /* config.module.rule('stylus').oneOf('vue-modules') */
          {
            resourceQuery: /module/,
            use: [
              /* config.module.rule('stylus').oneOf('vue-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('stylus').oneOf('vue-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('stylus').oneOf('vue-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('stylus').oneOf('vue-modules').use('stylus-loader') */
              {
                loader: 'stylus-loader',
                options: {
                  sourceMap: false,
                  preferPathResolver: 'webpack'
                }
              }
            ]
          },
          /* config.module.rule('stylus').oneOf('vue') */
          {
            resourceQuery: /\?vue/,
            use: [
              /* config.module.rule('stylus').oneOf('vue').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('stylus').oneOf('vue').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('stylus').oneOf('vue').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('stylus').oneOf('vue').use('stylus-loader') */
              {
                loader: 'stylus-loader',
                options: {
                  sourceMap: false,
                  preferPathResolver: 'webpack'
                }
              }
            ]
          },
          /* config.module.rule('stylus').oneOf('normal-modules') */
          {
            test: /\.module\.\w+$/,
            use: [
              /* config.module.rule('stylus').oneOf('normal-modules').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('stylus').oneOf('normal-modules').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2,
                  modules: true,
                  localIdentName: '[name]_[local]_[hash:base64:5]'
                }
              },
              /* config.module.rule('stylus').oneOf('normal-modules').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('stylus').oneOf('normal-modules').use('stylus-loader') */
              {
                loader: 'stylus-loader',
                options: {
                  sourceMap: false,
                  preferPathResolver: 'webpack'
                }
              }
            ]
          },
          /* config.module.rule('stylus').oneOf('normal') */
          {
            use: [
              /* config.module.rule('stylus').oneOf('normal').use('vue-style-loader') */
              {
                loader: 'vue-style-loader',
                options: {
                  sourceMap: false,
                  shadowMode: false
                }
              },
              /* config.module.rule('stylus').oneOf('normal').use('css-loader') */
              {
                loader: 'css-loader',
                options: {
                  sourceMap: false,
                  importLoaders: 2
                }
              },
              /* config.module.rule('stylus').oneOf('normal').use('postcss-loader') */
              {
                loader: 'postcss-loader',
                options: {
                  sourceMap: false
                }
              },
              /* config.module.rule('stylus').oneOf('normal').use('stylus-loader') */
              {
                loader: 'stylus-loader',
                options: {
                  sourceMap: false,
                  preferPathResolver: 'webpack'
                }
              }
            ]
          }
        ]
      },
      /* config.module.rule('js') */
      {
        test: /\.m?jsx?$/,
        exclude: [
          function () { /* omitted long function */ }
        ],
        use: [
          /* config.module.rule('js').use('cache-loader') */
          {
            loader: 'cache-loader',
            options: {
              cacheDirectory: '/home/eqemu/server/spire/frontend/vue2/node_modules/.cache/babel-loader',
              cacheIdentifier: '23cb18cd'
            }
          },
          /* config.module.rule('js').use('babel-loader') */
          {
            loader: 'babel-loader'
          }
        ]
      },
      /* config.module.rule('ts') */
      {
        test: /\.ts$/,
        use: [
          /* config.module.rule('ts').use('cache-loader') */
          {
            loader: 'cache-loader',
            options: {
              cacheDirectory: '/home/eqemu/server/spire/frontend/vue2/node_modules/.cache/ts-loader',
              cacheIdentifier: '04f3d4a2'
            }
          },
          /* config.module.rule('ts').use('babel-loader') */
          {
            loader: 'babel-loader'
          },
          /* config.module.rule('ts').use('ts-loader') */
          {
            loader: 'ts-loader',
            options: {
              transpileOnly: true,
              appendTsSuffixTo: [
                '\\.vue$'
              ],
              happyPackMode: false
            }
          }
        ]
      },
      /* config.module.rule('tsx') */
      {
        test: /\.tsx$/,
        use: [
          /* config.module.rule('tsx').use('cache-loader') */
          {
            loader: 'cache-loader',
            options: {
              cacheDirectory: '/home/eqemu/server/spire/frontend/vue2/node_modules/.cache/ts-loader',
              cacheIdentifier: '04f3d4a2'
            }
          },
          /* config.module.rule('tsx').use('babel-loader') */
          {
            loader: 'babel-loader'
          },
          /* config.module.rule('tsx').use('ts-loader') */
          {
            loader: 'ts-loader',
            options: {
              transpileOnly: true,
              happyPackMode: false,
              appendTsxSuffixTo: [
                '\\.vue$'
              ]
            }
          }
        ]
      }
    ]
  },
  plugins: [
    /* config.plugin('vue-loader') */
    new VueLoaderPlugin(),
    /* config.plugin('define') */
    new DefinePlugin(
      {
        'process.env': {
          NODE_ENV: '"development"',
          BASE_URL: '"/"'
        }
      }
    ),
    /* config.plugin('case-sensitive-paths') */
    new CaseSensitivePathsPlugin(),
    /* config.plugin('friendly-errors') */
    new FriendlyErrorsWebpackPlugin(
      {
        additionalTransformers: [
          function () { /* omitted long function */ }
        ],
        additionalFormatters: [
          function () { /* omitted long function */ }
        ]
      }
    ),
    /* config.plugin('hmr') */
    new HotModuleReplacementPlugin(),
    /* config.plugin('progress') */
    new ProgressPlugin(),
    /* config.plugin('html') */
    new HtmlWebpackPlugin(
      {
        templateParameters: function () { /* omitted long function */ },
        template: '/home/eqemu/server/spire/frontend/vue2/public/index.html'
      }
    ),
    /* config.plugin('preload') */
    new PreloadPlugin(
      {
        rel: 'preload',
        include: 'initial',
        fileBlacklist: [
          /\.map$/,
          /hot-update\.js$/
        ]
      }
    ),
    /* config.plugin('copy') */
    new CopyWebpackPlugin(
      [
        {
          from: '/home/eqemu/server/spire/frontend/vue2/public',
          to: '/home/eqemu/server/spire/frontend/vue2/dist',
          toType: 'dir',
          ignore: [
            '.DS_Store',
            {
              glob: 'index.html',
              matchBase: false
            },
            'eq-asset-preview-master/**/*'
          ]
        }
      ]
    ),
    /* config.plugin('fork-ts-checker') */
    new ForkTsCheckerWebpackPlugin(
      {
        vue: true,
        tslint: false,
        formatter: 'codeframe',
        checkSyntacticErrors: false
      }
    )
  ],
  performance: {
    maxEntrypointSize: 40000000,
    maxAssetSize: 40000000
  },
  entry: {
    app: [
      './src/main.ts'
    ]
  }
}
