# GoWood
Encryption Software which supports existing algorithms with GO


# introduction

The goal was to write a tool in Golang that Supports Modern Encryption Algorithms.
Supported Methods: Hashing, Symmetric Encryption, Asymmetric Encryption
Supported Algorithms: MD5, SHA-256, SHA-512, AS-128 , AES-192, AES-256, BLOWFISH, DES, RSA
 

# installation

### from source code

```bash
git clone github.com/haliliceylan/gowood.git
cd gowood
make build
./gowood
```
**if you want to install as binary in your /usr/local/bin run ** `make install-usr-local` instead of `make install`


### from release


# Algorithms

## Hashing
- [x] MD5
- [x] SHA-256
- [x] SHA-512

## Asymmetric Encryption
- [x] RSA

## Symmetric Encryption
- [X] AES-256, AES-192, AES-128
- [x] Blowfish
- [x] DES

# Usage
## Hashing
### MD5 Algorithm
[![asciicast](https://asciinema.org/a/AHP5AEpsbEWjXTolpSD11f7nU.svg)](https://asciinema.org/a/AHP5AEpsbEWjXTolpSD11f7nU)

### SHA256 Algorithm
[![asciicast](https://asciinema.org/a/5fBN5P1c2mG9YWIkjViTV1Go2.svg)](https://asciinema.org/a/5fBN5P1c2mG9YWIkjViTV1Go2)

### SHA512 Algorithm
[![asciicast](https://asciinema.org/a/Rp6i4NafurUmmTfwtMCNAejiH.svg)](https://asciinema.org/a/Rp6i4NafurUmmTfwtMCNAejiH)

## Asymmetric Encryption

### RSA Create Keys
[![asciicast](https://asciinema.org/a/Fy8l8t52J1kZJkxoccmsQM1zf.svg)](https://asciinema.org/a/Fy8l8t52J1kZJkxoccmsQM1zf)

### RSA Encryption/Decryption
[![asciicast](https://asciinema.org/a/4RZ0kGLPtH2N2qjF6DAMxnXWo.svg)](https://asciinema.org/a/4RZ0kGLPtH2N2qjF6DAMxnXWo)

### RSA Sign/Verify
[![asciicast](https://asciinema.org/a/Lri56t4of0I50DpB5vS9PLyHs.svg)](https://asciinema.org/a/Lri56t4of0I50DpB5vS9PLyHs)

## Symmetric Encryption

### AES-256, AES-192, AES-128
[![asciicast](https://asciinema.org/a/jZqzuudfunKW3uTWfK51Q8vwp.svg)](https://asciinema.org/a/jZqzuudfunKW3uTWfK51Q8vwp)