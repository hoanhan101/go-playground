# mux

These 2 files are essential the same. Both use mux to create routes.

However, one use Handle() function and one use HandleFunc().
- Handle() takes a string route and a function that implements Handler
  interface
- HandleFunction() takes a string route and a function in its own type, that we
  define, that take a ResponseWriter and a pointer to Request. This will use
  the DefaultServeMux, instead of calling the NewServeMux.

Look at Go doc for more details.
