
<p>A fibonacci with modulo sequence is defined for a given modulo, mod, as follows:</p>

<ul>
	<li>fib[x] = (fib[x - 1] + fib[x - 2]) % mod, x &gt; 2</li>
	<li>fib[1] = 1, fib[2] = 2</li>
</ul>

<p>&nbsp;</p>

<p>[1, 2, 3, 5, 8, 13, 21 ...] are the first few terms of the Fibonacci sequence. In this task, you are required to write a function that runs a server that generates a Fibonacci sequence.</p>

<p>&nbsp;</p>

<p>It accepts boolean requests and returns the next number in the sequence. It should have rate limiter which should make the server send every response no earlier than 10 milliseconds after the previous one. The server should be created with two arguments - a boolean channel that accepts requests and a result int channel through which every result will be returned. The main function will also accept two arguments: the number of results that will be skipped from the beginning, and the number of results that will be printed to the console.</p>

<p>&nbsp;</p>

<p>Note: Since the numbers can be large, the server must return the number modulo 10<sup>9</sup>. For example, for the 50<sup>th</sup>&nbsp;Fibonacci number - 12586269025, the server should return 586269025.</p>
&nbsp;

<p>For example, using the arguments 2 and 5 results in <i>[3, 5, 8, 13, 21].</i></p>
&nbsp;

<p class="section-title">Function Description</p>

<p>Complete the function <i>ModuloFibonacciSequence</i> in the editor below. The function will return nothing as all its results will be returned through a channel.</p>

<p>&nbsp;</p>

<p><i>ModuloFibonacciSequence</i> has the following parameter(s):</p>

<p>&nbsp;&nbsp;&nbsp;&nbsp;<em>requestChan:</em>&nbsp; a channel of booleans (chan bool)</p>

<p>&nbsp;&nbsp;&nbsp;&nbsp;<em>resultChan:</em>&nbsp; a channel of integers (chan int)</p>

<p>&nbsp;</p>

<p class="section-title">Constraints</p>

<ul>
	<li><em>0 ≤ skip&nbsp;≤ 100</em></li>
	<li><em>1 ≤ total ≤ 200</em></li>
</ul>
<!-- <StartOfInputFormat> DO NOT REMOVE THIS LINE-->

<details open=""><summary class="section-title">Input Format For Custom Testing</summary>

<div class="collapsable-details">
<p>The first line contains an integer, <em>skip</em>, the number of elements that should be skipped from the beginning of the Fibonacci sequence.<br>
The second line contains an integer, <em>total</em>, the number of elements that should be printed as a result of FibonacciServer working.</p>
</div>
</details>
<!-- </StartOfInputFormat> DO NOT REMOVE THIS LINE-->

<details open=""><summary class="section-title">Sample Case 0</summary>

<div class="collapsable-details">
<p class="section-title">Sample Input For Custom Testing</p>

<pre>STDIN&nbsp;&nbsp;&nbsp;&nbsp;Function
-----&nbsp;&nbsp;&nbsp;&nbsp;--------
0&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;skip = 0
6&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;total = 6
</pre>

<p class="section-title">Sample Output</p>

<pre>1
2
3
5
8
13
</pre>

<p class="section-title">Explanation</p>

<p>0 means no elements will be skipped, so the sequence starts with the first element, 1. The 6 elements in the sequence are [1, 2, 3, 5, 8, 13]. None of these is large enough for the modulo to change.</p>
</div>
</details>

<details><summary class="section-title">Sample Case 1</summary>

<div class="collapsable-details">
<p class="section-title">Sample Input For Custom Testing</p>

<pre>50
4
</pre>

<p class="section-title">Sample Output</p>

<pre>951280099
316291173
267571272
583862445
</pre>

<p class="section-title">Explanation</p>

<p>&nbsp;</p>

<p>Skip the first 50 numbers in the sequence and print the modulo of the next 4 integers.</p>
</div>
</details>
</div>
</div></div></section></div></div>