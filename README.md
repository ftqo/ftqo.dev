# ftqo.dev

this project requires node.js (u should use nvm) and go, then install dependencies with:
```bash
npm i && go mod tidy
```

to generate the tailwind, populate templates, and zip the static files, run:
```bash
go generate
```

then you can simply run:
```bash
go run cmd/server/main.go
```

if you have air installed, you can hot reload using:
```bash
air
```

snippets taken from [efron licht](https://gitlab.com/efronlicht/blog) with personal permission.
huge shoutout to his [backend basics posts](https://eblog.fly.dev/backendbasics.html). 
all of the good code in here is probably from him; the bad code is from me
