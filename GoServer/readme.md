<div class="ps-content-wrapper-v0">
<p>Implement an HTTP server that has three&nbsp;routes and maintains a database of the world's largest lakes.</p>

<ul>
	<li>The first route runs a handler <em>postHandler</em>&nbsp;which&nbsp;accepts a POST request with JSON-encoded lake information and posts it to the database.

	<pre><code class="language-json">{
    "type": "post",
    "payload": {"id": "id00", "name": "Name of the lake", "area": 452001}
}</code></pre>

	<p>&nbsp;</p>
	</li>
	<li>The second route runs a handler <em>deleteHandler</em>&nbsp;which deletes the lake from the database by id. If the payload id is not present in the data set, the server returns a 404 response.
	<pre><code class="language-json">{
    "type": "delete",
    "payload": "id00"
}</code></pre>
	</li>
	<li>The third route runs a handler <em>getHandler</em>&nbsp;which takes a lake from the database by id and returns it to the caller which prints the name and the area of the lake<em>. </em>If the id is not found in the database, the server returns a string message "404 Not Found". Otherwise, it returns the payload object corresponding to the id.
	<pre><code class="language-json">{
    "type": "get",
    "payload": "id00"
}</code></pre>
	</li>
</ul>

<p>Implement the server which will be queried by the program and print the output to STDOUT.</p>

<p>&nbsp;</p>

<p>Note:</p>

<p>The program uses these structs:</p>

<pre><code class="language-cpp">type Lake struct {

&nbsp; &nbsp; Name string

&nbsp; &nbsp; Area int32

}



type Action&nbsp;struct {

&nbsp; &nbsp; Type string

&nbsp; &nbsp; Payload string

}</code></pre>

<p>&nbsp;</p>

<p>The base URL is contained in the global variable <em>address</em>. Data is stored under the <em>store</em>&nbsp;variable which is&nbsp;the map[string] <em>Lake</em>.</p>

<p>&nbsp;</p>

<p class="section-title">Constraints</p>

<ul>
	<li>The total number of queries does not exceed 1000.&nbsp;</li>
	<li>The name and id strings consist of no more than 100 lowercase and uppercase English characters only.&nbsp;</li>
	<li>All the post ids are unique.</li>
</ul>

<p>&nbsp;</p>
<!-- <StartOfInputFormat> DO NOT REMOVE THIS LINE-->

<details><summary class="section-title">Input Format For Custom Testing</summary>

<div class="collapsable-details">
<p>The first line contains an integer, <em>n</em>, the number of elements in <em>actions</em>.<br>
Each line <em>i</em> of the <em>n</em> subsequent lines (where <em>0 ≤ i &lt; n</em>) contains a string, <em>actions[i]</em>.</p>
</div>
</details>
<!-- </StartOfInputFormat> DO NOT REMOVE THIS LINE-->

<details open="open"><summary class="section-title">Sample Case 0</summary>

<div class="collapsable-details">
<p class="section-title">Sample Input For Custom Testing</p>

<pre>4
{type:'post', payload:`{id:'id0000', name:'Malawi', area:29500}`}
{type:'post', payload:`{id:'id0001', name:'Great Bear Lake', area:31000}`}
{type:'get', payload:'id0001'}
{type:'get', payload:'id0000'}</pre>

<p class="section-title">Sample Output</p>

<pre>Great Bear Lake
31000
Malawi
29500</pre>

<p class="section-title">Explanation</p>

<p>The first two actions call postHandler and put two objects in the <em>store</em>. The 3rd and the 4th actions get items from the <em>store</em>&nbsp;by&nbsp;corresponding ids.</p>
</div>
</details>

<details><summary class="section-title">Sample Case 1</summary>

<div class="collapsable-details">
<p class="section-title">Sample Input For Custom Testing</p>

<pre>5
{type:'post', payload:`{id:'id0000', name:'Malawi', area:29500}`}
{type:'post', payload:`{id:'id0001', name:'Great Bear Lake', area:31000}`}
{type:'delete', payload:'id0000'}
{type:'get', payload:'id0001'}
{type:'get', payload:'id0000'}</pre>

<p class="section-title">Sample Output</p>

<pre>Great Bear Lake
31000
404 Not Found</pre>

<p class="section-title">Explanation</p>

<p>The first two actions call postHandler and put two objects in the store. The 3rd action deletes the object with id "id0000". The 4th and the 5th actions get items from the store by corresponding ids, but since the object "id0000" was deleted, the 5th message leads to printing a "404 Not Found" error.</p>
</div>
</details>
</div>