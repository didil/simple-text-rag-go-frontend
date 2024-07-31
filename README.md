# Simple Text RAG Go Frontend
Go Frontend App for [https://github.com/didil/simple-text-rag](https://github.com/didil/simple-text-rag). Receives requests via a REST API and sends them to the simple-text-rag app via gRPC.

## Development Setup

### Install protoc
```bash
brew install protobuf
```

### Install the protocol compiler plugins for Go
```bash
make protoc-plugins
```

### Install linter
```bash
make install-linter
```

### Setup .env file
```bash
cp .env.example .env
```

### Regenerate Protobuf output files
Only needed if making changes to the proto files
```bash
make gen-protos
```

### Lint
```bash
make lint
```

### Run server
```bash
make run
```

## Querying via curl
*Make sure [https://github.com/didil/simple-text-rag](https://github.com/didil/simple-text-rag) is also running.*

Create a collection using the Alice In Wonderland online book text:
```bash
$ curl --request POST \
  --url http://localhost:8080/api/v1/collections \
  --header 'Content-Type: application/json' \
  --data '{
	"name": "alice_in_wonderland",
	"fileUrl": "https://www.gutenberg.org/cache/epub/11/pg11.txt"
}'
# output
{}
```
Ask a question about the collection:
```bash
curl --request POST \
  --url http://localhost:8080/api/v1/questions \
  --header 'Content-Type: application/json' \
  --data '{
	"collectionName": "alice_in_wonderland",
	"question": "Who are the top 3 characters in this story ?"
}'
# output
{"text":"Based on the context provided, the top 3 characters in this story appear to be:\n\n1. Alice: The protagonist of the story, a young girl who finds herself in a fantastical world.\n2. The Queen of Hearts: A central figure in the story, known for her short temper and tendency to shout \"Off with her head!\"\n3. The King of Hearts: The ruler of the land, who also serves as the judge in a trial that takes place in the story.\n\nThese three characters seem to be the most prominent and influential in the story, although other characters like the White Rabbit, the Knave of Hearts, and the Mouse also play significant roles."}
```