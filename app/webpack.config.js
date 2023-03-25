/* eslint-disable @typescript-eslint/no-var-requires */
const path = require('path');
// const webpack = require('webpack');

const env = process.env.BABEL_ENV || 'dev';
const isProd = env === 'production';

const wd = path.resolve(__dirname);
const srcPath = path.join(wd, 'src');
const buildDest = path.join(__dirname, 'dist');

module.exports = {
  devServer: {
    static: {
      directory: buildDest,
    },
    compress: true,
    port: 9000,
  },

  mode: isProd ? 'production' : 'development',
  entry: { main: path.join(srcPath, './main.ts') },
  output: {
    path: buildDest,
    filename: '[name].js',
  },

  module: {
    rules: [
      { test: /\.([cm]?ts|tsx)$/, loader: 'ts-loader' },
      {
        test: /\.css$/i,
        include: srcPath,
        use: ['style-loader', 'css-loader', 'postcss-loader'],
      },
    ],
  },

  resolve: {
    // modules: [path.join(wd, 'node_modules'), srcPath],
    extensions: ['.ts', '.js'],
  },

  devtool: undefined, // !isProd ? 'inline-source-map' : '',
  context: srcPath,
  target: 'web',

  plugins: [].filter((val) => !!val),
};
