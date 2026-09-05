# password-generator (`pwg`)

**A simple, secure password generator**

`password-generator` is a lightweight CLI tool written in Go that generates cryptographically secure passwords. I built this as a side project to replace my reliance on online password generators, and to get hands-on experience with Go and practical security implementations.

## Why?

I’ve been using online tools to generate passwords for a while ([passwordsgenerator.net](https://passwordsgenerator.net/) - which is now shut down, [LastPass](https://www.lastpass.com/features/password-generator), etc), but I was always really skeptical of the security measures and practices of these websites. I didn't know if they were truly client side (and the passwords couldn't be tracked by malicious actors) or if they were truly safe from interception/interference, whether they were truly random, etc.

As a CS grad, equipped with a bunch of knowledge from [my final year Computer Security class](https://www.up.ac.za/yearbooks/2025/EBIT-faculty/UG-modules/view/COS%20330), I figured it was time to build my own. I started with a CLI so I could make sure I got the fundamentals right.

When I tried researching how password generators are actually written under the hood, Google just kept spitting out websites that generate passwords, which wasn't helpful at all. So I had to figure a bunch of things out myself.

I'll honest and say I'm not super passionate about computer security (the course at uni was mostly reading), but I wanted to engage with the security concerns to make this thorough and safe, even if realistically it'll probably only be used by me.

I also chose to write this in Go. I'm trying to familiarise myself with Go more and more. I did consider using Rust because I'd have more control over the memory and I wouldn't have to worry about Go's garbage collector possibly getting in the way but Go was the right choice for my learning goals and getting it down quicker.

I'm aware the *most* secure generation methods would involve doing math straight on the CPU or GPU registers, but I wasn't sure where I'd start with that, and realisitically, I'd probably make a new mistakes trying to re-invent the wheel, so I stuck with Go's `crypto/rand` and focused on what I could control: memory management and clipboard hygiene.

## Installation

### Via Releases
Coming soon.

### Via Go Install
If you have Go installed, you can easily install the CLI globally using:
```bash
go install github.com/lkekana/password-generator@latest
```

### From Source
Alternatively, you can clone the repository and run it directly:
```bash
git clone https://github.com/lkekana/password-generator.git
cd password-generator

# to install it in your path
go build -trimpath -ldflags -s
cp passwords-generator /usr/local/bin/

# (optional) to install the 'pwd' alias
cp passwords-generator /usr/local/bin/

# to run the code, without installing
go run .
```

## Usage

```bash
$ pwg --help
A simple password generator CLI

Usage:
  pwg [flags]

Flags:
  -y, --copy         Copy to clipboard (only works when generating a single password)
  -c, --count int    Number of passwords to generate (default 1)
  -d, --debug        Enable debug mode
  -h, --help         help for pwg
  -l, --length int   Length of the password (default 16)
      --lower        Include lowercase letters (default true)
      --newline      Print a newline after generating a single password (useful for piping output & does not apply when generating multiple passwords)
      --num          Include numbers (default true)
      --special      Include special characters
      --upper        Include uppercase letters (default true)
```

### Examples
```bash
# Generate a standard 16-character password
pwg

# Generate a 24-character password with special characters and copy it to the clipboard
pwg -l 24 --special -y

# Generate 5 passwords and print them to stdout
pwg -c 5
```

![picture alt](./assets/example.png)

## Design Choices & Security Considerations

- **Memory Zeroing:** Because Go uses a Garbage Collector, sensitive data like passwords can linger in memory longer than you'd like. To address this, I implemented a `zeroOutPassword` function paired with `runtime.KeepAlive`. This ensures the password is explicitly overwritten with zeros after it's printed / copied and prevents the compiler from optimizing the zeroing operation away before the GC sweeps it.
- **Regeneration & Modulo Bias (Rejection Sampling):** To ensure the password contains all required character types (upper, lower, numbers, special), the generator creates the full password and checks it against your requirements. If it fails, it throws it out and regenerates. This avoids the predictable patterns and modulo bias that come from forcing specific characters into specific indexes.
- **No Web Version (Yet):** This is strictly a CLI tool for now. Building a web version introduces a whole new can of worms. Maybe I'll get to that after this.

## Performance

A sample run for 50 passwords:
```bash
$ go run . --debug -c 50 -l 24
Debug mode enabled.

=== GENERATOR CONFIGURATION ===
Number of passwords to generate: 50
Length of each password: 24
Include uppercase: true
Include lowercase: true
Include numbers: true
Include special characters: false
===============================

=== CLIPBOARD INFORMATION ===
Size of clipboard.FmtText: 8 bytes
Size of clipboard.FmtImage: 8 bytes
Clipboard does not contain image data.
Clipboard contains text data of size: 9 bytes
=============================

Password 1: TJFu3bX3B5CXxSp6DTGLVO7W (Execution took 247.395µs)
Password 2: Gm8Lmdkc0P4iuyWlq2jlZC8y (Execution took 95.016µs)
Password 3: bep1qtdjjd1QvP9NUNe3xudN (Execution took 55.482µs)
Password 4: y1Y1P0DusfBlnGCeabwSD4D9 (Execution took 51.42µs)
Password 5: k15NOVNbMuH0U3uyv2fCqi8K (Execution took 57.399µs)
Password 6: MmwrV7ngPIGoSK9v9H395UyP (Execution took 51.3µs)
Password 7: GFDhbE16Ct9cefPLmJcUXfmU (Execution took 53.726µs)
Password 8: CoogiAlLnM8z0tLwqTPd10aM (Execution took 71.065µs)
Password 9: A69ZvMQK82ipXtYI0y97Rn2A (Execution took 56.385µs)
Password 10: Q3ccq8m8gB57XaEL0VRqpQ8l (Execution took 32.851µs)
Password 11: EuNU35ZFfwZXVplvDXtvgBTa (Execution took 67.903µs)
Password 12: fd37pwysxPISlVFDYFgmMht4 (Execution took 30.975µs)
Password 13: MKYUon7QIPOKOcENcgcx1OcL (Execution took 50.682µs)
Password 14: lWBd3HLPmXyhAQZvr7lGELVw (Execution took 50.307µs)
Password 15: 3Z51lgVAHDjDyKss9dz7Q0hW (Execution took 51.467µs)
Password 16: H8kwBjUEPaNxqoRiar6B83jU (Execution took 74.904µs)
Password 17: FF9s35WrghkHMnPoS9SqWMCy (Execution took 73.305µs)
Password 18: yDy2DK7T4J6uHF07obWJaFW7 (Execution took 59.063µs)
Password 19: Z9Y1hUHywCA34SM0BhaWfOl4 (Execution took 52.29µs)
Password 20: lPIKGJWDr3a0AxoOjn7ynxJN (Execution took 119.467µs)
Password 21: Ek9x7sd6Baj9iywq0KtJ6603 (Execution took 107.44µs)
Password 22: T4NJRU4XmzpGc8Bg8R4CmieI (Execution took 65.972µs)
Password 23: 0yWUc9Tr10cowqNlEDEpwRly (Execution took 36.284µs)
Password 24: YehEDFioH1XNk8WuaYotYYW2 (Execution took 83.687µs)
Password 25: wwBJCQyVI8287fYl00IzRTQt (Execution took 39.838µs)
Password 26: XgKY79QqKOahFjwetWWlzSBA (Execution took 43.615µs)
Password 27: ND8FLeQ8MrrkBFXLnKAv8SyS (Execution took 51.965µs)
Password 28: mBhVVhBybG0i0PZw3DJSCLNt (Execution took 77.604µs)
Password 29: HJ9j06QpanePCzBXvVgoTIs5 (Execution took 78.844µs)
Password 30: lmRqXIx2oX1yK3ewMuQvSi5i (Execution took 87.402µs)
Password 31: lphot0a4k4ezG24EU6B6lHIS (Execution took 75.049µs)
Password 32: A2bXcqsnJN3vD64OsoQw1qRm (Execution took 54.763µs)
Password 33: NJAzVNIBLXiwhCNThNEJv3Ux (Execution took 32.363µs)
Password 34: eiayVQG8xKb8TisEeNBPaVTq (Execution took 32.22µs)
Password 35: 31Fv1ARp35sTKM4XRkEEAwV2 (Execution took 53.03µs)
Password 36: 5AIOFs9CcciZRc4DqLT5BGhK (Execution took 50.754µs)
Password 37: DB7sYz6sKsoZyinw4BYOapLP (Execution took 33.845µs)
Password 38: UTH0bKEYj029j1a7c9p6yvl5 (Execution took 29.756µs)
Generated password did not meet requirements, regenerating...
Password 39: tn8OUTx1D02eZUiH7bGVchjs (Execution took 106.306µs)
Password 40: NmsOqr8HBVoqTaCJ7NGK8XEe (Execution took 44.523µs)
Password 41: zITzuwsusnyEYHqYD0W5Vohy (Execution took 44.463µs)
Password 42: N4gZXjocGlurvIv23EiY2yD8 (Execution took 39.378µs)
Password 43: EYxKpcCSc7IixIR3b1k67Ajh (Execution took 48.866µs)
Password 44: HaBYWdRqdaFB7mrICO80EpBq (Execution took 33.796µs)
Password 45: maNqqXZPqFN2slery7Tue5dI (Execution took 32.153µs)
Password 46: YjprmYYWVbFbUn07RagtpB91 (Execution took 31.035µs)
Password 47: 5MDgmcIRjvzBT6QKa6ddF3dG (Execution took 50.234µs)
Password 48: bqhlWt9OG1efnvEaqoDZHpfm (Execution took 34.043µs)
Password 49: nAZNkvNPxgtvDyYL3zJd5UjT (Execution took 46.913µs)
Password 50: 07XoikIuSOYtRhkV2tsiaeAD (Execution took 30.356µs)
Total execution time for 50 passwords: 3.471205ms
```

First password takes the longest to generate because of the time needed to create the charset, though it seems that the Golang compiler reuses the charset making subsequent passwords a lot faster.

## Roadmap

- [ ] **Clipboard Integration:** Possibly re-enable the 10-second clipboard timeout and restore logic across all supported OS environments. I had it implemented but removed it because I just didn't like the UX.
- [ ] **Web Version:** Eventually port the core logic to a secure, fully client-side web application (maybe via WebAssembly?) & PWA.
- [ ] **Cross-Platform Notifications:** Add native OS notifications (via `beeep` or similar) when the clipboard is successfully restored. Also tried this out but it was annoying so I disabled it.

## License
This project is licensed under the GNU GPLv3 License. See the [LICENSE](LICENSE) file for details.
