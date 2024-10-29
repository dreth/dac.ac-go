# dac.sg
My site built using the Fiber web framework (Go)

## CompileDaemon or air

```shell
CompileDaemon -command="./dac.sg"
```

```shell
air
```


## Tailwind

Autocompile tailwindcss files to css

```shell
NODE_ENV=production npx tailwindcss -i ./static/assets/tailwind/src/styles.css -o ./static/assets/tailwind/styles.css --watch
```
