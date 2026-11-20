> [!WARNING]
> **This project is archived and no longer maintained.** It was an early experiment, and existing tools now cover what it does:
> - [flog](https://github.com/mingrammer/flog): fake log generator (Apache, syslog, JSON)
> - [elastic/stream](https://github.com/elastic/stream): replays log files to TCP/UDP/TLS/HTTP inputs
> - [Redpanda Connect](https://github.com/redpanda-data/connect) (formerly Benthos): templated fake events sent to almost any output
> - [telemetrygen](https://github.com/open-telemetry/opentelemetry-collector-contrib/tree/main/cmd/telemetrygen): OpenTelemetry logs, traces and metrics
>
> To test SIEM detection rules, replay recorded attack data such as [Security-Datasets](https://github.com/OTRF/Security-Datasets) or [EVTX-ATTACK-SAMPLES](https://github.com/sbousseaden/EVTX-ATTACK-SAMPLES).

# Conductor

**Conductor** is a tool used to **send fake data** to log collectors such as **SIEM** (IBM QRadar, RSA Security Analytics, Splunk, etc...), it is built with **Golang**.

## Why ?

You may want to test the rules or the alerts you created in your SIEM, but you need logs to come up. Or you may want to perform a PoC of a SIEM solution but you do not want your systems to sends logs to it.

# Generators

The **generator** is used to generate fake payloads, here are the provided generators :

- SimpleGenerator : very simple, it repeat a template, X number of times, with a latency between each payload.
- FileReaderGenerator : also simple, it reads a file, and read line by line to send playload, with a latency.

# Contribute

You want to add more generators ? Please send a PR.