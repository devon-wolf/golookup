# golookup

Command line tool to look up words in [Merriam-Webster's](https://dictionaryapi.com/) Collegiate Dictionary and Thesaurus.

## Intended Commands
- `golookup WORD` - general lookup, will include multiple types of information
- `golookup WORD --def` - dictionary definition lookup
- `golookup WORD --syns` - thesaurus synonym lookup
- `golookup WORD --ants` - thesaurus antonym lookup

That's not how it works yet. See implementation for details.

## Setup
To run this locally you will need to include API keys for the dictionary and thesaurus APIs in a `.env` file - you can get your own keys by signing up at [dictionaryapi.com](https://dictionaryapi.com/). The `.env-example` file shows you what variables the app will be expecting; copy this over to a file named `.env`, paste in your API keys where the keys go, and all should be well. If you haven't included this, upon running the app you will see the following output:
```
Error loading .env file
exit status 1
```

You will also need to have [Go installed](https://go.dev/doc/install) to run it from the project files.

## Running
This will change, but right now you can run the app from the project directory with `go run . -word [some word you want to look up]`. You can optionally give it a -src flag of t (`go run . -word [some word] -src t) to use the thesaurus rather than the dictionary. This is not the eventual intended behavior but it'll work for now.