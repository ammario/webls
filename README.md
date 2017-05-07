<!-- START doctoc generated TOC please keep comment here to allow auto update -->
<!-- DON'T EDIT THIS SECTION, INSTEAD RE-RUN doctoc TO UPDATE -->
**Table of Contents**

- [Webls](#webls)
- [Features](#features)
- [Install](#install)
- [.index](#index)
- [TODO](#todo)

<!-- END doctoc generated TOC please keep comment here to allow auto update -->

# Webls

Standalone HTTP File Index

[Example](http://f.ammar.io)

# Features
- Easy to use
- Disable directory listing
- Set maximum listing count for directory
- Range support

# Install

```bash
go get -u github.com/ammario/webls
```

# .index

`.index` files are JSON and configure how the directory it's in is displayed.

A `.index` file may look like (without the comments)

```js
{
    //only show first 1000 files
    "max_files": 1000,
    //don't index directory
    "hide": true,
}
```

# TODO
- Simple theming
- More `.index` parameters
    - Optional basic auth
- Optional JSON/CSV listing