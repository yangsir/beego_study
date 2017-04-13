{{define "navbar"}}

<a class="navbar-brand" href="http://webyang.net" target="_blank">WebYang.NET</a>
<div>
	<ul class="nav navbar-nav">
		<li {{if eq "home" .isActive}}class="active"{{end}}><a href="/">首页</a></li>
		<li {{if eq "category" .isActive}}class="active"{{end}}><a href="/category">分类</a></li>
		<li {{if eq "topic" .isActive}}class="active"{{end}}><a href="/topic">文章</a></li>
	</ul>
</div>

<div class="pull-right">
	<ul class="nav navbar-nav">
		{{if .IsLogin}}
		<li><a href="/login?exit=true">退出登录</a></li>
		{{else}}
		<li><a href="/login">管理员登录</a></li>
		{{end}}
	</ul>
</div>
{{end}}