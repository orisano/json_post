# json_post
Benchmark of way of json post on go.

```
goos: darwin
goarch: amd64
pkg: github.com/orisano/json_post
BenchmarkSmall/MarshalE______________-4         	   10000	    132701 ns/op	    3253 B/op	      56 allocs/op
BenchmarkSmall/EncodeDefaultBuffer___-4         	   10000	    129992 ns/op	    3232 B/op	      56 allocs/op
BenchmarkSmall/EncodeDefaultBufferE__-4         	   10000	    125287 ns/op	    3231 B/op	      56 allocs/op
BenchmarkSmall/EncodeNewBuffer_______-4         	   10000	    122913 ns/op	    3231 B/op	      56 allocs/op
BenchmarkSmall/EncodeNewBufferE______-4         	   10000	    127185 ns/op	    3231 B/op	      56 allocs/op
BenchmarkSmall/EncodeNewNilBuffer____-4         	   10000	    121043 ns/op	    3231 B/op	      56 allocs/op
BenchmarkSmall/EncodeNewNilBufferE___-4         	   10000	    128071 ns/op	    3232 B/op	      56 allocs/op
BenchmarkSmall/EncodeReservedBuffer__-4         	   10000	    136702 ns/op	    4143 B/op	      56 allocs/op
BenchmarkSmall/EncodeReservedBufferE_-4         	   10000	    136559 ns/op	    4143 B/op	      56 allocs/op
BenchmarkSmall/EncodeCheatBuffer_____-4         	   10000	    142124 ns/op	    3247 B/op	      56 allocs/op
BenchmarkSmall/EncodeCheatBufferE____-4         	   10000	    133541 ns/op	    3248 B/op	      56 allocs/op
BenchmarkSmall/EncodePipe____________-4         	   10000	    179280 ns/op	   36220 B/op	      63 allocs/op
BenchmarkSmall/EncodePipeE___________-4         	   10000	    148650 ns/op	   36219 B/op	      63 allocs/op
BenchmarkSmall/EncodeBufferPool______-4         	   10000	    122229 ns/op	    3071 B/op	      54 allocs/op
BenchmarkSmall/EncodeBufferPoolE_____-4         	   10000	    122518 ns/op	    3071 B/op	      54 allocs/op
BenchmarkSmall/EncodeBPipe___________-4         	   10000	    143932 ns/op	   44571 B/op	      67 allocs/op
BenchmarkSmall/SharedBufferEncode____-4         	   10000	    131563 ns/op	   35493 B/op	      53 allocs/op
BenchmarkSmall/FastMarshal___________-4         	   10000	    123931 ns/op	    3528 B/op	      58 allocs/op
BenchmarkMiddle/MarshalE_____________-4         	    5000	    232813 ns/op	   35786 B/op	      51 allocs/op
BenchmarkMiddle/EncodeDefaultBuffer__-4         	    5000	    240336 ns/op	   35777 B/op	      51 allocs/op
BenchmarkMiddle/EncodeDefaultBufferE_-4         	    5000	    238095 ns/op	   35771 B/op	      51 allocs/op
BenchmarkMiddle/EncodeNewBuffer______-4         	    5000	    231755 ns/op	   35771 B/op	      51 allocs/op
BenchmarkMiddle/EncodeNewBufferE_____-4         	    5000	    226533 ns/op	   35771 B/op	      51 allocs/op
BenchmarkMiddle/EncodeNewNilBuffer___-4         	    5000	    241330 ns/op	   35770 B/op	      51 allocs/op
BenchmarkMiddle/EncodeNewNilBufferE__-4         	    5000	    237522 ns/op	   35770 B/op	      51 allocs/op
BenchmarkMiddle/EncodeReservedBuffer_-4         	    5000	    232825 ns/op	   60547 B/op	      51 allocs/op
BenchmarkMiddle/EncodeReservedBufferE-4         	   10000	    245071 ns/op	   60530 B/op	      51 allocs/op
BenchmarkMiddle/EncodeCheatBuffer____-4         	    5000	    232470 ns/op	   35802 B/op	      51 allocs/op
BenchmarkMiddle/EncodeCheatBufferE___-4         	    5000	    229806 ns/op	   35803 B/op	      51 allocs/op
BenchmarkMiddle/EncodePipe___________-4         	    5000	    262919 ns/op	   36071 B/op	      57 allocs/op
BenchmarkMiddle/EncodePipeE__________-4         	    5000	    267350 ns/op	   36065 B/op	      57 allocs/op
BenchmarkMiddle/EncodeBufferPool_____-4         	   10000	    223286 ns/op	    2805 B/op	      49 allocs/op
BenchmarkMiddle/EncodeBufferPoolE____-4         	    5000	    223139 ns/op	    2824 B/op	      49 allocs/op
BenchmarkMiddle/EncodeBPipe__________-4         	    5000	    285824 ns/op	   44460 B/op	      61 allocs/op
BenchmarkMiddle/SharedBufferEncode___-4         	    5000	    242880 ns/op	   35363 B/op	      47 allocs/op
BenchmarkMiddle/FastMarshal__________-4         	    5000	    305470 ns/op	  189443 B/op	      69 allocs/op
BenchmarkLarge/MarshalE______________-4         	      30	  42960238 ns/op	53506129 B/op	      92 allocs/op
BenchmarkLarge/EncodeDefaultBuffer___-4         	      30	  45497595 ns/op	48261589 B/op	      90 allocs/op
BenchmarkLarge/EncodeDefaultBufferE__-4         	      30	  47037591 ns/op	51408238 B/op	      91 allocs/op
BenchmarkLarge/EncodeNewBuffer_______-4         	      30	  58784422 ns/op	51408356 B/op	      91 allocs/op
BenchmarkLarge/EncodeNewBufferE______-4         	      20	  54064577 ns/op	53506226 B/op	      93 allocs/op
BenchmarkLarge/EncodeNewNilBuffer____-4         	      30	  50964893 ns/op	51408273 B/op	      91 allocs/op
BenchmarkLarge/EncodeNewNilBufferE___-4         	      30	  50164134 ns/op	49310568 B/op	      91 allocs/op
BenchmarkLarge/EncodeReservedBuffer__-4         	      30	  50053409 ns/op	66081374 B/op	      92 allocs/op
BenchmarkLarge/EncodeReservedBufferE_-4         	      30	  45217957 ns/op	71325837 B/op	      93 allocs/op
BenchmarkLarge/EncodeCheatBuffer_____-4         	      30	  47436967 ns/op	57692932 B/op	      91 allocs/op
BenchmarkLarge/EncodeCheatBufferE____-4         	      30	  47570246 ns/op	55595292 B/op	      92 allocs/op
BenchmarkLarge/EncodePipe____________-4         	      30	  41600816 ns/op	22073331 B/op	     105 allocs/op
BenchmarkLarge/EncodePipeE___________-4         	      30	  40070407 ns/op	26268860 B/op	     105 allocs/op
BenchmarkLarge/EncodeBufferPool______-4         	      30	  40588810 ns/op	10502117 B/op	      82 allocs/op
BenchmarkLarge/EncodeBufferPoolE_____-4         	      30	  43285747 ns/op	11551325 B/op	      83 allocs/op
BenchmarkLarge/EncodeBPipe___________-4         	      30	  49959702 ns/op	31520816 B/op	     105 allocs/op
BenchmarkLarge/SharedBufferEncode____-4         	      30	  40172115 ns/op	 4241193 B/op	      86 allocs/op
BenchmarkLarge/FastMarshal___________-4         	      10	 157478017 ns/op	207868536 B/op	     134 allocs/op
BenchmarkSimple/MarshalE_____________-4         	   10000	    122242 ns/op	    2700 B/op	      44 allocs/op
BenchmarkSimple/EncodeDefaultBuffer__-4         	   10000	    121896 ns/op	    2683 B/op	      44 allocs/op
BenchmarkSimple/EncodeDefaultBufferE_-4         	   10000	    119651 ns/op	    2683 B/op	      44 allocs/op
BenchmarkSimple/EncodeNewBuffer______-4         	   10000	    123248 ns/op	    2683 B/op	      44 allocs/op
BenchmarkSimple/EncodeNewBufferE_____-4         	   10000	    121143 ns/op	    2683 B/op	      44 allocs/op
BenchmarkSimple/EncodeNewNilBuffer___-4         	   10000	    122921 ns/op	    2683 B/op	      44 allocs/op
BenchmarkSimple/EncodeNewNilBufferE__-4         	   10000	    122606 ns/op	    2683 B/op	      44 allocs/op
BenchmarkSimple/EncodeReservedBuffer_-4         	   10000	    128893 ns/op	    3612 B/op	      44 allocs/op
BenchmarkSimple/EncodeReservedBufferE-4         	   10000	    120220 ns/op	    3612 B/op	      44 allocs/op
BenchmarkSimple/EncodeCheatBuffer____-4         	   10000	    120302 ns/op	    2716 B/op	      44 allocs/op
BenchmarkSimple/EncodeCheatBufferE___-4         	   10000	    130089 ns/op	    2716 B/op	      44 allocs/op
BenchmarkSimple/EncodePipe___________-4         	   10000	    137578 ns/op	   35682 B/op	      51 allocs/op
BenchmarkSimple/EncodePipeE__________-4         	   10000	    145168 ns/op	   35681 B/op	      51 allocs/op
BenchmarkSimple/EncodeBufferPool_____-4         	   10000	    126810 ns/op	    2539 B/op	      42 allocs/op
BenchmarkSimple/EncodeBufferPoolE____-4         	   10000	    124524 ns/op	    2539 B/op	      42 allocs/op
BenchmarkSimple/EncodeBPipe__________-4         	   10000	    140407 ns/op	   44032 B/op	      55 allocs/op
BenchmarkSimple/SharedBufferEncode___-4         	   10000	    133282 ns/op	   34960 B/op	      41 allocs/op
BenchmarkSimple/FastMarshal__________-4         	   10000	    127011 ns/op	    2699 B/op	      44 allocs/op
BenchmarkNested/MarshalE_____________-4         	   10000	    131572 ns/op	    3004 B/op	      47 allocs/op
BenchmarkNested/EncodeDefaultBuffer__-4         	   10000	    133251 ns/op	    2988 B/op	      47 allocs/op
BenchmarkNested/EncodeDefaultBufferE_-4         	   10000	    133596 ns/op	    2988 B/op	      47 allocs/op
BenchmarkNested/EncodeNewBuffer______-4         	   10000	    129522 ns/op	    2988 B/op	      47 allocs/op
BenchmarkNested/EncodeNewBufferE_____-4         	   10000	    131604 ns/op	    2987 B/op	      47 allocs/op
BenchmarkNested/EncodeNewNilBuffer___-4         	   10000	    135543 ns/op	    2988 B/op	      47 allocs/op
BenchmarkNested/EncodeNewNilBufferE__-4         	   10000	    146598 ns/op	    2988 B/op	      47 allocs/op
BenchmarkNested/EncodeReservedBuffer_-4         	   10000	    140895 ns/op	    3724 B/op	      47 allocs/op
BenchmarkNested/EncodeReservedBufferE-4         	   10000	    143000 ns/op	    3724 B/op	      47 allocs/op
BenchmarkNested/EncodeCheatBuffer____-4         	   10000	    130834 ns/op	    3852 B/op	      48 allocs/op
BenchmarkNested/EncodeCheatBufferE___-4         	   10000	    128229 ns/op	    3851 B/op	      48 allocs/op
BenchmarkNested/EncodePipe___________-4         	   10000	    133040 ns/op	   35793 B/op	      53 allocs/op
BenchmarkNested/EncodePipeE__________-4         	   10000	    145393 ns/op	   35793 B/op	      53 allocs/op
BenchmarkNested/EncodeBufferPool_____-4         	   10000	    130878 ns/op	    2652 B/op	      45 allocs/op
BenchmarkNested/EncodeBufferPoolE____-4         	   10000	    130850 ns/op	    2652 B/op	      45 allocs/op
BenchmarkNested/EncodeBPipe__________-4         	   10000	    149349 ns/op	   44145 B/op	      57 allocs/op
BenchmarkNested/SharedBufferEncode___-4         	   10000	    134332 ns/op	   35072 B/op	      43 allocs/op
BenchmarkNested/FastMarshal__________-4         	   10000	    129807 ns/op	    3052 B/op	      48 allocs/op
```

## Author
Nao Yonashiro (@orisano)

## License
MIT
