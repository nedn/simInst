simInst
=================
A simple compiler that compile a simple assembly-like language to brainfuck.

Instruction sets:

+ There are 16 registers: r1 -> r16.
+ Labels are strings that contains only capital alphabetical character followed by a colon.
+ r denotes register, k denotes integer, L denotes label.

inc     r   k     : r -> r + k
dec     r   k     : r -> r - k
jz      r   L     : jump to label L if r == 0
putmem  r1  r2    : put the value in r1 to the address r2
getmem  r1  r2    : copy the value in address r2 to r1
disp    r         : display the content of r.

