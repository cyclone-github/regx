# RegX
*A Flexible Potfile Parsing Tool*

*Pronounced "Reg-X" as in Regex eXtractor*

### Why use RegX instead of grep, ripgrep, etc?
Unlike general-purpose tools like grep or ripgrep, RegX offers a more nuanced approach that was developed specifically for parsing hashcat potfiles that contain multiple hash algo's. While this can be done with grep, compiling and testing regular expressions, especially for more complex hashes, can be cumbersome and error prone. See examples below.

### RegX vs grep (egrep):

- Parse md5 (hex32) hashes:
  - grep
    - `egrep '[a-f0-9]{32}' file.txt`
  - RegX
    - `./regx.bin -f file.txt -m hex32`
- Parse bcrypt hashes:
  - grep
    - `egrep '\$2[a-zA-Z]{1}\$[0-9]{2}\$[[:print:]]{53}' file.txt`
  - RegX
    - `./regx.bin -f file.txt -m 3200`
- Parse Django (PBKDF2-SHA256) hashes:
  - grep
    - `egrep 'pbkdf2_sha256\$[0-9]{1,6}\$[[:print:]]{57}' file.txt`
  - RegX
    - `./regx.bin -f file.txt -m 10000`
- Parse Bitcoin hashes:
  - grep
    - `egrep '\$bitcoin\$[0-9]{1,3}\$[a-f0-9]{40,}\$[0-9]{2}\$[a-f0-9]{2,}\$[0-9]{2,}\$[0-9]{1,}\$[0-9]{1,}\$[0-9]{1,}\$[0-9]{1,}' file.txt`
  - RegX
    - `./regx.bin -f file.txt -m 11300`

As shown in the examples above, RegX's built-in support for popular hashcat algorithms makes parsing hashes seamless.

RegX also supports a wide range of regex patterns compatible with RE2, by using option: `-r {regex_pattern}`

More info on RE2: https://github.com/google/re2/wiki/Syntax

### Usage Instructions:
- Parse all hex 32 hashes (md5, md4, ntlm, etc), both salted and non-salted:
  - `./regx.bin -f file.txt -m hex32`
- Parse bcrypt hashes by hashcat mode {-m 3200}:
  - `./regx.bin -f file.txt -m 3200`
- Parse bcrypt hashes by algo name {-m bcrypt}:
  - `./regx.bin -f file.txt -m bcrypt`
- Parse a custom hex length hash (where {nth} equals length):
  - `./regx.bin -f file.txt -m hex{nth}`
- Use custom RE2 regex with -r {regex}:
  - `./regx.bin -f file.txt -r '[a-fA-F0-9]{32}'`
- Run `./regx.bin -help` to see a list of all options

### Supported hash algorithms (more will be added):
- All HEX algos:
  - `-m hex32` covers all HEX32 hashes such as md5, md5, ntlm, etc
  - `-m hex40` covers all HEX40 hashes such as sha1, mysql5, ripemd-160, etc
  - custom hex lengths can be given as well, `-m hex56` would cover sha224 hashes, etc

| Mode: | Hashcat Mode: | HEX |
|--------|--------|--------|
| crc32 | 11500 | hex8 |
| crc64 | 28000 | hex16 |
| md4 | 900 | hex32 |
| md5 | 0 | hex32 |
| ntlm | 1000 | hex32 |
| ripemd-160 | 6000 | hex40 |
| sha1 | 100 | hex40 |
| mysql5 | 300 | hex40 |
| sha224 | 1300 | hex56 |
| sha256 | 1400 | hex64 |
| sha384 | 10800 | hex96 |
| sha512 | 1700 | hex128 |
| metamask | 26600 |
| bitcoin | 11300 |
| pbkdf2sha256 | 10000 |
| bcrypt | 3200 |
| sha512crypt | 1800 |
| md5crypt | 500 |
| phpass | 400 |

### Compile from source:
- This assumes you have Go and Git installed
  - `git clone https://github.com/cyclone-github/regx.git`
  - `cd regx`
  - `go mod init regx`
  - `go mod tidy`
  - `go build .`
- More info on compiling from source:
  - https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt

### Change Log:
- https://github.com/cyclone-github/regx/blob/main/CHANGELOG.md

### Antivirus False Positives:
- Several antivirus programs on VirusTotal incorrectly detect compiled Go binaries as a false positive. This issue primarily affects the Windows executable binary, but is not limited to it. If this concerns you, I recommend carefully reviewing the source code, then proceed to compile the binary yourself.
- Uploading your compiled binaries to https://virustotal.com and leaving an up-vote or a comment would be helpful as well.