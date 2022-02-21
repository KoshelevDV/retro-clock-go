# RETRO CLOCK

Retro clock show your local time and smoothly slides from right to left

<pre>
 ███  ███       ██   ███       ███  ███       ███
 █ █  █ █        █   █ █       █      █       █ █
 █ █  ███        █   ███       ███  ███       ███
 █ █    █        █     █         █  █         █ █
 ███  ███       ███  ███       ███  ███       ███
</pre>

## Run

```bash
git clone https://github.com/KoshelevDV/retro-clock-go.git
cd retro-clock-go
go build -o target/retro-clock cmd/clock/main.go
./target/retro-clock
```

## Docker

Repo contains scripts to build and run retro clock docker image

```bash
./scripts/build.sh
./scripts/run.sh
```
