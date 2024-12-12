# sha256-simd-test

```
time openssl dgst -sha256 < ~/data-8G
(stdin)= ff803171b009108a1a3da6978df30529442b3c5cc32350ccdb9138e9fde3727b
openssl dgst -sha256 < ~/data-8G  5.09s user 1.03s system 99% cpu 6.124 total
```

```
time ./sha256-simd-test ~/data-8G
ff803171b009108a1a3da6978df30529442b3c5cc32350ccdb9138e9fde3727b./sha256-simd-test ~/data-8G  5.02s user 1.11s system 100% cpu 6.084 total
```

Presumably the AVX512 implementation benefit is the greatest when digesting different messages in parallel.
