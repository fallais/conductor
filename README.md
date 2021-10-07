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