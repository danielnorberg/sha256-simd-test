# sha256-simd-test

* Python hashlib is fast 🚀
* OpenSSL is fast 🚀
* `minio/sha256-simd` is fast 🚀
* The `minio/sha256-simd` avx512 server does not offer faster single file performance

## Python 3.11 Hashlib
```
python --version
Python 3.11.11

time python -c 'import hashlib; hashlib.file_digest(open("/home/dano/data-8G", "rb"), "sha256")'
python -c   4.95s user 1.08s system 99% cpu 6.032 total
```

## OpenSSL
```
time openssl dgst -sha256 < ~/data-8G
(stdin)= ff803171b009108a1a3da6978df30529442b3c5cc32350ccdb9138e9fde3727b
openssl dgst -sha256 < ~/data-8G  5.09s user 1.03s system 99% cpu 6.124 total
```

## https://github.com/minio/sha256-simd
```
time ./sha256-simd-test simple ~/data-8G
ff803171b009108a1a3da6978df30529442b3c5cc32350ccdb9138e9fde3727b
./sha256-simd-test ~/data-8G  5.02s user 1.11s system 100% cpu 6.084 total
```

```
time ./sha256-simd-test server-avx512 ~/data-8G
e1a8a32281e3bbeaf178ce831b472b9dc0338da6d121f65990d89b4a09fb9ca0
./sha256-simd-test server ~/data-8G  39.18s user 1.63s system 111% cpu 36.482 total
```

## sha256sum
```
time sha256sum ~/data-8G
ff803171b009108a1a3da6978df30529442b3c5cc32350ccdb9138e9fde3727b  /home/dano/data-8G
sha256sum ~/data-8G  24.55s user 1.02s system 99% cpu 25.570 total
```

### Hardware

* AWS EC2 m7i.8xlarge

```
$ lscpu
Architecture:                         x86_64
CPU op-mode(s):                       32-bit, 64-bit
Byte Order:                           Little Endian
Address sizes:                        46 bits physical, 48 bits virtual
CPU(s):                               32
On-line CPU(s) list:                  0-31
Thread(s) per core:                   2
Core(s) per socket:                   16
Socket(s):                            1
NUMA node(s):                         1
Vendor ID:                            GenuineIntel
CPU family:                           6
Model:                                143
Model name:                           Intel(R) Xeon(R) Platinum 8488C
Stepping:                             8
CPU MHz:                              2400.000
BogoMIPS:                             4800.00
Hypervisor vendor:                    KVM
Virtualization type:                  full
L1d cache:                            768 KiB
L1i cache:                            512 KiB
L2 cache:                             32 MiB
L3 cache:                             105 MiB
NUMA node0 CPU(s):                    0-31
Vulnerability Gather data sampling:   Not affected
Vulnerability Itlb multihit:          Not affected
Vulnerability L1tf:                   Not affected
Vulnerability Mds:                    Not affected
Vulnerability Meltdown:               Not affected
Vulnerability Mmio stale data:        Not affected
Vulnerability Reg file data sampling: Not affected
Vulnerability Retbleed:               Not affected
Vulnerability Spec rstack overflow:   Not affected
Vulnerability Spec store bypass:      Mitigation; Speculative Store Bypass disabled via prctl and seccomp
Vulnerability Spectre v1:             Mitigation; usercopy/swapgs barriers and __user pointer sanitization
Vulnerability Spectre v2:             Mitigation; Enhanced / Automatic IBRS; IBPB conditional; RSB filling
                                      ; PBRSB-eIBRS SW sequence; BHI BHI_DIS_S
Vulnerability Srbds:                  Not affected
Vulnerability Tsx async abort:        Not affected
Flags:                                fpu vme de pse tsc msr pae mce cx8 apic sep mtrr pge mca cmov pat ps
                                      e36 clflush mmx fxsr sse sse2 ss ht syscall nx pdpe1gb rdtscp lm con
                                      stant_tsc arch_perfmon rep_good nopl xtopology nonstop_tsc cpuid ape
                                      rfmperf tsc_known_freq pni pclmulqdq monitor ssse3 fma cx16 pdcm pci
                                      d sse4_1 sse4_2 x2apic movbe popcnt tsc_deadline_timer aes xsave avx
                                       f16c rdrand hypervisor lahf_lm abm 3dnowprefetch invpcid_single ssb
                                      d ibrs ibpb stibp ibrs_enhanced fsgsbase tsc_adjust bmi1 avx2 smep b
                                      mi2 erms invpcid avx512f avx512dq rdseed adx smap avx512ifma clflush
                                      opt clwb avx512cd sha_ni avx512bw avx512vl xsaveopt xsavec xgetbv1 x
                                      saves avx_vnni avx512_bf16 wbnoinvd ida arat avx512vbmi umip pku osp
                                      ke waitpkg avx512_vbmi2 gfni vaes vpclmulqdq avx512_vnni avx512_bita
                                      lg tme avx512_vpopcntdq rdpid cldemote movdiri movdir64b md_clear se
                                      rialize amx_bf16 avx512_fp16 amx_tile amx_int8 flush_l1d arch_capabi
                                      lities
```
