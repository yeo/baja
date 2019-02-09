# baja

A minimal static site generator

# Concept

Baja simply looks at a directory and find file recursive and render
individual page in a layout plus an index page of each directory.

To look up the template for render, if it's a single page, it use
`themes/name/node.html`. If it's an index, it used `list.html`.

Plus, if it has an `index.html` page, that template are used to render
index of whole site. otherwise it used `list.html`.


# Getting started

The API is similar to git

```
# First time, bootstrap
baja new

# Add a post
baja add directory/post/

# Import from Gist or Evernote
baja import url

# Spit out static file
baja build

# You need to have 
baja deploy s3

# of if you're using Github page
baja deploy github
```

# Why the name

When my daughter started to speak, `baja` was one the word she kept
saying for no reason. No one in family know what it means. It remains
unknow nowsaday and she eventually stop saying when she know more words

But the word `baja` light up my days a lot.

# Release

- 2019/02/09: 0.0.2
