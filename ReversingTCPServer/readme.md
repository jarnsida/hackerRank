<div class="ps-content-wrapper-v0">
<p>Implement a TCP server that accepts messages in the form of a string. Reverse the messages and return them to the client.</p>

<p>&nbsp;</p>

<p>The stub function accepts an array of strings, <em>messages</em>, runs the server, waits for it to be ready, then sends the messages one by one to the server and prints responses. The server address is given in the constant <em>address </em>along with<em> maxBufferSize </em>for reading/writing purposes.</p>

<p>&nbsp;</p>

<p class="section-title">Example</p>

<p><em>messages= ["adi", "glm", "tuv"]</em></p>

<p>&nbsp;</p>

<p>The TCP server should respond with <em>["ida","mlg","vut"]</em>.</p>
&nbsp;

<p class="section-title">Function Description</p>

<p>Complete the function <em>TCPServer</em> in the editor below.</p>

<p>TCPServer has the following parameter:</p>

<p>&nbsp;&nbsp;&nbsp;&nbsp;<em>bool ready:</em>&nbsp; a channel that indicates whether the server is ready</p>

<p>&nbsp;</p>

<p class="section-title">Returns</p>

<ul>
</ul>

<p>&nbsp;&nbsp;&nbsp;&nbsp;<em>None:</em>&nbsp;&nbsp;&nbsp;all responses should be read from the TCP connection initiated by a TCP client&nbsp;</p>

<p>&nbsp;</p>

<p class="section-title">Constraints</p>

<ul>
	<li>The total number of messages does not exceed 500.</li>
	<li>Each message contains a maximum of 1000 lower case English letters only.</li>
</ul>

<p>&nbsp;</p>
<!-- <StartOfInputFormat> DO NOT REMOVE THIS LINE-->

<details><summary class="section-title">Input Format For Custom Testing</summary>

<div class="collapsable-details">
<p>The first line contains an integer, <em>n</em>, the number of elements in <em>messages</em>.<br>
Each of the next <em>i</em> lines (where <em>0 ≤ i &lt; n</em>) contains a string, <em>messages[i]</em>.</p>
</div>
</details>
<!-- </StartOfInputFormat> DO NOT REMOVE THIS LINE-->

<details open="open"><summary class="section-title">Sample Case 0</summary>

<div class="collapsable-details">
<p class="section-title">Sample Input For Custom Testing</p>

<pre>STDIN&nbsp;&nbsp;&nbsp;&nbsp;Function
-----&nbsp;&nbsp;&nbsp;&nbsp;--------
3&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;→&nbsp;&nbsp;messages[] size n = 3
abc&nbsp;&nbsp;&nbsp;→&nbsp;&nbsp;messages = ['abc', 'def', 'ghi']
def
ghi
</pre>

<p class="section-title">Sample Output</p>

<pre>cba
fed
ihg</pre>

<p class="section-title">Explanation</p>

<p>The server reverses the strings and sends them back to the main function.</p>
</div>
</details>

<details><summary class="section-title">Sample Case 1</summary>

<div class="collapsable-details">
<p class="section-title">Sample Input For Custom Testing</p>

<pre>4
eno
owt
eerht
ruof
</pre>

<p class="section-title">Sample Output</p>

<pre>one
two
three
four</pre>

<p class="section-title">Explanation</p>

<p>The server reverses the strings and sends them back to the main function.</p>
</div>
</details>
</div>