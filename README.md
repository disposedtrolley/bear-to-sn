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

## The Battle Plan

1. Export all Bear notes as `.md` files.
2. Use the [plaintext to Standard Notes Converter](https://dashboard.standardnotes.org/tools) tool to convert the `.md` files to a big JSON array.
3. Parse the JSON with this utility, which will extract any tags found in the body of the note and add the relevant Tag items to the JSON array, ensuring they're properly referenced back to the note.

## The Data Structure

For our purposes, an item in Standard Notes can either be a `Note` or a `Tag`.

### Notes

Notes look like this (we can tell from the `content_type`):

```
{
  "uuid": "8a1b67db-2773-47d2-a206-2b642fabbf5b",
  "content_type": "Note",
  "created_at": "2018-06-30T12:22:26.733Z",
  "content": {
    "title": "",
    "text": "",
    "references": [
      {
        "uuid": "642c94d4-70a7-422e-a70a-d60263d7f54f",
        "content_type": "Tag"
      }
    ],
    "appData": {
      "org.standardnotes.sn": {
        "client_updated_at": "2018-07-01T07:28:33.420Z"
      }
    }
  },
  "updated_at": "2018-07-01T07:28:33.830Z"
}
```

`Note`s can have references to one or more `Tag`s, shown in the `content.references` array. A reference object appears to simply consist of a `uuid` and the `content_type` being referenced:

```
{
  "uuid": "642c94d4-70a7-422e-a70a-d60263d7f54f",
  "content_type": "Tag"
} 
```

### Tags

