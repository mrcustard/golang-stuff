package templates
import "html/template"
var HTML = template.New("html")
func init() {
  template.Must(HTML.New("index.html").Parse("<html>\n  <head>\n    <title>\n      Go Database! Manager\n    </title>\n    <link href=\"/css/{{.T}}/all.css\" rel=\"stylesheet\" media=\"screen\">\n    <script type=\"text/template\" id=\"result_templ\">\n			<pre><%= JSON.stringify(data, null, \"  \") %></pre>\n    <button id=\"decode\">Decode</button>\n  </script>\n  <script type=\"text/template\" id=\"api_endpoint_item_templ\">\n    <li data-endpoint-name=\"<%= api_endpoint.name %>\"><%= api_endpoint.name %></li>\n  </script>\n  <script type=\"text/template\" id=\"api_endpoint_templ\">\n    <textarea id=\"code\"></textarea>\n    <button id=\"execute\">Execute</button>\n  </script>\n  <script type=\"text/template\" id=\"node_link_templ\">\n    <tr data-addr=\"<%= node.json_addr %>\" class=\"node\"><td><%= node.gob_addr %></td><td><%= node.hexpos %></td></tr>	\n  </script>\n  <script src=\"/js/{{.T}}/all.js\" type=\"text/javascript\"></script>\n</head>\n<body>		\n  <div id=\"chord_container\">\n    <canvas width=\"3000\" height=\"2000\" id=\"chord\"></canvas>\n  </div>\n  <div id=\"nodes_container\">\n    <table class=\"table table-striped\" id=\"nodes\">\n      <caption>nodes</caption>\n      <tr>\n	<th>address</th>\n	<th>position</th>\n      </tr>\n    </table>\n    <p><a href=\"http://zond.github.com/god/\">Architectural documentation</a></p>\n    <p><a href=\"http://godoc.org/github.com/zond/god/client\">Go client API documentation</a></p>\n    <form class=\"form-horizontal\">\n      <div class=\"control-group\">\n	<label class=\"control-label\" for=\"meth\">call method</label>\n	<div class=\"controls\">\n	  <div class=\"btn-group\">\n	    <a class=\"btn dropdown-toggle\" data-toggle=\"dropdown\" href=\"#\">\n	      Endpoint\n	      <span class=\"caret\"></span>\n	    </a>\n	    <ul id=\"endpoints\" class=\"dropdown-menu\">\n	    </ul>\n	  </div>\n	</div>\n      </div>\n    </form>\n    <div id=\"code_container\"></div>\n    <div id=\"result_container\"></div>\n  </div>\n  <div id=\"node_container\">\n    <a class=\"close\" id=\"hide_node_container\" href=\"#\">&times;</a>\n    <table class=\"table table-condensed\">\n      <caption>node</caption>\n      <tr>\n	<td>gob rpc address</td>\n	<td id=\"node_gob_addr\"></td>\n      </tr>\n      <tr>\n	<td>JSON/HTTP rpc address</td>\n	<td id=\"node_json_addr\"></td>\n      </tr>\n      <tr>\n	<td>position</td>\n	<td id=\"node_pos\"></td>\n      </tr>\n      <tr>\n	<td>owned keys</td>\n	<td id=\"node_owned_keys\"></td>\n      </tr>\n      <tr>\n	<td>held keys</td>\n	<td id=\"node_held_keys\"></td>\n      </tr>\n      <tr>\n	<td>load</td>\n	<td id=\"node_load\"></td>\n      </tr>\n    </table>\n  </div>\n</body>\n</html>\n"))
}
