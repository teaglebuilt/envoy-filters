# Envoy Filters

This repository is an example of a monorepo sandbox of envoy filters. Envoy is a powerful L7 proxy and designed for modernized scalable architectures. For further details into Envoy and how envoy filters can be used, check out [this article]() I wrote.

Both custom filters and pre built existing filters are setup so that you can leverage different options as you explore the landscape of possibilities.

In the folder `filters`, is a variety of different builtin filters that you can apply to envoy to test as needed.
The `packages` folder contains a custom filter by language, currently only `Go` and `Rust`.

## Testing

A `TODO:` is created to use a go cli to handle all operations with [Dagger]() for handling containerization operations. This will allow for a more streamlined a dynamic approach for deciding which filters you want to test with. In the meantime, refer to the commands in the Makefile for operational use.
