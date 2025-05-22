const path = require('path')
const VueLoaderPlugin = require('vue-loader/lib/plugin')
const webpack = require('webpack');
require('dotenv').config();

module.exports = {
    mode: 'production',
    entry: './src/mount.js',
    output: {
        filename: 'vue2-app.umd.js',
        path: path.resolve(__dirname, 'dist'),
        // library: 'Vue2App',
        // libraryTarget: 'umd',
        globalObject: `(typeof self !== 'undefined' ? self : this)`,
        publicPath: '/legacy/',
    },
    optimization: {
        splitChunks: false,
        runtimeChunk: false
    },
    resolve: {
        extensions: ['.ts', '.tsx', '.js', '.vue', '.json'],
        alias: {
            '@': path.resolve(__dirname, 'src'),
            'vue$': 'vue/dist/vue.esm.js',
        }
    },
    // externals: {
    //     vue: 'Vue',
    // },
    module: {
        rules: [
            {
                test: /\.vue$/,
                use: [
                    'cache-loader',
                    {
                        loader: 'vue-loader',
                        options: {
                            compilerOptions: { preserveWhitespace: false },
                        }
                    }
                ]
            },
            {
                test: /\.ts$/,
                exclude: /node_modules/,
                use: [
                    'cache-loader',
                    {
                        loader: 'ts-loader',
                        options: {
                            appendTsSuffixTo: [/\.vue$/],
                            transpileOnly: true,
                            configFile: path.resolve(__dirname, 'tsconfig.json')
                        }
                    }
                ]
            },
            {
                test: /\.js$/,
                exclude: /node_modules/,
                use: ['babel-loader']
            },
            {
                test: /\.css$/,
                use: [
                    'vue-style-loader',
                    'css-loader',
                    'postcss-loader'
                ]
            },
            {
                test: /\.scss$/,
                use: [
                    'vue-style-loader',
                    'css-loader',
                    'postcss-loader',
                    'sass-loader'
                ]
            },
            {
                test: /\.(png|jpe?g|gif|svg|woff2?|eot|ttf|otf)(\?.*)?$/,
                loader: 'file-loader',
                options: {
                    name: 'assets/[name].[hash:8].[ext]'
                }
            }
        ]
    },
    plugins: [
        new VueLoaderPlugin(),
        new webpack.DefinePlugin({
            'process.env.VUE_APP_BACKEND_BASE_URL': JSON.stringify(process.env.VUE_APP_BACKEND_BASE_URL),
            'process.env.NODE_ENV': JSON.stringify(process.env.NODE_ENV || 'development')
        })
    ]
}
