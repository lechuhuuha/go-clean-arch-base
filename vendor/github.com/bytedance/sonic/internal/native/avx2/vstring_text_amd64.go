// +build amd64
// Code generated by asm2asm, DO NOT EDIT.

package avx2

var _text_vstring = []byte{
	// .p2align 5, 0x00
	// LCPI0_0
	0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, // QUAD $0x2222222222222222; QUAD $0x2222222222222222  // .space 16, '""""""""""""""""'
	0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, 0x22, //0x00000010 QUAD $0x2222222222222222; QUAD $0x2222222222222222  // .space 16, '""""""""""""""""'
	//0x00000020 LCPI0_1
	0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, //0x00000020 QUAD $0x5c5c5c5c5c5c5c5c; QUAD $0x5c5c5c5c5c5c5c5c  // .space 16, '\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\'
	0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, 0x5c, //0x00000030 QUAD $0x5c5c5c5c5c5c5c5c; QUAD $0x5c5c5c5c5c5c5c5c  // .space 16, '\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\\'
	//0x00000040 LCPI0_2
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, //0x00000040 QUAD $0x2020202020202020; QUAD $0x2020202020202020  // .space 16, '                '
	0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, 0x20, //0x00000050 QUAD $0x2020202020202020; QUAD $0x2020202020202020  // .space 16, '                '
	//0x00000060 .p2align 4, 0x90
	//0x00000060 _vstring
	0x55, //0x00000060 pushq        %rbp
	0x48, 0x89, 0xe5, //0x00000061 movq         %rsp, %rbp
	0x41, 0x57, //0x00000064 pushq        %r15
	0x41, 0x56, //0x00000066 pushq        %r14
	0x41, 0x55, //0x00000068 pushq        %r13
	0x41, 0x54, //0x0000006a pushq        %r12
	0x53, //0x0000006c pushq        %rbx
	0x48, 0x83, 0xec, 0x28, //0x0000006d subq         $40, %rsp
	0x4c, 0x8b, 0x16, //0x00000071 movq         (%rsi), %r10
	0xf6, 0xc1, 0x20, //0x00000074 testb        $32, %cl
	0x48, 0x89, 0x75, 0xc0, //0x00000077 movq         %rsi, $-64(%rbp)
	0x0f, 0x85, 0x5b, 0x01, 0x00, 0x00, //0x0000007b jne          LBB0_12
	0x48, 0x8b, 0x4f, 0x08, //0x00000081 movq         $8(%rdi), %rcx
	0x48, 0x89, 0xc8, //0x00000085 movq         %rcx, %rax
	0x4c, 0x29, 0xd0, //0x00000088 subq         %r10, %rax
	0x0f, 0x84, 0x33, 0x04, 0x00, 0x00, //0x0000008b je           LBB0_42
	0x48, 0x89, 0x4d, 0xd0, //0x00000091 movq         %rcx, $-48(%rbp)
	0x4c, 0x8b, 0x0f, //0x00000095 movq         (%rdi), %r9
	0x4b, 0x8d, 0x0c, 0x11, //0x00000098 leaq         (%r9,%r10), %rcx
	0x48, 0x83, 0xf8, 0x40, //0x0000009c cmpq         $64, %rax
	0x0f, 0x82, 0x2a, 0x04, 0x00, 0x00, //0x000000a0 jb           LBB0_43
	0x49, 0x89, 0xd4, //0x000000a6 movq         %rdx, %r12
	0x41, 0x89, 0xc5, //0x000000a9 movl         %eax, %r13d
	0x41, 0x83, 0xe5, 0x3f, //0x000000ac andl         $63, %r13d
	0x48, 0x8d, 0x50, 0xc0, //0x000000b0 leaq         $-64(%rax), %rdx
	0x48, 0x83, 0xe2, 0xc0, //0x000000b4 andq         $-64, %rdx
	0x4c, 0x01, 0xd2, //0x000000b8 addq         %r10, %rdx
	0x4d, 0x8d, 0x7c, 0x11, 0x40, //0x000000bb leaq         $64(%r9,%rdx), %r15
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x000000c0 movq         $-1, %r11
	0x45, 0x31, 0xc0, //0x000000c7 xorl         %r8d, %r8d
	0xc5, 0xfe, 0x6f, 0x05, 0x2e, 0xff, 0xff, 0xff, //0x000000ca vmovdqu      $-210(%rip), %ymm0  /* LCPI0_0+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x0d, 0x46, 0xff, 0xff, 0xff, //0x000000d2 vmovdqu      $-186(%rip), %ymm1  /* LCPI0_1+0(%rip) */
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x000000da .p2align 4, 0x90
	//0x000000e0 LBB0_4
	0xc5, 0xfe, 0x6f, 0x11, //0x000000e0 vmovdqu      (%rcx), %ymm2
	0xc5, 0xfe, 0x6f, 0x59, 0x20, //0x000000e4 vmovdqu      $32(%rcx), %ymm3
	0xc5, 0xed, 0x74, 0xe0, //0x000000e9 vpcmpeqb     %ymm0, %ymm2, %ymm4
	0xc5, 0x7d, 0xd7, 0xf4, //0x000000ed vpmovmskb    %ymm4, %r14d
	0xc5, 0xe5, 0x74, 0xe0, //0x000000f1 vpcmpeqb     %ymm0, %ymm3, %ymm4
	0xc5, 0xfd, 0xd7, 0xdc, //0x000000f5 vpmovmskb    %ymm4, %ebx
	0xc5, 0xed, 0x74, 0xd1, //0x000000f9 vpcmpeqb     %ymm1, %ymm2, %ymm2
	0xc5, 0xfd, 0xd7, 0xfa, //0x000000fd vpmovmskb    %ymm2, %edi
	0xc5, 0xe5, 0x74, 0xd1, //0x00000101 vpcmpeqb     %ymm1, %ymm3, %ymm2
	0xc5, 0xfd, 0xd7, 0xd2, //0x00000105 vpmovmskb    %ymm2, %edx
	0x48, 0xc1, 0xe3, 0x20, //0x00000109 shlq         $32, %rbx
	0x48, 0xc1, 0xe2, 0x20, //0x0000010d shlq         $32, %rdx
	0x48, 0x09, 0xd7, //0x00000111 orq          %rdx, %rdi
	0x49, 0x83, 0xfb, 0xff, //0x00000114 cmpq         $-1, %r11
	0x0f, 0x85, 0x09, 0x00, 0x00, 0x00, //0x00000118 jne          LBB0_6
	0x48, 0x85, 0xff, //0x0000011e testq        %rdi, %rdi
	0x0f, 0x85, 0x2f, 0x00, 0x00, 0x00, //0x00000121 jne          LBB0_9
	//0x00000127 LBB0_6
	0x4c, 0x09, 0xf3, //0x00000127 orq          %r14, %rbx
	0x48, 0x89, 0xfa, //0x0000012a movq         %rdi, %rdx
	0x4c, 0x09, 0xc2, //0x0000012d orq          %r8, %rdx
	0x0f, 0x85, 0x3c, 0x00, 0x00, 0x00, //0x00000130 jne          LBB0_10
	//0x00000136 LBB0_7
	0x48, 0x85, 0xdb, //0x00000136 testq        %rbx, %rbx
	0x0f, 0x85, 0x85, 0x00, 0x00, 0x00, //0x00000139 jne          LBB0_11
	//0x0000013f LBB0_8
	0x48, 0x83, 0xc0, 0xc0, //0x0000013f addq         $-64, %rax
	0x48, 0x83, 0xc1, 0x40, //0x00000143 addq         $64, %rcx
	0x48, 0x83, 0xf8, 0x3f, //0x00000147 cmpq         $63, %rax
	0x0f, 0x87, 0x8f, 0xff, 0xff, 0xff, //0x0000014b ja           LBB0_4
	0xe9, 0x7c, 0x02, 0x00, 0x00, //0x00000151 jmp          LBB0_30
	//0x00000156 LBB0_9
	0x48, 0x89, 0xca, //0x00000156 movq         %rcx, %rdx
	0x4c, 0x29, 0xca, //0x00000159 subq         %r9, %rdx
	0x4c, 0x0f, 0xbc, 0xdf, //0x0000015c bsfq         %rdi, %r11
	0x49, 0x01, 0xd3, //0x00000160 addq         %rdx, %r11
	0x4c, 0x09, 0xf3, //0x00000163 orq          %r14, %rbx
	0x48, 0x89, 0xfa, //0x00000166 movq         %rdi, %rdx
	0x4c, 0x09, 0xc2, //0x00000169 orq          %r8, %rdx
	0x0f, 0x84, 0xc4, 0xff, 0xff, 0xff, //0x0000016c je           LBB0_7
	//0x00000172 LBB0_10
	0x4c, 0x89, 0xc2, //0x00000172 movq         %r8, %rdx
	0x48, 0xf7, 0xd2, //0x00000175 notq         %rdx
	0x48, 0x21, 0xfa, //0x00000178 andq         %rdi, %rdx
	0x4c, 0x8d, 0x34, 0x12, //0x0000017b leaq         (%rdx,%rdx), %r14
	0x4d, 0x09, 0xc6, //0x0000017f orq          %r8, %r14
	0x4c, 0x89, 0xf6, //0x00000182 movq         %r14, %rsi
	0x48, 0xf7, 0xd6, //0x00000185 notq         %rsi
	0x48, 0x21, 0xfe, //0x00000188 andq         %rdi, %rsi
	0x48, 0xbf, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, //0x0000018b movabsq      $-6148914691236517206, %rdi
	0x48, 0x21, 0xfe, //0x00000195 andq         %rdi, %rsi
	0x45, 0x31, 0xc0, //0x00000198 xorl         %r8d, %r8d
	0x48, 0x01, 0xd6, //0x0000019b addq         %rdx, %rsi
	0x41, 0x0f, 0x92, 0xc0, //0x0000019e setb         %r8b
	0x48, 0x01, 0xf6, //0x000001a2 addq         %rsi, %rsi
	0x48, 0xba, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, //0x000001a5 movabsq      $6148914691236517205, %rdx
	0x48, 0x31, 0xd6, //0x000001af xorq         %rdx, %rsi
	0x4c, 0x21, 0xf6, //0x000001b2 andq         %r14, %rsi
	0x48, 0xf7, 0xd6, //0x000001b5 notq         %rsi
	0x48, 0x21, 0xf3, //0x000001b8 andq         %rsi, %rbx
	0x48, 0x85, 0xdb, //0x000001bb testq        %rbx, %rbx
	0x0f, 0x84, 0x7b, 0xff, 0xff, 0xff, //0x000001be je           LBB0_8
	//0x000001c4 LBB0_11
	0x48, 0x0f, 0xbc, 0xc3, //0x000001c4 bsfq         %rbx, %rax
	0x4c, 0x29, 0xc9, //0x000001c8 subq         %r9, %rcx
	0x48, 0x8d, 0x7c, 0x01, 0x01, //0x000001cb leaq         $1(%rcx,%rax), %rdi
	0x4c, 0x89, 0xe2, //0x000001d0 movq         %r12, %rdx
	0x48, 0x8b, 0x75, 0xc0, //0x000001d3 movq         $-64(%rbp), %rsi
	0xe9, 0xc8, 0x01, 0x00, 0x00, //0x000001d7 jmp          LBB0_28
	//0x000001dc LBB0_12
	0x48, 0x8b, 0x4f, 0x08, //0x000001dc movq         $8(%rdi), %rcx
	0x48, 0x89, 0xc8, //0x000001e0 movq         %rcx, %rax
	0x4c, 0x29, 0xd0, //0x000001e3 subq         %r10, %rax
	0x0f, 0x84, 0xd8, 0x02, 0x00, 0x00, //0x000001e6 je           LBB0_42
	0x48, 0x89, 0x4d, 0xd0, //0x000001ec movq         %rcx, $-48(%rbp)
	0x48, 0x89, 0x55, 0xc8, //0x000001f0 movq         %rdx, $-56(%rbp)
	0x48, 0x8b, 0x17, //0x000001f4 movq         (%rdi), %rdx
	0x4a, 0x8d, 0x0c, 0x12, //0x000001f7 leaq         (%rdx,%r10), %rcx
	0x48, 0x83, 0xf8, 0x40, //0x000001fb cmpq         $64, %rax
	0x0f, 0x82, 0xef, 0x02, 0x00, 0x00, //0x000001ff jb           LBB0_45
	0x41, 0x89, 0xc5, //0x00000205 movl         %eax, %r13d
	0x41, 0x83, 0xe5, 0x3f, //0x00000208 andl         $63, %r13d
	0x48, 0x8d, 0x78, 0xc0, //0x0000020c leaq         $-64(%rax), %rdi
	0x48, 0x83, 0xe7, 0xc0, //0x00000210 andq         $-64, %rdi
	0x4c, 0x01, 0xd7, //0x00000214 addq         %r10, %rdi
	0x48, 0x8d, 0x74, 0x3a, 0x40, //0x00000217 leaq         $64(%rdx,%rdi), %rsi
	0x48, 0x89, 0x75, 0xb0, //0x0000021c movq         %rsi, $-80(%rbp)
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x00000220 movq         $-1, %r11
	0x31, 0xff, //0x00000227 xorl         %edi, %edi
	0xc5, 0xfe, 0x6f, 0x05, 0xcf, 0xfd, 0xff, 0xff, //0x00000229 vmovdqu      $-561(%rip), %ymm0  /* LCPI0_0+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x0d, 0xe7, 0xfd, 0xff, 0xff, //0x00000231 vmovdqu      $-537(%rip), %ymm1  /* LCPI0_1+0(%rip) */
	0xc5, 0xfe, 0x6f, 0x15, 0xff, 0xfd, 0xff, 0xff, //0x00000239 vmovdqu      $-513(%rip), %ymm2  /* LCPI0_2+0(%rip) */
	0xc5, 0xe5, 0x76, 0xdb, //0x00000241 vpcmpeqd     %ymm3, %ymm3, %ymm3
	0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, 0x90, //0x00000245 .p2align 4, 0x90
	//0x00000250 LBB0_15
	0xc5, 0xfe, 0x6f, 0x21, //0x00000250 vmovdqu      (%rcx), %ymm4
	0xc5, 0xfe, 0x6f, 0x69, 0x20, //0x00000254 vmovdqu      $32(%rcx), %ymm5
	0xc5, 0xdd, 0x74, 0xf0, //0x00000259 vpcmpeqb     %ymm0, %ymm4, %ymm6
	0xc5, 0x7d, 0xd7, 0xfe, //0x0000025d vpmovmskb    %ymm6, %r15d
	0xc5, 0xd5, 0x74, 0xf0, //0x00000261 vpcmpeqb     %ymm0, %ymm5, %ymm6
	0xc5, 0xfd, 0xd7, 0xde, //0x00000265 vpmovmskb    %ymm6, %ebx
	0xc5, 0xdd, 0x74, 0xf1, //0x00000269 vpcmpeqb     %ymm1, %ymm4, %ymm6
	0xc5, 0x7d, 0xd7, 0xc6, //0x0000026d vpmovmskb    %ymm6, %r8d
	0xc5, 0xd5, 0x74, 0xf1, //0x00000271 vpcmpeqb     %ymm1, %ymm5, %ymm6
	0xc5, 0x7d, 0xd7, 0xce, //0x00000275 vpmovmskb    %ymm6, %r9d
	0xc5, 0xed, 0x64, 0xf4, //0x00000279 vpcmpgtb     %ymm4, %ymm2, %ymm6
	0xc5, 0xdd, 0x64, 0xe3, //0x0000027d vpcmpgtb     %ymm3, %ymm4, %ymm4
	0xc5, 0xdd, 0xdb, 0xe6, //0x00000281 vpand        %ymm6, %ymm4, %ymm4
	0xc5, 0x7d, 0xd7, 0xe4, //0x00000285 vpmovmskb    %ymm4, %r12d
	0xc5, 0xed, 0x64, 0xe5, //0x00000289 vpcmpgtb     %ymm5, %ymm2, %ymm4
	0xc5, 0xd5, 0x64, 0xeb, //0x0000028d vpcmpgtb     %ymm3, %ymm5, %ymm5
	0xc5, 0xd5, 0xdb, 0xe4, //0x00000291 vpand        %ymm4, %ymm5, %ymm4
	0xc5, 0x7d, 0xd7, 0xf4, //0x00000295 vpmovmskb    %ymm4, %r14d
	0x48, 0xc1, 0xe3, 0x20, //0x00000299 shlq         $32, %rbx
	0x49, 0xc1, 0xe1, 0x20, //0x0000029d shlq         $32, %r9
	0x4d, 0x09, 0xc8, //0x000002a1 orq          %r9, %r8
	0x49, 0x83, 0xfb, 0xff, //0x000002a4 cmpq         $-1, %r11
	0x0f, 0x85, 0x09, 0x00, 0x00, 0x00, //0x000002a8 jne          LBB0_17
	0x4d, 0x85, 0xc0, //0x000002ae testq        %r8, %r8
	0x0f, 0x85, 0x9a, 0x00, 0x00, 0x00, //0x000002b1 jne          LBB0_22
	//0x000002b7 LBB0_17
	0x49, 0xc1, 0xe6, 0x20, //0x000002b7 shlq         $32, %r14
	0x4c, 0x09, 0xfb, //0x000002bb orq          %r15, %rbx
	0x4c, 0x89, 0xc6, //0x000002be movq         %r8, %rsi
	0x48, 0x09, 0xfe, //0x000002c1 orq          %rdi, %rsi
	0x0f, 0x85, 0x2c, 0x00, 0x00, 0x00, //0x000002c4 jne          LBB0_21
	0x4d, 0x09, 0xe6, //0x000002ca orq          %r12, %r14
	0x48, 0x85, 0xdb, //0x000002cd testq        %rbx, %rbx
	0x0f, 0x85, 0x8d, 0x00, 0x00, 0x00, //0x000002d0 jne          LBB0_23
	//0x000002d6 LBB0_19
	0x4d, 0x85, 0xf6, //0x000002d6 testq        %r14, %r14
	0x0f, 0x85, 0x57, 0x01, 0x00, 0x00, //0x000002d9 jne          LBB0_35
	0x48, 0x83, 0xc0, 0xc0, //0x000002df addq         $-64, %rax
	0x48, 0x83, 0xc1, 0x40, //0x000002e3 addq         $64, %rcx
	0x48, 0x83, 0xf8, 0x3f, //0x000002e7 cmpq         $63, %rax
	0x0f, 0x87, 0x5f, 0xff, 0xff, 0xff, //0x000002eb ja           LBB0_15
	0xe9, 0x58, 0x01, 0x00, 0x00, //0x000002f1 jmp          LBB0_36
	//0x000002f6 LBB0_21
	0x49, 0x89, 0xff, //0x000002f6 movq         %rdi, %r15
	0x49, 0xf7, 0xd7, //0x000002f9 notq         %r15
	0x4d, 0x21, 0xc7, //0x000002fc andq         %r8, %r15
	0x4f, 0x8d, 0x0c, 0x3f, //0x000002ff leaq         (%r15,%r15), %r9
	0x49, 0x09, 0xf9, //0x00000303 orq          %rdi, %r9
	0x4c, 0x89, 0x4d, 0xb8, //0x00000306 movq         %r9, $-72(%rbp)
	0x49, 0xf7, 0xd1, //0x0000030a notq         %r9
	0x4d, 0x21, 0xc1, //0x0000030d andq         %r8, %r9
	0x48, 0xbf, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, 0xaa, //0x00000310 movabsq      $-6148914691236517206, %rdi
	0x49, 0x21, 0xf9, //0x0000031a andq         %rdi, %r9
	0x31, 0xff, //0x0000031d xorl         %edi, %edi
	0x4d, 0x01, 0xf9, //0x0000031f addq         %r15, %r9
	0x40, 0x0f, 0x92, 0xc7, //0x00000322 setb         %dil
	0x4d, 0x01, 0xc9, //0x00000326 addq         %r9, %r9
	0x48, 0xbe, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, 0x55, //0x00000329 movabsq      $6148914691236517205, %rsi
	0x49, 0x31, 0xf1, //0x00000333 xorq         %rsi, %r9
	0x4c, 0x23, 0x4d, 0xb8, //0x00000336 andq         $-72(%rbp), %r9
	0x49, 0xf7, 0xd1, //0x0000033a notq         %r9
	0x4c, 0x21, 0xcb, //0x0000033d andq         %r9, %rbx
	0x4d, 0x09, 0xe6, //0x00000340 orq          %r12, %r14
	0x48, 0x85, 0xdb, //0x00000343 testq        %rbx, %rbx
	0x0f, 0x84, 0x8a, 0xff, 0xff, 0xff, //0x00000346 je           LBB0_19
	0xe9, 0x12, 0x00, 0x00, 0x00, //0x0000034c jmp          LBB0_23
	//0x00000351 LBB0_22
	0x49, 0x89, 0xc9, //0x00000351 movq         %rcx, %r9
	0x49, 0x29, 0xd1, //0x00000354 subq         %rdx, %r9
	0x4d, 0x0f, 0xbc, 0xd8, //0x00000357 bsfq         %r8, %r11
	0x4d, 0x01, 0xcb, //0x0000035b addq         %r9, %r11
	0xe9, 0x54, 0xff, 0xff, 0xff, //0x0000035e jmp          LBB0_17
	//0x00000363 LBB0_23
	0x48, 0x0f, 0xbc, 0xc3, //0x00000363 bsfq         %rbx, %rax
	0x4d, 0x85, 0xf6, //0x00000367 testq        %r14, %r14
	0x0f, 0x84, 0x16, 0x00, 0x00, 0x00, //0x0000036a je           LBB0_25
	0x49, 0x0f, 0xbc, 0xfe, //0x00000370 bsfq         %r14, %rdi
	0x48, 0x8b, 0x75, 0xc0, //0x00000374 movq         $-64(%rbp), %rsi
	0x48, 0x39, 0xc7, //0x00000378 cmpq         %rax, %rdi
	0x0f, 0x83, 0x17, 0x00, 0x00, 0x00, //0x0000037b jae          LBB0_26
	0xe9, 0xe9, 0x03, 0x00, 0x00, //0x00000381 jmp          LBB0_82
	//0x00000386 LBB0_25
	0xbf, 0x40, 0x00, 0x00, 0x00, //0x00000386 movl         $64, %edi
	0x48, 0x8b, 0x75, 0xc0, //0x0000038b movq         $-64(%rbp), %rsi
	0x48, 0x39, 0xc7, //0x0000038f cmpq         %rax, %rdi
	0x0f, 0x82, 0xd7, 0x03, 0x00, 0x00, //0x00000392 jb           LBB0_82
	//0x00000398 LBB0_26
	0x48, 0x29, 0xd1, //0x00000398 subq         %rdx, %rcx
	0x48, 0x8d, 0x7c, 0x01, 0x01, //0x0000039b leaq         $1(%rcx,%rax), %rdi
	//0x000003a0 LBB0_27
	0x48, 0x8b, 0x55, 0xc8, //0x000003a0 movq         $-56(%rbp), %rdx
	//0x000003a4 LBB0_28
	0x48, 0x85, 0xff, //0x000003a4 testq        %rdi, %rdi
	0x0f, 0x88, 0x14, 0x04, 0x00, 0x00, //0x000003a7 js           LBB0_86
	0x48, 0x89, 0x3e, //0x000003ad movq         %rdi, (%rsi)
	0x4c, 0x89, 0x52, 0x10, //0x000003b0 movq         %r10, $16(%rdx)
	0x48, 0xc7, 0x02, 0x07, 0x00, 0x00, 0x00, //0x000003b4 movq         $7, (%rdx)
	0x49, 0x39, 0xfb, //0x000003bb cmpq         %rdi, %r11
	0x48, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x000003be movq         $-1, %rax
	0x49, 0x0f, 0x4c, 0xc3, //0x000003c5 cmovlq       %r11, %rax
	0x48, 0x89, 0x42, 0x18, //0x000003c9 movq         %rax, $24(%rdx)
	0xe9, 0xf9, 0x03, 0x00, 0x00, //0x000003cd jmp          LBB0_88
	//0x000003d2 LBB0_30
	0x4d, 0x89, 0xde, //0x000003d2 movq         %r11, %r14
	0x4c, 0x89, 0xe8, //0x000003d5 movq         %r13, %rax
	0x4c, 0x89, 0xf9, //0x000003d8 movq         %r15, %rcx
	0x4c, 0x89, 0xe2, //0x000003db movq         %r12, %rdx
	0x48, 0x83, 0xf8, 0x20, //0x000003de cmpq         $32, %rax
	0x0f, 0x82, 0x03, 0x01, 0x00, 0x00, //0x000003e2 jb           LBB0_44
	//0x000003e8 LBB0_31
	0xc5, 0xfe, 0x6f, 0x01, //0x000003e8 vmovdqu      (%rcx), %ymm0
	0xc5, 0xfd, 0x74, 0x0d, 0x0c, 0xfc, 0xff, 0xff, //0x000003ec vpcmpeqb     $-1012(%rip), %ymm0, %ymm1  /* LCPI0_0+0(%rip) */
	0xc5, 0xfd, 0xd7, 0xd9, //0x000003f4 vpmovmskb    %ymm1, %ebx
	0xc5, 0xfd, 0x74, 0x05, 0x20, 0xfc, 0xff, 0xff, //0x000003f8 vpcmpeqb     $-992(%rip), %ymm0, %ymm0  /* LCPI0_1+0(%rip) */
	0xc5, 0xfd, 0xd7, 0xf8, //0x00000400 vpmovmskb    %ymm0, %edi
	0x85, 0xff, //0x00000404 testl        %edi, %edi
	0x0f, 0x85, 0x07, 0x01, 0x00, 0x00, //0x00000406 jne          LBB0_46
	0x4d, 0x85, 0xc0, //0x0000040c testq        %r8, %r8
	0x48, 0x8b, 0x75, 0xc0, //0x0000040f movq         $-64(%rbp), %rsi
	0x0f, 0x85, 0x1c, 0x01, 0x00, 0x00, //0x00000413 jne          LBB0_48
	0x45, 0x31, 0xc0, //0x00000419 xorl         %r8d, %r8d
	0x48, 0x85, 0xdb, //0x0000041c testq        %rbx, %rbx
	0x0f, 0x84, 0x5b, 0x01, 0x00, 0x00, //0x0000041f je           LBB0_50
	//0x00000425 LBB0_34
	0x48, 0x0f, 0xbc, 0xc3, //0x00000425 bsfq         %rbx, %rax
	0x4c, 0x29, 0xc9, //0x00000429 subq         %r9, %rcx
	0x48, 0x8d, 0x7c, 0x01, 0x01, //0x0000042c leaq         $1(%rcx,%rax), %rdi
	0xe9, 0x6e, 0xff, 0xff, 0xff, //0x00000431 jmp          LBB0_28
	//0x00000436 LBB0_35
	0x48, 0xc7, 0xc7, 0xfe, 0xff, 0xff, 0xff, //0x00000436 movq         $-2, %rdi
	0x4c, 0x8b, 0x55, 0xd0, //0x0000043d movq         $-48(%rbp), %r10
	0x48, 0x8b, 0x55, 0xc8, //0x00000441 movq         $-56(%rbp), %rdx
	0x48, 0x8b, 0x75, 0xc0, //0x00000445 movq         $-64(%rbp), %rsi
	0xe9, 0x77, 0x03, 0x00, 0x00, //0x00000449 jmp          LBB0_87
	//0x0000044e LBB0_36
	0x4d, 0x89, 0xde, //0x0000044e movq         %r11, %r14
	0x4c, 0x89, 0xe8, //0x00000451 movq         %r13, %rax
	0x48, 0x8b, 0x4d, 0xb0, //0x00000454 movq         $-80(%rbp), %rcx
	0x48, 0x8b, 0x75, 0xc0, //0x00000458 movq         $-64(%rbp), %rsi
	0x48, 0x83, 0xf8, 0x20, //0x0000045c cmpq         $32, %rax
	0x0f, 0x82, 0x4d, 0x02, 0x00, 0x00, //0x00000460 jb           LBB0_70
	//0x00000466 LBB0_37
	0xc5, 0xfe, 0x6f, 0x01, //0x00000466 vmovdqu      (%rcx), %ymm0
	0xc5, 0xfd, 0x74, 0x0d, 0x8e, 0xfb, 0xff, 0xff, //0x0000046a vpcmpeqb     $-1138(%rip), %ymm0, %ymm1  /* LCPI0_0+0(%rip) */
	0xc5, 0x7d, 0xd7, 0xe1, //0x00000472 vpmovmskb    %ymm1, %r12d
	0xc5, 0xfd, 0x74, 0x0d, 0xa2, 0xfb, 0xff, 0xff, //0x00000476 vpcmpeqb     $-1118(%rip), %ymm0, %ymm1  /* LCPI0_1+0(%rip) */
	0xc5, 0x7d, 0xd7, 0xc1, //0x0000047e vpmovmskb    %ymm1, %r8d
	0xc5, 0xfe, 0x6f, 0x0d, 0xb6, 0xfb, 0xff, 0xff, //0x00000482 vmovdqu      $-1098(%rip), %ymm1  /* LCPI0_2+0(%rip) */
	0xc5, 0xf5, 0x64, 0xc8, //0x0000048a vpcmpgtb     %ymm0, %ymm1, %ymm1
	0xc5, 0xed, 0x76, 0xd2, //0x0000048e vpcmpeqd     %ymm2, %ymm2, %ymm2
	0xc5, 0xfd, 0x64, 0xc2, //0x00000492 vpcmpgtb     %ymm2, %ymm0, %ymm0
	0xc5, 0xfd, 0xdb, 0xc1, //0x00000496 vpand        %ymm1, %ymm0, %ymm0
	0xc5, 0x7d, 0xd7, 0xf8, //0x0000049a vpmovmskb    %ymm0, %r15d
	0x45, 0x85, 0xc0, //0x0000049e testl        %r8d, %r8d
	0x0f, 0x85, 0x71, 0x01, 0x00, 0x00, //0x000004a1 jne          LBB0_60
	0x48, 0x85, 0xff, //0x000004a7 testq        %rdi, %rdi
	0x0f, 0x85, 0x8a, 0x01, 0x00, 0x00, //0x000004aa jne          LBB0_62
	0x31, 0xff, //0x000004b0 xorl         %edi, %edi
	0x4d, 0x85, 0xe4, //0x000004b2 testq        %r12, %r12
	0x0f, 0x84, 0xc7, 0x01, 0x00, 0x00, //0x000004b5 je           LBB0_64
	//0x000004bb LBB0_40
	0x49, 0x0f, 0xbc, 0xdc, //0x000004bb bsfq         %r12, %rbx
	0xe9, 0xc3, 0x01, 0x00, 0x00, //0x000004bf jmp          LBB0_65
	//0x000004c4 LBB0_42
	0x48, 0xc7, 0xc7, 0xff, 0xff, 0xff, 0xff, //0x000004c4 movq         $-1, %rdi
	0xe9, 0xf5, 0x02, 0x00, 0x00, //0x000004cb jmp          LBB0_87
	//0x000004d0 LBB0_43
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x000004d0 movq         $-1, %r11
	0x45, 0x31, 0xc0, //0x000004d7 xorl         %r8d, %r8d
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x000004da movq         $-1, %r14
	0x48, 0x83, 0xf8, 0x20, //0x000004e1 cmpq         $32, %rax
	0x0f, 0x83, 0xfd, 0xfe, 0xff, 0xff, //0x000004e5 jae          LBB0_31
	//0x000004eb LBB0_44
	0x48, 0x8b, 0x75, 0xc0, //0x000004eb movq         $-64(%rbp), %rsi
	0xe9, 0x94, 0x00, 0x00, 0x00, //0x000004ef jmp          LBB0_51
	//0x000004f4 LBB0_45
	0x49, 0xc7, 0xc3, 0xff, 0xff, 0xff, 0xff, //0x000004f4 movq         $-1, %r11
	0x31, 0xff, //0x000004fb xorl         %edi, %edi
	0x49, 0xc7, 0xc6, 0xff, 0xff, 0xff, 0xff, //0x000004fd movq         $-1, %r14
	0x48, 0x83, 0xf8, 0x20, //0x00000504 cmpq         $32, %rax
	0x0f, 0x83, 0x58, 0xff, 0xff, 0xff, //0x00000508 jae          LBB0_37
	0xe9, 0xa0, 0x01, 0x00, 0x00, //0x0000050e jmp          LBB0_70
	//0x00000513 LBB0_46
	0x49, 0x89, 0xd4, //0x00000513 movq         %rdx, %r12
	0x49, 0x83, 0xfe, 0xff, //0x00000516 cmpq         $-1, %r14
	0x0f, 0x85, 0x18, 0x00, 0x00, 0x00, //0x0000051a jne          LBB0_49
	0x48, 0x89, 0xca, //0x00000520 movq         %rcx, %rdx
	0x4c, 0x29, 0xca, //0x00000523 subq         %r9, %rdx
	0x4c, 0x0f, 0xbc, 0xdf, //0x00000526 bsfq         %rdi, %r11
	0x49, 0x01, 0xd3, //0x0000052a addq         %rdx, %r11
	0x4d, 0x89, 0xde, //0x0000052d movq         %r11, %r14
	0xe9, 0x03, 0x00, 0x00, 0x00, //0x00000530 jmp          LBB0_49
	//0x00000535 LBB0_48
	0x49, 0x89, 0xd4, //0x00000535 movq         %rdx, %r12
	//0x00000538 LBB0_49
	0x44, 0x89, 0xc2, //0x00000538 movl         %r8d, %edx
	0xf7, 0xd2, //0x0000053b notl         %edx
	0x21, 0xfa, //0x0000053d andl         %edi, %edx
	0x44, 0x8d, 0x3c, 0x12, //0x0000053f leal         (%rdx,%rdx), %r15d
	0x45, 0x09, 0xc7, //0x00000543 orl          %r8d, %r15d
	0x44, 0x89, 0xfe, //0x00000546 movl         %r15d, %esi
	0xf7, 0xd6, //0x00000549 notl         %esi
	0x21, 0xfe, //0x0000054b andl         %edi, %esi
	0x81, 0xe6, 0xaa, 0xaa, 0xaa, 0xaa, //0x0000054d andl         $-1431655766, %esi
	0x45, 0x31, 0xc0, //0x00000553 xorl         %r8d, %r8d
	0x01, 0xd6, //0x00000556 addl         %edx, %esi
	0x41, 0x0f, 0x92, 0xc0, //0x00000558 setb         %r8b
	0x01, 0xf6, //0x0000055c addl         %esi, %esi
	0x81, 0xf6, 0x55, 0x55, 0x55, 0x55, //0x0000055e xorl         $1431655765, %esi
	0x44, 0x21, 0xfe, //0x00000564 andl         %r15d, %esi
	0xba, 0xff, 0xff, 0xff, 0xff, //0x00000567 movl         $4294967295, %edx
	0x31, 0xf2, //0x0000056c xorl         %esi, %edx
	0x21, 0xd3, //0x0000056e andl         %edx, %ebx
	0x4c, 0x89, 0xe2, //0x00000570 movq         %r12, %rdx
	0x48, 0x8b, 0x75, 0xc0, //0x00000573 movq         $-64(%rbp), %rsi
	0x48, 0x85, 0xdb, //0x00000577 testq        %rbx, %rbx
	0x0f, 0x85, 0xa5, 0xfe, 0xff, 0xff, //0x0000057a jne          LBB0_34
	//0x00000580 LBB0_50
	0x48, 0x83, 0xc1, 0x20, //0x00000580 addq         $32, %rcx
	0x48, 0x83, 0xc0, 0xe0, //0x00000584 addq         $-32, %rax
	//0x00000588 LBB0_51
	0x4d, 0x85, 0xc0, //0x00000588 testq        %r8, %r8
	0x0f, 0x85, 0xf2, 0x01, 0x00, 0x00, //0x0000058b jne          LBB0_84
	0x48, 0xc7, 0xc7, 0xff, 0xff, 0xff, 0xff, //0x00000591 movq         $-1, %rdi
	0x48, 0x85, 0xc0, //0x00000598 testq        %rax, %rax
	0x0f, 0x84, 0x20, 0x02, 0x00, 0x00, //0x0000059b je           LBB0_86
	//0x000005a1 LBB0_53
	0x4d, 0x89, 0xc8, //0x000005a1 movq         %r9, %r8
	0x49, 0xf7, 0xd0, //0x000005a4 notq         %r8
	//0x000005a7 LBB0_54
	0x4c, 0x8d, 0x61, 0x01, //0x000005a7 leaq         $1(%rcx), %r12
	0x0f, 0xb6, 0x19, //0x000005ab movzbl       (%rcx), %ebx
	0x80, 0xfb, 0x22, //0x000005ae cmpb         $34, %bl
	0x0f, 0x84, 0x56, 0x00, 0x00, 0x00, //0x000005b1 je           LBB0_59
	0x4c, 0x8d, 0x78, 0xff, //0x000005b7 leaq         $-1(%rax), %r15
	0x80, 0xfb, 0x5c, //0x000005bb cmpb         $92, %bl
	0x0f, 0x84, 0x14, 0x00, 0x00, 0x00, //0x000005be je           LBB0_57
	0x4c, 0x89, 0xf8, //0x000005c4 movq         %r15, %rax
	0x4c, 0x89, 0xe1, //0x000005c7 movq         %r12, %rcx
	0x4d, 0x85, 0xff, //0x000005ca testq        %r15, %r15
	0x0f, 0x85, 0xd4, 0xff, 0xff, 0xff, //0x000005cd jne          LBB0_54
	0xe9, 0xe9, 0x01, 0x00, 0x00, //0x000005d3 jmp          LBB0_86
	//0x000005d8 LBB0_57
	0x4d, 0x85, 0xff, //0x000005d8 testq        %r15, %r15
	0x0f, 0x84, 0x69, 0x02, 0x00, 0x00, //0x000005db je           LBB0_93
	0x4d, 0x01, 0xc4, //0x000005e1 addq         %r8, %r12
	0x49, 0x83, 0xfe, 0xff, //0x000005e4 cmpq         $-1, %r14
	0x4d, 0x0f, 0x44, 0xdc, //0x000005e8 cmoveq       %r12, %r11
	0x4d, 0x0f, 0x44, 0xf4, //0x000005ec cmoveq       %r12, %r14
	0x48, 0x83, 0xc1, 0x02, //0x000005f0 addq         $2, %rcx
	0x48, 0x83, 0xc0, 0xfe, //0x000005f4 addq         $-2, %rax
	0x49, 0x89, 0xc7, //0x000005f8 movq         %rax, %r15
	0x48, 0x8b, 0x75, 0xc0, //0x000005fb movq         $-64(%rbp), %rsi
	0x4d, 0x85, 0xff, //0x000005ff testq        %r15, %r15
	0x0f, 0x85, 0x9f, 0xff, 0xff, 0xff, //0x00000602 jne          LBB0_54
	0xe9, 0xb4, 0x01, 0x00, 0x00, //0x00000608 jmp          LBB0_86
	//0x0000060d LBB0_59
	0x4d, 0x29, 0xcc, //0x0000060d subq         %r9, %r12
	0x4c, 0x89, 0xe7, //0x00000610 movq         %r12, %rdi
	0xe9, 0x8c, 0xfd, 0xff, 0xff, //0x00000613 jmp          LBB0_28
	//0x00000618 LBB0_60
	0x49, 0x89, 0xf5, //0x00000618 movq         %rsi, %r13
	0x49, 0x83, 0xfe, 0xff, //0x0000061b cmpq         $-1, %r14
	0x0f, 0x85, 0x18, 0x00, 0x00, 0x00, //0x0000061f jne          LBB0_63
	0x48, 0x89, 0xce, //0x00000625 movq         %rcx, %rsi
	0x48, 0x29, 0xd6, //0x00000628 subq         %rdx, %rsi
	0x4d, 0x0f, 0xbc, 0xd8, //0x0000062b bsfq         %r8, %r11
	0x49, 0x01, 0xf3, //0x0000062f addq         %rsi, %r11
	0x4d, 0x89, 0xde, //0x00000632 movq         %r11, %r14
	0xe9, 0x03, 0x00, 0x00, 0x00, //0x00000635 jmp          LBB0_63
	//0x0000063a LBB0_62
	0x49, 0x89, 0xf5, //0x0000063a movq         %rsi, %r13
	//0x0000063d LBB0_63
	0x89, 0xfe, //0x0000063d movl         %edi, %esi
	0xf7, 0xd6, //0x0000063f notl         %esi
	0x44, 0x21, 0xc6, //0x00000641 andl         %r8d, %esi
	0x44, 0x8d, 0x0c, 0x36, //0x00000644 leal         (%rsi,%rsi), %r9d
	0x41, 0x09, 0xf9, //0x00000648 orl          %edi, %r9d
	0x44, 0x89, 0xcb, //0x0000064b movl         %r9d, %ebx
	0xf7, 0xd3, //0x0000064e notl         %ebx
	0x44, 0x21, 0xc3, //0x00000650 andl         %r8d, %ebx
	0x81, 0xe3, 0xaa, 0xaa, 0xaa, 0xaa, //0x00000653 andl         $-1431655766, %ebx
	0x31, 0xff, //0x00000659 xorl         %edi, %edi
	0x01, 0xf3, //0x0000065b addl         %esi, %ebx
	0x40, 0x0f, 0x92, 0xc7, //0x0000065d setb         %dil
	0x01, 0xdb, //0x00000661 addl         %ebx, %ebx
	0x81, 0xf3, 0x55, 0x55, 0x55, 0x55, //0x00000663 xorl         $1431655765, %ebx
	0x44, 0x21, 0xcb, //0x00000669 andl         %r9d, %ebx
	0xbe, 0xff, 0xff, 0xff, 0xff, //0x0000066c movl         $4294967295, %esi
	0x31, 0xde, //0x00000671 xorl         %ebx, %esi
	0x41, 0x21, 0xf4, //0x00000673 andl         %esi, %r12d
	0x4c, 0x89, 0xee, //0x00000676 movq         %r13, %rsi
	0x4d, 0x85, 0xe4, //0x00000679 testq        %r12, %r12
	0x0f, 0x85, 0x39, 0xfe, 0xff, 0xff, //0x0000067c jne          LBB0_40
	//0x00000682 LBB0_64
	0xbb, 0x40, 0x00, 0x00, 0x00, //0x00000682 movl         $64, %ebx
	//0x00000687 LBB0_65
	0x4d, 0x85, 0xe4, //0x00000687 testq        %r12, %r12
	0x0f, 0x84, 0x12, 0x00, 0x00, 0x00, //0x0000068a je           LBB0_68
	0x45, 0x85, 0xff, //0x00000690 testl        %r15d, %r15d
	0x0f, 0x84, 0xad, 0x00, 0x00, 0x00, //0x00000693 je           LBB0_78
	0x49, 0x0f, 0xbc, 0xc7, //0x00000699 bsfq         %r15, %rax
	0xe9, 0xa9, 0x00, 0x00, 0x00, //0x0000069d jmp          LBB0_79
	//0x000006a2 LBB0_68
	0x45, 0x85, 0xff, //0x000006a2 testl        %r15d, %r15d
	0x0f, 0x85, 0xc4, 0x00, 0x00, 0x00, //0x000006a5 jne          LBB0_82
	0x48, 0x83, 0xc1, 0x20, //0x000006ab addq         $32, %rcx
	0x48, 0x83, 0xc0, 0xe0, //0x000006af addq         $-32, %rax
	//0x000006b3 LBB0_70
	0x48, 0x85, 0xff, //0x000006b3 testq        %rdi, %rdi
	0x0f, 0x85, 0x21, 0x01, 0x00, 0x00, //0x000006b6 jne          LBB0_89
	0x48, 0xc7, 0xc7, 0xff, 0xff, 0xff, 0xff, //0x000006bc movq         $-1, %rdi
	0x48, 0x85, 0xc0, //0x000006c3 testq        %rax, %rax
	0x0f, 0x84, 0xaa, 0x00, 0x00, 0x00, //0x000006c6 je           LBB0_83
	//0x000006cc LBB0_72
	0x0f, 0xb6, 0x19, //0x000006cc movzbl       (%rcx), %ebx
	0x80, 0xfb, 0x22, //0x000006cf cmpb         $34, %bl
	0x0f, 0x84, 0x89, 0x00, 0x00, 0x00, //0x000006d2 je           LBB0_81
	0x80, 0xfb, 0x5c, //0x000006d8 cmpb         $92, %bl
	0x0f, 0x84, 0x26, 0x00, 0x00, 0x00, //0x000006db je           LBB0_76
	0x80, 0xfb, 0x20, //0x000006e1 cmpb         $32, %bl
	0x0f, 0x82, 0x85, 0x00, 0x00, 0x00, //0x000006e4 jb           LBB0_82
	0x49, 0xc7, 0xc0, 0xff, 0xff, 0xff, 0xff, //0x000006ea movq         $-1, %r8
	0xbb, 0x01, 0x00, 0x00, 0x00, //0x000006f1 movl         $1, %ebx
	0x48, 0x01, 0xd9, //0x000006f6 addq         %rbx, %rcx
	0x4c, 0x01, 0xc0, //0x000006f9 addq         %r8, %rax
	0x0f, 0x85, 0xca, 0xff, 0xff, 0xff, //0x000006fc jne          LBB0_72
	0xe9, 0x6f, 0x00, 0x00, 0x00, //0x00000702 jmp          LBB0_83
	//0x00000707 LBB0_76
	0x49, 0x89, 0xf1, //0x00000707 movq         %rsi, %r9
	0x48, 0x83, 0xf8, 0x01, //0x0000070a cmpq         $1, %rax
	0x0f, 0x84, 0x43, 0x01, 0x00, 0x00, //0x0000070e je           LBB0_94
	0x48, 0x89, 0xce, //0x00000714 movq         %rcx, %rsi
	0x48, 0x29, 0xd6, //0x00000717 subq         %rdx, %rsi
	0x49, 0x83, 0xfe, 0xff, //0x0000071a cmpq         $-1, %r14
	0x4c, 0x0f, 0x44, 0xde, //0x0000071e cmoveq       %rsi, %r11
	0x4c, 0x0f, 0x44, 0xf6, //0x00000722 cmoveq       %rsi, %r14
	0x49, 0xc7, 0xc0, 0xfe, 0xff, 0xff, 0xff, //0x00000726 movq         $-2, %r8
	0xbb, 0x02, 0x00, 0x00, 0x00, //0x0000072d movl         $2, %ebx
	0x4c, 0x89, 0xce, //0x00000732 movq         %r9, %rsi
	0x48, 0x01, 0xd9, //0x00000735 addq         %rbx, %rcx
	0x4c, 0x01, 0xc0, //0x00000738 addq         %r8, %rax
	0x0f, 0x85, 0x8b, 0xff, 0xff, 0xff, //0x0000073b jne          LBB0_72
	0xe9, 0x30, 0x00, 0x00, 0x00, //0x00000741 jmp          LBB0_83
	//0x00000746 LBB0_78
	0xb8, 0x40, 0x00, 0x00, 0x00, //0x00000746 movl         $64, %eax
	//0x0000074b LBB0_79
	0x48, 0x39, 0xd8, //0x0000074b cmpq         %rbx, %rax
	0x0f, 0x82, 0x1b, 0x00, 0x00, 0x00, //0x0000074e jb           LBB0_82
	0x48, 0x29, 0xd1, //0x00000754 subq         %rdx, %rcx
	0x48, 0x8d, 0x7c, 0x19, 0x01, //0x00000757 leaq         $1(%rcx,%rbx), %rdi
	0xe9, 0x3f, 0xfc, 0xff, 0xff, //0x0000075c jmp          LBB0_27
	//0x00000761 LBB0_81
	0x48, 0x29, 0xd1, //0x00000761 subq         %rdx, %rcx
	0x48, 0xff, 0xc1, //0x00000764 incq         %rcx
	0x48, 0x89, 0xcf, //0x00000767 movq         %rcx, %rdi
	0xe9, 0x31, 0xfc, 0xff, 0xff, //0x0000076a jmp          LBB0_27
	//0x0000076f LBB0_82
	0x48, 0xc7, 0xc7, 0xfe, 0xff, 0xff, 0xff, //0x0000076f movq         $-2, %rdi
	//0x00000776 LBB0_83
	0x4c, 0x8b, 0x55, 0xd0, //0x00000776 movq         $-48(%rbp), %r10
	0x48, 0x8b, 0x55, 0xc8, //0x0000077a movq         $-56(%rbp), %rdx
	0xe9, 0x42, 0x00, 0x00, 0x00, //0x0000077e jmp          LBB0_87
	//0x00000783 LBB0_84
	0x48, 0x89, 0xd6, //0x00000783 movq         %rdx, %rsi
	0x48, 0x85, 0xc0, //0x00000786 testq        %rax, %rax
	0x0f, 0x84, 0x8d, 0x00, 0x00, 0x00, //0x00000789 je           LBB0_91
	0x4c, 0x89, 0xca, //0x0000078f movq         %r9, %rdx
	0x48, 0xf7, 0xd2, //0x00000792 notq         %rdx
	0x48, 0x01, 0xca, //0x00000795 addq         %rcx, %rdx
	0x49, 0x83, 0xfe, 0xff, //0x00000798 cmpq         $-1, %r14
	0x4c, 0x0f, 0x44, 0xda, //0x0000079c cmoveq       %rdx, %r11
	0x4c, 0x0f, 0x44, 0xf2, //0x000007a0 cmoveq       %rdx, %r14
	0x48, 0xff, 0xc1, //0x000007a4 incq         %rcx
	0x48, 0xff, 0xc8, //0x000007a7 decq         %rax
	0x48, 0x89, 0xf2, //0x000007aa movq         %rsi, %rdx
	0x48, 0x8b, 0x75, 0xc0, //0x000007ad movq         $-64(%rbp), %rsi
	0x48, 0xc7, 0xc7, 0xff, 0xff, 0xff, 0xff, //0x000007b1 movq         $-1, %rdi
	0x48, 0x85, 0xc0, //0x000007b8 testq        %rax, %rax
	0x0f, 0x85, 0xe0, 0xfd, 0xff, 0xff, //0x000007bb jne          LBB0_53
	//0x000007c1 LBB0_86
	0x4c, 0x8b, 0x55, 0xd0, //0x000007c1 movq         $-48(%rbp), %r10
	//0x000007c5 LBB0_87
	0x4c, 0x89, 0x16, //0x000007c5 movq         %r10, (%rsi)
	0x48, 0x89, 0x3a, //0x000007c8 movq         %rdi, (%rdx)
	//0x000007cb LBB0_88
	0x48, 0x83, 0xc4, 0x28, //0x000007cb addq         $40, %rsp
	0x5b, //0x000007cf popq         %rbx
	0x41, 0x5c, //0x000007d0 popq         %r12
	0x41, 0x5d, //0x000007d2 popq         %r13
	0x41, 0x5e, //0x000007d4 popq         %r14
	0x41, 0x5f, //0x000007d6 popq         %r15
	0x5d, //0x000007d8 popq         %rbp
	0xc5, 0xf8, 0x77, //0x000007d9 vzeroupper   
	0xc3, //0x000007dc retq         
	//0x000007dd LBB0_89
	0x48, 0x89, 0xf3, //0x000007dd movq         %rsi, %rbx
	0x48, 0x85, 0xc0, //0x000007e0 testq        %rax, %rax
	0x0f, 0x84, 0x4a, 0x00, 0x00, 0x00, //0x000007e3 je           LBB0_92
	0x48, 0x89, 0xd6, //0x000007e9 movq         %rdx, %rsi
	0x48, 0xf7, 0xd6, //0x000007ec notq         %rsi
	0x48, 0x01, 0xce, //0x000007ef addq         %rcx, %rsi
	0x49, 0x83, 0xfe, 0xff, //0x000007f2 cmpq         $-1, %r14
	0x4c, 0x0f, 0x44, 0xde, //0x000007f6 cmoveq       %rsi, %r11
	0x4c, 0x0f, 0x44, 0xf6, //0x000007fa cmoveq       %rsi, %r14
	0x48, 0xff, 0xc1, //0x000007fe incq         %rcx
	0x48, 0xff, 0xc8, //0x00000801 decq         %rax
	0x48, 0x89, 0xde, //0x00000804 movq         %rbx, %rsi
	0x48, 0xc7, 0xc7, 0xff, 0xff, 0xff, 0xff, //0x00000807 movq         $-1, %rdi
	0x48, 0x85, 0xc0, //0x0000080e testq        %rax, %rax
	0x0f, 0x85, 0xb5, 0xfe, 0xff, 0xff, //0x00000811 jne          LBB0_72
	0xe9, 0x5a, 0xff, 0xff, 0xff, //0x00000817 jmp          LBB0_83
	//0x0000081c LBB0_91
	0x48, 0xc7, 0xc7, 0xff, 0xff, 0xff, 0xff, //0x0000081c movq         $-1, %rdi
	0x4c, 0x8b, 0x55, 0xd0, //0x00000823 movq         $-48(%rbp), %r10
	0x48, 0x89, 0xf2, //0x00000827 movq         %rsi, %rdx
	0x48, 0x8b, 0x75, 0xc0, //0x0000082a movq         $-64(%rbp), %rsi
	0xe9, 0x92, 0xff, 0xff, 0xff, //0x0000082e jmp          LBB0_87
	//0x00000833 LBB0_92
	0x48, 0xc7, 0xc7, 0xff, 0xff, 0xff, 0xff, //0x00000833 movq         $-1, %rdi
	0x4c, 0x8b, 0x55, 0xd0, //0x0000083a movq         $-48(%rbp), %r10
	0x48, 0x8b, 0x55, 0xc8, //0x0000083e movq         $-56(%rbp), %rdx
	0x48, 0x89, 0xde, //0x00000842 movq         %rbx, %rsi
	0xe9, 0x7b, 0xff, 0xff, 0xff, //0x00000845 jmp          LBB0_87
	//0x0000084a LBB0_93
	0x4c, 0x8b, 0x55, 0xd0, //0x0000084a movq         $-48(%rbp), %r10
	0x48, 0x8b, 0x75, 0xc0, //0x0000084e movq         $-64(%rbp), %rsi
	0xe9, 0x6e, 0xff, 0xff, 0xff, //0x00000852 jmp          LBB0_87
	//0x00000857 LBB0_94
	0x4c, 0x8b, 0x55, 0xd0, //0x00000857 movq         $-48(%rbp), %r10
	0x48, 0x8b, 0x55, 0xc8, //0x0000085b movq         $-56(%rbp), %rdx
	0x4c, 0x89, 0xce, //0x0000085f movq         %r9, %rsi
	0xe9, 0x5e, 0xff, 0xff, 0xff, //0x00000862 jmp          LBB0_87
	0x00, //0x00000867 .p2align 2, 0x00
	//0x00000868 _MASK_USE_NUMBER
	0x02, 0x00, 0x00, 0x00, //0x00000868 .long 2
}
 