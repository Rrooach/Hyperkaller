package html
const style = `
#topbar {
	padding: 5px 10px;
	background: #E0EBF5;
}

#topbar a {
	color: #375EAB;
	text-decoration: none;
}

h1, h2, h3, h4 {
	margin: 0;
	padding: 0;
	color: #375EAB;
	font-weight: bold;
}

table {
	border: 1px solid #ccc;
	margin: 20px 5px;
	border-collapse: collapse;
	white-space: nowrap;
	text-overflow: ellipsis;
	overflow: hidden;
}

table caption {
	font-weight: bold;
}

table td, table th {
	vertical-align: top;
	padding: 2px 8px;
	text-overflow: ellipsis;
	overflow: hidden;
}

.namespace {
	font-weight: bold;
	font-size: large;
	color: #375EAB;
}

.position_table {
	border: 0px;
	margin: 0px;
	width: 100%;
	border-collapse: collapse;
}

.position_table td, .position_table tr {
	vertical-align: center;
	padding: 0px;
}

.position_table .search {
	text-align: right;
}

.list_table td, .list_table th {
	border-left: 1px solid #ccc;
}

.list_table th {
	background: #F4F4F4;
}

.list_table tr:nth-child(2n) {
	background: #F4F4F4;
}

.list_table tr:hover {
	background: #ffff99;
}

.list_table .namespace {
	width: 100pt;
	max-width: 100pt;
}

.list_table .title {
	width: 350pt;
	max-width: 350pt;
}

.list_table .commit_list {
	width: 500pt;
	max-width: 500pt;
}

.list_table .tag {
	font-family: monospace;
	font-size: 8pt;
	width: 40pt;
	max-width: 40pt;
}

.list_table .opts {
	width: 40pt;
	max-width: 40pt;
}

.list_table .status {
	width: 250pt;
	max-width: 250pt;
}

.list_table .patched {
	width: 60pt;
	max-width: 60pt;
	text-align: center;
}

.list_table .kernel {
	width: 80pt;
	max-width: 80pt;
}

.list_table .maintainers {
	width: 150pt;
	max-width: 150pt;
}

.list_table .result {
	width: 60pt;
	max-width: 60pt;
}

.list_table .stat {
	width: 55pt;
	max-width: 55pt;
	font-family: monospace;
	text-align: right;
}

.list_table .date {
	width: 60pt;
	max-width: 60pt;
	font-family: monospace;
	text-align: right;
}

.list_table .stat_name {
	width: 150pt;
	max-width: 150pt;
	font-family: monospace;
}

.list_table .stat_value {
	width: 120pt;
	max-width: 120pt;
	font-family: monospace;
}

.bad {
	color: #f00;
	font-weight: bold;
}

.inactive {
	color: #888;
}

.plain {
	text-decoration: none;
}

textarea {
	width:100%;
	font-family: monospace;
}

.mono {
	font-family: monospace;
}

.info_link {
	color: #25a7db;
	text-decoration: none;
}
`
const js = `
// Copyright 2018 syzkaller project authors. All rights reserved.
// Use of this source code is governed by Apache 2 LICENSE that can be found in the LICENSE file.

function sortTable(item, colName, conv, desc = false) {
	table = item.parentNode.parentNode.parentNode.parentNode;
	rows = table.rows;
	col = findColumnByName(rows[0].getElementsByTagName("th"), colName);
	values = [];
	for (i = 1; i < rows.length; i++)
		values.push([conv(rows[i].getElementsByTagName("td")[col].textContent), rows[i]]);
	if (desc)
		desc = !isSorted(values.slice().reverse())
	else
		desc = isSorted(values);
	values.sort(function(a, b) {
		if (a[0] == b[0]) return 0;
		if (desc && a[0] > b[0] || !desc && a[0] < b[0]) return -1;
		return 1;
	});
	for (i = 0; i < values.length; i++)
		table.tBodies[0].appendChild(values[i][1]);
	return false;
}

function findColumnByName(headers, colName) {
	for (i = 0; i < headers.length; i++) {
		if (headers[i].textContent == colName)
			return i;
	}
	return 0;
}

function isSorted(values) {
	for (i = 0; i < values.length - 1; i++) {
		if (values[i][0] > values[i + 1][0])
			return false;
	}
	return true;
}

function textSort(v) { return v.toLowerCase(); }
function numSort(v) { return -parseInt(v); }
function floatSort(v) { return -parseFloat(v); }
function reproSort(v) { return v == "C" ? 0 : v == "syz" ? 1 : 2; }
function patchedSort(v) { return v == "" ? -1 : parseInt(v); }

function timeSort(v) {
	if (v == "now")
		return 0;
	m = v.indexOf('m');
	h = v.indexOf('h');
	d = v.indexOf('d');
	if (m > 0 && h < 0)
		return parseInt(v);
	if (h > 0 && m > 0)
		return parseInt(v) * 60 + parseInt(v.substring(h + 1));
	if (d > 0 && h > 0)
		return parseInt(v) * 60 * 24 + parseInt(v.substring(d + 1)) * 60;
	if (d > 0)
		return parseInt(v) * 60 * 24;
	return 1000000000;
}
`
