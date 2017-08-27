package templates

import (
	"html/template"
)
var(
	IndexTemplate = template.Must(template.New("index").Parse(`
	<style>
		a {color:blue};
	</style>
	<div>
		<h2>A HTTP test server for clients</h2>
	</div>
	<div>
	<h3>ENDPOINTS</h3>
		<ul>
		<li><a href = "/">/</a>  Returns home page.</li>
		<li><a href = "/ip">/ip</a>  Returns origin ip.</li>
		<li><a href = "/uuid">/uuid</a>  Returns UUID.</li>		
		<li><a href = "/user-agent">/user-agent</a>  Returns user-agent.</li>		
		<li><a href = "/headers">/headers</a>  Return headers map.</li>
		<li><a href = "/get">/get</a> Returns GET data.</li>		
		</ul>
	</div>
	`))
)