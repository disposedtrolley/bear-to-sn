# Bear to Standard Notes

I've recently switched to Standard Notes to remain platform agnostic while maintaining a great note taking experience, and have struggled in getting my Bear notes to import the way I'd like to. The [plaintext to Standard Notes Converter](https://dashboard.standardnotes.org/tools) works a charm in bringing in the content of my notes, but doesn't take into account Bear's interesting tagging system.

## Tagging in Bear

Bear notes are written in Markdown, and support inline tagging by prefixing words with an octothorpe (`#`). For example, the below Bear snippet will be tagged in this folder structure:

```
--travel
  |
  |--asia
     |
     |--japan
```

```
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.

#travel/asia/japan
```

The note appears in all three folders; `travel`, `asia`, and `japan`, rather than being dumped into a composite folder `travel/asia/japan`. This is a cool aspect of Bear as it saves you some time in having to create multiple tags, each going a level deeper than the last.

## Tagging in Standard Notes

The Basic version of Standard Notes supports non-heirarchical tagging (the tagging is also done in a separate section of the note), so the equivalent of the above in SN would look like this in the sidebar:

```
--travel/asia/japan
```

```
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
```

*The tags section*
```
#travel/asia/japan
```

Thankfully, the `Folders` plugins unlocked with Standard Notes extended provides a folder structure which respects tags, which a similar appeareance to the Bear implementation. It expects tags to be dot-notated for nesting purposes, and repeated for each level of nesting, like this:

```
--travel
  |
  |--asia
     |
     |--japan
```

```
Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua.
```

*The tags section*
```
#travel #travel.asia #travel.asia.japan
```