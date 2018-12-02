# GopherJS Webpack loader

[![npm](https://img.shields.io/npm/v/gopherjs-loader.svg)](https://www.npmjs.com/package/gopherjs-loader)

### Usage

This is a simple Webpack loader that shells out to [gopherjs](https://github.com/gopherjs/gopherjs).

To use it, first install the package:

```bash
$ npm install --save gopherjs-loader
```

then configure the loader in your Webpack config:

```js
module.exports = {
  // ...
  module: {
    rules: [
      { test: /\.go$/, loader: 'gopherjs-loader' },
      // ...
    ]
  }
}
```

Make sure you have the `gopherjs` binary somewhere in your `PATH`.

### Example

Check out the [example](example) directory for a simple Hello World example.
