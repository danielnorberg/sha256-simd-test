# sha256-simd-test

## OpenSSL
```
time openssl dgst -sha256 < ~/data-8G
(stdin)= ff803171b009108a1a3da6978df30529442b3c5cc32350ccdb9138e9fde3727b
openssl dgst -sha256 < ~/data-8G  5.09s user 1.03s system 99% cpu 6.124 total
```

## https://github.com/minio/sha256-simd
```
time ./sha256-simd-test ~/data-8G
ff803171b009108a1a3da6978df30529442b3c5cc32350ccdb9138e9fde3727b
./sha256-simd-test ~/data-8G  5.02s user 1.11s system 100% cpu 6.084 total
```

## sha256sum
```
time sha256sum ~/data-8G
ff803171b009108a1a3da6978df30529442b3c5cc32350ccdb9138e9fde3727b  /home/dano/data-8G
sha256sum ~/data-8G  24.55s user 1.02s system 99% cpu 25.570 total
```
