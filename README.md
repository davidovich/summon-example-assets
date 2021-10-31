Summon Data Provider Repo
=========================

This repository hosts a [summon](https://github.com/davidovich/summon) data
provider.

You install this provider as `summon`:

```shell
go install github.com/davidovich/summon-example-assets/summon@latest
```

And then use the `summon` executable to summon assets.

Summon `some-asset` like so:

```shell
summon some-asset
```

By default, summon will instantiate the asset in the `.summoned/` directory and
return its path. This can be overriden in the `asssets/summon.config.yaml` file
or by using the `-o` flag.

Get more help with `summon -h`.

Updating assets
---------------

1) Make modifications (additions, removals) in the assets/ dir
2) Invoke `make`
3) Commit changes
4) Tag with a semantic version (prefix with `v`)
5) git push --tags
