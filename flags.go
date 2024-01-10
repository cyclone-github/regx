package main

import (
	"encoding/base64"
	"fmt"
	"os"
)

// last modified 2024-01-10.0945

func cycloneFunc() {
	codedBy := "Q29kZWQgYnkgY3ljbG9uZSA7KQo="
	codedByDecoded, _ := base64.StdEncoding.DecodeString(codedBy)
	fmt.Fprintln(os.Stderr, string(codedByDecoded))
}

func versionFunc() {
	version := "Q3ljbG9uZSdzIFJlZ1ggdjAuMS4wCg=="
	versionDecoded, _ := base64.StdEncoding.DecodeString(version)
	fmt.Fprintln(os.Stderr, string(versionDecoded))
}

// help function
func helpFunc() {
	versionFunc()
	str := "Supported Options:\n-f {file} -m {mode} \n-f {file} -r {regex}\n" +
		"\nExample Usage:\n" +
		"\nParse all hex 32 hashes (md5, md4, ntlm, etc), both salted and non-salted:\n./regx.bin -f file.txt -m hex32\n" +
		"\nParse bcrypt hashes by hashcat mode {-m 3200}:\n./regx.bin -f file.txt -m 3200\n" +
		"\nParse bcrypt hashes by algo name {-m bcrypt}\n./regx.bin -f file.txt -m bcrypt\n" +
		"\nParse a custom hex length hash (where {nth} equals length):\n./regx.bin -f file.txt -m hex{nth}\n" +
		"\nUse custom RE2 regex with -r {regex}:\n./regx.bin -f file.txt -r '[a-fA-F0-9]{32}'\n" +
		"\nMore info on RE2: https://github.com/google/re2/wiki/Syntax\n" +
		"\nModes:\t\tHashcat Mode:\tHEX:\n" +
		"crc32\t\t11500\t\thex8\n" +
		"crc64\t\t28000\t\thex16\n" +
		"md4\t\t900\t\thex32\n" +
		"md5\t\t0\t\thex32\n" +
		"ntlm\t\t1000\t\thex32\n" +
		"ripemd-160\t6000\t\thex40\n" +
		"sha1\t\t100\t\thex40\n" +
		"mysql5\t\t300\t\thex40\n" +
		"sha224\t\t1300\t\thex56\n" +
		"sha256\t\t1400\t\thex64\n" +
		"sha384\t\t10800\t\thex96\n" +
		"sha512\t\t1700\t\thex128\n" +
		"metamask\t26600\n" +
		"bitcoin\t\t11300\n" +
		"pbkdf2sha256\t10000\n" +
		"bcrypt\t\t3200\n" +
		"sha512crypt\t1800\n" +
		"md5crypt\t500\n" +
		"phpass\t\t400\n"
	fmt.Fprintln(os.Stderr, str)
}
